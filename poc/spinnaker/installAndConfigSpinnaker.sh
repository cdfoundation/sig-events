#!/bin/bash

set -e -o pipefail

## This script is used to install configure spinnaker to run with the sig-events poc
## 1. Create hal stable deployment and update with Spinnaker service account
## 2. Deploy distributed spinnaker in K8S cluster using halyard
## 3. Update spin-echo with the cdevents API and build new image
## 3. Create a trigger to subscribe spinnaker event API
## 4. Launch spinnaker deployment and create Application/Pipeline using spin CLI
## 5. Create Ingress to access spinnaker UI from host


BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

## Declare variables
GIT_PATH_SIGEVENTS=$BASE_DIR/../../
SPIN_ECHO_GIT_URL="git@github.com:rjalander/echo.git -b cdevent_consume"

function deployHalAndUpdateServiceAccount() {
    kubectl delete deployment hal || true
    kubectl create deployment hal --image gcr.io/spinnaker-marketplace/halyard:stable
    cd $GIT_PATH_SIGEVENTS/poc/spinnaker
    kubectl create -f spinnakerServiceAccount.yaml
    kubectl patch deployment hal --patch '{"spec": {"template": {"spec": {"serviceAccountName":"spinnaker-service-account"}}}}'
    echo "Sleep for 10Sec for hal deployment to start"
    sleep 10
}

function checkHalPodIsRunning() {
    HAL_POD_NAME=$(kubectl get pods -l app=hal --field-selector=status.phase==Running -o jsonpath="{.items[0].metadata.name}")
    count=0
    max_count=20
    echo "Check hal pod is running."
    while [ $count -lt $max_count ]; do
        echo "HAL_POD_NAME ==> $HAL_POD_NAME"
        if [[ $HAL_POD_NAME != "" ]]; then
            break
        fi
        count=$[$count +1]
        echo "Sleep for 5Sec before next attempt."
        sleep 5
        HAL_POD_NAME=$(kubectl get pods -l app=hal --field-selector=status.phase==Running -o jsonpath="{.items[0].metadata.name}")
    done
    if [[ $HAL_POD_NAME == "" ]]; then
            echo "Hal pod is not running, exiting."
            exit 1
    fi
}

function deploySpinnakerWithHalPod() {
    checkHalPodIsRunning
    HAL_POD_NAME=$(kubectl get pods -l app=hal --field-selector=status.phase==Running -o jsonpath="{.items[0].metadata.name}")
    chmod +x $GIT_PATH_SIGEVENTS/poc/spinnaker/deploySpinnakerWithHalyard.sh
    kubectl cp $GIT_PATH_SIGEVENTS/poc/spinnaker/deploySpinnakerWithHalyard.sh default/$HAL_POD_NAME:/home/spinnaker/
    kubectl exec -it $HAL_POD_NAME -- /bin/bash -c "/home/spinnaker/deploySpinnakerWithHalyard.sh"
}

function updateSpinEchoWithCDEventsAPI() {
    rm -rf ~/dev/spinnaker/echo
    mkdir -p ~/dev/spinnaker/echo
    git clone $SPIN_ECHO_GIT_URL ~/dev/spinnaker/echo
    cd ~/dev/spinnaker/echo
    rm -rf ./echo.yml ./spinnaker.yml
    checkHalPodIsRunning
    HAL_POD_NAME=$(kubectl get pods -l app=hal --field-selector=status.phase==Running -o jsonpath="{.items[0].metadata.name}")
    kubectl cp default/$HAL_POD_NAME:/home/spinnaker/.hal/default/staging/spinnaker.yml ./spinnaker.yml
    kubectl cp default/$HAL_POD_NAME:/home/spinnaker/.hal/default/staging/echo.yml ./echo.yml
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
    ## connect with nohup - before pipeline creation..
    nohup kubectl -n spinnaker port-forward service/spin-deck 9000:9000 &
    nohup kubectl -n spinnaker port-forward service/spin-gate 8084:8084 &
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
        echo "Spinnaker API gateway is still NOT running, please check the logs from spin-gate micro service pod.")
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
echo "Checking if spin CLI is installed"
spin --version >> /dev/null

deployHalAndUpdateServiceAccount
deploySpinnakerWithHalPod
updateSpinEchoWithCDEventsAPI
createTriggerToSubscribeSpinnakerEvent
createApplicationAndPipeline
createIngressSpinnakerUI
exit 0
