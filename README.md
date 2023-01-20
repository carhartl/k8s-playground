# GoLang K8s Playground

[![Pipeline](https://github.com/carhartl/golang-k8s-playground/actions/workflows/pipeline.yml/badge.svg)](https://github.com/carhartl/golang-k8s-playground/actions/workflows/pipeline.yml)
[![Vulnerability Scan](https://github.com/carhartl/golang-k8s-playground/actions/workflows/vulnerability-scan.yml/badge.svg)](https://github.com/carhartl/golang-k8s-playground/actions/workflows/vulnerability-scan.yml)

## Prerequisites

[Minikube](https://minikube.sigs.k8s.io/docs/start/) (with Kubernetes 1.25 or later), [Skaffold](https://skaffold.dev), [Helm](https://helm.sh) (3.4 or later), [Kubectl](https://kubectl.docs.kubernetes.io/guides/introduction/kubectl/), [Kustomize](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/):

```bash
brew install helm kubectl kustomize minikube skaffold
```

(Installing `kubectl` is not strictly necessary, Minikube comes with kubectl as well...)

Optional, when using the [Lefthook](https://github.com/evilmartians/lefthook) based Git hooks setup:

```bash
brew install golangci-lint lefthook prettier yamllint && lefthook install
```

## Development

Start local K8s cluster (make sure the resources to start minikube with are in line with what is configured in Docker Desktop):

```bash
minikube start --memory=4096 --cpus=4 --disk-size=30g
```

Start continuous local development with Skaffold:

```bash
skaffold dev
```

## Database

Access the database:

```bash
kubectl --namespace yugabytedb-system exec -it yb-tserver-0 -- sh -c "cd /home/yugabyte && ysqlsh -h yb-tserver-0 --echo-queries"
```
