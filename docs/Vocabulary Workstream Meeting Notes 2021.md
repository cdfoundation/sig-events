# Vocabulary Discussion Meeting Notes (2021)



## Meeting 14th December
* KubeCon CfP ends this week
* Quotes for the CD Events launch - Kara will provide a link to a doc for it
* Use cases should be evolved - https://hackmd.io/ZCS2KYKZTpKBqhU9PMuCew
* cdCon 2022
    * CfP will be published early January
    * Tracks are being set right now. Tracks could change later depending on what submissions come in
    * The con will be hybrid. Not yet set if presenters are requested to be in person or if remote is ok

Participants:
- Emil Bäckmark, Ericsson
- Mattias Linnér, Ericsson
- Kara de la Marck, CDF



## Meeting 30th November
* FYI: FOSDEM Call for participation opened today? https://fosdem.org/2022/news/2021-11-02-devroom-cfp/
    * Not open yet - we should keep an eye on it
    * Erik and Andrea are interested in presenting
    * Emil happy to review at least
    * We should have a chat on slack once the CfP is open
    * Timing is a good fit when we want to have more publicity on CDEvents
    * 
    * https://fosdem.org/2022/schedule/track/continuous_integration_and_continuous_deployment/
    * 
* PRs to the CDEvents/spec repo:
    * [Code of Conduct](https://github.com/cdevents/spec/pull/2): Merged
    * [MarkDown Linter](https://github.com/cdevents/spec/pull/3): Merged
    * [Bootstrap Governance](https://github.com/cdevents/spec/pull/4):
        * We need to give a rough timeline for the bootstrap governance work
        * [OpenTelemetry Charter](https://github.com/open-telemetry/community/blob/main/governance-charter.md) for reference
        * We need to allow a certain amount of time for folks to build up to the requirements that we set to be part of the governance / owners team
        * 
    * [Import Draft Spec](https://github.com/cdevents/spec/pull/5):
        * Do we want to merge this one or change this first?
        * Let's mark it clearly that this is not yet ready for implementation
        * Update the README in this PR
        * 
    * [Primer and Binding](https://github.com/cdevents/spec/pull/6):
        * Someone would like to work on the CloudEvents binding?
        * The binding is the input document to update the SDK
        * ...
* Proof of Concept
    * Do we have new use cases?
        * If not, how do we get them?
        * If yes, we could start focussing part of the meeting on this
        * Four keys
            * Remediation events are not in the spec yet
        * Let's keep them documented here: [List of Use Cases](https://hackmd.io/ZCS2KYKZTpKBqhU9PMuCew)
        * It could eventually be part of the roadmap/primer in GitHub

* Versions of the spec
    * Clarify that the current version is draft and shall not be implemented yet
    * Need to define roadmap of what we need to have for a first version
    * Need to write a version policy document

* Discussions Review
    * [GitHub Discussions](https://github.com/cdfoundation/sig-events/discussions)

    * Environments
        * https://github.com/cdfoundation/sig-events/discussions/28
        * Different kind of environments, ephemeral or long running, different scale too
        * We should look at tools that make use of environments and their model
        * It might not be needed for the 4 keys PoC
        * Environment referred to by name in the initial PoC
        * 
    * 

Participants:
- Mattias Linnér, Ericsson
- Andrea Frittoli, IBM
- Erik Sternerson, doWhile
- Emil Bäckmark, Ericsson
- Alois Reitbauer, Dynatrace

## Meeting 16th November

* Review https://github.com/cdfoundation/sig-events/discussions/22
    * Idea: Add TaskRunQueued, and let TaskRuns be parents of other TaskRuns. (Seems to collide with the Generic Activity Events discussed earlier)
    * PipelineRunFinished is an issue for non-orchestrated pipelines, as there is may not be any authority able to say that a pipeline run is finished.


* [LF/CDF identifies CD Events as a focus area.](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g?view#Meeting-November-8th) Idea is to reach out more broadly Jan/Feb 2021. What do we want to have in place before this?
    * Vision: What makes events successful are concrete use cases connecting tools together.
        * Current situation is that most orgs have a mix of tools in their end-to-end continuous delivery.
        * If we can get 3-or-4 good use cases we have a good base for getting more attention
            * We should probably have use cases that don't immediately duplicate what Argo and Keptn do.
            * E.g. Docker build (Jenkins) -> Signature (Jenkins) -> Deploy (Argo) -> Approval gate. (Keptn)
        * "Give us your pipeline so we can se what's going on"
            * We probably have quite a wide set of pipeline categories already in our own workgroup, but we have opportunities to reach out as well.
        * Use cases / pipeline types drive the spec work.
            * Primarily diagram it out so we can present ideas and the roadmap.
        * PoCs that exist, and possibly some new ones, can be used to provide "some proof".
    * We should also make sure to make progress on the protocol, regardless of when we find use cases.
        * Get the basics in place so we can show something.
    * Prepare for objections
        * The main objections against events are:
            * How do we get an overview of the whole workflow when everything is event-based?
            * How do we manage security in a distributed system?
        * A common protocol probably helps resolve both these issues.
        * We need a stronger/more clear story around the above objections ready.
    * Is there a good CloudEvents visualization solution available that we can use for demos?
        * If we have links, it would be nice for the solution to go fetch linked entities
        * Can Keptn or Argo provide anything here? (they don't speak CloudEvents today, but maybe we can copy/rejigger their solution to use our events)
        * https://github.com/visjs/vis-network could perhaps be used
        * https://www.d3-graph-gallery.com/ is also a contender
    * Move repos to new GitHub org.
        * Action: Look at Steve's PR before next week
        * https://github.com/cdevents/cdf-toc/pull/1
    * 

Participants:
- (Add yourself)
- Steve Taylor (DeployHub)
- Erik Sternerson (doWhile)
- Mattias Linnér (Ericsson)
- Andrea Frittoli (IBM)
- Emil Bäckmark (Ericsson)

## Meeting 2nd November

* Created https://twitter.com/_cdevents (placeholder)
* ...
* Review [Vocabulary Actions & Discussions](https://github.com/cdfoundation/sig-events/wiki/Vocabulary-Actions-&-Discussions)
    * #15 Generic vs. concrete activity events? 
        * Proposal: Generic plus concrete events
        * Generic events act as a "base class" for activity events
        * Generic events can also be used for events related to activites that do not have an appropriate subclass.
        * Concrete events provide for stronger alignment, e.g. "if you want to send events for a build, here is what you send".
        * Concrete events extend generic events with additional parameters.
        * Allowing or generic events may lead to platforms prefering generic events over concrete events as they don't feel that the concrete events are "a perfect fit", which would make us miss out on alignment opportunities.
            * "If you want the consumer to understand that this is a build, you need to use the build activity event, not a generic activity event".
        * 

Participants:
- Andrea Frittoli (IBM)
- Mattias Linnér (Ericsson)
- Oleg Nenashev (Jenkins/Keptn/Dynatrace)
- Emil Bäckmark (Ericsson)
- Erik Sternerson (doWhile)



## Meeting 19th October

* Decision based on discussion from SIG:
    * Vocabulary meetings should focus on spec content and PoCs, not on "project and admin" stuff.

* https://github.com/cdfoundation/sig-events/pull/87
    * Inside data object vs. top-level attributes
    * Existing extensions https://github.com/cloudevents/spec/blob/master/documented-extensions.md
    * https://github.com/cloudevents/spec/blob/v1.0.1/primer.md#cloudevent-attribute-extensions seems to indicate that top level attributes should be few and far between, and should have support from several "voting organizations" of the relevant CNCF working group.
    * We should set up a vocabulary that covers non-cloud-native scenarios as well, for those scenarios CloudEvents may not be optimal.
    * Current PoC is defined more as an extension, should be re-written to use the data attribute instead.
    * Proposal from work group: 
      > The Events SIG defines a specification ("CD Events") that provides a set of JSON object schemas (one for each event type, covering mandatory and optional attributes etc.)
      >
      > When used with CloudEvents, the JSON schema is passed via the `dataschema` attribute and a corresponding JSON object is sent through the `data` attribute.
      > 
      > Existing CloudEvents extension attributes (e.g. `partitionkey` from the [Partitioning](https://github.com/cloudevents/spec/blob/v1.0.1/extensions/partitioning.md) extension) should be considered before defining own extensions. If no appropriate extension attributes exists, the goal should be to make an official CloudEvents extension that can be included in the CloudEvents spec.

    * Next steps:
        * Protocol Format
            * Document how to vocabulary draft maps to a JSON format
                * JSONSchema (used in Eiffel)
                * OpenAPI
                * 
            * Recreate the SDK based on the defined schema
                * Investigate whether it's possible to generate the code
            * Re-run the PoC with the updated SDK
        * Events links, context, extensions
            * We need dedicated discussions

    * How do we take decisions?
        * Previous discussion: https://github.com/cdfoundation/sig-events/discussions/40
        * We should document decisions in the spec / documentation
        * This groups should be empowered to take decisions on the protocol spec / vocabulary
        * Review cycles are required after decisions
        * Potential process:
            * Large features on the spec require a larger review audience
            * We can work in a doc/pr/hackmd until a feature is ready for review from a larger audience
            * Small changes are done in PRs and approved by the small group
            * 

Participants:
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Emil Bäckmark (Ericsson)
- Andrea Frittoli (IBM)
- Oleg Nenashev (Jenkins)



## Meeting 5th October

* Protocol Name "CDEvents"
    * DNS name still tbd - cd.events / events.cd ?
        * events.cd might be available -> types would be cd.events.*
    * Website
        * Hugo and docsy works well for Ortelius and Tekton
    * Changing events name depends on DNS name
    * New GitHub org https://github.com/cdevents
    * Logo design: [creative brief](https://docs.google.com/document/d/1wJySvZFPetKTXEa_VD_gbUyDVnV6O272H-n__0FoG2E/edit) 
    
* Relocation to `cdevents` org:
    * Hosting: GitHub or GitLab or else?
        * GitHub makes sense as most CDF projects are there
    * Repos:
        * `spec`: current `vocabulary-draft`
            * Same format as in https://github.com/cloudevents/spec? 
            * `community.md` in `spec` repo (or `community` repo?)
            * we could have experimental extensions
            * 
        * `sdk-XYZ`: one repo per SDK
        * eventually one repo for the website
            * docsy
        * License
            * Let's use Apache2 everywhere 
    * What about PoC and related code?
        * 1 repo per PoC or 1 repo for multiple PoCs?
            * we could have an `experimental` repo
        * Shall we host adapater code or try to host in target orgs?
            * github.com/salaboy/cdf-events-keptn-adapter
            * github.com/salaboy/keptn-cdf-translator
            * github.com/tektoncd/experimental/cloudevents
            * We could use the `experimental` repo to host experimental code
            * Having a repo makes it easier to clone everything at once
            * It feels like a catalog of adapters
                * It makes sense to have all adapters discoverable in one place
        * CloudEvents have a [primer](https://github.com/cloudevents/spec/blob/v1.0.1/primer.md) document in the `spec` repo
            * 
    
    * Add people to the org:
        * sign-in here to join
        * 

* Jenkins + CDEvents + Tekton? PoC
    * What kind of events does Jenkins send
        * Queue enter -> Pipeline Queued
        * Queue leave -> Pipeline Start
        * Job completed / failed / finalised -> Pipeline Finished
        * New Job created / deleted
            * Should we have thing kind events?
        * Node events
            * Needs to be clarified with Jenkins community
            * No perfect match with current cdevents
            * Maybe similar to standing up k8s as test infra

        * Plan for pipeline stage events
    * How does it match to existing cdevents event buckets?
        * Events from the core bucket can be sent by the plugin
        * More specific events could be sent by the job
        * Could be offloaded to the user:
            * Which extra events to be sent 

    * Use this meeting to sync on progress / questions

* We should add a recurring agenda item to review the [AP](https://github.com/cdfoundation/sig-events/wiki/Vocabulary-Actions-&-Discussions) on the wiki

* AP from the wiki
    * #15 Generic vs. concrete activity events?
        * Spec provides terminology for activity events
        * We could have a translation matrix 
            * cdevents activity names -> platform spefic names

* Spec format and SDK code generation
    * 

* From draft towards v0.1
    * We could start at version 12 :P 

* CloudEvents discovery / subscription
    * https://github.com/cloudevents/spec/blob/master/discovery.md
    * https://github.com/cloudevents/spec/blob/master/subscriptions-api.md
    * 

Participants:
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Andrea Frittoli (IBM)
- Shruti Chaturvedi (GSoC CloudEvents plugin for Jenkins)

## Meeting 21st September
* Discussion about how to generate code from a specification
    * https://github.com/michaelawyu/cloudevents-generator
        * Downside to this is that it doesn't use any common know format for building the events but own "home coked" solution
        * Looks like it can specify types but it would be good if you could also generete doc strings
    * Can we use https://mustache.github.io/ to generate the code for us? (Github https://github.com/mustache)
    * Could we model the events in the same way as SPDX https://github.com/spdx/spdx-spec/blob/development/v2.2.1/model/SPDX-UML-Class-Diagram.jpg
    * What are the "best practice" ways of generating API code based on schema for CloudEvents (or related)
    * There are libraries in the cloudevents go-sdk to help you test e.g. https://pkg.go.dev/github.com/cloudevents/sdk-go/v2@v2.5.0/test
* Summary: We would like to have a schema so that other people that want to generate a sdk in their language can do with ease

Participants:
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)

## Meeting 7th September

- Discussion about routing:
    - CloudEvents include a subject <https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#subject>
    - We could re-use the subject to carry routing / filtering information
    - We could analyse existing events and check what info could be mapped to subject
- Context
    - In the PoC we needed Keptn specific code in the sender and receiver to store / extract the keptn context
    - We could have a common facility to transport context and app specific data
- Protocol Stack
    - We follow OSI stack approach 
    - It would be good to draw a picture that represents our protocol stack (Mattias)
    - We can use the stack as a framework to decide where to store information in the protocol
- Unit tests for the SDK
    - Code duplication in the SDK
    - We could use code generation
    - https://github.com/cdfoundation/sig-events/pull/92
- (A) Mattias to create GitHub Issues to track missing docs and code duplication in the SDK
    - https://github.com/cdfoundation/sig-events/issues/94
    - https://github.com/cdfoundation/sig-events/issues/95


Participants:
- Mattias Linnér (Ericsson)
- Andrea Frittoli (IBM)
- Erik Sternerson (doWhile)

## Meeting 24th August

- Continued discussion on events from artifact repositories (Docker, PyPi, Nexus etc.)
    - How to find out that event needs to be sent
        - Polling service as a fallback solutions, can be used for most repo types
        - Plugin to send events (e.g. Eiffel Nexus plugin)
        - Webhooks
    - How to transmit events
        - Webhooks is a common solution in the Git world.
        - In corporate settings, a managed message bus is common.
        - How to deal with this in the public open source world?
        - Keptn uses a message bus.
    - "Webhook routing"?
        - Repo owner sets up webhooks that fires on interesting events.
        - Webhook sends to "eventrouter.cd.foundation" (or some name, but hosted by CDF)
        - Anyone can set up a listener for this webhook that in turn triggers additional webhook __or__ (event better) sends CD-style cloudevents to whoever wants to listen.
        - This would remove the burden from the repo owner to add new listeners, who could instead register themeselves in "eventrouter.cd.foundation"
- TOC update
    - Events group want to become a project under CDF
    - TOC response: Become a project or a standard
        - Standard is a Linux Foundation group type
        - We should be a bit careful about becoming a formal standard.
        - Does the standardization org come with frameworks we can use for support? Or perhaps frameworks that will cause issues for us?
        - Meeting participants leaning somewhat towards the "project" group type.
        - CloudEvents is a project group, which hosts a specification. Not a standard.
- Spec work
    - Getting the CD events spec into at least a "v1 alpha" that can form at least part of the spec that is unlikely to change would be beneficial to projects e.g. Keptn who want to start using the spec without having to re-write a lot of code later on because the spec is fluid.
- PoC
    - Mattias has checked the PoC a bit, with some focus on the prerequisites for getting things running.
    - If we want the PoC to be an introduction to people looking into our spec, we should look into making the "getting started" procedure more well documented and possibly more automated.
    - Can something be done with like Docker Compose? Kind? Helm Chart? To create something more like a one-click-setup on a local machine or VM.
        - Erik (not a committment) wants to learn more in this area
- Event layout and structure
    - We looked at https://github.com/cdfoundation/sig-events/pull/87/files?short_path=79997d9#diff-79997d96fcc0fe274554038a7dcff7fbb000640508b50a77c5ecdd3eaa9ed0a7
    - Current structure is quite flat, with most attributes on top level
    - CloudEvents spec recommends that attributes that are __not__ used for routing and filtering events should go under `data`.
    - For instance, for ServiceDeployed, we have three atrributes: env ID, name and version. Where do we draw the line between what is used for filtering/routing and what is not.
        - We could probably come up with a scenario where _someone_ will want to filter/route on a specific attribute.
        - Probably parameters that change for every release may be less likely to be used for filtering/routing.
        - Some clients may not need to understand the content of the data block, which should primarily for parameters that certain actions or activities need
        - Comment from Alois: Everything generic for for instance all services should be on top level. Anything specific to one type of service or one instance should be in the `data` segment. E.g. Keptn would not want to care about the content of the `data` attribute.
        - Comment from Mattias: Putting a lot of attributes on top level would make our CloudEvents extension quite large (rather than making the format of the `data` content large).
            - c.f. https://github.com/cloudevents/spec/blob/master/documented-extensions.md
        - Being forced to listen to events that you don't want to listen to just because there isn't enough filtering fields is a scalability issue.
            - Eiffel has a "domain" concept to help solve the scalability issue, where events can stay within one domain and/or be published to other domains.
            - https://eiffel-community.github.io/eiffel-sepia/rabbitmq-message-broker.html
            - "It SHOULD use a routing key on the form eiffel.  `<family>.<type>.<tag>.<domainid>`, where: " ... "`domainid` is the non-empty string representing the domain the event applies to. It corresponds to the `meta.source.domainId` member of an Eiffel event."
        - Not being able to restrict certain clients to certain events just because there isn't enough filtering fields is a security issue.
- How do write the cd-events SDK
    - Design that reduces duplication and boiler-plate would be nice. What can be done?
    - Design guidelines for sig-events code?
    - Mattias has a PR up that should be looked at 
        - Erik will have a look
- Bring to SIG-meeting: We should get more implementers, more big players involved. Anyone with contacts into any project related to events can help bring them into the SIG and the workgroup.

Participants:
 - Erik Sternerson (doWhile)
 - Mattias Linnér (Ericsson)
 - Steve Taylor (DeployHub)
 - Alois Reitbauer (Dynatrace)

## Meeting 10th August

- Discuss https://github.com/cdfoundation/sig-events/discussions/56 

_No further notes_

## Meeting 27th July

Topics:
- Wiki Page
    - https://github.com/cdfoundation/sig-events/wiki 
    - No concurrent editing possible
    - One person to update the Wiki
    - We could review open actions in the main SIG meeting too
    - (A) Erik to make the current wiki up-to-date before the next SIG
- PR with samples 

- Event-specific data as attributes vs. in data
    - https://github.com/cdfoundation/sig-events/discussions/60#discussioncomment-1026664
    - Events should not be tied to CloudEvents as their only carrier
    - What format do we want to move forward with in the PR?
    - Meeting decision: We should have as much as possible under `data`

- Format of keys in events
    - CloudEvents states "all lowercase, no spaces" for attributes
    - For keys/attributes under `data`, we should decide on a format
        - Common formats seem to include camelCase and dash-case
        - See https://github.com/cloudevents/spec/blob/v1.0.1/primer.md#existing-event-formats for examples of other event formats
        - (A) Erik: Start a discussion on aligning around a format for keys/attributes under `data`

- Zipkin and distibuted tracing
    - (A) Andrea to add more context about zipkin to https://github.com/cdfoundation/sig-events/pull/86 

Participants:
- Mattias Linnér (Ericsson)
- Andrea Frittoli (IBM)
- Erik Sternerson (doWhile)

## Meeting 13th July

Topics:
- Long term home for the spec
    - Specification repo
    - Own opensource project
    - Dedicated governace, independent from the SIG
    - Inspiration
        - OpenTelemetry
        - CloudEvents
            - CloudEvents
            - Spec
            - SDKs
        - Eiffel
            - Spec and other repos
            - Eiffel community repo in the eiffel org
    - CDF SIG Events mission beyond the protocol?
        - Presentations
        - Education
    - Name discussion ongoing
        - We'd like the name to be indepentent from SIG/CDF
        - Form: https://docs.google.com/forms/d/1CVPooDG16B6JaBqQysH7V6UehYcZcSw_PM1HrtxwNns/edit 
        - Tentative deadline Oct 1st
- PoC Presentation
    - Jenkins is working on CloudEvents
        - It could be added to the demo
        - We could also build different versions of the demo
    - Convert to OTLP OpenTelemetry Line Protocol
        - Then can use Jeager to visualise the overall flow
        - Alois to connect with OTLP experience engineers @ Dynatrace
    - We could add Spinnaker to the PoC
        - (A) Andrea to reach out to Spinnaker folks

Participants: 
- Mattias Linnér (Ericsson)
- Andrea Frittoli (IBM)
- Alois Reitbauer (Dynatrace)
- Erik Sternerson (doWhile)


## Meeting 15th June

Topics:
- Progress on Go Library github.com/cdfoundation/sig-events/cde/sdk/go 
- Progress on PoCs
    - Keptn Translator 
    - Tekton Controller
    - Jenkins Plugin (Optional)

PoC Use Case: 
- Two tekton pipelines (Build & Deploy) that emits CD events
    - The first tekton pipeline emits a CD Event: "New Artifact Event"
- A Ketpn Translation component that does two things:
    - Exposes an endpoint to receive CD Events, translate the events to Keptn and forwards them to keptn /events API
    - A Keptn Service handle the Sequence Triggered event, translate it to a CD event to trigger a Tekton Pipeline
- The Tekton Controller accept a CD Event "Pipeline Queued" or similar that comes from Ketpn Sequence. This pipeline can emit a CD Event "New Service Deployed" and the translator can forward that to the ketpn sequence to see if we can show it in the keptn User Interface. 

Notes: 
- K3s can keptn and Tekton
- Keptn needs context / event ID back on deployment completion
- Put all the YAML file under the PoC folder
- One extra repo for the translation layer CDE/Keptn
- Keptn Service receives events (dedicated repo)
- Andrea to setup the Tekton pipelines and triggers

Participants:
- Mauricio Salatino (LearnK8s)
- Andrea Frittoli (IBM)
- Erik Sternerson (doWhile)
- Mattias Linnér (Ericsson)
- Jürgen Etzlstorfer (Dynatrace / Keptn)
- Emil Bäckmark (Ericsson)
- Johannes Bräuer (Dynatrace / Keptn)

## Meeting 1st June

Topics:
- PoC: 
    - Library PR (https://github.com/cdfoundation/sig-events/pull/55)
        - When to merge this PR? 
    - Tekton Use Case / Controller (https://github.com/tektoncd/community/issues/435)
    - Keptn Use Case
        - Actions: Translation Component CD Events to Keptn Events (bidirectional)
        - https://github.com/waveywaves/cdf-cloudevents-controller
    - Jenkins GSoC student 

Use Case: 
- Build Pipeline (Tekton)
- New Artifact Event (new docker image),  (Trigger keptn sequence) Keptn Decides which next:
    - Deploy Pipeline (Tekton)
    - New Artifact Event (new service running)
        - Which Environment should I deploy this (Keptn)
        - Quality gates next




Participants: 
- Mauricio Salatino (LearnK8s)
- Vibhav Bobade (Red Hat)
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Jürgen Etzlstorfer (Dynatrace)


## Meeting May 18 

- Buckets (or phases) of events [Discussion](https://github.com/cdfoundation/sig-events/discussions/44)
    - Pull Requests submitted
- Creating a small library for mapping to CloudEvents and share in a directory in sig-events
- Check if keptn is lightweight enough to run in KIND
- Tekton events for reference: https://github.com/tektoncd/pipeline/tree/main/pkg/reconciler/events/cloudevent
- We'll need a translation layer in front of the change events (github, gitlab, bitbucket, gerrit)

Participants:
- Mauricio Salatino (Learnk8s)
- Mattias Linnér (Ericsson)
- Emil Bäckmark (Ericsson)

Actions:
- 

## Meeting April 20

Went through a few open GitHub discussion and added comments.

- Buckets (or phases) of events [Discussion](https://github.com/cdfoundation/sig-events/discussions/44)
- GitOps specific events - [Discussion](https://github.com/cdfoundation/sig-events/discussions/30)
- 4keys specific events / mappings - [Discussion](https://github.com/cdfoundation/sig-events/discussions/42)
- What is a "change", granularity of events (PR, commit) - [Discussion](https://github.com/cdfoundation/sig-events/discussions/34)

Discussion on our short-term ambitions. Proposal: Have a few buckets ready, and a PoC or two,
by CDCon (2 months). We should think about how we want to demo things to make it appealing.

Discussion: Use the term "sequence" instead of "pipeline" in events.

Participants:

- Erik Sternerson (doWhile)
- Alois Reitbauer (Dynatrace)
- Mattias Linnér (Ericsson)
- Andrea Frittoli (IBM)
- Steve Taylor (DeployHub/Ortelius)
- Tracy Ragan (DeployHub/Ortlius)

Actions:

- [x] Andrea: Create PR for removing bucket names from event types (done by Mauricio)
- [ ] Tracy: Bring GitOps events up with GitOps working group in CNCF
- [x] Mattias: Start a GitHub discussion on Incident events - [Discussion](https://github.com/cdfoundation/sig-events/discussions/49)
- [ ] Tracy: Request people from GitLab and GitHub to join the vocabulary workgroup
- [ ] Alois: On next Vocab WG meeting, resent ideas to merge [Keptn spec](https://github.com/keptn/spec) with CD Events spec
- [ ] Andrea: Add some thoughts to the Change discussion
- [ ] Erik: Start a GitHub discussion on what to aim for us being able to present at CDCon
- [ ] Erik: Start a GitHub discussion on how to introduce to what an event-based system would look like and what the benefits are.
- [ ] Tracy: Request people from Jenkins to join the vocabulary workgroup
- [ ] Erik: Dig up Jenkins Summer-of-Code Events plugin contacts and see if they are interested in being involved

## Meeting April 6

Check-in, mostly. No new actions?

Participants:
- Mauricio Salatino (Camunda/LearnK8s)
- Emil Bäckmark (Ericsson)
- Erik Sternerson (doWhile)
- Alois Reitbauer (Dynatrace)
- Gopinath Rebala


## Meeting March 23rd
Meeting time [in your timezone](https://time.is/4pm_24_March_2021_in_UTC). You're welcome to join!

Participants:
- Mauricio Salatino (Camunda/LearnK8s)
- Steve Taylor (DeployHub/Ortelius)
- Tracy Ragan (DeployHub/Ortlius)
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Alois Reitbauer (Dynatrace)
- Emil Bäckmark (Ericsson)
- Andrea Frittoli (IBM)

### Agenda and Notes

Welcome Alois to the SIG!

During the meeting today we reviewed and merged open PRs in
the [SIG events](https://github.com/cdfoundation/sig-events) repo:

- [#5](https://github.com/cdfoundation/sig-events/pull/5) Introduction to Vocabulary PR 
- [#7](https://github.com/cdfoundation/sig-events/pull/7) Source Code Version Control Events section
- [#8](https://github.com/cdfoundation/sig-events/pull/8) Continuous Integration pipeline events section
- [#9](https://github.com/cdfoundation/sig-events/pull/9) Continuous Deployment pipelines events section

The group agrees on merging the PRs as they are, and track open discussion
points from pull requests, plus those identified today in the meeting, as
dedicated GitHub discussions, each with a dedicated owner.

GitHub discussions (tick once created):

- [x] (Andrea) Buckets (or phases) of events [Discussion](https://github.com/cdfoundation/sig-events/discussions/44)
- [x] (Tracy) GitOps specific events - [Discussion](https://github.com/cdfoundation/sig-events/discussions/30)
- [x] (Mattias) 4keys specific events / mappings - [Discussion](https://github.com/cdfoundation/sig-events/discussions/42)
- [x] (Erik) What is a "change", granularity of events (PR, commit) - [Discussion](https://github.com/cdfoundation/sig-events/discussions/34)
- [x] (Steve) DevSecOps specific events and bucket - [Discussion](https://github.com/cdfoundation/sig-events/discussions/26)
- [ ] (Tracy) Continuous Tests as its own bucket
- [ ] (Alois) Deployments and versions (e.g. blue/green)
- [x] (Steve) Environments - [Discussion](https://github.com/cdfoundation/sig-events/discussions/28)
- [x] (Andrea) Lightweigth vs. heavyweigth events (artifacts, logs) - [Discussion](https://github.com/cdfoundation/sig-events/discussions/47)
- [x] (Erik) New [CD Bucket events](https://github.com/cdfoundation/sig-events/pull/23)
  - [Artifact / Component discussion](https://github.com/cdfoundation/sig-events/discussions/24)
  - [Product Compositions discussion](https://github.com/cdfoundation/sig-events/discussions/35)
  - [Confidence Labels discussion](https://github.com/cdfoundation/sig-events/discussions/37)
  - [Environment Allocation discussion](https://github.com/cdfoundation/sig-events/discussions/38)
- [ ] (Mauricio) Core / base / orchestration events - [Pull Request?](https://github.com/cdfoundation/sig-events/pull/31)
- [ ] (Alois) Status changed / heartbeat events

Action Points:

- [x] (Mauricio) PR to link bucket docs to introduction doc - [Pull Request](https://github.com/cdfoundation/sig-events/pull/32)
- [ ] (Steve) Add deployment start/finish events
- [x] (Andrea) Create issue about removing the bucket from the event type - [Issue](https://github.com/cdfoundation/sig-events/issues/45)