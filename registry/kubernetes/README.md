# Usage

## Pre-requirement

- kubernetes installed
- docker installed or other building tools like podman
- $HOME/.kube/config configured

## Server

```bash
cd registry/kubernetes/server
docker build . -t kratos-helloworld:v1.0.0
kubectl create -f deploy/deployment.yaml
```

## Client

```bash
go mod tidy
cd registry/kubernetes
go run client/main.go
```
