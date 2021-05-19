# CDE - CloudEvents for Continuous Delivery

Simple CLI and library to emit Cloud Events related to Continuous Delivery.

The idea behind this framework and command-line interface is to empower you  emitting events related to a Continuous Delivery to enable interoperability between different tools, measure performance and gain visibility of your CD processes.


## Usage

Download the binary `cde` or build from source, by clonning this repository and run `make build`

Set `CDE_SINK` environment variable to define where CloudEvents will be sent. You can do this by running:

`export CDE_SINK=http://my-cloudevent-broker` 

You can use [SockEye](https://github.com/n3wscott/sockeye) to visualize the cloud events.

Integrate with your existing pipelines, repositories, and environments.

## Supported Events

Use `cde --help`

The following events are supported:
- [Environment Events](https://github.com/salaboy/cd-flow/#environment-events)
- [Service Events](https://github.com/salaboy/cd-flow/#service-events)
            

# CDE  Events

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




# References
- [CDFoundation SIG Events Vocabulary Draft](https://github.com/cdfoundation/sig-events/tree/main/vocabulary-draft)
