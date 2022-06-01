#!/bin/bash

set -e -o pipefail

## This script is used to deploy spinnaker in halyard K8S pod
## 1. Install helm CLI 
## 2. Configure default K8S context inside halyard pod 
## 3. Install Minio Server as service in k8s cluster and configure as storage service for spinnaker
## 4. Configure kubernetes account with spinnaker
## 5. deploy spinnaker using halyard command


function installHelmCLI() {
    cd /home/spinnaker/
    curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
    sed -i 's/\/usr\/local\/bin/\/home\/spinnaker/g' get_helm.sh
    sed -i 's/sudo //g' get_helm.sh
    export PATH=/home/spinnaker:$PATH
    chmod 700 get_helm.sh
    ./get_helm.sh
}
function configureDefaultK8SContext() {
    kubectl config set-cluster default --server=https://kubernetes.default --certificate-authority=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    kubectl config set-context default --cluster=default
    token=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
    kubectl config set-credentials user --token=$token
    kubectl config set-context default --user=user
    kubectl config use-context default
}

function installMinioService() {
    helm repo add minio https://helm.min.io/
    helm install my-release minio/minio || true

    ACCESS_KEY=$(kubectl get secret -n default my-release-minio -o jsonpath="{.data.accesskey}" | base64 -d)
    SECRET_KEY=$(kubectl get secret -n default my-release-minio -o jsonpath="{.data.secretkey}" | base64 -d)
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

function deploySpinnaker() {
    hal config version edit --version 1.26.6
    hal deploy apply
    echo "Sleep for 20Sec to initialize spinnaker micro services"
    sleep 20
}

########
## Main
#######
installHelmCLI
echo "Checking if required CLIs are installed inside hal pod"
helm version > /dev/null
hal -v > /dev/null
configureDefaultK8SContext
installMinioService
configureK8SAccountWithSpinnaker
deploySpinnaker
exit 0
