This PoC was built following the Keptn Quickstart into K3s: https://keptn.sh/docs/quickstart/

If you have keptn cli installed already you can run these two commands: 

1) `k3d cluster create mykeptn -p "8082:80@agent[0]" --k3s-server-arg '--no-deploy=traefik' --agents 1` 
2) `keptn install --use-case=continuous-delivery`

And instead of creating a project with the quickstart instructions the shipyard.yaml file is provided to run the following steps: 


Then you need to export the two following Environment Variables: 

**Note**: If you don't install istio for ingress you need to port-forward with:
`kubectl -n keptn port-forward service/api-gateway-nginx 8080:80`

1) `export KEPTN_ENDPOINT=http://localhost:8080/api` I

2) `export KEPTN_API_TOKEN=$(kubectl get secret keptn-api-token -n keptn -ojsonpath={.data.keptn-api-token} | base64 --decode)

Then run
`keptn auth --endpoint=$KEPTN_ENDPOINT --api-token=$KEPTN_API_TOKEN`

Then configure the bridge (this will give you the credentials for the user interface located at http://localhost:8080)
`keptn configure bridge -o`

Run the CLI Command: `keptn create project fmtok8s --shipyard=your_shipyard.yaml` 

Run the CLI command: `keptn create service agenda-service --project=fmtok8s` 

Now you can start the CDEvents to Ketpn Translator component with from inside the tranlator go project: 

```
go build
go run main.go
```

Using the CDE binary you can emit an CD CloudEvent to the translator layer. Make sure that you export the CDE_SINK for CDE to point to the translator endpoint: 
`export CDE_SINK=http://localhost:8081/events`


`cde artifact packaged --id agenda-service-rest --name agenda-service -v 0.0.20 --data artifact=hash`

This will hit the translator layer which will create a Keptn Event to trigger a Keptn Sequence
