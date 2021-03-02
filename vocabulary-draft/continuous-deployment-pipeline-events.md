# Continuous Deployment Pipelines Events 

These events are related to continous deployment pipelines and their target environments. 
These events can be emitted by environments to report where software artifacts such as services, binaries, deamons, jobs or embeded software are running. 

The term Service is used to represent a running Artifact. This service can represent a binary that is running, a deamon, an application, a docker container, etc.
The term Environment represent any platform which has all the means to run a Service. 

- **Environment Created**: an environment has been created and it can be used to deploy Services
- **Environment Modified**: an environment has been modified, this event advertise the changes made in the environment
- **Environment Deleted**: an environment has been deleted and cannot longer be used
- **Service Deployed**: a new instance of the Service has been deployed
- **Service Upgraded**: an existing instance of a Service has been upgraded to a new version
- **Service Undeployed**: an existing instance of a Service has been terminated an it is not longer present in an environment

Continous Deployment Events MUST include the following attributes:
- **Event Type**: the type is restricted to include `CD.Environment**` prefix. For example `CD.Environment.Service.Upgraded` or `CD.Environment.Created`
- **Environment ID**: unique identifier for the Environment


Optional attributes: 

- **Environment Name**: user-friendly name for the environment, to be displayed in tools or User Interfaces
- **Environment URL**: URL to reference where the environment is located