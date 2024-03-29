# K8s Playground

[![Pipeline](https://github.com/carhartl/k8s-playground/actions/workflows/pipeline.yml/badge.svg)](https://github.com/carhartl/k8s-playground/actions/workflows/pipeline.yml)
[![Vulnerability Scan](https://github.com/carhartl/k8s-playground/actions/workflows/vulnerability-scan.yml/badge.svg)](https://github.com/carhartl/k8s-playground/actions/workflows/vulnerability-scan.yml)

## Prerequisites

[Minikube](https://minikube.sigs.k8s.io/docs/start/) (with Kubernetes 1.25 or later), [Skaffold](https://skaffold.dev), [Helm](https://helm.sh) (3.4 or later), [Kubectl](https://kubectl.docs.kubernetes.io/guides/introduction/kubectl/), [Kustomize](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/):

```bash
brew install helm kubectl kustomize minikube skaffold
```

Use [Devbox](https://www.jetpack.io/devbox/docs/installing_devbox/) to install instead of Homebrew:

```bash
devbox install
devbox shell
```

## Development

Start local K8s cluster (make sure the resources to start minikube with are in line with what is configured in Docker Desktop):

```bash
minikube start --memory=4096 --cpus=4 --disk-size=30g --kubernetes-version 1.26.3
```

Start continuous local development with Skaffold:

```bash
skaffold dev
```

## Monitoring

Grafana dashboard:

```bash
kubectl port-forward -n monitoring svc/kube-prometheus-stack-grafana 3000:80
```

(user: admin, password: prom-operator)

_Note: skaffold sets up port forwarding in `dev` mode._

Kuberhealthy:

```bash
kubectl port-forward -n kuberhealthy svc/kuberhealthy 3001:80
```

_Note: skaffold sets up port forwarding in `dev` mode._

## Database

Access the database:

```bash
kubectl -n yugabyte exec -it yb-tserver-0 -- sh -c "cd /home/yugabyte && ysqlsh -h yb-tserver-0 --echo-queries"
```

## Provision Cluster on GKE

Requires configured gcloud SDK: https://developer.hashicorp.com/terraform/tutorials/kubernetes/gke#prerequisites

```bash
cd infra/terraform
terraform init
terraform apply
```

## Deploy to GKE

```bash
skaffold run
```
