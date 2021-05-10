# Events SIG Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g/badge)](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g)

This SIG meets the 1st and 3rd Mondays of every month at 3pm UTC (during summer time, in winter time it is 4pm UTC). See your timezone [here](https://time.is/3pm_in_UTC). The meeting can be accessed through [this zoom link](https://zoom.us/j/97660712600?pwd=Z3BqYTE5YzNsbEhmck16cjdZNEFIUT09).

The forming of this workstream was suggested on a [recent SIG Interoperability meeting]( https://github.com/cdfoundation/sig-interoperability/blob/master/docs/meetings.md#may-28-2020) and its first meeting was held on June 8th.

Meeting notes for the workstream are managed on HackMD [here](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g), and published to GitHub [here](https://github.com/cdfoundation/sig-interoperability/blob/master/workstreams/events_in_cicd/meetings.md).


## Agenda Template

- AP From Last week
- SIG Events Updates
- Workstream Updates
  - Dictionary
- Meetups / conferences
- Blog Posts / Links
- Presentations
- PR / Design Discussions

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
            - Eduction material
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
- Presentations
    - Eiffel (Mattias)
- PR / Design Discussions


## Meeting April 19th
Meeeting time in your timezone [here](https://time.is/3pm_19_April_2021_in_UTC). You're welcome to join!

Participants:
- Andrea Frittoli (IBM)
- Mattias LInnér (Ericsson)
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