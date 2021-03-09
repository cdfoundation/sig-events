# Continuous Integration Pipeline Events

These events are related to **Continuous Integration(CI) pipelines**, these pipelines usually include building, testing, packaging and releasing software artifacts. 
Due the dynamic nature of Pipelines, most of actual work needs to be queued to happen in a distributed way, hence Queued events are added. 
Adopters can choose to ignore these events if they don't apply to their use cases. 

A pipeline, in the context of Continuous Integration, is the definition of a set of tasks that needs to be performed to build, test, package and release software artifacts. A pipeline can be instanciated multiple times, for example to build different versions of the same artifact. We are refering to this instance as PipelineRun. It will have a unique Id and it will help us to track the build and release progress on a particular software artifact. 

- **PipelineRun Queued**: a PipelineRun has been schedule to run
- **PipelineRun Started**: a PipelineRun has started and it is running
- **PipelineRun Finished**: a PipelineRun has finished it execution, the event will contain the finished status, success, error or failure

Each pipeline is defined as a set of Tasks to be performed in sequence, hence tracking this sequence might be important for some cases. A TaskRun is an instance of the Task defined inside the pipeline, as you can expect multiple execution of the pipelines (each a PipelineRun) you can also expect multiple execution of the Tasks, for that reason we use TaskRun to refer to one of these instances.  

- **TaskRun Started**: a TaskRun inside a PipelineRun has started.
- **TaskRun Finished**: a TaskRun inside a PipelineRun has finished. 

The following events represent concrete Tasks that are associated with the execution of CI pipelines:

- **Build Queued**: a Build task has been queued, this build process usually is in charge of producing a binary from source code
- **Build Started**: a Build task has started 
- **Build Finished**: a Build task has finished, the event will contain the finished status, success, error or failure

The following Test events are defined in two separate categories **Test Case** and **Test Suite**. A **Test Case** is the smallest unit of testing that the user wants to track. A **Test Suite** is a collection of test case executions and/or other test suite executions. **Test Cases** executed, and Test Suites are for grouping purposes. For this reason, **Test Cases** can be queued. 

- **Test Case Queued**: a Test task has been queued, and it is waiting to be started
- **Test Case Started**: a Test task has started
- **Test Case Finished**: a Test task has finished, the event will contain the finished status, success, error or failure
- **Test Suite Started**: a Test task has started
- **Test Suite Finished**: a Test task has finished, the event will contain the finished status, success, error or failure

Finally, events needs to be generated for the output of the pipeline such as the artifacts that were packaged and released for others to use. 

- **Artifact Packaged**: an artifact has been packaged for distribution, this artifact is now versioned with a fixed version
- **Artifact Released**: an artifact has been released and it can be advertised for others to use

Pipeline Events MUST include the following attributes:
- **Event Type**: the type is restricted to include `CD.Pipeline**` prefix. For example `CD.Pipeline.PipelineRun.Queued` or `CD.Pipeline.Tests.Started`
- **PipelineRun Id**: unique identifier for a pipeline execution
- **Pipeline Name**: unique identifier for the pipeline, not for the instance. A pipeline can have multiple instances/runs.  
- **PipelineRun Status**: current status of the PipelineRun at the time when the event was emitted. If the pipeline is finished, this attribute should reflect if it finished successfully or if there was an error on the execution. Possible statuses: [Running, Finished, Error]

Optional attributes: 
- **PipelineRun URL**: URL to locate where  pipeline instances are running
- **PipelineRun Errors**: error field to indicate possible compilation, test, build and package errors.
