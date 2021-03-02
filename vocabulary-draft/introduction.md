This document intends to describe a set of events related to different phases of a Continuous Delivery process. 
These events are categorized by meaning and the phase where they are intenteded to be used. 
These events are agnostic from any specific tool and are designed to fit a wide range of scenarios. 

The phases covered by this proposal are:

- **Source Code Version Control Events**: Events emitted by changes in source code or by the creation, modification or deletion of new repositories that hold source code.
- **Continuous Integration Pipelines Events**: includes events related to building, testings, packaging and releasing software artifacts, usually binaries.
- **Continuous Deployment Pipelines Events**: include events related with environments where the artifacts produced by the integration pipelines actually run. These are usually services running in a specific environment (dev, QA, production), or embedded software running in a specific hardware. 

These phases can also be considered as different profiles of the vocabulary that can be adopted independently. 

# Required Metadata for CD Events

The following attributes are REQUIRED to be present in all the Events defined in this vocabulary:

- **Event ID**: defines a unique identifier for the event
- **Event Type**: defines a textual description of the event type, only event types described in this document are supported. All event types should be prefixed with `CD.`
- **Event Source**: defines the context in which an event happened
- **Event Timestamp**: defines the time when the event was produced

The following sections list the events in different phases, allowing adopters to choose the events that they are more interested in.