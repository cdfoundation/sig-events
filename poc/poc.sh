#!/bin/bash
set -e -o pipefail

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

declare TEKTON_PIPELINE_VERSION TEKTON_TRIGGERS_VERSION TEKTON_DASHBOARD_VERSION TEKTON_CLOUDEVENTS_PATH
declare KNATIVE_VERSION KNATIVE_NET_CONTOUR_VERSION KNATIVE_EVENTING_VERSION
declare KEPTN_CDF_TRANSLATOR_PATH CDF_EVENTS_KEPTN_ADAPTER_PATH KEPTN_PROJECT KEPTN_SERVICE
declare BROKER_NAME

export KO_DOCKER_REPO=kind.local

# This script deploys all the software components required for the PoC
# on a local laptop running docker.
# It creates a kind cluster named cde and deploys to it:
# - Tekton
# - Tekton experiment cloudevents controller
# - Keptn (v0.9.x+)
# - Keptn inbound translation layer
# - Keptn output adapter layer
# - Knative + Courier Ingress
# It deploys a container registry available at localhost:5000 and reachable from
# the kind cluster.

# Prerequisites:
# - go 1.15+
# - docker (recommended 8GB memory config)
# - ko (https://github.com/google/ko)
# - kind 0.11.1 (https://kind.sigs.k8s.io/docs/user/quick-start/)
# - tkn CLI (https://github.com/tektoncd/cli)
# - keptn CLI (`curl -sL https://get.keptn.sh | bash`)
# - keptn inbound/outbound adapters and tekton cloudevent controller
#   https://github.com/salaboy/cdf-events-keptn-adapter - inbound
#   https://github.com/salaboy/keptn-cdf-translator - outbound
#   https://github.com/tektoncd/experimental - contains tekton cloudevent controller
#   should be cloned under $GOROOT/src/github.com/<org>/<repo> or alternatively
#   the corresponding PATH environment variables must be set (see the declare section above)
# - Spin CLI (https://spinnaker.io/docs/setup/other_config/spin/)


# Notes:
# - Latest versions will be installed if not specified
# - The script is written to be mostly idempotent

get_latest_release() {
  curl --silent "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

retry_command() {
  for counter in {1..5}; do
    $1 && break
    sleep $((counter + 1))
    [[ $counter -eq 5 ]] && echo "Failed!" && exit 1
    echo "Trying again. Try #$counter"
  done
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

KNATIVE_VERSION=${KNATIVE_VERSION:-1.3.0}

echo "Checking if needed repos can be found"
if [ ! -d "${KEPTN_CDF_TRANSLATOR_PATH:-${GOPATH}/src/github.com/salaboy/keptn-cdf-translator}" ]
then
    echo "Can not find required repo: github.com/salaboy/keptn-cdf-translator"
    exit 1 # die with error code 1
fi

if [ ! -d "${CDF_EVENTS_KEPTN_ADAPTER_PATH:-${GOPATH}/src/github.com/salaboy/cdf-events-keptn-adapter}" ]
then
    echo "Can not find required repo: github.com/salaboy/cdf-events-keptn-adapter"
    exit 1 # die with error code 1
fi

if [ ! -d "${TEKTON_CLOUDEVENTS_PATH:-${GOPATH}/src/github.com/tektoncd/experimental/cloudevents}" ]
then
    echo "Can not find required repo: github.com/tektoncd/experimental/cloudevents"
    exit 1 # die with error code 1
fi

# As we have "set -e" we only need to run the commands. If not found the script will exit.
echo "Checking if needed commands can be found"
go version > /dev/null
ko version > /dev/null
kind version > /dev/null
keptn > /dev/null
helm version > /dev/null
hal -v > /dev/null
spin --version >> /dev/null

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
      - containerPort: 31080 # expose port 31380 of the node to port 80 on the host, later to be use by contour ingress
        listenAddress: 127.0.0.1
        hostPort: 80
      - containerPort: 443
        hostPort: 443
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

# Wait for the cluster to be ready
kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n kube-system > /dev/null

# connect the registry to the cluster network
# (the network may already be connected)
docker network connect "kind" "${reg_name}" || true

# Knative (serving only for now)
echo "===> Deploying Knative controller"
export KNATIVE_VERSION
curl -sL https://raw.githubusercontent.com/csantanapr/knative-kind/master/02-serving.sh | bash

echo "===> Deploying Contour"
export KNATIVE_NET_CONTOUR_VERSION=${KNATIVE_NET_CONTOUR_VERSION:-$KNATIVE_VERSION}
curl -sL https://raw.githubusercontent.com/csantanapr/knative-kind/master/02-contour.sh | bash

echo "===> Deploying Knative Serving"
export KNATIVE_EVENTING_VERSION=${KNATIVE_EVENTING_VERSION:-$KNATIVE_VERSION}
export BROKER_NAME=${BROKER_NAME:-"events-broker"}
kubectl delete "broker/${BROKER_NAME}" > /dev/null || true
curl -sL https://raw.githubusercontent.com/csantanapr/knative-kind/master/04-eventing.sh | bash

EXTERNAL_IP="127.0.0.1"
KNATIVE_DOMAIN="knative-$EXTERNAL_IP.nip.io"
kubectl patch configmap -n knative-serving config-domain -p "{\"data\": {\"$KNATIVE_DOMAIN\": \"\"}}"

echo "===> RBAC and secrets"
# Install some basic RBAC and secrets needed by triggers
kubectl create -f tekton/rbac.yaml || true

echo "===> Install Tekton"

kubectl create namespace --save-config tekton-pipelines || true

# Deploy an ingress for the Tekton Dashboard
cat <<EOF | kubectl create -f - || true
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: tekton-dashboard
  namespace: tekton-pipelines
  annotations:
    kubernetes.io/ingress.class: "contour-external"
spec:
  rules:
  - host: tekton-127.0.0.1.nip.io
    http:
      paths:
      - path: /
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
kubectl apply -f "https://storage.googleapis.com/tekton-releases/pipeline/previous/${TEKTON_PIPELINE_VERSION}/release.yaml" > /dev/null
kubectl apply -f "https://storage.googleapis.com/tekton-releases/triggers/previous/${TEKTON_TRIGGERS_VERSION}/release.yaml" > /dev/null
kubectl wait --for condition=established --timeout=60s crd -l app.kubernetes.io/part-of=tekton-triggers
kubectl apply -f "https://storage.googleapis.com/tekton-releases/triggers/previous/${TEKTON_TRIGGERS_VERSION}/interceptors.yaml" > /dev/null
kubectl apply -f "https://github.com/tektoncd/dashboard/releases/download/${TEKTON_DASHBOARD_VERSION}/tekton-dashboard-release.yaml" > /dev/null
# Wait until all pods are ready
kubectl wait -n tekton-pipelines --for=condition=ready pods --all --timeout=120s

# Configure Tekton
retry_command "kubectl patch cm feature-flags -n tekton-pipelines -p {\"data\":{\"enable-tekton-oci-bundles\":\"true\"}}"
kubectl create namespace production || true

# Install keptn

keptn install --yes -q > /dev/null

# Deploy an ingress for Keptn Bridge and API
kubectl create -f - <<EOF || true
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: keptn
  namespace: keptn
  annotations:
    kubernetes.io/ingress.class: "contour-external"
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

# Keptn Auth
export KEPTN_ENDPOINT=http://keptn-127.0.0.1.nip.io/api
export KEPTN_API_TOKEN=$(kubectl get secret keptn-api-token -n keptn -ojsonpath={.data.keptn-api-token} | base64 --decode)
retry_command "keptn auth --endpoint=$KEPTN_ENDPOINT --api-token=$KEPTN_API_TOKEN"

# Keptn Create Project and Service (always re-create if exists)
KEPTN_PROJECT=${KEPTN_PROJECT:-cde}
KEPTN_SERVICE=${KEPTN_SERVICE:-poc}
keptn delete project "$KEPTN_PROJECT" > /dev/null || true
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
docker build --tag localhost:${reg_port}/cdevents/cdf-events-keptn-adapter:latest .
docker push localhost:${reg_port}/cdevents/cdf-events-keptn-adapter:latest
kubectl apply -f deploy/service.yaml
# Set the correct pub-sub topics
kubectl set env deployment/cdf-events-keptn-adapter -n keptn PUBSUB_TOPIC=sh.keptn.event.deployment.triggered -c distributor
pushd

echo "===> Install Tekton CloudEvents controller"
# Install the Tekton cloudevents experimental controller
TEKTON_CLOUDEVENTS_PATH=${TEKTON_CLOUDEVENTS_PATH:-${GOPATH}/src/github.com/tektoncd/experimental/cloudevents}
pushd "$TEKTON_CLOUDEVENTS_PATH"
ko apply -f config/
popd

# Configure the cloudevents controller to talk to the keptn inbound adapter
BROKER_SINK="http://broker-ingress.knative-eventing.svc.cluster.local/default/$BROKER_NAME"
kubectl patch cm config-defaults -n tekton-cloudevents -p '{"data": {"default-cloud-events-sink": "'$BROKER_SINK'"}}'

# Install Tekton Resources
kubectl create -f "$BASE_DIR/tekton/resources/" || true

# Install cloud-player
kubectl create -f "$BASE_DIR/cloudplayer/deploy.yaml"

# Create the triggers to get the events to the keptn inbound service
kubectl create -f - <<EOF
apiVersion: eventing.knative.dev/v1
kind: Trigger
metadata:
  name: cd-artifact-packaged-to-keptn-in
spec:
  broker: $BROKER_NAME
  filter:
    attributes:
      type: cd.artifact.packaged.v1
  subscriber:
    uri: http://keptn-cdevents.keptn:8080/events
EOF
kubectl create -f - <<EOF
apiVersion: eventing.knative.dev/v1
kind: Trigger
metadata:
  name: cd-service-deployed-to-keptn-in
spec:
  broker: $BROKER_NAME
  filter:
    attributes:
      type: cd.service.deployed.v1
  subscriber:
    uri: http://keptn-cdevents.keptn:8080/events
EOF
kubectl create -f - <<EOF
apiVersion: eventing.knative.dev/v1
kind: Trigger
metadata:
  name: cd-artifact-published-to-keptn-in
spec:
  broker: $BROKER_NAME
  filter:
    attributes:
      type: cd.artifact.published.v1
  subscriber:
    uri: http://keptn-cdevents.keptn:8080/events
EOF
kubectl create -f - <<EOF
apiVersion: eventing.knative.dev/v1
kind: Trigger
metadata:
  name: cd-artifact-published-to-tekton-triggers
spec:
  broker: $BROKER_NAME
  filter:
    attributes:
      type: cd.artifact.published.v1
  subscriber:
    uri: http://el-cdevent-listener.cdevents:8080
EOF

## Run the script to install and configure Spinnaker with poc
chmod +x $BASE_DIR/spinnaker/installAndConfigSpinnaker.sh
$BASE_DIR/spinnaker/installAndConfigSpinnaker.sh

# Echo relevant environment
env | egrep '(KO|KIND|KEPTN|^TEKTON|BROKER|KNATIVE|reg_)' > poc.env

# Echo endpoints and demo
echo "Tekton Dashboard available at http://tekton-127.0.0.1.nip.io"
echo "Keptn Bridge available at http://keptn-127.0.0.1.nip.io"
echo "-> for the login creds, use keptn configure bridge -o"
echo "CloudEvents player available at http://cloudevents-player.default.knative-127.0.0.1.nip.io"
echo "Spinnaker UI available at http://spin-ui-127.0.0.1.nip.io"

echo "To kick off the demo, from the poc folder, run tkn:"
echo "tkn pipeline start build-artifact -w name=sources,volumeClaimTemplateFile=./tekton/workspace-template.yaml"
