# CDE - CloudEvents for Continuous Delivery

Simple CLI and library to emit Cloud Events related to Continuous Delivery.

The framework and command-line interface enables you to emit and receive events related to Continuous Delivery, with the purpose of enabling features such as interoperability between different tools/services, performance measurements, increased visibility into your Continuous Delivery processes, etc.


## Usage

Download the binary `cde` or build from source, by cloning this repository and run `make build`

Set `CDE_SINK` environment variable to define where CloudEvents will be sent. You can do this by running:

`export CDE_SINK=http://my-cloudevent-broker` 

You can use [SockEye](https://github.com/n3wscott/sockeye) to visualize the cloud events.

Integrate with your existing pipelines, repositories, and environments.

## Supported Events

Use `cde --help`

The following events are supported:
- [Environment Events](#environment-events)
- [Service Events](#service-events)
- [PipelineRun Events](#pipelinerun-events)
- [TaskRun Events](#taskrun-events)
- [Repository Events](#repository-events)
            

# CDE Events

The following events are currently supported: 

## Environment Events
- Properties
    - Id
    - Name
    - URL
      
Examples: 

- **cd.environment.created.v1** `./cde env created --id staging --name "Staging Environment" --repo http://github.com/user/myrepo --data event=data`
- **cd.environment.modified.v1** `./cde env modified --id staging --name "Staging Environment" --repo http://github.com/user/myrepo --data event=data`
- **cd.environment.deleted.v1** `./cde env deleted --id staging --name "Staging Environment" --repo http://github.com/user/myrepo --data event=data`


## Service Events
- Properties
    - Name
    - Version
    - Environment Id

Examples:

- **cd.service.deployed.v1** `./cde service deployed --envId staging --name my-service --version 1.0.2 --data service=data`
- **cd.service.upgraded.v1** `./cde service upgraded --envId staging --name my-service --version 1.0.3 --data service=data`
- **cd.service.rolledback.v1** `./cde service upgraded --envId staging --name my-service --version 1.0.2 --data service=data`
- **cd.service.removed.v1** `./cde service upgraded --envId staging --name my-service --version 1.0.2 --data service=data`

## PipelineRun Events
- Properties
  - Id
  - Name
  - URL
  - Status  
  - Errors
    
Examples:
- **cd.pipelinerun.started.v1**  `./cde pipelinerun queued --id pipe1 --name "My Pipeline" --status "Starting" --url "http://my-pipelinerunner" --errors "Hopfully no errors" --data pipeline=data`
- **cd.pipelinerun.finished.v1** `./cde pipelinerun started --id pipe1 --name "My Pipeline" --status "Starting" --url "http://my-pipelinerunner" --errors "Hopfully no errors" --data pipeline=data`
- **cd.pipelinerun.queued.v1**   `./cde pipelinerun finished --id pipe1 --name "My Pipeline" --status "Starting" --url "http://my-pipelinerunner" --errors "Hopfully no errors" --data pipeline=data`

## TaskRun Events
- Properties
  - Id
  - Name
  - PipelineId
  - Status
  - Errors

Examples:
- **cd.taskrun.started.v1**  `./cde taskrun queued --id task1 --name "My Task Run" --pipeid pipe1 --status "Running" --url "http://my-pipelinerunner/task1" --errors "Hopfully no errors" --data task=data`
- **cd.taskrun.finished.v1** `./cde taskrun started --id task1 --name "My Task Run" --pipeid pipe1 --status "Successful" --url "http://my-pipelinerunner/task1" --errors "Hopfully no errors" --data task=data`
- **cd.taskrun.queued.v1**   `./cde taskrun finished --id task1 --name "My Task Run" --pipeid pipe1 --status "Queued" --url "http://my-pipelinerunner/task1" --errors "Hopfully no errors" --data task=data`

## Repository Events
- Properties
  - Id
  - Name
  - URL
    
Examples: 
- **cd.repository.created.v1** `./cde repository created --id repository_id --name "Name of repository" --url "http://my-repository" --data repository=data`
- **cd.repository.modified.v1** `./cde repository modified --id repository_id --name "Name of repository" --url "http://my-repository" --data repository=data`
- **cd.repository.deleted.v1** `./cde repository deleted --id repository_id --name "Name of repository" --url "http://my-repository" --data repository=data`


# References
- [CDFoundation SIG Events Vocabulary Draft](https://github.com/cdfoundation/sig-events/tree/main/vocabulary-draft)
