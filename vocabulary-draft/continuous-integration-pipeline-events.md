# Continuous Integration Pipeline Events

These events are related to continuous integration pipelines, this pipelines usually include building, testing, packaging and releasing software artifacts. 
Due the dynamic nature of Pipelines, most of actual work needs to be queued to happen in a distributed way, hence Queued events are added. 
Adopters can choose to ignore these events if they don't apply to their use cases. 

A pipeline, in the context of Continuous Integration, is the definition of a set of tasks that needs to be performed to build, test, package and release software artifacts. A pipeline can be instanciated multiple times, for example to build different versions of the same artifact. We are refering to this instance as PipelineRun. It will have a unique Id and it will help us to track the build and release progress on a particular software artifact. 

- **PipelineRun Queued**: a PipelineRun has been schedule to run
- **PipelineRun Started**: a PipelineRun has started and it is running
- **PipelineRun Finished**: a PipelineRun has finished it execution, the event will contain the finished status, success, error or failure
- **Build Queued**: a Build task has been queued, this build process usually is in charge of producing a binary from source code
- **Build Started**: a Build task has started 
- **Build Finished**: a Build task has finished, the event will contain the finished status, success, error or failure
- **Tests Queued**: a Test task has been queued, and it is waiting to be started
- **Tests Started**: a Test task has started
- **Tests Finished**: a Test task has finished, the event will contain the finished status, success, error or failure
- **Artifact Packaged**: an artifact has been packaged for distribution, this artifact is now versioned with a fixed version
- **Artifact Released**: an artifact has been released and it can be advertised for others to use

Pipeline Events MUST include the following attributes:
- **Event Type**: the type is restricted to include `CD.Pipeline**` prefix. For example `CD.Pipeline.PipelineInstance.Queued` or `CD.Pipeline.Tests.Started`
- **Pipeline Instance Id**: unique identifier for a pipeline execution
- **Pipeline Name**: unique identifier for the pipeline, not for the instance. A pipeline can have multiple instances/runs.  
- **Pipeline Status**: current status of the pipeline at the time when the event was emitted. If the pipeline is finished, this attribute should reflect if it finished successfully or if there was an error on the execution.  

Optional attributes: 
- **Pipeline URL**: URL to locate where  pipeline instances are running
- **PipelineRun Errors**: error field to indicate possible compilation, test, build and package errors.
