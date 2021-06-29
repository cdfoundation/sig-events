#!/bin/bash
set -e -o pipefail

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

declare TEKTON_PIPELINE_VERSION TEKTON_TRIGGERS_VERSION TEKTON_DASHBOARD_VERSION TEKTON_CLOUDEVENTS_PATH
declare KNATIVE_VERSION KNATIVE_NET_KOURIER_VERSION
declare KEPTN_CDF_TRANSLATOR_PATH CDF_EVENTS_KEPTN_ADAPTER_PATH KEPTN_PROJECT KEPTN_SERVICE

export KO_DOCKER_REPO=kind.local

# This script deploys all the software components required for the PoC
# on a local laptop running docker.
# It creates a kind cluster named cde and deploys to it:
# - Tekton
# - Tekton experiment cloudevents controller (TBD)
# - Keptn (TBD)
# - Keptn inbound translation layer (TBD)
# - Keptn output adapter layer (TBD)
# - Nginx ingress controller
# - Knative (optional - partly implemented)
# - Istio (optional - not implemented)
# It deploys a container registry available at localhost:5000 and reachable from
# the kind cluster.

# Prerequisites:
# - go 1.15+
# - docker (recommended 8GB memory config)
# - kind 0.11.1
# - tkn CLI
# - keptn CLI (`curl -sL https://get.keptn.sh | bash`)

# Notes:
# - Latest versions will be installed if not specified
# - If a kind cluster named "cde" already exists this will fail

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
      echo "Usage:  tekton_in_kind.sh [-c cluster-name -p pipeline-version -t triggers-version -d dashboard-version -k knative-version]"
      ;;
    : )
      echo "Invalid option: $OPTARG requires an argument" 1>&2
      ;;
  esac
done
shift $((OPTIND -1))

# Check and defaults input params
export KIND_CLUSTER_NAME=${CLUSTER_NAME:-"cde"}

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

echo "===> Creating a local Container Registry"
# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi
echo
echo "Accessing your local container registry:"
echo "- From you local laptop: localhost:5000 (docker tag <image> localhost:5000/path/to/<image>)"
echo "- From within the cluster:"
echo "  - The container runtime can pull images from localhost:5000"
echo "  - The nodes / containers can reach the registry at kind-registry:5000"
echo "    Since it's not localhost, most tools will assume https. Use an \"--insecure\" flag if available"
echo

echo "===> Creating a Kind Cluster (1 control plan, 2 worker nodes"
# Create the kind cluster
# create a cluster with the local registry enabled in containerd
running_cluster=$(kind get clusters | grep "$KIND_CLUSTER_NAME" || true)
if [ "${running_cluster}" != "$KIND_CLUSTER_NAME" ]; then
 cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    image: kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6
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
    image: kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6
  - role: worker
    image: kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6
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

# Knative (serving only for now)
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
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.47.0/deploy/static/provider/kind/deploy.yaml
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=160s

echo "===> RBAC and secrets"
# Install some basic RBAC and secrets needed by triggers
kubectl create -f tekton/rbac.yaml || true

echo "===> Install Tekton"

kubectl create namespace tekton-pipelines || true

# Deploy an ingress for the Tekton Dashboard
cat <<EOF | kubectl create -f - || true
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

echo export TEKTON_PIPELINE_VERSION="$TEKTON_PIPELINE_VERSION"
echo export TEKTON_TRIGGERS_VERSION="$TEKTON_TRIGGERS_VERSION"
echo export TEKTON_DASHBOARD_VERSION="$TEKTON_DASHBOARD_VERSION"

# Install Tekton Pipeline, Triggers and Dashboard
kubectl apply -f "https://storage.googleapis.com/tekton-releases/pipeline/previous/${TEKTON_PIPELINE_VERSION}/release.yaml"
kubectl apply -f "https://storage.googleapis.com/tekton-releases/triggers/previous/${TEKTON_TRIGGERS_VERSION}/release.yaml"
kubectl wait --for condition=established --timeout=60s crd -l app.kubernetes.io/part-of=tekton-triggers
kubectl apply -f "https://storage.googleapis.com/tekton-releases/triggers/previous/${TEKTON_TRIGGERS_VERSION}/interceptors.yaml" || true
kubectl apply -f "https://github.com/tektoncd/dashboard/releases/download/${TEKTON_DASHBOARD_VERSION}/tekton-dashboard-release.yaml"

# Wait until all pods are ready
kubectl wait -n tekton-pipelines --for=condition=ready pods --all --timeout=120s

# Configure Tekton
kubectl patch cm feature-flags -n tekton-pipelines -p '{"data": {"enable-tekton-oci-bundles": "true"}}'
kubectl create namespace production || true

echo Tekton Dashboard available at http://localhost:"${TEKTON_PORT}"/dashboard/ after installation

# Install keptn

kubectl create namespace keptn || true

# Deploy an ingress for Keptn Bridge and API
cat <<EOF | kubectl create -f - || true
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: keptn
  namespace: keptn
spec:
  rules:
  - host: keptn-127.0.0.1.nip.io
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: api-gateway-nginx
            port:
              number: 80
EOF

keptn install --use-case=continuous-delivery --yes -q

# Keptn Auth
export KEPTN_ENDPOINT=http://keptn-127.0.0.1.nip.io/api
export KEPTN_API_TOKEN=$(kubectl get secret keptn-api-token -n keptn -ojsonpath={.data.keptn-api-token} | base64 --decode)
keptn auth --endpoint=$KEPTN_ENDPOINT --api-token=$KEPTN_API_TOKEN

# Keptn Create Project and Service
KEPTN_PROJECT=${KEPTN_PROJECT:-cde}
KEPTN_SERVICE=${KEPTN_SERVICE:-poc}
keptn create project "$KEPTN_PROJECT" --shipyard="$BASE_DIR/resources/shipyard.yaml"
keptn create service "$KEPTN_SERVICE" --project="$KEPTN_PROJECT"

echo "To obtain bridge credentials, run:"
echo "keptn configure bridge -o"

# Install the keptn inbound and outbound adapters
export GO111MODULE=on

echo "===> Install Keptn Inboud"
KEPTN_CDF_TRANSLATOR_PATH=${KEPTN_CDF_TRANSLATOR_PATH:-${GOPATH}/src/github.com/salaboy/keptn-cdf-translator}
pushd "$KEPTN_CDF_TRANSLATOR_PATH"
ko apply -f config/
popd

echo "===> Install Keptn Outbound"
CDF_EVENTS_KEPTN_ADAPTER_PATH=${CDF_EVENTS_KEPTN_ADAPTER_PATH:-${GOPATH}/src/github.com/salaboy/cdf-events-keptn-adapter}
pushd "$CDF_EVENTS_KEPTN_ADAPTER_PATH"
docker build --tag localhost:${reg_port}/cdevents/cdf-events-keptn-keptn-cdeventsadapter:latest .
docker push localhost:${reg_port}/cdevents/cdf-events-keptn-adapter:latest
kubectl apply -f deploy/service.yaml
# Set the correct pub-sub topics
kubectl set env deployment/helm-service -n keptn PUBSUB_TOPIC=sh.keptn.event.service.create.finished,sh.keptn.event.rollback.triggered,sh.keptn.event.release.triggered,sh.keptn.event.action.triggered,sh.keptn.event.service.delete.finished -c distributor
kubectl set env deployment/cdf-events-keptn-adapter -n keptn PUBSUB_TOPIC=sh.keptn.event.deployment.triggered -c distributor
kubectl set env deployment/cdf-events-keptn-adapter -n keptn PUBSUB_URL="" -c distributor
pushd

echo "===> Install Tekton CloudEvents controller"
# Install the Tekton cloudevents experimental controller
TEKTON_CLOUDEVENTS_PATH=${TEKTON_CLOUDEVENTS_PATH:-${GOPATH}/src/github.com/tektoncd/experimental/cloudevents}
pushd "$TEKTON_CLOUDEVENTS_PATH"
ko apply -f config/
popd

# Configure the cloudevents controller to talk to the keptn inbound adapter
kubectl patch cm config-defaults -n tekton-cloudevents -p '{"data": {"default-cloud-events-sink": "http://keptn-cdevents.keptn:8080/events"}}'

# Install Tekton Resources
kubectl create -f "$BASE_DIR/tekton/" || true