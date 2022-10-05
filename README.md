# go-rabbit-demo

In this demo app we will:
1. Spin up a rabbit cluster on a local k8n
2. Create an exchange
3. Build a sample golang http server which produces to the exchange
4. Dockerize and Kubernize your app
5. Build a sample golang consumer
6. Dockerize and Kubernize your app


## What is kubernetes? - A container orchestrator!
https://kubernetes.io/docs/concepts/overview/

## K8S Components
https://kubernetes.io/docs/concepts/overview/components/

## RabbitMQ
https://www.rabbitmq.com/

### Demo Steps:

1. Start minikube

```minikube start --cpus 4 --driver=docker```

2. Load kubernetes UI

```minikube dashboard```

GOTCHA: if there it stalls on "checking proxy health" check deployment "kubectl get deploy -n kubernetes-dashboard"

3. Ensure we're using the correct context!

```minikube update-context```

4. Install Operator

```kubectl apply -f "https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml"```

K8s Operators are controllers for packaging, managing, and deploying applications on Kubernetes. In order to do these things, the Operator uses Custom Resources (CR) that define the desired configuration and state of a specific application through Custom Resource Definitions (CRD)

5. Deploy operator

```kubectl apply -f https://raw.githubusercontent.com/rabbitmq/cluster-operator/main/docs/examples/hello-world/rabbitmq.yaml```

6. Lets have a look at the rabbitmq ui

```kubectl port-forward "service/hello-world" 15672```

linux:
```
username="$(kubectl get secret hello-world-default-user -o jsonpath='{.data.username}' | base64 --decode)" && password="$(kubectl get secret hello-world-default-user -o jsonpath='{.data.password}' | base64 --decode)" && echo "username: $username, password: $password"
```
windows:
```
$username="$([Text.Encoding]::Utf8.GetString([Convert]::FromBase64String($(kubectl get secret hello-world-default-user -o jsonpath='{.data.username}'))) )"

$password="$([Text.Encoding]::Utf8.GetString([Convert]::FromBase64String($(kubectl get secret hello-world-default-user -o jsonpath='{.data.password}'))) )"

Write-Output "username: $username, password: $password"
```

7. Add an exchange, add a user

```exchange: exchange1```
```user: guest```

8. Build consumer image and load to minikube

```docker build -f cmd/consumer/Dockerfile -t go-rabbit-consumer .```

```minikube image load go-rabbit-consumer```

9. Apply k8n deploy of new consumer

```kubectl apply -f deployments/k8s/base/consumer/resources/deployment.yaml```

10. Build server and load image to minikube
 
```docker build -f cmd/server/Dockerfile -t go-rabbit-server .```

```minikube image load go-rabbit-server```

11. Apply k8n deploy of new server

```kubectl apply -f deployments/k8s/base/server/resources/deployment.yaml```

12. Apply k8n deploy of new service for server

```kubectl apply -f deployments/k8s/base/server/resources/service.yaml```

13. Proxy new service

```kubectl port-forward "service/go-rabbit-server-service" 8090```

14. Hit endpoint in browser

```http://localhost:8090/result?type=andy4```

15. Ramp up requests with k6

```k6 run script.js```
