# Vocabulary Discussion Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA/badge)](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)

This document contains the notes from the of the Events SIG meetings focused on [vocabulary discussion](https://hackmd.io/lBlDCrL7TvmtNOjxdopJ5g).

## Next

- Issue [Consider taking advantage of this list of task/step types when qualifying event types](https://github.com/cdevents/spec/issues/18)


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