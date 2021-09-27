# Events SIG Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g/badge)](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g)

This SIG meets bi-weekly on Mondays at 3pm UTC (during summer time, in winter time it is 4pm UTC). See your timezone [here](https://time.is/3pm_in_UTC). The meeting can be accessed through [this zoom link](https://zoom.us/j/97660712600?pwd=Z3BqYTE5YzNsbEhmck16cjdZNEFIUT09). For older meetings please see our [playlist on YouTube](https://www.youtube.com/playlist?list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC)

The SIG was initiated as a [workstream under SIG Interoperability ]( https://github.com/cdfoundation/sig-interoperability/blob/master/docs/meetings_2020.md#May-28-2020) and its first meeting was held on June 8th 2020.

Meeting notes for the workstream are managed on HackMD [here](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g), and published to GitHub [here](https://github.com/cdfoundation/sig-events/blob/main/docs/meetings.md).


## Agenda Template

- Action Items From Last week
- CDF & SIG Events Updates
- Workstream Updates
    - Vocabulary
- Meetups / conferences
- Blog Posts / Links
- Presentations
- PR / Design Discussions


## Meeting September 27nd

Meeeting time in your timezone [here](https://time.is/3pm_27_September_2021_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Andrea Frittoli, IBM
- Steve Taylor, DeployHub
- \<add yourself\>

### Agenda
- Action Items From Last week
- Protocol Naming
    - [Results](https://docs.google.com/forms/d/e/1FAIpQLSdTYx_1jNio67CX1H3GQAy7gn7gikCo3wKCYS6cohCHGNSFfA/viewanalytics)
        - Decided: "CD Events" it is
    - When we agree on a name, what are the next steps?
        - Domain name proposal: events.cd
        - GitHub organization: "github.com/cdevents"
            - Registered by Emil just now: https://github.com/cdevents
        - Event type is reverse of DNS: cd.events would be nice
        - CD Foundation project proposal
        - Icon
            - Would be great if the icon would be possible to add to arbitrary "CD Events supporting" components/projects, like "Evie". A tone/note?
        - (A) Andrea to ask CDF if the can buy the domain for us
        - (A) Steve to bring up on TOC tomorrow about the name and how to propose a new project and if there are any artists available to draw an icon
        - (A) Andrea to register a Twitter handle related to CD Events
- CDF & SIG Events Updates
    - [Jenkins plugin with Cloudevents from GSoC](https://github.com/jenkinsci/cloudevents-plugin)
        - Plugin is published. Inspired by our work but not yet aligned fully.
        - It would be great to have PoC with this plugin and e.g. Tekton using the CD Events. A use case should be written for such a PoC first.
        - That PoC discussion can continue on the Vocabulary meeting. Shruti should be invited there.
    - [SPDX becomes a BOM standard](https://www.linuxfoundation.org/press-release/spdx-becomes-internationally-recognized-standard-for-software-bill-of-materials/)
        - Postpone this discussion for next meeting
        - (A) Emil to put it on next SIG meeting agenda
- Workstream Updates
    - [Vocabulary](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- Meetups / conferences
    - We have a booth at Kubecon. Not sure what is expected of us there. Tracy will check with T Miranda
- Blog Posts / Links
- Presentations
- PR / Design Discussions
  - [Open actions and discussions](https://github.com/cdfoundation/sig-events/wiki/Vocabulary-Actions-&-Discussions)
      - We should look at these items, where decisions are needed, during our SIG meetings.
      - (A) Emil to put it on next SIG meeting agenda

## Meeting September 13th

Meeeting time in your timezone [here](https://time.is/3pm_13_September_2021_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark (Ericsson)
- Andrea Frittoli (IBM)
- Steve Taylor (DeployHub)
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- \<add yourself\>

### Agenda
- Action Items From Last week
- Presentations
  - Argo Events, by Derek Wang. 15-20 minutes
- Protocol Naming
    - Requirements for a good name (10 min)
        - Be related to our topic, e.g. continuous delivery, supply chain, messages, events
        - It is pronounceable (either as a word or as an abbreviation of at most three letters) (e.g. Signal or CDF is fine, SPDX and Sgehgre is a but much)
        - It is available (i.e. it doesn't already mean something else in our domain, or something negative in any other domain)
        - NOT Bind us to any organization or group (e.g. Events SIG)
        - NOT Bind us to any underlying transport or protocol (e.g. CloudEvents)
    - Other comments:
        - Provide a nice "brand name"
        - Should represent the disruption we want to bring
        - It is spelt as it sounds (not super important when most communication is written, so it is not a strong wish)
        - We might want to align it to something that people already know
    - Name Structure
        - Descriptive name
            - Subject + "Events"
            - Subject could be either of: CD, (Software) Supply Chain, Lifecycle, DevOps
        - Evocative name
    - Name selection (10 min)
        - [Current proposals](https://docs.google.com/spreadsheets/d/14Q9X_330YNGK1Naq22TN0iA2q1mVSSXIusGysCGmcVU/edit?resourcekey#gid=1286431544)
        - Proposals that meet the requirements:
            - To be added...
    - We could have two names, one for the protocol, one for the project
        - CDE
          CD Events
          CDEvents
          ~~CDCE CD CloudEvents~~
          ~~CICD Cloudevents (Continuous Integration Continuous Delivery Cloudevents)~~ 
          events.cd ("Events for Continous Delivery)
          Semio/Semios (some startups already used that)
          ~~CSE (CDF SIG Events)~~
          ~~CDFE (Continious Delivery Foundation Events)~~
          ECD (Events for Continuous Delivery)
          ~~Open Lifecycle Orchestration~~
          LEAP (Lifecycle Events and Analytics Protocol)
          SPELL (Sandardized Protocol for Event Lifecycle)
          ~~CICDEP "CICDEP is Continuous Delivery Events Protocol"~~
          LifecycleEvents: what the events are about as opposed to a trend
             LiEve
          (A) Andrea to make a poll from the list
          (A) Andrea to check with CDF for branding check
        - https://github.com/cdfoundation/sig-events/discussions/70 
- SIG Events Updates
- Workstream Updates
    - Vocabulary
- Meetups / conferences
- Blog Posts / Links
- PR / Design Discussions


## Meeting August 31st

Meeeting time in your timezone [here](https://time.is/3pm_31_August_2021_in_UTC). You're welcome to join!

Participants:
- \<add yourself\>
- Emil Bäckmark (Ericsson)
- Steve Taylor (DeployHub)
- Tracy Ragan (DeployHub)

### Agenda
- Action Items From Last week
  - (A) Emil will prepare for the next SIG together with Andrea to try to set a name on the next SIG meeting. Next SIG meeting should be dedicated to the name discussion.
      - Andrea is still on PTO, so this is postponed until next meeting
- SIG Events Updates
    - [Protocol Naming discussion](https://github.com/cdfoundation/sig-events/discussions/70)
        - Tracy:
            - Protocol name should be "CD Events", for reasons of simplicity and clarity
            - Listener tool could be called "Kinetic"?
            - The event robot "EV" (pronounced "Evie") listens to events. Simple, easy to describe, easy to market, ie logo. Approachable.  Trying to simpifiy complex topic.
    - [CDF Podcast about Shipwright](https://podcasts.google.com/feed/aHR0cHM6Ly9mZWVkcy5idXp6c3Byb3V0LmNvbS8xMDA4Njk3LnJzcw/episode/QnV6enNwcm91dC04OTY5NTQz?hl=sv&ved=2ahUKEwjxq7WHps7yAhUUmWoFHScBB_MQieUEegQIDBAL&ep=6) mentions the need for artifact events
- Workstream Updates
    - [Vocabulary workstream MoM](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA?view)
        - WebHooks vs Event sending from upstream services, such as SCMs and Artifact Repositories
        - Event protocol to become a standard or a project group? Compare with Cloudevents, SPDX, etc
            - Steve: Standard vs specification vs project. Cloudevents is probably a specification
        - Bring to SIG-meeting: We should get more implementers, more big players involved. Anyone with contacts into any project related to events can help bring them into the SIG and the workgroup.
            - Emil: We've invited Shipwright to the SIG
- Meetups / conferences
  - [KubeCon NA 2021](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/)
      - We have a slot on KubeCon for SIG Events - Events SIG CDF Virtual Booth @ KubeCon 2021, on October 13th. 12pm Eastern, 6pm CET.
      - Andrea will give a [presentation about SIG Events](https://kccncna2021.sched.com/event/lV1H/sig-events-using-cloudevents-to-create-an-interoperable-cicd-ecosystem-andrea-frittoli-ibm?iframe=no&w=100%&sidebar=yes&bg=no)
- Blog Posts / Links
    - We should write a blog post as soon as we've settled on a name for the protocol (and listener)
- Presentations
  - Argo Events presentation aimed for September 13th, by Derek Wang. 15-20 minutes
- PR / Design Discussions


## Meeting August 16th

Meeeting time in your timezone [here](https://time.is/3pm_16_August_2021_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Kara de la Marck, CloudBees
- Mattias Linnér, Ericsson
- Tracy Ragan (DeployHub)
- Steve Taylor (DeployHub)
- Shruti Chaturvedi (GSoC CloudEvents plugin for Jenkins)

### Agenda
- Action Items From Last week
- SIG Events Updates
- Workstream Updates
    - Vocabulary
        - We need a name on the protocol! It should be reflected in a new repo where the protocol will be developed
        - Should the protocol repo be located under [cdfoundation](https://github.com/cdfoundation) or should it have its own Github organization?
        - The protocol name should then not conflict with an existing Github org...
        - [Current results](https://docs.google.com/spreadsheets/d/14Q9X_330YNGK1Naq22TN0iA2q1mVSSXIusGysCGmcVU/edit?resourcekey#gid=1286431544) from [poll](https://docs.google.com/forms/d/1CVPooDG16B6JaBqQysH7V6UehYcZcSw_PM1HrtxwNns/viewform?edit_requested=true)
        - (A) Emil will prepare for the next SIG together with Andrea to try to set a name on the next SIG meeting. Next SIG meeting should be dedicated to the name discussion.
        - Should the event protocol be its own project proposed to CDF?
- Meetups / conferences
    - Mauricio joins [Keptn user group](https://keptn.sh/community/meetings/#keptn-user-group) tomorrow to show the Tekton+Keptn interoperability PoC
    - We have a slot on KubeCon for SIG Events - Events SIG CDF Virtual Booth @ KubeCon 2021, on October 13th. 12pm Eastern, 6pm CET.
        - If you don't have the invite and whish to participate please ask someone that is in, for example Emil.
- Blog Posts / Links
- Presentations
    - Shruti presents [Cloudevents plugin for Jenkins](https://github.com/jenkinsci/cloudevents-plugin) (originally a Google summer of code project) 🎉
        - Support for CloudEvents, and for the CloudEvents plugin in particular, is part of [Jenkins Roadmap](https://www.jenkins.io/project/roadmap/)
- PR / Design Discussions
    - There are a number of PRs that should be reviewed, for example regarding unit tests.

## Meeting July 19th


Meeeting time in your timezone [here](https://time.is/3pm_21_June_2021_in_UTC). You're welcome to join!


Participants:
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Andrea Frittoli (IBM)(IBM)(IBM)
- Kara de la Marck (CloudBees)
- Shruti Chaturvedi (GSoC CloudEvents plugin for Jenkins 🎉)
- \<add me\>


### Agenda
- Action Items From Last week
    - (A) Erik to make a PR about "What is a change": https://github.com/cdfoundation/sig-events/discussions/34
        - Done: https://github.com/cdfoundation/sig-events/pull/85
    - (A) Andrea to make a PR about Lightweigth vs heavyweight: https://github.com/cdfoundation/sig-events/discussions/47
        - Done: https://github.com/cdfoundation/sig-events/pull/77 
- SIG Events Updates
    - none
- Workstream Updates
    - PoC Updates
    - Jenkins CloudEvents
        - Slack channel gsoc-2021-jenkins-cloudevents-plugin
        - Demo on 20/07: https://www.jenkins.io/blog/2021/07/16/gsoc-midterm-presentation/
        - Jenkins Cloud Native SIG YT channel posts some of our GSoC meetings on Jenkins CloudEvents plugin, see here: https://www.youtube.com/playlist?list=PLN7ajX_VdyaOFG9hTrswbO-ZK_n4B8CaG
    - Dedicate open source project for spec and artifacts
        - With own repo, governance
- Meetups / conferences
    - DevOpsWorld Presentation Mauricio / Andrea
- Blog Posts / Links
    - none
- Presentations
    - Jenkins CloudEvents on August 16th
- PR / Design Discussions
    - https://github.com/cdfoundation/sig-events/pull/
        - RedHat uses the "Code contribution" terminology
- Uploading of the recordings



## Meeting June 21st

Meeeting time in your timezone [here](https://time.is/3pm_21_June_2021_in_UTC). You're welcome to join!


Participants:
- Emil Bäckmark (Ericsson)
- Mauricio Salatino (LearnK8s)
- Mattias Linnér (Ericsson)
- Jürgen Etzlstorfer (Dynatrace, Keptn)
- Erik Sternerson (doWhile)
- Tracy Ragan (DeployHub, Ortelius)
- Andrea Frittoli (IBM)
- Derek Wang (Intuit, Argo Events)
- Steve Taylor (DeployHub/Ortelius)
- Gopinath Rebala (OpsMx)

### Agenda
- Action Items From Last week
    - (A) Erik to make a PR about "What is a change": https://github.com/cdfoundation/sig-events/discussions/34
    - (A) Andrea to make a PR about Lightweigth vs heavyweight: https://github.com/cdfoundation/sig-events/discussions/47
- SIG Events Updates
    - SIG meetings during the summer period
        - Emil is away July 5th to Aug 6th
        - Andrea is away 26/07 -> 26/08 (possibly)
        - The workstream should continue if enough people are available. The SIG could have fewer meetings during the summer and wait with needed decisions until most people are back.
        - Decided: SIG meeting July 19th and August 16th
            - (A) Emil will update the calendar invite
                - Emil: Done - Meeting Aug 2nd is cancelled
- Workstream Updates - Vocabulary
    - Event protocol name poll
        - Short-term we use the name "CD Events", but we should have a better name
        - [Name Poll](https://docs.google.com/forms/d/1CVPooDG16B6JaBqQysH7V6UehYcZcSw_PM1HrtxwNns/edit)
        - [Poll Answers so far](https://docs.google.com/spreadsheets/d/14Q9X_330YNGK1Naq22TN0iA2q1mVSSXIusGysCGmcVU/edit?resourcekey#gid=1286431544)
        - (A) Erik to start a discussion about the name and criteria that we want to consider for a name
        - Target Date: let's aim for Oct 1st for the name
    - PoC status
        - PoC hacking very much ongoing
        - In the BoF we will present the ongoing work, probably without showing a live run. We could show some code examples and the PoC readme/description
- Meetups / conferences
    - Event related presentations etc on cdCon
        - [Behind the Scenes of Keptn: Event-Driven Delivery & Ops Orchestration - Andreas Grabner, Dynatrace](https://cdcon2021.sched.com/event/ipny/behind-the-scenes-of-keptn-event-driven-delivery-ops-orchestration-andreas-grabner-dynatrace)
        - [Linking Events to Achieve Traceability and Data Aggregation in CI/CD Systems - Emil Bäckmark, Ericsson](https://cdcon2021.sched.com/event/iori/linking-events-to-achieve-traceability-and-data-aggregation-in-cicd-systems-emil-backmark-ericsson)
        - [Event-driven Distributed CI for 20000 Projects - Aleksandra Fedorova, Red Hat](https://cdcon2021.sched.com/event/iosg/event-driven-distributed-ci-for-20000-projects-aleksandra-fedorova-red-hat)
        - [Lightning Talk: Events of Screwdriver - Alan Dong & Jithin Emmanuel, Verizon Media](https://cdcon2021.sched.com/event/ioui/lightning-talk-events-of-screwdriver-alan-dong-jithin-emmanuel-verizon-media)
        - [BoF Session: Event Based CD: Conversations with the CDF Events SIG - Tracy Ragan, DeployHub; Emil Backmark, Ericsson; Andrea Frittoli, IBM & Erik Sternerson, doWhile](https://cdcon2021.sched.com/event/iv29/bof-session-event-based-cd-conversations-with-the-cdf-events-sig-tracy-ragan-deployhub-emil-backmark-ericsson-andrea-frittoli-ibm-erik-sternerson-dowhile)
            - Questions: https://docs.google.com/document/d/1LgEp0ZuhhNsvp7oW-M0Uzdpsyu0uai8-YoB6L4S42Q8/edit
            - Tracy to facilitate
        - [Putting Chaos into Continuous Delivery – Evaluate & Increase the Resilience of your Applications - Juergen Etzlstorfer, Dynatrace & Karthik Satchitanand, ChaosNative](https://cdcon2021.sched.com/event/iouH/putting-chaos-into-continuous-delivery-evaluate-increase-the-resilience-of-your-applications-juergen-etzlstorfer-dynatrace-karthik-satchitanand-chaosnative)
    - DevOpsWorld (Mauricio / Andrea) - Using CloudEvents to Create an Interoperable CI/CD Ecosystem - accepted
- Blog Posts / Links
    - Meeting recordings https://www.youtube.com/playlist?list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC 
- Presentations
    - Argo Events presentation aimed for August 30th, by Derek Wang. 15-20 minutes
- PR / Design Discussions
    - [Experimental buckets](https://github.com/cdfoundation/sig-events/discussions/68)



## Meeting June 7th

Meeeting time in your timezone [here](https://time.is/3pm_7_June_2021_in_UTC). You're welcome to join!


Participants:
- Emil Bäckmark (Ericsson)
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Jürgen Etzlstorfer (Dynatrace, Keptn)
- Johannes Bräuer (Dynatrace, Keptn)
- Steve Taylor (DeployHub/Ortelius)
- Andrea Frittoli (IBM)
- Gopinath Rebala (OpsMx)
- Tracy Ragan (DeployHub/Ortelius)
- Vibhav Bobade (Red Hat)
- \<add yourself\>

### Agenda
- Action Items From Last week
    - (A) Erik to make a PR about "What is a change": https://github.com/cdfoundation/sig-events/discussions/34
    - (A) Andrea to make a PR about Lightweigth vs heavyweight: https://github.com/cdfoundation/sig-events/discussions/47
    - ~~(A) Tracy to write a blog post to push for our BoF in cdCon ~~Done!
- SIG Events Updates
    - Event Protocol name poll
        - Most common (a few variations): CDEvents, CD Events, events.cd, CDE, ECD
        - CDCE CD CloudEvents
        - Semio(s)
        - CICD Cloudevents
        - CSE (CDF SIG Events)
        - CDFE (Continious Delivery Foundation Events)
        - LEAP (Lifecycle Events and Analytics Protocol)
        - SPELL (Sandardized Protocol for Event Lifecycle)
        - CICDEP "CICDEP is Continuous Delivery Events Protocol"
        - Full [Results](https://docs.google.com/spreadsheets/d/14Q9X_330YNGK1Naq22TN0iA2q1mVSSXIusGysCGmcVU/edit?usp=sharing)
    - Should have
        - Related to the topic
            - continuous delivery
            - supply chain
        - Andrea: nice brand name
        - Tracy: we should take time to get it right
        - Tracy: should represent the disruption we want to bring
    - Should not have
        - Emil: not related to the group / foundation
        - Emil: CloudEvents is not only about Cloud
        - Erik: CloudEvents itself should not be part of the name as we may used a different base protocol in future
    - Proposal 1 (Andrea)
        - Name: CDEvents
        - Domain: events.cd
        - Event types: cd.events.*
    - Pros
        - Similar name structure as cloud events
        - Not bound to a foundation or group name
        - events.cd domain is available
        - Intuitive and descriptive (at least with the acronym expanded)
        - github.com/cdevents is available
    - Cons
        - twitter/cdevents is taken
        - perhaps too generic? ContinousDeliverEvents would be more specific (but very long too)
    - Ideas
        - Lifecycle 
        - Infinity symbol
        - It is not sequential / linear
        - Tibetan Knot is not linear
        - Event Mesh
        - Continue on GitHub Discussion
        - Temporary Name for PoC can be cd.events
        - GNU = Gnu's Not Unix - recursive acronym?
        - YAML = YAML Ain't Markup Language
        - 
- Workstream Updates
    - Vocabulary
        - Good discussions on the PoC
        - We can dedicate the vocabulary event for now to the PoC
        - CDF Calendar: https://calendar.google.com/calendar/u/0/embed?src=linuxfoundation.org_mhf0kmgedn67ihni8r129avp24@group.calendar.google.com 
- Meetups / conferences
    - BoF Questions: https://docs.google.com/document/d/1LgEp0ZuhhNsvp7oW-M0Uzdpsyu0uai8-YoB6L4S42Q8/edit
        - There is an opportunity to present
        - We could prepare a presentation about the PoC
        - Erik started working on a diagram
        - Time conflict with other BOF (already reported by the interop SIG) - 11:15 EDT
        - Erik: What do we want from the BoF?
            - Discuss about the benefits of events based CI/CD
            - 
    - KubeCon Submission
        - Mauricio / Andrea - Using CloudEvents to Create an Interoperable CI/CD Ecosystem
        - Tracy - Events in CI/CD
- Presentations
    - Does anyone know anyone from Argo/Intuit that could present Argo Events (https://argoproj.github.io/argo-events/#what-is-argo-events)?
        - Juergen will ask Argo CD contacts
        - Tracy will ask Gale for contacts on CNCF side
    - PoC should be presented here (before cdCon, next meeting 21 June)
    - 
- Blog Posts / Links
    - BoF Blog Post: https://cd.foundation/blog/2021/05/27/meet-the-cdf-events-sig-at-cdcon-2021/
        - It's not possible to comment on the blog
        - Promoted via twitter
        - 
- PR / Design Discussions
    - Protocol PR: https://github.com/cdfoundation/sig-events/pull/55/files
        - Please review


## Meeting May 24th

Meeeting time in your timezone [here](https://time.is/3pm_24_May_2021_in_UTC). You're welcome to join!

Participants:
- Mattias Linnér (Ericsson)
- Emil Bäckmark (Ericsson)
- Mauricio Salatino (LearnK8s)
- Erik Sternerson (doWhile)
- Steve Taylor (Ortelius/DeployHub)
- Vibhav Bobade (Red Hat)
- Kara de la Marck (CloudBees)
- Gopinath Rebala (OpsMx)
- \<add yourself\>

### Agenda
- Action Items From Last week
- SIG Events Updates
    - PoC for protocol draft
        - Provide a simple diagram for the specific PoC - use https://whimsical.com/? @erkist
        - Goal: https://github.com/cdfoundation/sig-events/discussions/48
        - Library impl: https://github.com/cdfoundation/sig-events/pull/55
        - Spec: https://github.com/cdfoundation/sig-events/tree/main/vocabulary-draft
- Workstream Updates
    - Dictionary
        - GitHub Discussions
            - What is a change: https://github.com/cdfoundation/sig-events/discussions/34
                - (A) Erik to make a PR about this
            - Lightweigth vs heavyweight: https://github.com/cdfoundation/sig-events/discussions/47
                - (A) Andrea to make a PR about this
- Meetups / conferences
    - [cdCon BoF](https://docs.google.com/document/d/1LgEp0ZuhhNsvp7oW-M0Uzdpsyu0uai8-YoB6L4S42Q8/edit)
        - We can share screen on the BoF so we ar able to show a PoC diagram for example
    - Mauricio/Andrea has submitted a talk to KubeCon NA
        - Could be resubmitted to DevOps World as well
        - Mauricio will do the submission
- Blog Posts / Links
    - Write a blog post to attract people to the cdCon BoF session?
        - (A) Tracy to write a blog post to push for our BoF in cdCon
- Presentations
    - [CDF cloudevents controller](https://github.com/waveywaves/cloudevents-controller) hack and todo
- PR / Design Discussions
    - [Initial CloudEvents mapping and Go Library](https://github.com/cdfoundation/sig-events/pull/55)


## Meeting May 10th
**Note: The meeting is postponed one week due to bank holiday in UK**

Meeeting time in your timezone [here](https://time.is/3pm_10_May_2021_in_UTC). You're welcome to join!

Participants:

- Emil Bäckmark, Ericsson
- Steve Taylor, DeployHub
- Andrea Frittoli, IBM
- Mattias Linnér, Ericsson
- Mauricio Salatino, LearnK8s
- Vibhav Bobade, RedHat
- Kara de la Marck, CloudBees
- Tracy Ragan (DeployHub)
- \<add yourself\>

### Agenda

- AP From Last week
    - From last TOC meeting: We can use Apache2 for our contributions
    - We are going to use [Apache2 License](https://github.com/cdfoundation/sig-events/blob/main/LICENSE)
    - [Google doc for BoF](https://docs.google.com/document/d/1LgEp0ZuhhNsvp7oW-M0Uzdpsyu0uai8-YoB6L4S42Q8/edit)
- SIG Events Updates
    - Meeting frequency change
        - Who wants to have the invite from the CDF Calendar?
    - Which user group should we start focusing on?
        - Who has capacity to perform a PoC?
            - Mauricio Salatino
            - Erik Sternerson
            - Andrea Frittoli
            - Mattias Linnér
            - Vibhav Bobade
        - Create a small library to generate / parse the events
        - 0. Diagram to demostrate what we can build with shared events
            - Education material
        - 1. Integration between a couple of project
            - We could use a translation layer to begin with
            - Demonstrate that two projects can interop through common events
        - 2. Remove translation layer
        - 3. Stretch goal: 4-keys (collecting metric)
            - It requires running services
- Workstream Updates
  - Dictionary
      - No updates
- Meetups / conferences
- Blog Posts / Links
    - 
- Presentations
    - Eiffel (Mattias)
- PR / Design Discussions


## Meeting April 19th
Meeeting time in your timezone [here](https://time.is/3pm_19_April_2021_in_UTC). You're welcome to join!

Participants:
- Andrea Frittoli (IBM)
- Mattias Linnér (Ericsson)
- Emil Bäckmark (Ericsson)
- Tracy Ragan (DeployHub)
- \<add yourself\>

### Agenda and Notes

- AP From Last week
    - Licenses for Events SIG
    - From TOC:
        - Any OSI approved license is ok
        - Keptn, Tekton, Eiffel and Ortelius use Apache2
        - (AP - Andrea) Open a GitHub Discussion, and send a mail to the mailing list
        - (AP - Andrea/Emil) Bring the Apache2 proposal to the TOC
- SIG Events Updates
  - Tracy joined the CNCF App Delivery SIG
  - Steve to look into SDPX/Events SIG interop (Pedigree / Provenance)
  - [SPDX Presentation]( https://docs.google.com/presentation/d/1YNtXPRGuj9hri_rGid_E7LZ9OODcNASbdNmNOOOzFAk/edit#slide=id.gd25d926f37_0_136) at TOC
- Workstream Updates
  - Dictionary
      - No updates from the last meeting
      - Some AP open from the [previous meeting](https://github.com/cdfoundation/sig-events/blob/main/docs/vocabulary_meetings.md#meeting-march-23rd) - [HackMD Version](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- Meetups / conferences
    - [cdCon Events SIG BoF](https://cdcon2021.sched.com/event/iv29/bof-session-event-based-cd-conversations-with-the-cdf-events-sig-tracy-ragan-deployhub-emil-backmark-ericsson-andrea-frittoli-ibm-erik-sternerson-dowhile)
        - (AP Tracy) Start a shared Google Docs for the questions - [https://docs.google.com/document/d/1LgEp0ZuhhNsvp7oW-M0Uzdpsyu0uai8-YoB6L4S42Q8/edit]
        - We can still make it panel-like within the BoF
        - 30min slot
    - [KubeCon NA CFP](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/cfp/#overview) until May 23rd
        - We should submit proposals from the SIG
        - Present work done on the protocol
        - Any PoC that we have until then
- Blog Posts / Links
- Presentations
    - (Postponed until May 3rd) Mattias: Presentation - Eiffel Protocol with focus on linked/referenced events
    - 
- PR / Design Discussions
- Open 
- \<To be added\>


## Meeting April 5th
***Meeting cancelled due to bank holiday in several countries***


## Meeting March 15th
Meeeting time in your timezone [here](https://time.is/1600_15_March_2021_in_UTC). You're welcome to join!

Participants:
- Steve Taylor (DeployHub/Ortelius)
- Tracy Ragan (DeployHub/Ortlius)
- Mattias LInnér (Ericsson)
- Mauricio Salatino (Camunda/LearnK8s)
- Andreas Grimmer (Keptn/Dynatrace)
- Fredrik Fristedt (Axis Communications)
- Ravi Lachhman (Harness)
- Erik Sternerson (doWhile)
- Emil Bäckmark (Ericsson)
- Tracy Miranda (CDF)

### Agenda and Notes

- APs

- SIG setup - Emil/Andrea
  - Recordings uploaded to Youtube, since November 2020, Ok?
      - That's ok with us :)
  - New HackMD team
  - GitHub repo
  - GitHub discussions
  - [Slack](https://cdeliveryfdn.slack.com/archives/C0151BTKEJX)
  - [CDF announcement for the SIG](https://docs.google.com/document/d/1-dITtHInAnwpdoYo2_idQP2IG1CvNI6whqSzE8VckZk/edit#heading=h.hgckty8ur8h9)
  - [CDF podcast presenting our SIG](https://cdeliveryfdn.buzzsprout.com/1008697/8123014-cdf-special-interest-group-events)

- Feedback from Vocabulary meeting - Mauricio
  - Proposal to Aligning terms with CNCF audience, even if that is not our full scope, for the first revisions I think that it will help us to focus
  - Perhaps focus on one use case at a time? (But make sure "not to close the door" for later use cases)
      - 1: Cloud service development and deployment
      - 2: Embedded / high-complexity product delivery
      - 3: ...?
  - We intend to merge the pull requests in there in a week or so. But it will be very preliminary, and the definitions will be up for changes as the discussions go on.

- What should our SIG meetings focus on as we now have complementing vocabulary meetings? - All
    - Presentations?
    - Blog posts and similar?
    - Recurring feedback from workgroups (e.g. Vocabulary workgroup)
    - Upcoming Conferences

- Daylight savings time and frequency - Emil/Andrea
    - What is [your time](https://time.is/1600_5_April_2021_in_UTC) on Monday in ***three*** weeks? Could we have it one hour earlier ([3pm UTC](https://time.is/3pm_5_April_2021_in_UTC)) instead?
    - Decided to go for 3pm UTC from next meeting on.

- Current status of [PRs in sig-events](https://github.com/cdfoundation/sig-events/pulls)
    - And [one in TOC repo](https://github.com/cdfoundation/toc/pull/92)

- Upcoming Presentations
    - [Events Vocabulary](https://hackmd.io/B-Z7mLh_Qc6frm51f3GIYA?view)
        - We should wait for the current PRs to be merged before we present this, and base this presentation on those pull requests mainly. Add some info from the old vocabulary document if any.
    - [CD Flow](https://github.com/salaboy/cd-flow/) - date to be decided
    - Eiffel Protocol - with focus on linked/referenced events?
        - Mattias could have that presentation, but not on next meeting. In 5 weeks probably.


- \<add your agenda point\>

### Action Points
- ~~Emil/Andrea: Set up new repo, Slack, HackMD, etc for our SIG~~
- ~~Mauricio: Set up a meeting series for vocabulary discussions~~
- Emil: Reach out to Dan L to check if CDF defines what licenses to use.
    - Dan wasn't aware of that, but we're welcome to bring this to TOC meeting
    - Steve will bring it to TOC meeting tomorrow (March 16th)


## Meeting March 1st
Meeeting time in your timezone [here](https://time.is/1600_1_March_2021_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Erik Sternerson, doWhile
- Ravi Lachhman, Harness
- Mattias Linnér, Ericsson
- Steve Taylor, DeployHub
- Tracy Ragan, DeployHub
- Kara de la Marck, CloudBees
- Andrea Frittoli, IBM
- Mauricio Salatino, Camunda / LearnK8s
- \<addme\>

### Agenda and Notes
- Tracy M: Promoting the SIG
    - Roxanne is here to prepare for the announcement
    - https://docs.google.com/document/d/1-dITtHInAnwpdoYo2_idQP2IG1CvNI6whqSzE8VckZk/edit#
        - We will provide some quotes to that. Exampel from Jenkins https://cd.foundation/announcement/2020/08/04/cd-foundation-announces-jenkins-graduation/
    - Podcast
        - Ravi, Mauricio and Andrea will join
        - Postcast template: https://docs.google.com/document/d/1K-taXA4k5wc82tFQMVXvMDJArCANT2D4nneSmeVldzs/edit
    - Blog posts
        - Can be sent at any time to Roxanne, and not only about the Events SIG
- Daylight Savings Time
    - US starts on March 14th already: https://www.timeanddate.com/time/change/usa/new-york
    - Europe two weeks later: https://www.timeanddate.com/time/change/sweden/stockholm
    - What is [your time](https://time.is/1600_15_March_2021_in_UTC) on Monday in two weeks? Ok to keep the current time slot (4pm UTC)?
    - Decision: We keep the same time UTC next occasion as well
- We're a SIG!
  - Tracy M has asked us (during last SIG Interop meeting) to provide quotes for a CDF announcement of SIG Events - deadline March 1st. Example of topics:
      - About the problem space
      - The connection to CloudEvents
      - How it as been a spinoff from SIG Interoperability (Kara/Fatih)
  - New repo created for SIG Events
      - https://github.com/cdfoundation/sig-events/ 
      - Everyone can create a PR to that repo to be automatically included as contributors to the repo, and can then be added to the team
  - Update the [sig interoperability page](https://github.com/cdfoundation/sig-interoperability/README.md) in the cdfoundation repo
      - (Andrea) Make a PR to move documents and update links
  - Move our stuff from https://github.com/cdfoundation/sig-interoperability/tree/master/workstreams/events_in_cicd
      - (Andrea) Make a PR to move documents into the new repo
  - Add the mailing list to the communication channel once allocated
      - Tracy M will create the maillist for us sig-events@lists.cd.foundation subscribe via  https://lists.cd.foundation/g/sig-events
      - Andrea will add info about it to our readmes
  - HackMD document updates (title, new hackmd team, github sync). Create new team in Hackmd
      - Emil to create a new hackmd team and connect its docs to the new SIG Events github
      - Emil put a note in the old document saying this is now archived and is continued in the new document
  - Slack channel renamed to sig-events, and change its description.
      - Tracy M will reach out to Gale to rename our channel to sig-events
      - Emil updates the description in the Slack channel
      - Andrea to update info in the README
  - Write a blog post on the CDF blog
      - Mauricio to start it
      - A follow up on the announcement
          - Mentioning specific initiatives if applicable
          - Use cases
  - CDF Podcast
      - See above
  - Setup GitHub repo
    - Created sig-events team, sig-events-admin team
    - Add chairs as admins, everyone else as maintainer
    - https://github.com/cdfoundation/sig-events/
- Initial Vocabulary Proposal
    - First vocabulary meeting brief
    - Initial Vocabulary Proposal PR https://github.com/cdfoundation/sig-interoperability/pull/64/files
    - Proposed inital workflow for working on the vocabulary
        - Break the PR into smaller chunks to be able to merge sooner. Mauricio will fix.
        - PR comments for discussions directly related to the PR
        - GitHub Discussions for "more involved" topics, consider announcing on Slack when there is a new discussion you want some attention on.
            - Who can enable GitHub Discussions for our repo?
    

## Meeting February 15th
Meeeting time in your timezone [here](https://time.is/1600_15_February_2021_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark - Ericsson
- Mauricio Salatino - Camunda / LearnK8s
- Mattias Linnér - Ericsson
- Andreas Grimmer, Keptn
- Erik Sternerson - doWhile
- Andrea Frittoli - IBM
- Tracy Ragan - DeployHub
- Aleksandra Fedorova - Red Hat

### Agenda and Notes

- APs
- (salaboy) Setting a time in the calendar for Vocabulary discussion and progress
    - Mauricio will set up a separate meeting to start that discussion. Proposal: same time, different day. AP salaboy
    - Categories of events added to our vocabulary document - https://hackmd.io/B-Z7mLh_Qc6frm51f3GIYA
- Becoming a SIG
    - Waiting for votes in TOC
- (afrittoli) Licenses for content produced by the group / SIG (see https://cdeliveryfdn.slack.com/archives/C0151BTKEJX/p1612211728022200)
    - Emil reached out to Dan, waiting for reply. Emil pings Dan again :)
    - Input to this discussion: https://github.com/CommunitySpecification/1.0
- [Enterprise-Scale CI](https://github.com/erkist/sig-interoperability/tree/erkist-esci/workstreams/events_in_cicd/input/enterprise-scale-ci) concepts - Erik Sternerson
    - Complementing info from Aleksandra on RedHat (Fedora): https://quantum-integration.org/posts/ci-system-vs-ci-pipeline.html
    - Also worth relating to the SIG Interoperability vocabulary document (Rosetta Stone): https://github.com/cdfoundation/sig-interoperability/blob/master/docs/vocabulary.md
- Upcoming Presentations
  - [Events Vocabulary](https://hackmd.io/B-Z7mLh_Qc6frm51f3GIYA?view)
  - [CD Flow](https://github.com/salaboy/cd-flow/) - date to be decided

### APs
- ~~Andrea: Reach out on the cloudevents Slack channel - https://cloud-native.slack.com/archives/C9DB5ABAA/p1612448038043100~~
- ~~All: We should define buckets of terms to define, e.g. artifacts, activities, deployments etc and continue diving into those to propose vocabulary for them - Let's discuss how to proceed on that on Slack~~
- Mauricio (Salaboy) to set up a meeting series for vocabulary discussions


## Meeting February 1st
Meeeting time in your timezone [here](https://time.is/1600_1_February_2021_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Erik Sternerson, doWhile
- Mattias Linnér, Ericsson
- Andrea Frittoli, IBM
- Mauricio Salatino, Camunda / LearnK8s
- Andreas Grimmer, Keptn
- Tracy Ragan, DeployHub / Ortelius
- Gareth Evans, CloudBees

### Agenda and Notes

#### Upcoming Conferences
- [FOSDEM 2021](https://fosdem.org/2021/) - February 7th - 8th
- [Istio-CON](https://sessionize.com/istiocon-2021/) - Feb 22 - 25, 2021 
- [Devops Online Summit](https://devopsonlinesummit.com/) - April 26th – 30th
    - [CFP](https://devopsonlinesummit.com/the-tremendous-devops-online-summit-2021-tracks-have-just-arrived/) closes on February 7th
    - Focus topics: Trends in CI/CD and landscape for CI/CD
- [cdCon 2021](https://events.linuxfoundation.org/cdcon/) - June 23rd - 24th
    - [CFP](https://events.linuxfoundation.org/cdcon/program/cfp/) closes March 5th (Early Bird closes February 19th)
    - TR: A panel proposal would be good. They are often well attended.
- FYI (partly relevant) [Scaling Continuous Delivery](https://continuousdeliveryatscale.splashthat.com/) February 2nd, Ravi and Andrea will be presenting

#### Becoming a SIG
- [Draft PR created](https://github.com/cdfoundation/toc/pull/91)
- Will probably be discussed on tomorrow's TOC meeting: https://docs.google.com/document/d/1uBHar55fTInWF9Li4t0lyG3tTC8BRLU0FfBfsgk_Jrs/edit?ts=5c9580be#
- Meeting cadence needs to be in sync with Best Practices SIG
    - Events SIG meet every 1st and 3rd Monday of the month
    - Best Practices SIG should meet every 2nd and 4th Monday of the month. @Mauricio to bring that to Best Practices SIG.
- Follow-up to be done once the SIG proposal is accepted
    - Update the [sig interoperability page](https://github.com/cdfoundation/sig-interoperability/README.md) in the cdfoundation repo
    - Move our stuff from https://github.com/cdfoundation/sig-interoperability/tree/master/workstreams/events_in_cicd
    - Add the mailing list to the communication channel once allocated
    - HackMD document updates (title, new hackmd team, github sync). Create new team in Hackmd
    - Slack channel renamed to sig-events, and change its description.
    - Write a blog post on the CDF blog. (Mauricio to start it)
    - CDF Podcast
    - Setup GitHub repo
        - Create/ask to create sig-events team, sig-events-admin team
        - Add chairs as admins, everyone else as maintainer

#### Contribute to Metadata workstream
- Metadata workstream is about to get started and we should provide any input we have that relates to metadata to it: [Standardized Metadata](https://hackmd.io/BYbkuR8uSlKt_w7Y4KE1OQ)
- We don't have anything to provide on 'artifacts' and 'commits' yet from our group
- We should define buckets of terms to define, e.g. artifacts, activities, deployments etc and continue diving into those to propose vocabulary for them
    - Let's discuss how to proceed on that on Slack. AP all

#### Upcoming Presentations
- [CD Flow](https://github.com/salaboy/cd-flow/) - date to be decided
- [Events Vocabulary](https://hackmd.io/B-Z7mLh_Qc6frm51f3GIYA?view)
- [Enterprise-Scale CI](https://github.com/erkist/sig-interoperability/tree/erkist-esci/workstreams/events_in_cicd/input/enterprise-scale-ci) concepts - Erik Sternerson

#### Cloud Events
- Question from the Interop WG. Did we connect (or plant to) with the CloudEvents team?
    - Reach out on the cloudevents Slack channel - AP Andrea

### APs
- Andrea: Reach out on the cloudevents Slack channel
- All: We should define buckets of terms to define, e.g. artifacts, activities, deployments etc and continue diving into those to propose vocabulary for them
    - Let's discuss how to proceed on that on Slack

## Meeting January 18th
Meeeting time in your timezone [here](https://time.is/1600_18_January_2021_in_UTC). You're welcome to join!

Participants:
- Andreas Grimmer, Dynatrace
- Emil Bäckmark, Ericsson
- Ravi Lachhman, Harness
- Mattias Linnér, Ericsson
- Steve Taylor, DeployHub/Ortelius
- Mauricio Salatino, Camunda / LearnK8s 
- Erik Sternerson, doWhile
- Tracy Ragan, DeployHub/Ortelius

### Agenda and Notes

#### Upcoming Conferences
- [FOSDEM 2021](https://fosdem.org/2021/) - February 7th - 8th
    - Andrea will present our workstream in the [CI/CD room](https://fosdem.org/2021/schedule/track/continuous_integration_and_continuous_deployment/)
    - https://afrittoli.github.io/events-in-cicd/
- [Devops Online Summit](https://devopsonlinesummit.com/) - April 26th – 30th
    - [CFP](https://devopsonlinesummit.com/the-tremendous-devops-online-summit-2021-tracks-have-just-arrived/) closes on February 7th
    - Held on Slack
    - Tracy Ragan will drive the CI/CD track
- [cdCon 2021](https://events.linuxfoundation.org/cdcon/) - June 23rd - 24th
    - [CFP](https://events.linuxfoundation.org/cdcon/program/cfp/) closes March 5th (Early Bird closes February 19th)
- [Istio-CON](https://sessionize.com/istiocon-2021/) - Feb 22 - 25, 2021 
    - CFS closes today

#### Becoming a SIG
- Will possibly be brought up on tomorrow's TOC meeting: https://docs.google.com/document/d/1uBHar55fTInWF9Li4t0lyG3tTC8BRLU0FfBfsgk_Jrs/edit?ts=5c9580be#
- [CDF Toc Invite for tomorrow](https://calendar.google.com/calendar/event?eid=N2tiaTNwamdlOW0xbnRiYzZodjgzYW10NXJfMjAyMTAxMTlUMTYwMDAwWiBsaW51eGZvdW5kYXRpb24ub3JnX21oZjBrbWdlZG42N2lobmk4cjEyOWF2cDI0QGc)
- https://zoom.us/j/262239698?pwd%3DNWhQWjg2UUdKT0ZxYW1CSElwbDVodz09&sa=D&source=calendar&ust=1611503983559000&usg=AOvVaw03zXFPiKEUEw29kp-VGUje
    - Emil could present our workstream there
- Events should be considered for all areas in the [CD Foundation Landspace](https://landscape.cd.foundation/)

#### Charter
[Charter document in progress](https://docs.google.com/document/d/1cuaOHmrvK20WxHxHZZuLif3CW3q1ICZ_RglJRp_IAQ0/edit?usp=sharing)
- Main focus areas for the Events SIG (as written on last meetings minutes)
  - Defining a protocol, and documenting it
  - Providing how to apply the protocol in different use cases, such as monitoring, triggering activities, audits, etc
  - Providing a reference implementation of services/libraries showcasing the protocol

#### Contribute to Metadata workstream
- Metadata workstream is about to get started and we should provide any input we have that relates to metadata to it: [Standardized Metadata](https://hackmd.io/BYbkuR8uSlKt_w7Y4KE1OQ)

#### Upcoming Presentations
- [CD Flow](https://github.com/salaboy/cd-flow/) - date to be decided
- [Events Vocabulary](https://hackmd.io/B-Z7mLh_Qc6frm51f3GIYA?view)

#### Cloud Events
- Question from the Interop WG. Did we connect (or plant to) with the CloudEvents team