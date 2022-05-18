#!/bin/bash

set -e -o pipefail

## This script is used for the following actions
## 1. Install Minio Server as service in k8s cluster and configure as storage service for spinnaker
## 2. Configure kubernetes account with spinnaker
## 3. Install spinnaker using halyard command
## 4. Create a trigger to subscribe spinnaker event API
## 5. Launch spinnaker deployment and create Application/Pipeline using spincli
## 6. Create Ingress to access spinnaker UI from host


BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

## Declare variables
GIT_PATH_SIGEVENTS=$BASE_DIR/../../
SPIN_ECHO_GIT_URL="git@github.com:rjalander/echo.git -b cdevent_consume"

function installMinioService() {
    helm repo add minio https://helm.min.io/
    helm install my-release minio/minio || true

    ACCESS_KEY=$(kubectl get secret -n default my-release-minio -o jsonpath="{.data.accesskey}" | base64 --decode)
    SECRET_KEY=$(kubectl get secret -n default my-release-minio -o jsonpath="{.data.secretkey}" | base64 --decode)
    MINIO_END_POINT=$(kubectl get endpoints my-release-minio -n default | grep my-release-minio | awk '{print $2}')

    count=0
    max_count=10
    echo "Check minio endpoint is configured."
    while [ $count -lt $max_count ]; do
        echo "MINIO_END_POINT ==> $MINIO_END_POINT"
        if [[ $MINIO_END_POINT != "<none>" ]]; then
            break
        fi
        count=$[$count +1]
        echo "Sleep for 5Sec before next attempt."
        sleep 5
        MINIO_END_POINT=$(kubectl get endpoints my-release-minio -n default | grep my-release-minio | awk '{print $2}')
    done
    if [[ $MINIO_END_POINT == "<none>" ]]; then
            echo "Unable to get the minio endpoint - $MINIO_END_POINT"
            exit 1
    fi

    echo $SECRET_KEY | \
         hal config storage s3 edit --endpoint http://$MINIO_END_POINT \
             --access-key-id $ACCESS_KEY \
             --secret-access-key

    hal config storage edit --type s3

    mkdir -p ~/.hal/default/profiles &&   touch ~/.hal/default/profiles/front50-local.yml
    echo 'spinnaker.s3.versioning: false' > ~/.hal/default/profiles/front50-local.yml
}

function configureK8SAccountWithSpinnaker() {
    hal config provider kubernetes enable
    K8S_ACCOUNT=poc-k8s-account
    hal config provider kubernetes account add $K8S_ACCOUNT \
        --provider-version v2 \
        --context $(kubectl config current-context) || true
    hal config deploy edit --type distributed --account-name $K8S_ACCOUNT
}

function installSpinnaker() {
    hal config version edit --version 1.26.6
    hal deploy apply
    echo "Sleep for 20Sec to initialize spinnaker micro services"
    sleep 20

    rm -rf ~/dev/spinnaker/echo
    mkdir -p ~/dev/spinnaker/echo
    git clone $SPIN_ECHO_GIT_URL ~/dev/spinnaker/echo
    cd ~/dev/spinnaker/echo
    rm -rf ./echo.yml ./spinnaker.yml
    ln ~/.hal/default/staging/echo.yml ./echo.yml
    ln ~/.hal/default/staging/spinnaker.yml ./spinnaker.yml
    ./gradlew echo-web:installDist -x test && docker build -f Dockerfile.slim --tag localhost:5000/cdevents/spinnaker-echo-poc .
    docker push localhost:5000/cdevents/spinnaker-echo-poc:latest

    cd $GIT_PATH_SIGEVENTS/poc/spinnaker
    kubectl delete -f spin-echo-deploy.yaml || true
    kubectl apply -f spin-echo-deploy.yaml
    echo "Sleep for 20Sec to initialize spinnaker poc echo service"
    sleep 20

}

function createTriggerToSubscribeSpinnakerEvent() {
kubectl create -f - <<EOF || true
apiVersion: eventing.knative.dev/v1
kind: Trigger
metadata:
  name: cd-artifact-published-to-spinnaker-poc
spec:
  broker: events-broker
  filter:
    attributes:
      type: cd.artifact.published.v1
  subscriber:
    uri: http://spin-echo.spinnaker:8089/cdevent/consume
EOF
}

function createApplicationAndPipeline() {
    ## Run hal deploy connect with nohup - before pipeline creation..
    nohup hal deploy connect &
    count=0
    max_count=20
    echo "Check Spinnaker API gateway is running."
    while [ $count -lt $max_count ]; do
        sudo netstat -tulpn | grep 8084 && break
        count=$[$count +1]
        echo "Spinnaker API gateway is NOT running, sleep for 5Sec before next attempt."
        sleep 5
    done

    sudo netstat -tulpn | grep 8084 && ( 
        echo "Spinnaker API gateway is running, proceeding with pipeline creation." 
            ##** Run hal deploy connect - from the VM before pipeline creation..
        spin application save --application-name cdevents-poc --owner-email someone@example.com --cloud-providers "kubernetes"
        cd $GIT_PATH_SIGEVENTS/poc/spinnaker
        spin pipeline save -f deploy-spinnaker-poc.json
        echo "Spinnaker Application and Pipeline created successfully" ) || (
        echo "Spinnaker API gateway is still NOT running, please run 'nohup hal deploy connect &' and run the below commands to create the pipeline"
        echo "-------------------------------------------"
        echo "spin application save --application-name cdevents-poc --owner-email someone@example.com --cloud-providers "kubernetes""
        echo "cd $GIT_PATH_SIGEVENTS/poc/spinnaker"
        echo "spin pipeline save -f deploy-spinnaker-poc.json" 
        echo "-------------------------------------------" )
}

function createIngressSpinnakerUI() {
kubectl create -f - <<EOF || true
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: spinnaker-ui
  namespace: spinnaker
  annotations:
    kubernetes.io/ingress.class: "contour-external"
spec:
  rules:
  - host: spin-ui-127.0.0.1.nip.io
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: spin-deck
            port:
              number: 9000
EOF
}

########
## Main
#######
echo "Checking if required CLIs are installed"
helm version > /dev/null
hal -v > /dev/null
spin --version >> /dev/null

installMinioService
configureK8SAccountWithSpinnaker
installSpinnaker
createTriggerToSubscribeSpinnakerEvent
createApplicationAndPipeline
createIngressSpinnakerUI
exit 0
