# Run the following commands

## Create the deployment
`kubectl apply -f deployment.yaml`

## Create the service to expose pods
`kubectl apply -f service.yaml`

## Get the Kind container's IP 

`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' kind-fountane-workshop-control-plane`

## Get the nginx service NodePort 
`kubectl get svc`

## Check if nginx pod is accesible
`curl http://<Kind container's IP>:<NodePort>`

## Generate load
`ab -n 10000 -c 100 http://<Kind container's IP>:<NodePort>/`

## Scale up the deployment
`kubectl scale deployment/nginx-deployment --replicas=5`

> Note: Tell them about Automatic Horizontal Pod Autoscaler, and how it works with metrics-server.
