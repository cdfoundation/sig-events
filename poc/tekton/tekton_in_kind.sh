#!/bin/bash
set -e -o pipefail

declare TEKTON_PIPELINE_VERSION TEKTON_TRIGGERS_VERSION TEKTON_DASHBOARD_VERSION
declare KNATIVE_VERSION KNATIVE_NET_KOURIER_VERSION

# This script deploys Tekton on a local kind cluster
# It creates a kind cluster and deploys pipeline, triggers and dashboard

# Prerequisites:
# - go 1.15+
# - docker (recommended 8GB memory config)
# - kind

# Notes:
# - Latest versions will be installed if not specified
# - If a kind cluster named "tekton" already exists this will fail
# - Local access to the dashboard requires port 9097 to be locally available

get_latest_release() {
  curl --silent "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

# Read command line options
while getopts ":c:p:t:d:k:" opt; do
  case ${opt} in
    c )
      CLUSTER_NAME=$OPTARG
      ;;
    p )
      TEKTON_PIPELINE_VERSION=$OPTARG
      ;;
    t )
      TEKTON_TRIGGERS_VERSION=$OPTARG
      ;;
    d )
      TEKTON_DASHBOARD_VERSION=$OPTARG
      ;;
    k )
      KNATIVE_VERSION=$OPTARG
      ;;
    \? )
      echo "Invalid option: $OPTARG" 1>&2
      echo 1>&2
      echo "Usage:  tekton_in_kind.sh [-c cluster-name -p pipeline-version -t triggers-version -d dashboard-version -i kourier|nginx -k knative-version]"
      ;;
    : )
      echo "Invalid option: $OPTARG requires an argument" 1>&2
      ;;
  esac
done
shift $((OPTIND -1))

# Check and defaults input params
export KIND_CLUSTER_NAME=${CLUSTER_NAME:-"tekton"}

if [ -z "$TEKTON_PIPELINE_VERSION" ]; then
  TEKTON_PIPELINE_VERSION=$(get_latest_release tektoncd/pipeline)
fi
if [ -z "$TEKTON_TRIGGERS_VERSION" ]; then
  TEKTON_TRIGGERS_VERSION=$(get_latest_release tektoncd/triggers)
fi
if [ -z "$TEKTON_DASHBOARD_VERSION" ]; then
  TEKTON_DASHBOARD_VERSION=$(get_latest_release tektoncd/dashboard)
fi
if [ -z "$KNATIVE_VERSION" ]; then
  MAIN_CONTAINER_HTTP_PORT=80
  MAIN_CONTAINER_HTTPS_PORT=443
  ALT_CONTAINER_HTTP_PORT=22222 # unused
  ALT_CONTAINER_HTTPS_PORT=22223 # unused
  TEKTON_PORT=80
else
  MAIN_CONTAINER_HTTP_PORT=31080
  MAIN_CONTAINER_HTTPS_PORT=31443
  ALT_CONTAINER_HTTP_PORT=80
  ALT_CONTAINER_HTTPS_PORT=443
  TEKTON_PORT=8080
fi

echo "===> Creating a Kind Cluster"
# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

# Create the kind cluster
# create a cluster with the local registry enabled in containerd
running_cluster=$(kind get clusters | grep tekton || true)
if [ "${running_cluster}" != "$KIND_CLUSTER_NAME" ]; then
 cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    kubeadmConfigPatches:
      - |
        kind: InitConfiguration
        nodeRegistration:
          kubeletExtraArgs:
            node-labels: "ingress-ready=true"
    extraPortMappings:
      - containerPort: $MAIN_CONTAINER_HTTP_PORT
        hostPort: 80
        protocol: TCP
      - containerPort: $MAIN_CONTAINER_HTTPS_PORT
        hostPort: 443
        protocol: TCP
      - containerPort: $ALT_CONTAINER_HTTP_PORT
        hostPort: 8080
        protocol: TCP
      - containerPort: $ALT_CONTAINER_HTTPS_PORT
        hostPort: 8443
        protocol: TCP
  - role: worker
featureGates:
  "EphemeralContainers": true
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
EOF
fi

# connect the registry to the cluster network
# (the network may already be connected)
docker network connect "kind" "${reg_name}" || true

if [ -n "$KNATIVE_VERSION" ]; then
  echo "===> Deploying Knative controller"
  export KNATIVE_VERSION
  curl -sL https://raw.githubusercontent.com/csantanapr/knative-kind/master/02-serving.sh | sh

  echo "===> Deploying Kourier"
  export KNATIVE_NET_KOURIER_VERSION=${KNATIVE_NET_KOURIER_VERSION:-$KNATIVE_VERSION}
  curl -sL https://raw.githubusercontent.com/csantanapr/knative-kind/master/02-kourier.sh | sh

  EXTERNAL_IP="127.0.0.1"
  KNATIVE_DOMAIN="$EXTERNAL_IP.nip.io"
  kubectl patch configmap -n knative-serving config-domain -p "{\"data\": {\"$KNATIVE_DOMAIN\": \"\"}}"
fi

echo "===> Deploying the Ingress controller"
# Deploy the ingress
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=160s

echo "===> RBAC and secrets"
# Install some basic RBAC and secrets needed by triggers
kubectl create -f rbac.yaml || true
if [ -f .secrets/icr.yaml ]; then
  kubectl create -f .secrets/icr.yaml || true
  kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "all-icr-io"}]}' || true
fi
if [ -f .secrets/config.json ]; then
  kubectl create secret generic regcred \
    --from-file=config.json=.secrets/config.json \
    --from-file=.dockerconfigjson=.secrets/config.json \
    --type=kubernetes.io/dockerconfigjson || true
fi

kubectl create namespace tekton-pipelines || true

cat <<EOF | kubectl create -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: tekton-dashboard
  namespace: tekton-pipelines
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /\$2
    nginx.ingress.kubernetes.io/configuration-snippet: |
      rewrite ^(/[a-z1-9\-]*)$ \$1/ redirect;
spec:
  rules:
  - http:
      paths:
      - path: /dashboard(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: tekton-dashboard
            port:
              number: 9097
EOF

echo "===> Install Tekton"

echo export TEKTON_PIPELINE_VERSION="$TEKTON_PIPELINE_VERSION"
echo export TEKTON_TRIGGERS_VERSION="$TEKTON_TRIGGERS_VERSION"
echo export TEKTON_DASHBOARD_VERSION="$TEKTON_DASHBOARD_VERSION"

# Install Tekton Pipeline, Triggers and Dashboard
kubectl apply -f "https://storage.googleapis.com/tekton-releases/pipeline/previous/${TEKTON_PIPELINE_VERSION}/release.yaml"
kubectl apply -f "https://storage.googleapis.com/tekton-releases/triggers/previous/${TEKTON_TRIGGERS_VERSION}/release.yaml"
kubectl apply -f "https://github.com/tektoncd/dashboard/releases/download/${TEKTON_DASHBOARD_VERSION}/tekton-dashboard-release.yaml"

if [ -f .secrets/icr.yaml ]; then
  kubectl create -f .secrets/icr.yaml -n tekton-pipelines || true
  kubectl patch serviceaccount tekton-pipelines-controller -n tekton-pipelines -p '{"imagePullSecrets": [{"name": "all-icr-io"}]}' || true
fi

# # Wait until all pods are ready
kubectl wait -n tekton-pipelines --for=condition=ready pods --all --timeout=120s

echo Tekton Dashboard available at http://localhost:"${TEKTON_PORT}"/dashboard/ after installation
