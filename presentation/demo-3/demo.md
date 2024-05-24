# Rolling updates

## Rollout to a specific version

`kubectl rollout undo deployment/go-app --to-revision=1`

## check the rollouts

`kubectl rollout history deployment/go-app`

## go to the previous version of the deployment

`kubectl rollout undo deployment/go-app`

## updating the image of the deployment

`kubectl set image deployment/go-app go-app=prateek041/go-app:v2`

## get the cluster ID

`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' kind-fountane-workshop-control-plane`
