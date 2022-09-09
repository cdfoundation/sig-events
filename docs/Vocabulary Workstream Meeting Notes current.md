---
tags: CDEvents
---

# CDEvents / Vocabulary Discussion Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA/badge)](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)

This document contains the notes from the of the Events SIG meetings focused on [vocabulary discussion](https://hackmd.io/lBlDCrL7TvmtNOjxdopJ5g).


## Sep 14th
*Meeting canceled due Open Source Summit*

## Sep 6th

Participants:

- Andrea Frittoli, he/him, IBM, UTC+1
- Emil Bäckmark, Ericsson, UTC+2
- Kara de la Marck, CDF, UTC-4
- Jalander Ramagiri, Ericsson Software Technology, UTC+1
- Erik Sternerson, doWhile

Agenda:

- Action items from last time:
    - spec: adding json schema: WIP https://github.com/cdevents/spec/pull/61
    - spec: adding `data` field: https://github.com/cdevents/spec/pull/63 -> merged
        - WIP adding `customData` to the `go-sdk` https://github.com/cdevents/sdk-go/pull/19
    - spec: Subscriber model considerations, add diagram to spec: TBD
    - document decision (from CDEvents community summit) about prescriptive events in the spec: TBD
    - poc / demo implementation: Create issue to track this work: TBD

- Versioning: https://github.com/cdevents/spec/issues/43
    - From SDK user pov, users are most likely interested to send an event by spec version (event version is then implied)
    - The receiving side needs to know the event version (spec version is not so relevant)
    - Proposed starting approach
        - Version in the context for the spec, free text
        - Semantic version for events in the event type string

- CDEventer demo (Andrea)

- Documentation format:
    - Initial improvements: https://github.com/cdevents/spec/pull/63

- Go SDK unit tests: https://github.com/cdevents/sdk-go/issues/2 -> Issue clused now

- Document general required features for sdk
    - Community repo
    - Unit test coverage
    - Conformance tests (hosted in the spec/jsonschema?)
    - We need to pin the spec version in the SDK


## Aug 31st

Participants:
- Andrea Frittoli, he/him, IBM, UTC+1
- Emil Bäckmark, Ericsson, UTC+2
- Mattias Linnér, Ericsson, UTC+2
- Brad McCoy, he/him, Basiq, UTC+11


Agenda:

- Action items from last time:
    - sdk-go: support for validating through json schema: DONE
    - spec: adding json schema: WIP https://github.com/cdevents/spec/pull/61
    - spec: adding `data` field: TBD https://github.com/cdevents/spec/issues/60
    - spec: add the version to the schema: DONE in the go-sdk, for the spec: https://github.com/cdevents/spec/pull/61
    - java-sdk: Check with Jalander if he plans to update the Java SDK with the new spec version (draft/v0.1): TBD
        - Java SDK presentation at OSS Mini summit with Eiffel
        - No plan on updating the SDK before v0.1
        - It should not be a requirement for v0.1, but it still could be done
        - We should work on it at least once v0.1 is released
    - spec: Subscriber model considerations, add diagram to spec: TBD
    - document decision (from CDEvents community summit) about prescriptive events in the spec: TBD
    - poc / demo implementation: Create issue to track this work: TBD

- sdk-go: discuss comments on https://github.com/cdevents/spec/pull/61

- spec/sdk: event versions
    - Old and new SDK include `v1` in event types:
        - dev.cdevents.artifact.packaged.v1
    - Spec does not include version in event types:
        - dev.cdevents.artifact.packaged
    - (Andrea) I think we should include version
    - What versioning schema do we use and how does it compare to the spec version? https://github.com/cdevents/spec/issues/43 

- poc / demo implementation
    - Mattias: Could the SDK serve this purpose?
    - Emil: SDK (one or more will be used), but do not rely on specific tools (Tekton, Jenkins, etc)
    - Mattias: we should lower the barrier for getting started with CDEvents
    - Emil: existing POCs can be complex to understand as they require specific tool knowledge
    - A simple CDEvent listener / display tool would be useful
    - We could use the TOC budget for CDEvents for that project
    - (A) Bring the topic to the next TOC meeting
    - (A) Find the issue with the description of the project
        - https://github.com/cdfoundation/sig-events/issues/126

- spec format:
    - Improve the event documentation in the spec: https://gist.github.com/afrittoli/372489774e1eed01756b54098c06eb76

- (Brad) Using k8s events to generate CloudEvents / CDEvents
    - Would be useful for Flux / CDEvents
    - Knative example: https://knative.dev/docs/eventing/sources/apiserversource/


## Aug 23rd

Participants:
- Andrea Frittoli, he/him, IBM, UTC+1
- Erik Sternerson, doWhile
- Tarek Badr, doWhile
- Ben Powell, Apple
- Emil Bäckmark, Ericsson


Agenda:

- GO SDK Updates
    - TaskRun, PipelineRun and Change events completed (matching spec)
    - All other events added, matching spec WIP
    - Added support for creating JSONschemas from Go types
    - Added support for validating an `interface{}` through a JSONschema
        - One pending issue on https://github.com/cdevents/sdk-go/pull/6

- Spec updates:
    - Formatting of the spec merged and live at https://cdevents.dev/docs
    - CDEvents binary mode removed
    - Adding json schema https://github.com/cdevents/spec/pull/61
    - Add issue to track the `data` field to be added https://github.com/cdevents/spec/issues/60
        - Add the version to the schema
            - Add default value for the "version" field
            - or enum with single value
            - example from Eiffel
                - ```json
                  "type": "string",
                  "enum": [
                    "3.1.0"
                  ],
                  "default": "3.1.0"
                  },
                  ```
        - See https://schema.org/docs/schemas.html 
        - No additional properties in spec fields, only in `data`
            - already in the schema
                - https://github.com/cdevents/spec/pull/61/files#diff-115d17a864e7c9f66d3f6f4b51282e803d161c93fcc2d0294b2cbf55179be200R24

- V0.1 roadmap updates
    - Actually closed issues marked as done (project automation missing there)
    - Removed PoC from v0.1 plan 
    - Added more details to https://github.com/cdevents/spec/issues/13
        - Created 4 sub-issues for the type of events required for the DORA PoC
    - Reviewed the v0.1 milestone to ensure it's aligned to with the project
    - Java SDK
        - Check with Jalander if he plans to update the Java SDK with the new spec version (draft/v0.1)

- Py SDK updates
    - DRAFT PR: https://github.com/cdevents/sdk-python/pull/1
    - Default values for context fields
        - Source
        - Id
        - Timestamp
        - reusable application specific context could be useful (across SDKs)

- Spec Consideration
    - Subscriber model considerations
        - Add a diagram like in the spec https://raw.githubusercontent.com/cdfoundation/sig-events/main/poc/sig-events-spinnaker-functional.png
        - Make it clearer in the spec that the subscriber model is the primary use case
        - It's in the primer today but we need to make it more relevant
            - <https://cdevents.dev/docs/primer/#why-not-point-to-point-communication>

- Should we provide PoCs or reference implementations aligned with each CDEvents release?
    - Build a publish/suscribe use case
    - It doesn't have to use an actual tool (tekton, keptn, etc)
    - Easy to demo, easy to update for each release
    - Should this be a requirement for v0.1?

- Migrate content from sig-events to cdevents


## Aug 17th

Participants:
- Mattias Linnér, Ericsson
- Emil Bäckmark, Ericsson
- Andrea Frittoli
- Brad McCoy


Agenda:

- Community repo
    - https://github.com/cdevents/community
    - comments welcome
	- Fill in the project description
	- Update links about creating issues
	- Moving issues across
	- Setup markdown linter
	- Drop the WIP

- GO SDK Update
    - Initial working SDK at https://github.com/cdevents/sdk-go/pull/1
    - Short walkthrough the SDK
    - CI pending
	- Check unit tests from the old SDK
	- Documentation:
        - Development.md
        - Contribute.md
    - OWNERS team

- [Release v0.1](https://github.com/orgs/cdevents/projects/1/views/1) reviews
    - Schema: https://github.com/cdevents/spec/issues/13
    - Milestone v0.1: https://github.com/cdevents/spec/milestone/1

## Aug 9th

Participants:
- Erik Sternerson, doWhile
- Tarek Badr, doWhile :)
- Andrea Frittoli, IBM
- Emil Bäckmark, Ericsson
- Ben Powell, Apple
- Kara de la Marck, CDF
- 

Agenda:

- Meeting recordings
    - Updated to youtube via Zapier (Andrea's personal account)
    - One manual step left, add video to playlist, set to public

- Python SDK demo
    - (Erik&Tarek) Think about minimum supported Python version
        - Cloudevents Python SDK is Python 3.6+

- Should we align across SDKs?
    - In terms of SDK, we shall follow standard expectation for each specific language
    - We need cross-compatibility / conformance tests across SDKs

- SDK "readiness" checklist
    - Send events?
    - Receive events?
    - "Classes" for all event types?
        - Classes can have strongly typed properties and documentation, which gives a better user experience in many cases.
    - Other?
        - Documentation
            - Generated docs
            - Basic README intro to the project
        - Test cases / conformance tests (cross-sdk)

- Shall we have one CLI?
    - We need to pick a language
        - Python: pros, no arch specific builds, cons: you need the correct python version
        - Go: multiplatform, but it requires multiple builds
        - (Andrea) create an issue to track this discussion

- Go SDK
    - Andrea started working on the SDK
    - Proposed API for the SDK:
        - Create generic CDEvent by source, subjectId
            - Version is set automatically, can be overwritten
            - Event ID is generated (?), can be overwritten
            - Timestamp is set to the even creation time, can be overwritten
            - Subject source and event source are set to source
        - SubjectContent lets read/write fields by key name
        - One go struct per event type, which meets the SubjectContent interface
        - Create the struct, set it to the event
            - ```golang
              cdevent := NewEvent(subjectId, source)
              sc := PipelineRunQueuedSubjectContent{
                    PipelineName: "pipeline",
                    URL: "someurl",
              }
              cdevent.SetContent(sc)
              ```
            - Sets the event type in the main event
            - Sets the subject content
        - Get the event as CloudEvent, send it through the CloudEvents SDK go client
            - ```golang
              ce := cdevent.AsCloudEvent()
              result := ceClient.Send(cloudevents.ContextWithRetriesExponentialBackoff(ctx, 10*time.Millisecond, 10), *ce);
              ```
    - Using generic in go?
    - We need to be able to receive cdevents into defined objects
    - Working with empty interfaces can be painful

- For the CLI probably keep the same structure as today?

- Binary vs Structured mode
    - Andrea to roll back existing part about binary/structured
    - Add new field for vendor data



## Aug 3rd

_Meeting canceled_

## July 26th

Participants:
- Andrea Frittoli, IBM
- Erik Sternerson, doWhile
- 


Agenda:

- Metrics POC events
    - Change events
        - Merged event
            - Data model missing
                - Sha of last commit
                - List of all shas
    - Artifact events
        - Source repos and last sha for each
            - Using a list from the start
            - Let's try with a list 
            - we can revert to a single repo if running into issues during the POC
        - Assumption
            - Many cases with a single repo
            - Some cases with many repo
            - Let's try use the latter to cover all cases
        - PR soon
    - Incident events
        - Add new stage?
        - Incident created, updated, resolved
        - PR with bare minium data model

- Binary Mode POC 
    - https://github.com/cdevents/spec/pull/54 
    - What about deeper levels
        - Other not support them in "binary" mode
        - Json serialised to string
        - Base64 encoded (to be signalled to receivers)
    - Why "binary" mode? 
        - We could have clearer naming like HTTP header mode
        - Andrea to check with CloudEvents team
        - Consistency of naming may be a plus

- Python SDK
    - Work progressing

- Meetings
    - Weekly cadence
        - Alternating weeks help us progress
    - How do we make sure people who can attend one meeting (and not both) do not feel left out
        - Recap in the beginning of the meeting?
        - Advice to read the notes
        - Upload for meetings
            - Andrea to look into recording automation

- Java SDK: https://github.com/cdevents/sdk-java

## July 20th

Participants:
- Erik Sternerson, doWhile
- Andrea Frittoli, IBM
- Mattias Linnér, Ericsson
- Meha Bhalodiya
- Vibhav Bobade, Uffizzi
- Jamie Plower, Fidelity Investments

Agenda:

- Updates
    - Python SDK
        - First goal is 1:1 parity with Go SDK, and evolve from there
        - Current ongoing work in forked repo at https://github.com/tarekbadrshalaan/sdk-python/tree/skeleton-project (very much work-in-progress status)
        - Actively under development
        - 
    - Spinnaker POC
        - PR merged on SIG events side
        - Variable image sha implemented
    - Website
        - [Nighty job to publish spec](https://github.com/cdevents/cdevents.dev/actions/workflows/nightly-deploy.yaml)
    - CDEvents presentation at Open Source Summit Europe
        - Friday Sept 16 @ 15:55-16:35
        - https://sched.co/15z6i
    - Reformatting of the spec, updates to lowerCamelCase
        - Working on the CI bucket (Andrea)
            - https://github.com/cdevents/spec/pull/53
        - Next up the CD bucket
        - Lots of missing parts in the data model
    - Java SDK
    - GitHub webhook to CDEvents

- Build events
    - https://github.com/cdevents/spec/issues/37
    - What do we mean by change?
        - PR/MR/etc or 
        - Commit sha
    - Eiffel does not have a build event
    - Eiffel has a concept of composition
        - A composition may combine multiple inputs, with a version
    - Jamie - SCM event correlates to clone event which correlates to artifact
    - We could use artifacts only for the metrics POC

- CDEvents content mode (on top of CloudEvents HTTP Binary mode):
    - Structured:
        - JSON payload with CDEvents meta, subject, links as "payload"
        - Content type for payload content?
        - ```http
          Content-Type: application/cdevents+json; charset=utf-8
          ce-headers: ...
          
          {
           "meta": {
              "version": "draft",
              "id": "A234-1234-1234",
              "source": "/staging/tekton/",
              "type": "dev.cdevents.taskrun.started",
              "timestamp": "2018-04-05T17:31:00Z",
              "datatype": "application/json+my-vendor-app"
           }
           "subject": {
              "id": "/namespace/taskrun-123",
              "type": "taskrun",
              "content": {
                 "task": "my-task",
                 "URL": "/apis/tekton.dev/v1beta1/namespaces/default/taskruns/my-taskrun-123"
                 "pipelinerun": {
                    "id": "/somewherelse/pipelinerun-123",
                    "source": "/staging/jenkins/"
                 }
              }
           }
           "data": { <vendor-data> }
           }
          ```
    - Binary:
        - CDEvents meta, subject, links all in HTTP headers
        - ```http
          Content-Type: application/xml; charset=utf-8
          ce-headers: ... 
          cde-meta:  version=draft; id=A234-1234-1234; source=/staging/tekton/; type=dev.cdevents.taskrun.started; timestamp=2018-04-05T17:31:00Z
            cde-subject: id=/namespace/taskrun-123; type=taskrun; content-task=my-task; content-URL=/apis/tekton.dev/v1beta1/namespaces/default/taskruns/my-taskrun-123; content-pipelinerun-id=/somewherelse/pipelinerun-123; content-pipelinerun-source=/staging/jenkins/

          <xml>
              <root>Some vendor content</root>
          </xml>
          ```

- Community repo (Andrea)
    - Plan to create a cdevents/community repo over the summer
    - Migrate existing relevant docs and links
    - Define governace docs in the new repo

- CDEvents visualisation
    - https://github.com/cdfoundation/sig-events/issues/126 

- Spec feature branches for POCs?


## July 12th

\<No meeting\>
    
## July 6th

Participants:
- Emil Bäckmark, Ericsson, UTC+2
- Andrea Frittoli, IBM, UTC+1
- Mattias Linnér, Ericsson, UTC+2
- Meha Bhalodiya
- Brad McCoy, UTC+10
- 

Agenda:

- Brad to go over Flux integration for CDEvents
- Introduce Subjects: https://github.com/cdevents/spec/pull/35
    - Type:
        - The first part is all reverse DNS
        - We can keep all lower case?
    - UpperCamelCase vs lowerCamelCase
        - lowerCamelCase seems more popular
        - we should use same format for keys and enum values
        - document the decision as part of the PR
    - Subjects and Predicates https://github.com/cdevents/spec/pull/35#discussion_r911184157 
        - When grouping events, having the static part of the subject separate could be useful
        - Split this into a different PR?
        - 
- CDEvents content mode (on top of CloudEvents HTTP Binary mode):
    - Structured:
        - JSON payload with CDEvents meta, subject, links as "payload"
        - Content type for payload content?
        - ```http
          Content-Type: application/cdevents+json; charset=utf-8
          ce-headers: ...
          
          {
           "meta": {
              "version": "draft",
              "id": "A234-1234-1234",
              "source": "/staging/tekton/",
              "type": "dev.cdevents.taskrun.started",
              "timestamp": "2018-04-05T17:31:00Z",
              "datatype": "application/json+my-vendor-app"
           }
           "subject": {
              "id": "/namespace/taskrun-123",
              "type": "taskrun",
              "content": {
                 "task": "my-task",
                 "URL": "/apis/tekton.dev/v1beta1/namespaces/default/taskruns/my-taskrun-123"
                 "pipelinerun": {
                    "id": "/somewherelse/pipelinerun-123",
                    "source": "/staging/jenkins/"
                 }
              }
           }
           "data": { <vendor-data> }
           }
          ```
        - Useful for storing events as JSON blobs in a DB
            - Do we need the CloudEvents part there too?
                - We could consider structured mode on top of CloudEvents structured, mostly for this kind of storage as JSON use cases
    - Binary:
        - CDEvents meta, subject, links all in HTTP headers
            ```http
            Content-Type: application/xml; charset=utf-8
            ce-headers: ... 
            cde-meta:  version=draft; id=A234-1234-1234; source=/staging/tekton/; type=dev.cdevents.taskrun.started; timestamp=2018-04-05T17:31:00Z
            cde-subject: id=/namespace/taskrun-123; type=taskrun; content-task=my-task; content-URL=/apis/tekton.dev/v1beta1/namespaces/default/taskruns/my-taskrun-123; content-pipelinerun-id=/somewherelse/pipelinerun-123; content-pipelinerun-source=/staging/jenkins/
              
            <xml>
                <root>Some vendor content</root>
            </xml>
            ```


## June 28th

Participants:
- \<add yourself\>
- Andrea Frittoli, he/him, IBM, UTC+1
- Emil Bäckmark, he/him, Ericsson, UTC+2

Agenda:

- Meetings
    - Poll: https://forms.gle/oC2qto7hjxBAGPvq8
    - Responses: https://docs.google.com/forms/d/1h2kmkcoV5CdeiFfFaLfGJ_dxuAIHb7I8D3hNy5R7bp0/edit#responses
    - Existing meeting time
    - (Andrea) Make proposals in CDEvents channel for specific times/dates and ask people to vote
        - Avoid conflicts within the CDF calendar
        - Options (added after the meeting)
            - :one: Monday July 4th, 10am UTC, and then every other week
            - :two: Monday July 4th, 11am UTC, and then every other week
            - :three: Tuesday July 5th, 10am UTC, and then every other week
            - :four: Tuesday July 5th, 11am UTC, and then every other week
            - :five: Wednesday July 6th, 10am UTC, and then every other week
            - :six: Wednesday July 6th, 11am UTC, and then every other week
    - We would have three meetings in total
        - SIG meeting
        - 2 CDEvent meetings
    - Who would be interested in joining from Pacific
        - Ben
        - Check who else
        - We could have a shorter follow-up meeting at a Pacific friendly time?

- Updates on https://github.com/cdevents/spec/pull/35
    - Updated data section, introduced content modes in `cloudevent-binding.md`
    - Changed format of fields on `core.md`
    - Addressed comments on `spec.md`
    - Schema updates pending - move schemas to different PR?
    - To be done - next PR?
        - CDEvents binary mode (all in headers, vendor data in payload) 
        - Full examples for various modes, including vendor payload

    - Up next: Updates to the rest of vocabulary docs, similar to `core.md`

- Aligning the vocabulary with SIG Interop
    - https://github.com/cdevents/spec/issues/18 
        - Interop WG: Pipeline, stage, step
        - CDEvents: Pipelinerun, taskrun
    - Make a statement in our docs
        - We strive to align with SIG interop names from the vocabulary
        - The final naming won't be ready in time for v0.1
        - People should be aware that names may/will change in later releases
    - (Andrea) Reach out / join Interop SIG to clarify plan for defining the names
    - Comment on issue https://github.com/cdevents/spec/issues/18 
        - Create a new issue for the docs part discussed today

- Roadmap: 
    - https://github.com/orgs/cdevents/projects/1/views/1
    - Consolidate PoCs back to cdevents org
    - Review and consolidate discussions to the cdevents org

- Support for links:
    - Anyone would like to propose spec changes?
    - Emil is interested

- How to collaborate on spec
    - It's hard to collaborate on PR/issues for spec
    - We could attempt working on an hackmd docs until it's ready and then PR

- SDKs
    - Python SDK updates
    - Java SDK targeted for v0.1

- PR to review:
    - Spinnaker PoC: https://github.com/cdfoundation/sig-events/pull/120
    - Events published: https://github.com/cdevents/spec/pull/31
    - Subjects: https://github.com/cdevents/spec/pull/35

- Create a community repo
    - Governance setup
    - PoC discussion
    - Overall cdevents issues / discussions (not spec specific)
    - Is it possible to move issues across orgs or export/import?

- Summer meetings
    - It looks like enough people might be working through the summer
        - Erik, Jalander, Andrea
        - July 4th holiday in the US
        - Emil away 4 weeks after the 4th
    - Let's keep meetings in July and August


## June 14th

Participants:
- \<add yourself\>
- Erik Sternerson, doWhile
- Andrea Frittoli, he/him, IBM, UTC+1
- Emil Bäckmark, Ericsson
- Ben Powell, he/him, Apple
- Kara de la Marck, CDF

Agenda:

- Updates:
    - Whitepaper: https://cd.foundation/blog/2022/06/07/cdevents-publishes-first-whitepaper/
    - Project Announcement:
        - https://sdtimes.com/devops/cd-foundation-announces-new-specification-for-defining-event-data-format/
        - https://www.martechcube.com/cd-foundation-announces-cdevents/
    - DevOps.com: https://devops.com/continuous-delivery-foundation-adds-interoperability-project/
    - DevOps.com: https://devops.com/cdevents-aims-to-standardize-ci-cd-interoperability/
    - TechTarget.com: https://www.techtarget.com/searchitoperations/news/252521316/OpenTelemetry-inspires-CDFs-event-driven-architecture-plan
    - TechTarget.com: https://www.techtarget.com/searchitoperations/news/252521264/New-CD-Foundation-GM-fights-CI-CD-pipeline-fragmentation


- Roadmap review:
    - Anything redundant?
    - Anything missing?
- SDK automations
    - Swagger, smithy (?) modeling languages to explore
    - We must rely on CloudEvents SDK as base layers for the SDKs
    - We would like to automate docs creation too where possible
    - 
- Payload structure:
    - [Initial PR](https://github.com/cdevents/spec/pull/35) (already discussed last time)
    - [Further discussion](https://github.com/cdevents/spec/issues/40) based on Ben's input
    - [Schema issue](https://github.com/cdevents/spec/issues/14) in the roadmap
- Dependency Updates:
    - [PoC Dependency Updates through events](https://github.com/cdevents/spec/issues/39)
- Versioning Strategies:
    - [#43](https://github.com/cdevents/spec/issues/43), [#40](https://github.com/cdevents/spec/issues/40)
    - Specification version
        - Semantic versioning
    - Event version
        - Add version numbers to events before v0.1
    - How and when do we provide backward compatibility?
        - e.g. Consumer v1 receives events from producer v2
- To review:
    - Spinnaker PoC: https://github.com/cdfoundation/sig-events/pull/120
    - Events published: https://github.com/cdevents/spec/pull/31


## May 31st
*Meeting canceled due to no participants joining*

## May 3rd

Participants:
- Emil Bäckmark, Ericsson
- Andrea Frittoli, IBM
- Mattias Linnér, Ericsson
- Erik Sternerson, doWhile
- \<add yourself\>

Agenda:
- Introduction of objects and subjects: https://github.com/cdevents/spec/pull/35
    - Spent the whole meeting discussing this
    - (Emil) add results from discussion on the PR
- "Requested" events: https://github.com/cdevents/spec/issues/36
    - We didn't have time to discuss this.
- About CDEventsCon
    - Someone needs to be responsible to handle the chat from virtual audience and bring it to the in-person speaker

## April 19th

Participants:
- Emil Bäckmark, Ericsson
- Liora Milbaum, RedHat

Agenda:
- https://github.com/cdevents/spec/issues/36
    - We didn't bring this up due to lack of participants in the meeting.
- Since it was only two people in the meeting we spent it on getting to know each others background and discussing the use/need of events in CI/CD and Software Supply Chains

## April 5th

Participants:
- Erik Sternerson, doWhile
- Kevin Chu, GitLab
- Emil Bäckmark, Ericsson
- Tracy Ragan, DeployHub / Ortelius OS
- Mattias Linnér, Ericsson

Agenda:

- Events:
    - CDEvents at cdCon
        - Half a day, June 9th afternoon
            - Texas Time -> that's going to be night in Europe
            - Morning time will enable European participants to take part virtually
        - Title: CDEvents Community Summit
        - Event description:
            - CDEvents v0.1 is coming, let’s implement it in the wild! We are building SDKs for go, java and python - let’s use them to add CDEvents to your project. Meet at cdCon, CDEvents users and developers, to learn about CDEvents and start using them!
            - 
    - CDEventsCon
        - Registration is now live
            - Still missing on [event page](https://events.linuxfoundation.org/cdeventscon/) though
        - Draft agenda available on dedicated [sched page](https://cdeventscon2022.sched.com/)
        - Content to be finalised / confirmed:
            -  Eiffel talk
            -  Keptn workshop (Oleg driving this)
            -  SIGs panel (Oleg helping to organise)
            -  Spinnaker (lighting) talk 
        -  Setup for hybrid event being clarified
            -  Which platform (zoom)?
            -  Who's managing the virtual room?
                -  Remote speakers
                -  Access to stream for remote viewers
        -  Kevin volunteered to help with the conference
        -  Promoting the event:
            -  Twitter
            -  Homepage (https://cdevents.dev/)
            -  README (org)
            -  Slack
-  SDKs
    -  First PRs merged for Java SDK!
        -  https://github.com/cdevents/sdk-java
    -  Eventually it would be good to use the SDKs to send CDEvents for our own automation
-  Spec
    -  Intro from Mattias merged https://github.com/cdevents/spec/pull/34 
    -  https://github.com/cdevents/spec/pull/35 

## March 22nd

Participants:

- Emil Bäckmark, Ericsson
- Andrea Frittoli, IBM
- Mattias Linnér, Ericsson
- \<add yourself\>

Agenda:

Events:
- cdCon Virtual Track
    - Panel / BoF session?
    - Andrea to ask Kara if panels are accepted
- CDF Online Meetup
    - Date to be decided
    - Panel for CDEvents
    - Tracy to start a google doc about this https://docs.google.com/document/d/12UNdVBO7_rOBIEtOsrjViEjVUdp8dyYl_tM5ljSP7Ps/edit?usp=sharing
- DevOps Institue
    - CfP still open https://www.devopsinstitute.com/cicd-2021/
- CDEventsCon Content
    - Event Page: https://events.linuxfoundation.org/cdeventscon/ 
    - Planning: https://hackmd.io/JMM-V7IlTwWoS1YcDK9m_A
    - Looking for speakers / demos / workshops
- cdCon CDEvent contributor summit
    - Hybrid Event
    - Planning: TBD
- KubeCon EU
    - CDF Booth is not going to be there

Workshops:
- Write an receive/send adapter with the SDK?

Proof of Concepts:
- [Knative PoC](https://github.com/cdfoundation/sig-events/issues/107)
- [Metrics PoC](https://hackmd.io/cFw0xO6XSwGwIoJ78Re91A)

SDKs:
- Defined schemas will help with SDK creations
    - Andrea is working on this
- Golang
- Python
- Java

Pull Requests:
- Introduction https://github.com/cdevents/spec/pull/34
- Service Published https://github.com/cdevents/spec/pull/31
- Subjects WIP https://github.com/cdevents/spec/pull/35

Interesting tool: https://github.com/direktiv/direktiv
    - Session for SIG Events

Andrea on PTO 07/04 -> 25/04

## 22nd February

Participants:

- Andrea Frittoli
- Mattias Linnér, Ericsson
- Mauricio Salatino (salaboy), VMware
- Ishan Khare
- \<add yourself\>


Agenda:

- Announcements
    - [CDEvents Day @ KubeCon](https://docs.google.com/forms/d/e/1FAIpQLSd7CSxSs_Q0mxromY-FXzEEWyJdKycniKMBHThycJsDXrA9nQ/viewform?usp=sf_link)
    - [Contributor Summit](https://docs.google.com/forms/d/1iIKp_xh3Sx3Mh9hp_VYAIAyR_EkF-nVoTp2JkGkdErQ/edit)
    - [v0.1 Planning](https://github.com/orgs/cdevents/projects/1/views/1)


- Issues to discuss:
    - Service Published event: 
        - https://github.com/cdfoundation/sig-events/issues/111 
        - We agree on `published`
        - Needs to be clearly documented when to use which event type
        - Andrea to create a separate issue about predicates
    - Using source in CDEvents: 
        - https://github.com/cdevents/spec/issues/29
        - https://cloud-native.slack.com/archives/C9DB5ABAA/p1645456569204809
        - https://cloud-native.slack.com/archives/C9DB5ABAA/p1645461986213429
        - The source should allow the received to distinguish different producers
        - Should the source be an object?
            - We could have an object in CDEvents
            - Serialize it to subject?
        - Eiffel: https://github.com/eiffel-community/eiffel/blob/master/eiffel-vocabulary/EiffelActivityTriggeredEvent.md#metasource
    - Update the SDK
        - https://github.com/cdevents/spec/issues/21
        - Create new issue for the SDK (Andrea)

- Pull Requests:
    - [Add a contributing.md and improve the README.md](https://github.com/cdevents/spec/pull/28)
        - Move to `.github` repo and add link into the README.md
    - [Add syntax details for attributes in CDEvents](https://github.com/cdevents/spec/pull/30)


- Proof of concepts
    - [Knative PoC](https://github.com/cdfoundation/sig-events/issues/107)
        - Good progress on Knative side
    - Crossplane PoC - issue to be created
    - [Metrics PoC](https://hackmd.io/cFw0xO6XSwGwIoJ78Re91A)
    - Supply Chain with Cartographer - Heads up - Issue to be created

- Project updates
    - Tekton
    - Jenkins
    - Knative
    - Spinnaker
    - Keptn
    - Cartographer
        - Mauricio reaching out
        - Interested in integrating with Knative and perhaps CDEvents
    - Crossplane
        - cdCon proposal by Mauricio & Ishan


## 8th February

Participants:

- Erik Sternerson, doWhile
- Kara de la Marck, CDF
- Ishan Khare
- Emil Bäckmark, Ericsson
- Vibhav Bobade
- \<add yourself\>

Agenda:

- Announcements
    - [Website](https://cdevents.dev)
    - [Mailing list](https://groups.google.com/g/cdevents-dev)
    - [Slack Channel](https://cdeliveryfdn.slack.com/archives/C030SKZ0F4K)
    - CDEvents [Issues](https://github.com/cdevents/spec)
    - [Contributor Summit](https://docs.google.com/forms/d/1iIKp_xh3Sx3Mh9hp_VYAIAyR_EkF-nVoTp2JkGkdErQ/edit)
    - [v0.1 Planning](https://github.com/orgs/cdevents/projects/1/views/1)
    - [Webinar](https://doodle.com/poll/34u9znbdmnayx7k4)

- Issues to discuss:
    - [Consider taking advantage of this list of task/step types when qualifying event types](https://github.com/cdevents/spec/issues/18)
        - The discussion is still ongoing on SIG interop side
        - There are two discussions going on, one about "[Steps](https://github.com/cdfoundation/sig-interoperability/pull/81)" and "[Stages](https://github.com/cdfoundation/sig-interoperability/pull/76)"
        - 
    - Usage of [subject](https://github.com/cdevents/spec/issues/11)
        - There is no subject equivalent in Eiffel
        - It should be an optional field on the CloudEvents binding
        - It could be an optional field on CDEvents
        - Andrea to make a PR
    - [Link produced events and consumed events](https://github.com/cdevents/spec/issues/10)
        - We should support different kind of links
        - The standardization work done in defining "objects" in CDEvents can be used to provide extra context (still specified by CDEvents)
        - Andrea to make a PR
    - [Add cd.service.published event to vocabulary](https://github.com/cdfoundation/sig-events/issues/111)

- SDKs
    - Golang, Python, Java
        - For Python SDK, any preference on project/dependency solution?
            - Python
                - Poetry, setuptools etc.
                - No preference stated, first person to start developing can decide.
            - We can continue to work async on the SKD for now
            - 

- Proof of Concept
    - [Knative PoC](https://github.com/cdfoundation/sig-events/issues/107)
    - Metrics PoC
        - Andrea is going to create an issue to track the work

- Project updates
    - Tekton
        - Resuming work on experimental controller with Vibhav
        - Setup a roadmap for the project
        - Plan to move out of experimental
    - Jenkins
    - Knative
    - Spinnaker
    - Keptn


## Meeting 25th January

Participants:

- Andrea Frittoli, IBM, UTC
- Mauricio Salatino, VMWare, UTC
- Erik Sternerson, doWhile
- Mattias Linnér, Ericsson
- Ishan Khare

Agenda:

- [CDEvents Marketing / PR Plan](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit#)
- CDEvents Roadmap [PR](https://github.com/cdevents/spec/pull/9)
- Use Cases:
    - Proposed for PoC: https://hackmd.io/ZCS2KYKZTpKBqhU9PMuCew
    - Knative: https://github.com/cdfoundation/sig-events/issues/107#issuecomment-1020995310 
    - (Andrea) KubeCon Talk submission focuses on the "Metrics" use case, so I would focus on building the spec around that area
        - Even if the talk is not accepted, it will make a good PoC - and there's probably a good overlap with the Knative use case too
    - Helping consuming services/tools understand what events to consume, actually consuming these events, and declaring what events it consumes, is something we should consider for the CDEvents project.
        - Related Interop PR https://github.com/cdfoundation/sig-interoperability/pull/81/files#
- Design discussions
    - https://github.com/cdevents/spec/issues/10
        - Erik to add note on sending type of cause event
        - Erik to sketch a simple picture of in-pipeline fan-out - fan-in using shared context.
    - https://github.com/cdevents/spec/issues/11

- CDEvents v0.1 planning: https://github.com/orgs/cdevents/projects/1/views/1

## Meeting 11th January



Participants:
- Mattias Linnér, Ericsson
- Steve Taylor, DeployHub/Ortelius
- Andrea Frittoli, IBM
- Emil Bäckmark, Ericsson


- FOSDEM Presentation
    - CI/CD Devroom - from developers, to developers
    - Key Takeways that people should get from the talk?
        - Interoperability
        - We want their INPUT
    - What does your CI/CD world look like?
        - What tools do you use together?
            - Perhaps a form with a pre-populated list of tools
        - Shall we get contact details / get in touch during the conference
        - Ideally we'd like a picture of a pipeline
            - Come give us a demo of your pipeline
        - Shortlink that we can re-point?
    - Key points
        - What do we do
        - What we need from you
        - How to contribute
    - Things to prepare
        - Update the https://github.com/cdevents/.github/profile/README.md to include to links to https://github.com/cdfoundation/sig-events#communication
    - Two use cases
        - Interop between tools talking to each other
        - Interop between tools by collecting data consistently across tools (metrics)

- CdCon 2022 CFP - https://cd.foundation/event/cdcon-2022/
    - https://events.linuxfoundation.org/cdcon/program/cfp/#topics-track-focus

- Logo:
    - https://github.com/cdfoundation/foundation/issues/348
    - Ideas for stacked versions:
    - ![](https://i.imgur.com/bG6Mker.png)
    - ![](https://i.imgur.com/iO1FTVw.png)
    - ![](https://i.imgur.com/cgpNf66.png)
    - ![](https://i.imgur.com/oRe1OHJ.png)
    - ![](https://i.imgur.com/B56iRg5.png)
    - Emil will update the issue with some of these proposals. Maybe also propose a "draft"/"preliminary" version to the landscape so we get that in? And post on Slack to SIG Events on the proposal





***Previous meeting notes can be found on [GitHub](https://github.com/cdfoundation/sig-events/tree/main/docs)***