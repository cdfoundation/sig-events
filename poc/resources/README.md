This PoC was built following the Keptn Quickstart into K3s: https://keptn.sh/docs/quickstart/

If you have keptn cli installed already you can run these two commands: 

1) `k3d cluster create mykeptn -p "8082:80@agent[0]" --k3s-server-arg '--no-deploy=traefik' --agents 1` 
2) `keptn install --use-case=continuous-delivery`

And instead of creating a project with the quickstart instructions the shipyard.yaml file is provided to run the following steps. 

You need to export the two following Environment Variables: 

**Note**: If you don't install istio for ingress you need to port-forward with:
`kubectl -n keptn port-forward service/api-gateway-nginx 8080:80`

1) `export KEPTN_ENDPOINT=http://localhost:8080/api` I

2) `export KEPTN_API_TOKEN=$(kubectl get secret keptn-api-token -n keptn -ojsonpath={.data.keptn-api-token} | base64 --decode)

Then run
`keptn auth --endpoint=$KEPTN_ENDPOINT --api-token=$KEPTN_API_TOKEN`

Then configure the bridge (this will give you the credentials for the user interface located at http://localhost:8080)
`keptn configure bridge -o`

Run the CLI Command: `keptn create project fmtok8s --shipyard=your_shipyard.yaml` 

Run the CLI command: `keptn create service podtato-server --project=fmtok8s` 

Now you can start the CDEvents to Ketpn Inbound Translator component () with from inside the translator go project (this will configure a server in port 8081 to avoid conflict with Keptn in 8080): 

```
go build
go run main.go
```

You can also run the translator component in the k8s cluster:
```bash
export KO_DOCKER_REPO=<your-local-repo>
ko apply -f config/
```

Using the CDE binary you can emit an CD CloudEvent to the translator layer. Make sure that you export the CDE_SINK for CDE to point to the translator endpoint: 
`export CDE_SINK=http://localhost:8081/events`

`cde artifact packaged --id agenda-service-rest --name agenda-service -v 0.0.20 --data artifact=hash`

This will hit the translator layer which will create a Keptn Event to trigger the Keptn Delivery Sequence (specified in the shipyard file)

Inside Keptn we need to have an outbound event translator to CD Events (https://github.com/salaboy/cdf-events-keptn-adapter). This component picks events emitted by Keptn, transform them into CD Events and send them to a configurable SINK. 

For this to work, you need to make sure that the Helm Service in Keptn which is automatically configured when Keptn is installed is not listening to delivery events. You can do that by editing the `helm-service` deployment in the `keptn` namespace and remove the value: `sh.keptn.event.deployment.triggered` from the `PUBSUB_TOPIC` env variable. 

You need to make sure that the outbound service deployed in Keptn is listening to this event: 

```
        - name: PUBSUB_TOPIC
          value: sh.keptn.event.deployment.triggered
```

Note: make sure that you remove the following variable for the routing to work: 
```
        - name: PUBSUB_URL
          value: nats://keptn-nats-cluster
```

