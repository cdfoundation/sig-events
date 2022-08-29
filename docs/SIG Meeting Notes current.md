---
tags: SIG Events
---

# CDF Events SIG Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g/badge)](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g)

This SIG meets bi-weekly on Mondays at [3pm UTC](https://time.is/3pm_in_UTC) during summer time and at [4pm UTC](https://time.is/4pm_in_UTC) during winter time). The meeting can be accessed through [this zoom link](https://zoom.us/j/98408983891?pwd=VXBxMjJTaThGRGFWRXFmdUxsRUZUdz09). For older meetings please see our [playlist on YouTube](https://www.youtube.com/playlist?list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC)

The SIG was initiated as a [workstream under SIG Interoperability ]( https://github.com/cdfoundation/sig-interoperability/blob/master/docs/meetings_2020.md#May-28-2020) and its first meeting was held on June 8th 2020.

Meeting notes for the workstream are managed on HackMD [here](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g), and published to GitHub [here](https://github.com/cdfoundation/sig-events/blob/main/docs/meetings.md).


## Agenda Template

- Action Items From Last meeting
- Updates From Other Groups
    - CDF/TOC
    - SIG Interoperability, Best Practices, Software supply Chain
    - CDEvents [Vocabulary/Spec Updates](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- CD Events Project Updates
- Upcoming/Recent Meetups & Conferences
- Today's Presentations
- Future Presentations


## Future Presentations

- Brett to present the event CI/CD architecture at SAS
- Mike to present/demo the work on CNCF Reference Architecture, OpenSFF and events
    - Frsca

## Meeting August 29th

### Participants
- \<add yourself\>


### Agenda
- Action Items From Last meeting
    - Zoom meeting link confusion sorted out
- Updates From Other Groups
    - CDF/TOC
    - SIG Interoperability, Best Practices, Software supply Chain
    - CDEvents [Vocabulary/Spec Updates](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- CD Events Project Updates
    - GitHub Migration from cdfoundation/sig-events to cdevents
    - New governance setup to be initiated
- Upcoming/Recent Meetups & Conferences
    - [CD mini summit](https://cd.foundation/cd-mini-summit-2022-dublin/), Sept 12th, Dublin
    - [Open Source Summit](https://events.linuxfoundation.org/open-source-summit-europe/), Sept 13th-16th, Dublin
    - [Spinnaker Summit](https://events.linuxfoundation.org/spinnaker-summit/program/schedule/), Oct 23rd-24th, Detroit
    - [CD summit](https://cd.foundation/blog/2022/08/23/continuous-delivery-summit-2022-call-for-papers-open/), Oct 25th, Detroit
        - CFP ends Aug 31st
    - [KubeCon](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/), Oct 24th-28th, Detroit
- Today's Presentations
    - Mike presents/demoes the work on CNCF Reference Architecture, OpenSFF and events
- Future Presentation Proposals


## Meeting August 15th

### Participants
- Emil Bäckmark, Ericsson
- Mike Lieberman, Kusari
- Brett Smith, SAS
- Tracy Ragan, DeployHub
- Kara de la Marck, CDF


### Agenda
- Seems the zoom link on the CDF calendar to this meeting doesn't match the link on GitHub and in this MoM. Emil will fix it.
- Presentation from Fidelity (Jamie Plower) on their use of Cloudevents in CI/CD pipelines.


## Meeting August 1st

### Participants

- Mike Lieberman, Kusari and CNCF Security TAG
- Tim Miller, Kusari
- Tracy Ragan, DeployHub, CDF TOC, OpenSSF Board
- Vibhav Bobade, Uffizzi
- Parth Patel, Kusari
- Brett Smith, SAS Institute

### Agenda

- Reference Architecture
    - It should be vendor neutral
    - How does it relate to the [vocabulary](https://github.com/cdfoundation/sig-events/blob/main/vocabulary-draft/introduction.md)
        - The reference architecture should contribute back into the vocabulary
            - And contribute back to the landscape
    - Reference architecture can help to:
        - train engineers
        - have a concrete example to explain events driven architecture
        - explain what to do/not to with CD tools
        - naming is hard: terms can be overloaded
    - There are not best practices that are valid for everyone
        - Different companies will need different practices
        - How do we target the reference architecture

- Landscape
    - Creating the CDF landscape is hard
        - It's relatively static
        - It's difficult to put a tool in a category since you need to select a primary feature 
        - The landscape is not mapped to the vocabulary
    - Primary category + relation to vocabulary

- Working group to define a vocabulary across SIGs and landscape
    - Bring a proposal for this WG to the TOC

- CDEvents Whitepaper: 
    - https://cd.foundation/blog/2022/06/07/cdevents-publishes-first-whitepaper/
    - [PDF](https://redpanda.com/)
    - We should convert it into markdown and add it to the repo
    - Continue to iterate on the whitepaper
    - We could use the whitepaper

- Kafka vs RedPanda
    - RedPanda drop-in replacent of Kafka written in C
    - [RedPanda](https://redpanda.com/)


## Meeting July 18th

Meeeting time in your timezone [here](https://time.is/3pm_4_July_2022_in_UTC). You're welcome to join!

### Participants
- Erik Sternerson, doWhile
- Andrea Frittoli, IBM
- Kara de la Marck, CDF

### Last Meeting's Actions

- Oleg to provide a link to more info about this
    - We still have budget to use for CDEvents during this year. It could be used for the 0.1 release or other activity related to CDEvents.
    - Andrea: create issue about CDEvents visualisation
        - https://github.com/cdfoundation/sig-events/issues/126 

### Agenda

- Updates From Other Groups
    - CDF/TOC
    - SIG Interoperability
    - SIG Best Practices
    - [Vocabulary workstream](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)

- Jenkins CloudEvents->CDEvents plugin transformation

- Community repo (Andrea)
    - Plan to create a cdevents/community repo over the summer
    - Migrate existing relevant docs and links
    - Define governace docs in the new repo

- Python SDK progressing
    - First goal is 1:1 parity with Go SDK, and evolve from there
    - Current ongoing work in forked repo at https://github.com/tarekbadrshalaan/sdk-python/tree/skeleton-project (very much work-in-progress status)

- CDEvents presentation at Open Source Summit Europe
    - Friday Sept 16 @ 15:55-16:35
    - https://sched.co/15z6i



## Meeting July 4th

Meeeting time in your timezone [here](https://time.is/3pm_4_July_2022_in_UTC). You're welcome to join!

### Participants
- Emil Bäckmark, Ericsson
- Erik Sternerson, doWhile
- Andrea Frittoli, IBM
- Jalander Ramagiri, Ericsson Software Technologies
- Oleg Nenashev, Dynatrace

### Last Meeting's Actions
- None

### Agenda
- Last meeting's actions
- [CDEvents 0.1 Planning](https://github.com/orgs/cdevents/projects/1/views/1)
    - [Release announcement](https://github.com/cdevents/spec/issues/47)
    - [Project governance plan](https://github.com/cdevents/spec/issues/48)
        - Should be started soon (after summer?)
    - CDF Outreach Committee had a meeting last week, and we should interact more to make sure the CDEvents release is published in relevant channels and with the relevant artifacts provided
        - Oleg to provide a link to more info about this
    - We still have budget to use for CDEvents during this year. It could be used for the 0.1 release or other activity related to CDEvents.
        - Technical writer or similar could be used to make our docs more professional. copy editing for the documentation or new docs
        - No need to hire an intern/student, LFX Mentorship support having experienced developers/writers who want to start in open source
        - Internships for extended features related to CDEvents
            - Adapters for established platforms/tools, such as GitHub etc. The money should be used neutrally to vendors though.
            - (Andrea) Providing a visualization frontend for CDEvents would be great. Maybe just a PoC one for internal/demo use, not a production tool
            - Automation for SDKs etc
         - Alternative: supporting organized events at a major conference, or maybe travel grants
         - Plan:
             - Oleg to launch a call for potential docs contributors.
             - Oleg to setup the LFX Mentorship project and the Crowdfunding backend, get the LF to transfer money there on behalf of the CDF.
             - Andrea to look in to mentoring for a visualization solution
- Presentations
        
    - Jalander: Demo of Spinnaker integration (https://github.com/cdfoundation/sig-events/pull/120)
        - Expands the current PoC, where Keptn and Tekton communicates, with a Spinnaker integration that does a similar deployment.
        - Patches were needed in Keptn. Oleg to follow-up on what's needed (CDEvents adoption) and to invite Jalander to the developer meeting
- Upcoming Presentations
    - Fidelity. Date being discussed. Probably August 1st, unless Jamie shows up on his first working day after vacations on July 18th.
- Takeaways from Conferences
    - June 21-24: [Open Source Summit NA](https://events.linuxfoundation.org/open-source-summit-north-america/)
- Upcoming Conferences
    - Sept 13-16: [Open Source Summit Europe ](https://events.linuxfoundation.org/open-source-summit-europe/)
        - Co-located events can still be requested
        - Andrea has submitted a CFP for a CDEvents intro, together with Erik
    - Sept 1-Oct 1?: [Spinnaker Summit](https://cd.foundation/event/spinnaker-summit-2/)
        - Not really launched yet
    - Oct 24-28: [KubeCon NA](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/)
        - Deadline for requesting co-located events is July 29
        - CFP for Spinnaker integration provided by Jalander
- Sync with SIG Interop about definition of terms
    - https://github.com/cdevents/spec/issues/18


## Meeting June 20th

Meeeting time in your timezone [here](https://time.is/3pm_20_June_2022_in_UTC). You're welcome to join!

Participants:
- \<add yourself\>
- Erik Sternerson, doWhile
- Fatih Degirmenci, CDF
- Emil Bäckmark, Ericsson
- Jalander Ramagiri, Ericsson Software Technologies

Agenda:
- Last meetings actions
- CDEventsCon takeaways
    - [Recordings](https://www.youtube.com/playlist?list=PL2KXbZ9-EY9QtshpDdxEyDTbSPMZzZZS1) and most [Presentations](https://github.com/cdevents/presentations/tree/main/CDEventsCon-2022) are available
    - Any feedback to bring to next CDEventsCon or cdCon/CDEvents summit?
    - We seem to need to re-iterate the basic CDEvents presentation over and over again. We should probably store an introduction. Primarily the presentation but also a recorded version would be beneficial. A short intro video would be suitable to publish on cdevents.dev as well. We should use the existing intro presentation as a base for it. https://github.com/cdevents/spec/issues/49
- CDEvents 0.1 Planning
    - [Release announcement](https://github.com/cdevents/spec/issues/47)
    - [Project governance plan](https://github.com/cdevents/spec/issues/48)
        - Should be started soon (after summer?)
- Upcoming Presentations
    - Fidelity? Emil to reach out to Jamie Plower and Ger McMahon
    - Jalander: Demo of Spinnaker integration (https://github.com/cdfoundation/sig-events/pull/120) - date to be set
- Upcoming Conferences
    - June 21-24: [Open Source Summit NA](https://events.linuxfoundation.org/open-source-summit-north-america/)
    - Sept 13-16: [Open Source Summit Europe ](https://events.linuxfoundation.org/open-source-summit-europe/)
        - Co-located events can still be requested
        - Andrea has submitted a CFP for a CDEvents intro, together with Erik
    - Sept 1-Oct 1?: [Spinnaker Summit](https://cd.foundation/event/spinnaker-summit-2/)
        - Not really launched yet
    - Oct 24-28: [KubeCon NA](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/)
        - Deadline for requesting co-located events is July 29
        - CFP for Spinnaker integration provided by Jalander
- Summer meeting schedule
    - SIG meetings: July 4th (No US participation?), July 18th, Aug 1st
    - Vocabulary meetings: June 28th, July 12th, July 26th, Aug 9th
    - Cancel July meetings?
    - Erik will be working through the summer. Emil will be away from second week in July - first week of August. Jalander plan to work all summer.
    - Let's discuss it on Slack
- Integrate Spinnaker with sig-events PoC
    - PR review https://github.com/cdfoundation/sig-events/pull/120
- RFC : Implementing CDEvents to Spinnaker using Java SDK
    - https://github.com/spinnaker/governance/pull/299


## Meeting May 23th

Meeeting time in your timezone [here](https://time.is/3pm_23_May_2022_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Erik Sternerson, doWhile
- Andrea Frittoli, IBM
- Tracy Ragan, DeployHub
- Kara de la Marck, CDF
- \<add yourself\>

Agenda:
- Last meetings actions
- Launch of CDEvents project
    - [Post](https://cd.foundation/announcement/2022/05/17/cd-foundation-announces-cdevents-a-vendor-neutral-specification-for-defining-the-format-of-event-data/)
    - Should we write a whitepaper accompanying the launch, providing the more high-level values of CDEvents?
        - Yes. We're probably not able to finish it before CDCon, but we should mention there that we're working on it.
        - https://docs.google.com/document/d/1SVVE-N2gzxQ9us-kDB0dBhkDZS9JsZAWjrB2doLnpok/edit
        - Everyone should start contributing to it now
- CDEvents Community Meeting at cdCon
    - https://hackmd.io/rzJPQL4nTlWORg-wwVBw1g
    - Virtual Platform?
        - Bevvy for the event and also for the actual online meeting
    - We have some ideas but the event needs to be organised
        - Volunteers?
    - Emil, Andrea and Erik to prepare for the topics to discuss (see hackmd doc above)
        - Create a hackmd doc for each and provide an intro to the discussion
- cdCon
    - Booth?
        - No, but a demo theater is available
- Takeaways from CDEventsCon/KubeCon
    - CDEventsCon presentations to be uploaded to Sched and [CDEvents presentations repo](https://github.com/cdevents/presentations)
        - [PR #6](https://github.com/cdevents/presentations/pull/6)
    - Photos: https://photos.app.goo.gl/18fSgCZVrqU1bRnw9
        - Andrea can add us as contributors to that folder to upload photos
    - Social: Twitter / Linked-In
- Open Source Summit EU
    - CFP Closes May 30th
- CDEvents Project Governance
    - Time to start setting it up?


## Meeting May 9th

Meeeting time in your timezone [here](https://time.is/3pm_9_May_2022_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Steve Taylor, DeployHub
- Tracy Ragan, DeployHub
- Mattias Linnér Ericsson
- Andrea Frittoli, IBM
- Erik Sternerson, doWhile
- \<add yourself\>

Agenda:
- Last meetings actions
    - (Emil) Reach out to Kara to find out whats left before the CDEvents project can be officially launched
        - Reached out, and it is ongoing.
- Conclusions from the webinar/meetup
    - https://www.youtube.com/watch?v=qTii8beEyBM
- CDEventsCon
    - RSVP for virtual attendees not yet in place. Emily is working on it.
    - Panel discussion
        - https://docs.google.com/document/d/1HZwgnsPFgBBl4gmZHqacJoWWBwFlxbICg5dJCUNsF8s/edit
    - [CDEventsCon Schedule](https://cdcon2022.sched.com/event/10fUW)
- CDEvents Contributor Summit @ cdCon
    - [cdCon full schedule](https://events.linuxfoundation.org/cdcon/program/schedule/)
    - [Planning doc](https://hackmd.io/rzJPQL4nTlWORg-wwVBw1g)
    - 3 hours only (half-day)
- Presentations to be held
    - Still waiting for response from Fidelity
    - Cartographer is to be presented on May 26th on SIG supply chain: https://hackmd.io/l1BfXp1kQKGKSaKLbl6xJw?view


## Meeting April 25th

Meeeting time in your timezone [here](https://time.is/3pm_25_April_2022_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Mattias Linnér, Ericsson
- Erik Sternerson, doWhile
- Tracy Ragan, DeployHub

Agenda:
- Action Items From Last meeting
    - (Andrea) Ask Roxanne to help announcing the webinar
        - Roxanne accepted to announce it
        - Emil asked Michelle today to announce it, but it might be too late to attract new people to it now
    - (All) Reach out to the companies/organizations that could be interested to join the webinar
    - (Emil) Reach out to Kara to find out whats left before the CDEvents project can be officially launched
        - Reached out, and it is ongoing.
- Webinar/meetup April 27th
    - https://community.cd.foundation/events/details/cd-foundation-cdevents-presents-cdevents-online-webinarmeetup/
    - Agenda:
        - Andrea (and Erik) presents an updated presentation from FOSDEM
            - In-progress update pushed to <https://github.com/cdevents/presentations/tree/main/2022-04-27-cdevents-online>
            - Q: What do we want to present as updates?
        - Panel or open discussion on the participants use cases
- CDEvents Panel
    - We had [this idea](https://docs.google.com/document/d/12UNdVBO7_rOBIEtOsrjViEjVUdp8dyYl_tM5ljSP7Ps/edit) some time ago, should we try to squeeze it is as well? 
    - We can revisit this idea for sometime early autumn unless the panel on CDEventsCon would replace it.
- CDEventsCon May 17th @ KubeCon
    - [Planning](https://hackmd.io/JMM-V7IlTwWoS1YcDK9m_A)
    - [Blog announcement](https://cd.foundation/blog/2022/04/11/attend-cdeventscon-this-may/)
    - (Andrea) to reach out to the speakers to see that it is all set
- CDEvents Day June 10th @ cdCon
    - [Full cdCon Schedule](https://events.linuxfoundation.org/cdcon/program/schedule/)
    - [CDEvents Community Summit Schedule](https://cdcon2022.sched.com/event/10fUW/cdevents-community-summit?iframe=no&w=100%&sidebar=yes&bg=no)
    - [Erik and Andrea will talk](https://cdcon2022.sched.com/event/10UYC/building-devops-metrics-for-your-choice-of-cd-tools-through-cdevents-andrea-frittoli-ibm-erik-sternerson-dowhile?iframe=yes&w=100%&sidebar=yes&bg=no)
- CD Events Project set up
    - [PR Document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit#)
    - (Kara) Marketing Clip
- Other items
    - Updates from other CDF SIGs, PRs/issues
    - https://github.com/cdevents/spec/issues/18
    - https://github.com/cdevents/spec/pull/35


## Meeting April 11th

Meeeting time in your timezone [here](https://time.is/3pm_11_April_2022_in_UTC). You're welcome to join!

Participants:
- Wilhelm Wonigkeit (Direktiv)
- Emil Bäckmark (Ericsson)
- Steve Taylor (DeployHub)
- Mattias Linnér (Ericsson)
- \<add yourself\>

Agenda:
- Action Items From Last meeting
    - (Emil) Plan for webinar/meetup on April 26th or 27th and add to CDF calendar
        - Done. April 27th is set.
    - (Andrea) Ask Roxanne to help announcing the webinar
        - Roxanne accepted to announce it
    - (All) Reach out to the companies/organizations that could be interested to join the webinar
- Presentations
    - [Direktiv](https://github.com/direktiv/direktiv) presented by Wilhelm Wonigkeit
        - See recording of this SIG for demo and discussions
- Meetups / Conferences
    - Webinar
        - Agenda:
            - Andrea (and Erik) presents an updated presentation from FOSDEM
            - Panel or open discussion on the participants use cases
    - CDF Online Meetup
        - https://docs.google.com/document/d/12UNdVBO7_rOBIEtOsrjViEjVUdp8dyYl_tM5ljSP7Ps/edit
    - KubeCon
        - CDEventsCon May 17th
            - https://hackmd.io/JMM-V7IlTwWoS1YcDK9m_A
            - https://cd.foundation/blog/2022/04/11/attend-cdeventscon-this-may/
        - CDF Booth
            - Canceled (?)
    - cdCon
        - https://events.linuxfoundation.org/cdcon/program/schedule/
        - (Andrea) cdCon contributor summit
            - June 10th
- Updates From Other Groups
- CD Events Project set up
    - [PR Document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit#)
    - (Kara) Marketing Clip
    - (Emil) Reach out to Kara to find out whats left
- Blog Posts / Links
- PR / Design Discussions


## Meeting March 28th

Meeeting time in your timezone [here](https://time.is/3pm_28_March_2022_in_UTC). You're welcome to join!

Participants:
- Steve Taylor (DeployHub)
- Tracy Ragan (DeployHub)
- Mattias Linér (Ericsson)
- Emil Bäckmark (Ericsson)
- Andrea Frittoli (IBM)
- \<add yourself\>

Agenda:
- Action Items From Last meeting
    - (Steve) Reach out to Direktiv for a presentation on their events usage
        - Contact in place (Emil et al). Possibly a presentation for April 25th
    - (Oleg) Add webinar on April 4th to CDF calendar
    - (Andrea) Ask Roxanne to help announcing the webinar
    - (All) Reach out to the companies/organizations that could be interested to join the webinar
- Meetups / Conferences
    - Webinar
        - Proposed to have it on April 4th, at 3pm-4pm UTC
            - April 4th is not feasible anymore
        - We should use the Bevy portal for this. CDEvents chapter - https://community.cd.foundation/cdevents/
        - We could join this initiative and the CDF Online Meetup discussed below. We should use the Bevy portal and announce it through the CDF Online Meetup forum
        - Agenda:
            - Andrea (and Erik) presents an updated presentation from FOSDEM
            - Panel or open discussion on the participants use cases
        - To try out Bevy we could set up one of our SIG meetings through Bevy (vocabulary meeting next week?)
        - Andrea will be on PTO from April 7th to 25th. We should set up this meetup in the week of 25th (earliest on 26th). To be discussed on slack.
    - CDF Online Meetup
        - https://docs.google.com/document/d/12UNdVBO7_rOBIEtOsrjViEjVUdp8dyYl_tM5ljSP7Ps/edit
    - KubeCon
        - CDEventsCon
            - May 17th, 9am-5pm
            - https://hackmd.io/JMM-V7IlTwWoS1YcDK9m_A
        - CDF Booth
            - Canceled (?)
    - cdCon
        - (Andrea) cdCon contributor summit
            - google form done for <10 in person partecipants, requested support for remote partecipants
            - Info to be provided until April 1st:
                - Event contact name and email address
                - Official event title - this will be displayed on the schedule and on print signage
                - 300 word (maximum) event description
                - Date of event - originally submitted: Friday, June 10 (Options: Thursday, June 9 / Friday, June 10 / Thursday, June 9 AND Friday, June 10)
                - Event time frame - originally submitted: 9:00 AM - 12:00 PM. Note that the convention center opens at 8:00 AM, so sessions should start closer to 9:00 AM.
                - Estimated number of attendees - originally submitted: <10
                - Link to external event site or schedule (optional)
                - Project high resolution logo (if applicable)
                - Email address or contact information for attendee questions (to be listed on the website)
- Updates From Other Groups
- CD Events Project set up
    - [PR Document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit#)
    - (Kara) Marketing Clip
- Blog Posts / Links
- Presentations
- PR / Design Discussions


## Meeting March 14th

Meeeting time in your timezone [here](https://time.is/4pm_14_March_2022_in_UTC). You're welcome to join!


Participants:
- Matias Linnér, Ericsson
- Emil Bäckmark, Ericsson
- Andrea Frittoli, IBM
- Oleg Nenashev, Dynatrace, Keptn/Jenkins
- \<add yourself\>

Agenda:
- Action Items From Last meeting
    - (Steve) Reach out to Direktiv for a presentation on their events usage
- General
    - Daylight savings time
        - US is already there. For most counrties in Europe it begins on March 27th (before next meeting)
        - Proposal: Move to 3pm UTC. Next meeting on [this time](https://time.is/3pm_28_March_2022_in_UTC)
        - Decided to accept the proposal
        - (Oleg) update the CDF calendar from next meeting
- Meetups / Conferences
    - Webinar
        - Postponed from March 2nd. New date?
        - Proposed to have it on April 4th, at 3pm-4pm UTC
        - We should use the Bevy portal for this. CDEvents chapter
        - (Oleg) Add to CDF calendar
        - (Andrea) Roxanne could help us announce this
        - Agenda:
            - Andrea (and Erik) presents an updated presentation from FOSDEM
            - Open discussion on the participants use cases
        - We should all reach out to the companies/organizations that could be interested to join
    - KubeCon
        - CDEventsCon
            - May 17th, 9am-5pm
            - https://hackmd.io/JMM-V7IlTwWoS1YcDK9m_A
        - CDF Booth
            - All days? CDEvents one day.
    - cdCon
        - (Andrea) cdCon contributor summit
            - google form done for <10 in person partecipants, requested support for remote partecipants
- Updates From Other Groups
    - CDF/TOC
    - SIG Interoperability
    - SIG Best Practices
    - SIG Software Supply Chain
    - [Vocabulary workstream](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- CD Events Project set up
    - [PR Document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit#)
    - (Kara) Marketing Clip
- Blog Posts / Links
- Presentations
- PR / Design Discussions


## Meeting February 28th

Meeeting time in your timezone [here](https://time.is/4pm_28_February_2022_in_UTC). You're welcome to join!


Participants:
- Steve Taylor (DeployHub)
- Erik Sternerson (doWhile)
- Mattias Linnér (Ericsson)
- Andrea Frittoli (IBM)
- 

Agenda:

- Action Items From Last meeting
    - (Andrea) [Contribution process added](https://github.com/cdevents/.github/blob/main/docs/CONTRIBUTING.md)
    - (Andrea) cdCon contributor summit
        - google form done for <10 in person partecipants 
            - requested support for remote partecipants
    - (Emil / Andrea) [PR Document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit#)
        - Comments merged in (mostly)
        - We have a few quotes now, it would be nice to have some more
    - (Kara) Marketing Clip
    - (Andrea) CDEvents at KubeCon
        - These's still time to reply to the [form](https://docs.google.com/forms/d/e/1FAIpQLSd7CSxSs_Q0mxromY-FXzEEWyJdKycniKMBHThycJsDXrA9nQ/viewform?usp=sf_link)
        - 9 responses so far
        - Proposed event names
            - CDEvents Day
            - Events in CI/CD
            - CDEventsCon
            - CDEvents contributor summit
            - How do we achieve Iteroperability in the CNCF ecosystem?
        - Ortelius could prepare a talk about GitOps events
- Updates From Other Groups
    - CDF/TOC
        - [CDFoundation Non-profit Sponsorship Application to HackMD](https://github.com/cdfoundation/foundation/issues/358)
        - Bevy now available at https://community.cd.foundation/cdevents/
        - Zoom client cannot be installed on laptop in a few companies
        - Bevy has a desktop client?
    - SIG Interoperability
    - SIG Best Practices
    - [CDEvents / Vocabulary workstream](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- CD Events Project set up
- Meetups / Conferences
    - Upcoming CDEvents webinar Wed 2nd
        - Needs to be added into the system
        - Content?
        - Advertisement?
        - (Andrea) Is there enough time between now and Wed or shall we reschedule?
        - We should probably reschedule
            - Reach out to Kara (Andrea)
- Blog Posts / Links
- Presentations
- PR / Design Discussions

- [Direktiv](https://docs.direktiv.io/v0.6.0/getting_started/events/)
    - They have a lot of events that might help CDEvents leaping forward
    - Steve to check if it possible to get a presentation

- Eiffel community is working on describing the entire git graph through events
    - It might be helpful for CDEvents
    - It would be great to have a presentation
    - It requires event persistence
        - Eiffel@Ericsson uses MongoDB to store events
        - Graph database: [ArangoDB](https://arangodb.com)

- [Metrics PoC](https://hackmd.io/cFw0xO6XSwGwIoJ78Re91A)
    - What can we reuse from 4-keys
    - Switch big-query for a different backend?
    - Re-use some of the queries?
    - Visualization of relationships:
        - D3S (js)
        - Vis.js
        - Hardest part is getting the data in the right (JSON) format

- Releases
    - In eiffel to group events togther: 
        - https://github.com/eiffel-community/eiffel/blob/master/eiffel-vocabulary/EiffelFlowContextDefinedEvent.md


## Meeting February 14th

Meeeting time in your timezone [here](https://time.is/4pm_14_February_2022_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Erik Sternerson, doWhile
- Steve Taylor, DeployHub
- Kevin Chu, GitLab
- Mattias Linnér, Ericsson
- Oleg Nenashev, Dynatrace, Keptn/Jenkins
- Kara de la Marck, CDF


### Agenda

- Action Items From Last meeting
    - Kara / Oleg: online meetup
        - Should be planned as soon as possible after the project is officially announced, to be announced before KubeCon
        - We'll try to arrange something a few weeks after FOSDEM. Time around 3 or 4 PM UTC. Duration 90 minutes.
        - Doodle poll: https://doodle.com/poll/34u9znbdmnayx7k4?utm_source=poll&utm_medium=link
    - Kara: contributor summit at cdCon - [Slack msg](https://cdeliveryfdn.slack.com/archives/C030SKZ0F4K/p1643901186679179)
         * For each Contributor Summit, we should normally be able to provide breakfast, snacks at breaks, and beverages throughout the day. :coffee: :croissant:  
        - So it is important for CDF to have an estimate of number of attendees
        - [Form](https://docs.google.com/forms/d/e/1FAIpQLSfJGFdQhFAbScBt3y3sW0WCQMl903X9T3ycmix0PpB2Zq8aoQ/viewform) to be filled in (deadline Feb 28th) - Andrea will fill it in once the below form is filled in
        - Related [form](https://docs.google.com/forms/d/e/1FAIpQLSeZGMEGldtkrj1viMcT9nQgAgHZVkIlzQZ45aTvN74992Mawg/viewform?usp=sf_link) from Andrea - gathering partecipants and ideas. All: Please answer this form
            - Only three responses so far
    - Emil & Andrea / all: [PR document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit?usp=sharing) - deadline? No fixed deadline. Can be shaped as the project evolves, SDKs? Also, currently need to source more supporting quotes.
        - Erik: Provide supporting quote from doWhile/VCC
    - Kara to check CDF/LF possibilities for a "Marketing clip" for CDEvents, to apply for being part of the $10K CDF/TOC budget. E.g. 30-90 second video presenting the project and the spec in an appealing manner.
        - (Kara) It should be possible. Kara will check what they can do and come back to us.
        
- Meetups / Conferences
* Kara: KubeCon update - CDEvents will have a full day co-located at KubeCon EU 🥳 
    - Please respond with your ideas for the full day event here: https://github.com/cdfoundation/sig-events/issues/115
    - We need to title the event and form a schedule to promote the full day event.
    - This will be officially announced soon. If you will be at [KubeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) please do join us!
- Presentations
    - Presentation from Fidelity (Jamie Plower) on their use of Cloudevents in CI/CD pipelines

- Updates From Other Groups
    - CDF/TOC
        - SIG Software Supply Chain seems to become approved
        - CDF Ambassadors. Representation of the CDEvents
        - Waiting for Bevy rollout. Needs to happen first
        - Migrate CDF members from Meetup to Bevy.
        - Kara: CDF cannot port the CDF meetup members to Bevy, but we can do an email blast to invite them through Meetup Pro. Roxanne will send this email.
        - CDF Community Builders will be announced soon.
    - SIG Interoperability
        - At our last meeting, Justin Abrahms raised the idea that the Interoperability SIG create a catalogue of quality gates.
        - GH Discussion: https://github.com/cdfoundation/sig-interoperability/discussions/83
        - This Thursday, 17 February, Anders Eknert, from Styra, will join us to discuss quality gates. This discussion will be a great way to further our catalogue of quality gates.
        - Note: the term "quality gates" probably comes from Keptn. It might not be the final name for it. Keptn is not happy about this term and want to change
    - SIG Best Practices
    - [Vocabulary workstream](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- CD Events Project set up
    - [Bootstrap Governance](https://github.com/cdevents/spec/blob/main/governance.md) to-do list:
        - [Contribution Process](https://github.com/cdevents/spec/blob/main/governance.md#contribution-process)
            - Andrea will draft a first version of it. Everyone should help out contributing to it.
        - [Project Charter](https://github.com/cdevents/spec/blob/main/governance.md#initial-charter)
            - (Andrea) One volunteer to draft (hackmd), discussion/approval in meeting?
        - [PR document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit?usp=sharing)
    - Review [0.1 project](https://github.com/orgs/cdevents/projects/1/views/1)
        - Event examples should probably be added
        - Swagger API to be provided? AsynchAPI/OpenAPI?
    - CDEvents / Keptn alignment. Keptn Events Spec is wider and includes app lifecycle events too, future is to be seen (new standard?). Oleg to follow-up
    - ON: Get CloudEvents folks participating too?
- Blog Posts / Links
- Presentations
- PR / Design Discussions
    - https://github.com/cdevents/spec/pull/26


## Meeting January 31st

Meeeting time in your timezone [here](https://time.is/4pm_31_January_2022_in_UTC). You're welcome to join!

Participants:
- Steve Taylor (DeployHub/Ortelius)
- Mattias Linnér (Ericsson)
- Erik Sternerson (doWhile)
- Kara de la Marck (CDF)
- Oleg Nenashev (Dynatrace/Jenkins/Keptn)
- Emil Bäckmark (Ericsson)

Agenda:
- Action Items From Last meeting
    - Andrea: `events.cd` domain
        - We have cdevents.dev now for us. 
        - Is this good enough or do we need `events.cd` ?
            - Example https://screwdriver.cd/ 
        - https://github.com/cdevents/spec/issues/21 
        - Let's use cdevents.dev instead. It's already in place :)
        - Action closed
    - Andrea: check [tweetdeck](https://tweetdeck.twitter.com/)
        - I setup a team for `_cdevents`
        - Anyone in the team can impersonate `_cdevents` with no need for pwd sharing
        - I would propose to have folks in the cdevents governance in the team (if they have a twitter account)
        - Everyone in the governance team should be invited to admin the tweetdeck team
        - Action closed
    - Andrea: prepare for the roadmap discussion
        - [PR](https://github.com/cdevents/spec/pull/9)
        - Then intention to release v 0.1 should be added
        - We should during the year sort out what is lacking to get to 1.0
        - PR can be merged once the above is added
        - Action closed
    - Andrea: move the governance docs to the .github repo
        - Only a [few documents](https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/creating-a-default-community-health-file#supported-file-types) can be added to `.github` and be shown in all repos
        - I would leave the doc in the `spec` repo where it is today
        - Agreed. Let's keep it in spec repo
        - Action closed
    - Kara / Oleg: online meetup
        - Should be planned as soon as possible after the project is officially announced, to be announced before KubeCon
        - We'll try to arrange something a few weeks after FOSDEM. Time around 3 or 4 PM UTC. Duration 90 minutes.
        - Doodle poll: https://doodle.com/poll/34u9znbdmnayx7k4?utm_source=poll&utm_medium=link
    - Kara: contributor summit at cdCon
        - Do we want to have it?
    - Emil & Andrea / all: [PR document](https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit?usp=sharing)
        - Erik: Provide supporting quote from doWhile/VCC
    - Emil: archive meeting minutes from 2021
        - [PR](https://github.com/cdfoundation/sig-events/pull/109)


- Updates From Other Groups
    - CDF/TOC
        - Budget exists if we want to request some
        - $10K. Could e.g. be used to fund an intern for ~3 months, but we'd need to have a task and a mentor.
            - SDK work? Erik could mentor a person working on a Python SDK, but it might not make sense at this stage of the protocol evolution
            - Webpage work?
            - "Marketing clip" for CDEvents? E.g. 30-90 second video presenting the project and the spec in an appealing manner
                - Kara to check CDF/LF possibilities
            - SWAG :)
    - SIG Interoperability
        - Fatih is stepping down from chairing. There is a need for a new chair.
        - PR about pipeline stages and steps was discussed. It will/should affect the naming of our events and properties in them
            - https://github.com/cdevents/spec/issues/18 
    - SIG Best Practices
        - .

- FOSDEM poll
    - .

- CD Events Project set up
    - We have a logo!
        > ![cdevents logo](https://github.com/cdfoundation/artwork/raw/main/cdevents/horizontal/color/cdevents_horizontal-color.png)

    - Website https://cdevents.dev :grinning_face_with_star_eyes:
        - Currently using google/docsy, a hugo theme (K8s, Tekton, Kubeflow, Spinnaker)
        - Alternatives
            - Nuxt JS (Sigstore)
            - Other hugo themes?
            - Jekyll (screwdriver.cd)
            - [mkdocs-material](https://squidfunk.github.io/mkdocs-material/) (Knative)
        - (Andrea) Any web designer in the team?
            - Fix color scheme / contrast in links
        - What content do we want on the cover page?
            - Diagrams of key use cases
            - Cards with projects involved (Tekton, Knative, Keptn, Jenkins)
            - Three Cards with highlights / benefits from cdevents
            - Companies involved?
            - Else?
        - Andrea working on adding the spec published on the website
            - https://staging.cdevents.dev/

    - CD Events should be listed as an incubating project on CDF homepage menu and https://cd.foundation/projects/
        - https://github.com/cdfoundation/foundation/issues/353
        - (Andrea) Would be nice to have it in time for the FOSDEM talk on Feb 6th

    - [Bootstrap Governance](https://github.com/cdevents/spec/blob/main/governance.md) to-do list:
        - [Contribution Process](https://github.com/cdevents/spec/blob/main/governance.md#contribution-process)
        - [Project Charter](https://github.com/cdevents/spec/blob/main/governance.md#initial-charter)
        - (Andrea) One volunteer to draft (hackmd), discussion/approval in meeting?

    - Questionnaire for FOSDEM
        - Erik or Andrea to present current proposal and get feedback.

- End user presentations
    - Jamie Plower from Fidelity can present their CI/CD with Cloudevents on next meeting
    - Hopefully on next meeting (Feb 14th)



## Meeting January 17th

Meeeting time in your timezone [here](https://time.is/4pm_17_January_2022_in_UTC). You're welcome to join!

Participants:
- Emil Bäckmark, Ericsson
- Andrea Frittoli, IBM
- Mattias Linnér, Ericsson
- Steve Taylor, DeployHub
- Kara de la Marck, CDF
- Tracy Ragan, DeployHub/Ortelius
- Erik Sternerson, doWhile
- <\add your name\>

### Agenda
- Action Items From Last meeting
    - Andrea to ask CDF if they can buy the domain events.cd for us
        - Done, waiting for response
    - Steve to share his ideas on SPDX - How should we relate to the SPDX standard? What are the next steps needed?
        - After the project proposal
    - (A)Tracy to reach out to CD Foundation online meetup organizers (CDF ambassadors) to see who is driving this for next year, so that we can get on to the calendar to present our work there
        - Tacy reached out to Garima, waiting for response
        - Moving to new platform Bevy early 2022
        - Oleg is happy to host some meetups
        - (A) Oleg to reach out to Garima
        - Unless we can join the online meetup we could self-host such a meetup
        - Should be about an hour, including 30 minutes presentation
        - 3rd week of January, in the US/EMEA timezone
        - (A) Oleg to create a poll for the timeslot to use
        - (A) Andrea will provide the autumn-21 Kubecon presentation as a base
        - (?) newsletter
        - Tracy will meet with Garima on Friday 21st to see how to proceed
    - Andrea to check [tweetdeck](https://tweetdeck.twitter.com/)
        - https://twitter.com/_cdevents
    - ~~Emil to consolidate the latest comments on the~~ [CD Events artwork PR](https://github.com/cdfoundation/foundation/issues/348)
        - ~~Done, waiting for the creative team~~
    - ~~Steve to take Code of Conduct v2.1 into the `spec` repo~~
    - ~~Andrea to add a governance.md into the `spec` repo~~
        - https://github.com/cdevents/spec/pull/4
        - ~~Andrea to address some comments~~
    - (A) Andrea to move the governance docs to the .github repo
    - (A) Andrea can prepare for the roadmap discussion
    - Emil & Andrea to start filling out the CD Events Public Relations doc: https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit?usp=sharing
    - **All**: to add quotes to the CD Events Public Relations document from the organization we belong to. At least add a note on if a quote is to be contributed
    - Emil to clone the SIG Events MoM to a 2022 document and archive the old one

- Updates From Other Groups
    - CDF/TOC
        - cdCon CFP is open: https://events.linuxfoundation.org/cdcon/program/cfp/#overview
        - Suggestion from Kara is to submit a talk even if we don’t know if we’re allowed to travel yet
    - SIG Interoperability
        - Next meeting this Thursday
        - Fatih will step down from being co-chair in the SIG
    - SIG Best Practices
    - [Vocabulary workstream](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
        - When creating the FOSDEM presentation it would have helped to have more real world use cases for events in CI/CD. In the presentation there will be a shout out to request such input from the audience
            - https://github.com/cdevents/presentations/blob/main/fosdem2022/talk-to-cdevents.md
- CD Events Project set up
    - CD Events should be listed as an incubating project on CDF homepage menu and https://cd.foundation/projects/
        - https://github.com/cdfoundation/foundation/issues/353
    - CD Events icon progressed: https://github.com/cdfoundation/foundation/issues/348
    - Where to create issues for requesting new repos under cdevents, or request changes to the github organization?
        - Info added to top github cdevents organization page 
    - Use Cases for CD Events updated: https://hackmd.io/ZCS2KYKZTpKBqhU9PMuCew
    - Work on the roadmap - https://hackmd.io/ja_hYsh8RHCkGb_LlA-REA
        - Proposed to be discussed during next SIG meeting
        - (A) Andrea can prepare for it
    - CD Events Public Relations doc: https://docs.google.com/document/d/1asktguMF_K4N5Vugn0EQ_dCyCgxZX69o2D02MRGzNjA/edit?usp=sharing
        - (A) Emil & Andrea: to start filling it out
        - (A) **All**: to add quotes to that document from the organization we belong to. At least add a note on if a quote is to be contributed
- Archive SIG Events MoMs from 2021?
    - (A) Emil to clone this to a 2022 document and archive the old one
- Meetups / Conferences
    - cdCon-22 BoF or panel online/virtual or just onsite?
        - At this time, we are planning all speakers to be in person. 
        - BoF/panel could probably be combined virtual/onsite
        - CFP:  https://events.linuxfoundation.org/cdcon/program/cfp/#overview
            - Suggestion from Kara is to submit a talk even if we don't know if we're allowed to travel yet
    - Potential for a half day event focused on CD Events at KubeCon Europe 2022
        - Great opportunity to promote CD Events -- let's start to form a schedule for the half day
        - The half day community event will likely be free to attend and independent of KubeCon ticket.



***Previous meeting notes can be found on [GitHub](https://github.com/cdfoundation/sig-events/tree/main/docs)***