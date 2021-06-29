# Tekton Resources

This folder contains all the Tekton resources required for the PoC.

## Prerequisites

### Install Tekton

Set up a kind cluster with Tekton using [tekton_in_kind.sh](../tekton_in_kind.sh).
Alternatively:

1. Install Tekton Pipeline:
```bash
kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/previous/v0.25.0/release.yaml
```

2. Install Tekton Triggers:
```bash
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/previous/v0.14.2/release.yaml
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/previous/v0.14.2/interceptors.yaml
```

3. Install Tekton Dashboard (optional):
```bash
kubectl apply -f https://github.com/tektoncd/dashboard/releases/download/v0.17.0/tekton-dashboard-release.yaml
```

### Configure Tekton Pipeline

The Tekton pipeline uses some alpha features (Tekton Bundles) which needs to
be enabled in the config. The `feature-flags` config map in the `tekton-pipelines` namespace
should look like:

```yaml
apiVersion: v1
data:
  enable-api-fields: alpha # <------- 
  disable-affinity-assistant: "false"
  disable-creds-init: "false"
  disable-home-env-overwrite: "true"
  disable-working-directory-overwrite: "true"
  enable-custom-tasks: "false"
  enable-tekton-oci-bundles: "false"
  require-git-ssh-secret-known-hosts: "false"
  running-in-environment-with-injected-sidecars: "true"
kind: ConfigMap
(...)
```

### Install Tekton CloudEvents Controller

Nightly builds are not available yet, so this has to be installed from source.
Requires:

- go 1.16+
- `GO111MODULE=on`
- [`ko`](https://github.com/google/ko)

Clone the [tekton experimental repo](https://github.com/tektoncd/experimental).

Setup `KO_DOCKER_REPO` to point to your container registry (or `kind.local` when using kind).
Note than when using a private container registry, the tekton cloudevents controller service account
will need to have access via `imagePullSecrets`.

Move to the cloned folder:
```
cd cloudevents
ko apply -f config/
```

### Configure Tekton CloudEvents Controller

The `config-defaults` config map in the `tekton-cloudevents` namespace should look like:

```yaml
apiVersion: v1
kind: ConfigMap
data:
  default-cloud-events-sink: http://<endpoint of cloudevents listener on keptn side>
(...)
```

## Install Tekton Resources

Install all the yaml files:

```bash
kubectl create -f poc/tekton
```

## Run the build pipeline

Prerequisites:
- `tkn` (github.com/tektoncd/cli)
- Push access to a container registry. The `dockerconfig` secret in the example below is called `regcred`

```bash
tkn pipeline start build-artifact -w name=sources,volumeClaimTemplateFile=poc/tekton/workspace-template.yaml -w name=dockerconfig,secret=regcred
```
