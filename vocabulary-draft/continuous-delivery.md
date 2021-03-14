# Continuous Delivery

Continuous Delivery related to activities and orchestration that needs to exist to be able to deterministically and continously being able to delivery software to users. 

## Pipeline run events

A pipeline, in the context of Continuous Delivery, is the definition of a set of tasks that needs to be performed to build, test, package, release and deploy software artifacts. A pipeline can be instanciated multiple times, for example to build different versions of the same artifact. We are refering to this instance as PipelineRun. It will have a unique Id and it will help us to track the build and release progress on a particular software artifact. 

Due the dynamic nature of Pipelines, most of actual work needs to be queued to happen in a distributed way, hence Queued events are added. Adopters can choose to ignore these events if they don't apply to their use cases. 

- **PipelineRun Queued**: a PipelineRun has been schedule to run
- **PipelineRun Started**: a PipelineRun has started and it is running
- **PipelineRun Finished**: a PipelineRun has finished it execution, the event will contain the finished status, success, error or failure

Each pipeline is defined as a set of Tasks to be performed in sequence, hence tracking this sequence might be important for some cases. A TaskRun is an instance of the Task defined inside the pipeline, as you can expect multiple execution of the pipelines (each a PipelineRun) you can also expect multiple execution of the Tasks, for that reason we use TaskRun to refer to one of these instances.  

- **TaskRun Started**: a TaskRun inside a PipelineRun has started.
- **TaskRun Finished**: a TaskRun inside a PipelineRun has finished. 

Pipeline Events MUST include the following attributes:
- **Event Type**: the type is restricted to include `CD.Pipeline**` prefix. For example `CD.Pipeline.PipelineRun.Queued` or `CD.Pipeline.Tests.Started`
- **PipelineRun Id**: unique identifier for a pipeline execution
- **Pipeline Name**: unique identifier for the pipeline, not for the instance. A pipeline can have multiple instances/runs.  
- **PipelineRun Status**: current status of the PipelineRun at the time when the event was emitted. If the pipeline is finished, this attribute should reflect if it finished successfully or if there was an error on the execution. Possible statuses: [Running, Finished, Error]

Optional attributes: 
- **PipelineRun URL**: URL to locate where  pipeline instances are running
- **PipelineRun Errors**: error field to indicate possible compilation, test, build and package errors.

## Product Composition Events

When the products delivered to users can be, or needs to be, broken down into components (hardware, software, configuration files etc.), the creation of versions of said components and the compositoning of components into products or product subsystems may warrant the use of separate events detailed in this section.

Product Composition Events MUST include the following attributes:

- **Event Type**: the type is restricted to include `CD.ProductComposition` prefix. For example `CD.ProductComposition.ComponentVersionCreated` or `CD.ProductComposition.CompositionCreated`

### Component Version Created

Announces that a new version of a Product Component has been created.

Examples of Component Versions include:

- New version of a software module or application.
- New version of a document or piece of configuration.
- New revision of a hardware component.

The purpose of this event would be for systems to start creating Product Compositions / Baseline Versions that include this Component Version.

Component Version Created Events MUST include the following attributes:

- **Component Version ID**: Unique identifier of component version.
- **Component ID**: Identifier of component of which this is a version.

Optional attributes:

- **Version Name**: Machine-friendly and human-friendly version tag/number/etc. for component version, e.g. “1.4.2-alpha2”.
- **Component Name**: Machine-friendly and human-friendly name of component, e.g. “c37-backbone-router”.
- **Component Display Name**: UX/human-friendly display name of component, e.g. “C37 Backbone Router”.
- **Artifacts**: List of artifacts (files, packages, Docker images etc.) part of this component version.
- **Component Track**: States what track (e.g. platform variant (x86_64, ARM etc.), project (e.g. X-DataCenter, X-Cloud etc.) or other) the component, of which this is a version, belongs to.

### Product Composition Created

Announces that a new composition of source-files, artifacts, component versions and other compositions/baseline versions, has been created.

__Note:__ This event description suffers from a missing piece of terminology to represent the “recipe” by that controls what is included in new compositions. E.g. “The latest stable version (LSV) of this application, the LSV of this library, and the LSV of this configuration file”. For now, this will be referred to as a “Recipe” until a better name is given. (ESCI calls this a Baseline, Eiffel likely has a term for this as well).

Product Composition Created Events MUST include the following attributes:

- **Composition ID**: Unique identifier of composition.
- **Recipe ID**: Identifier of recipe of which this is a composition.
- **Content**: List of sources, artifacts, components and other compositions part of this composition

Optional attributes:

- **Version Name**: Machine-friendly and human-friendly version tag/number/etc. for composition, e.g. “1.4.2-alpha2”.
- **Recipe Name**: Machine-friendly and human-friendly name of recipe, e.g. “gemeni-backbone-gen2”.
- **Recipe Display Name**: UX/human-friendly display name of recipe, e.g. "Gemini Backbone, Gen2”.
- **Artifacts**: List of artifacts (files, packages, Docker images etc.) part of this component version.
- **Component Track**: States what track (e.g. platform variant (x86_64, ARM etc.), project (e.g. X-DataCenter, X-Cloud etc.) or other) the component, of which this is a version, belongs to.

### Confidence Label Verdict

Announces that a Confidence Label Verdict (a.k.a. Confidence Level) for one or more source changes, artifacts, component versions or compositions has been set or updated.

Confidence Label Verdict Events MUST include the following attributes:

- **Confidence Label Verdict ID**: Unique identifier of confidence label verdict.
- **Name**: Machine-friendly and human-friendly name of confidence label
- **Verdict**: The verdict of this confidence label result (Success, Failure, Inconclusive)
- **Targets**: List of sources, artifacts, component versions and compositions this confidence label verdict applies to

Optional attributes:

- **Causes**: List of e.g. pipeline runs or test suites that lead to this confidence label verdict.
- **Sub Confidence Label Verdicts**: List of other confidence label verdicts that resulted in this confidence label verdict. Used for “summarizing” (e.g. “parent”) confidence labels.

### Environment Allocation Status Changed

__Open issue:__ Move to Continouous Deployment bucket? Bake into Environment Changed?

Announces that the allocation status of an environment has changed (e.g. it has been allocated, de-allocated or taken offline).

Environment Allocation Status Changed Events MUST include the following attributes:

- **Environment ID**: Unique identifier for the environment.
- **Allocation Status**: The new allocation status for the environment (Allocated, Idle, Offline)

Optional attributes:

- **Environment Name**: Machine-friendly and human-friendly name of recipe, e.g. "gemini-backbone-hil-4".
- **Environment Display Name**: UX/human-friendly display name of recipe, e.g. "Gemini Backbone HIL 4”.
- **Environment URL**: URL to reference where the environment is located.
