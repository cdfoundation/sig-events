---
tags: CDEvents
---

# CDEvents / Vocabulary Discussion Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA/badge)](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)

This document contains the notes from the of the Events SIG meetings focused on [vocabulary discussion](https://hackmd.io/lBlDCrL7TvmtNOjxdopJ5g).


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