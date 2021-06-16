This PoC was built following the Keptn Quickstart into K3s: https://keptn.sh/docs/quickstart/

And instead of creating a project with the quickstart instructions the shipyard.yaml file is provided to run the following steps: 

Download the Keptn CLI: https://keptn.sh/docs/0.8.x/operate/install/#install-keptn-cli
Authenticate the Keptn CLI: https://keptn.sh/docs/0.8.x/operate/install/#authenticate-keptn-cli (using KEPTN_API_TOKEN and KEPTN_ENDPOINT)
Run the CLI Command: keptn create project podtatohead --shipyard=your_shipyard.yaml --git-user=GIT_USER --git-token=GIT_TOKEN --git-remote-url=GIT_REMOTE_URL
Run the CLI command: keptn create service helloservice --project=podtatohead 
