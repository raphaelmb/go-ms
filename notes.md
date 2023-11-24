`docker-compose-postgres-only.yaml` is a separate file to simulate a database outside the k8s cluster.

`k8s/auth-service.yaml` DSN env value "host=host.minikube.internal" is pointing to a specific minikube feature. Change it accordingly.

### minikube specific stuff
`minikube addons enable ingress`
`minikube addons enable ingress-dns`

### /etc/hosts
`127.0.0.1   front-end.info broker-service.info`
