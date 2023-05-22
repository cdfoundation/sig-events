---
tags: SIG Events
---

# CDF Events SIG Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g/badge)](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g)

This SIG meets bi-weekly on Mondays at [3pm UTC](https://time.is/3pm_in_UTC) during summer time and at [4pm UTC](https://time.is/4pm_in_UTC) during winter time). The meeting can be accessed through [this zoom link](https://zoom.us/j/98408983891?pwd=VXBxMjJTaThGRGFWRXFmdUxsRUZUdz09). For older meetings please see our [playlist on YouTube](https://www.youtube.com/playlist?list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC)

The SIG was initiated as a [workstream under SIG Interoperability ]( https://github.com/cdfoundation/sig-interoperability/blob/master/docs/meetings_2020.md#May-28-2020) and its first meeting was held on June 8th 2020.

Meeting notes for the workstream are managed on HackMD [here](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g), and published to GitHub [here](https://github.com/cdfoundation/sig-events/blob/main/docs/meetings.md).

## Future Presentations/Meetings

### TOC Project Nominations
- Nominations to be provided between May 22 and June 5. See [process](https://github.com/cdfoundation/toc/pull/202). Voters registration open until May 26.
- Are you on [the voters list](https://protect2.fireeye.com/v1/url?k=31323334-501d5122-313273af-454445555731-4b8c207147c35781&q=1&e=14a79fae-dd7b-404d-85c3-fd3782645f4d&u=https%3A%2F%2Fgithub.com%2Fcdfoundation%2Ftoc%2Fblob%2Ff4accdde9248981c80507b982f4214b098c780d6%2Felections%2F2023%2Fvoters.md)? If not, and if you consider yourself a contributor to CDF or any of its projects, you could sign up on [this form](https://protect2.fireeye.com/v1/url?k=31323334-501d5122-313273af-454445555731-7c934222755ef6a8&q=1&e=14a79fae-dd7b-404d-85c3-fd3782645f4d&u=https%3A%2F%2Fdocs.google.com%2Fforms%2Fd%2Fe%2F1FAIpQLSdsxCadKwauQ4EHjaNihGtVQJQ027ECsXBZACU7WqXm4dBMCw%2Fviewform).

## May 22

### Participants
* Emil Bäckmark, Ericsson, UTC+2
* \<please add yourself\>

### Agenda

- Action items from previous meetings
    - (Emil) prepare SIG Events/CDEvents project info needed by TOC

- Updates from other CDF Groups
    - CDEvents Workgroup
        - Lot's of good input from cdCon. See [CDEvents minutes](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA?view#May-22-2023)
    - TOC
        - No meeting recently
        - [TOC Project representatives elections](https://github.com/cdfoundation/toc/pull/202/files?short_path=ff37c30#diff-ff37c3064e1c5d7bef2269062e40afd924448d1df50c63cbedcd3f1af1328742)
            - Emil/Andrea to drive this for CDEvents and provide info through Slack/Email

- Interoperability initiative together with VSMI
    - Emil to reach out to Steve to see if he plans to join next SIG meeting or if we should better set up a dedicated meeting about the VSMI collaboration

- PoCs
    - Time to update the old PoC for v0.3? Use same tools?
        - https://github.com/cdevents/community/issues/18
    - Tweak it to include event aggregator to subscribe to? A.k.a. "evidence store"
        - Include Fidelity (Jamie Plower) and SAS (Brett Smith) in that proposal
        - https://github.com/cdevents/community/issues/27
    - Additional PoC with GitHub Actions and some other tool?
        - "Iteratively expanding CI/CD System", for example:
            - Initially just a GitHub Actions pipeline, with a set of consumers (e.g. deployer, visualizer, DORA calculator, raw event viewer)
            - Secondly break out the pipeline execution to a Tekton pipeline, with no change to the consumers
            - Thirdly break out test activities from Tekton pipeline to Testkube, with no change to the consumers
        - https://github.com/cdevents/community/issues/28
    - PoC for DORA metrics initialized by Andrea
        - https://github.com/afrittoli/cdevents-metrics-poc
        - https://github.com/afrittoli/cdeventer
        - https://github.com/cdevents/community/issues/29

- Collaboration with other communities
    - Argo CD
        - Andrea will talk to Michael that he met on cdCon
    - CNCF TAG App Delivery
        - Meet them again on June 7?
        - https://calendar.google.com/calendar/u/0/embed?src=linuxfoundation.org_o5avjlvt2cae9bq7a95emc4740@group.calendar.google.com
        - https://github.com/cncf/tag-app-delivery/issues/397



## May 8

### Participants
* Steve Pereira
* Emil Bäckmark, Ericsson, UTC+2
* \<please add yourself\>

### Agenda

- Action items from last meeting
    - (Emil) prepare SIG Events/CDEvents project info needed by TOC
    - ~~(Andrea) follow up on needed updates of the CDEvents Whitepaper~~

- Updates from other CDF Groups
    - CDEvents Workgroup
        - CDEvents v0.3 released
            - Testing events reiterated and placed in new bucket
            - Artifact signed event
    - TOC
        - Tracy Ragan is now our SIG sponsor
        - [TOC Project representatives elections](https://github.com/cdfoundation/toc/pull/202/files?short_path=ff37c30#diff-ff37c3064e1c5d7bef2269062e40afd924448d1df50c63cbedcd3f1af1328742)
            - CDEvents should nominate a candidate between May 22 and June 5. Whom to nominate will be decided on a coming CDEvents meeting.
            - 4 seats available (from 9 projects)

- CNCF TAG App Delivery
    - We participated on their meeting on May 3
    - CDEvents was discussed to be implemented in [Podtato-head](https://github.com/podtato-head/podtato-head)
    - Collaboration should continue. We should meet again in about a month.

- CDEvents Whitepaper for cdCon
    - To be published on cdCon this week (?)

- Who is currently using CDEvents?
    - We got this question from somewhere - don't know where
    - No one is *actually* using CDEvents in production yet

- Spinnaker implementation of CDEvents
    - What types of events should be mandatory/possible to send?
    - No such list of mandatory events exist at the moment
    - Spinnaker is free to implement support for any type of event types
    - Some should probably be implicitly sent (if configured to do so), like PipelineRun and TaskRun events, others should probably require an explicit "step" in the Spinnaker pipeline config

- Interoperability initiative together with VSMI
    - Dive deeper into the VSMI collaboration & finish reviewing VSMI deck
    - https://docs.google.com/presentation/d/16-c-Kui3LKmiIc54N7f_4tebTt-zcbCnKDypksT5WRo/edit#slide=id.p
    - Where to keep all data? Some data lake will be needed
    - VSM is mostly interested in historical data, to see trends on the value in a system of work, while CDEvents is both interested in historical data (visualization, metrics, etc) AND realtime data for triggering and such
    - There is an overlap between the data for VSM and the data for CI/CD. CI/CD data is probably not a pure subset of VSM data.
    - VSM is more about inspecting how a "system of work" is evolving over time and with new/improved/changed value added to it
    - CI/CD for historical data is more about individual pipelines/workflows and their trends
    - A long discussion was held but not all documented here
    - Discussion to be continued on our next SIG meeting, if possible. Until then we could dig into Steve's presentation and also the joint Charter.


## April 24th

### Participants
* Emil Bäckmark, Ericsson, UTC+2
* Andrea Frittoli, IBM, UTC+1
* Steve Pereira, VSMC, UTC-4
* Tracy Ragan, DeployHub / Ortelius

### Agenda

- Action items from last meeting

- Updates from other CDF Groups
    - CDEvents Workgroup
        - [Releases](https://github.com/cdevents/spec/milestones)
            - v0.3 planned for late April
                - Focus on artifact signing event and improved test events
    - TOC
        - SIG sponsors being appointed
            - Tracy Ragan volunteered to sponsor SIG Events. To be brought up on TOC tomorrow.
        - TOC will probably request from each CDF SIG:
            - Roadmap
            - Quarterly updates
            - End-of-year report
            - We should describe our intended actions for the coming months/year to the TOC
                - Emil to prepare something about this for the next SIG meeting
        - [TOC Project representatives elections](https://github.com/cdfoundation/foundation/blob/main/CDF%20Election%20Process.md#project-representative-4-representatives-elections-managed-by-projects) to be started beginning of May (9 projects for 4 seats)
            - Emil to send a message on Slack/mail to invite everyone in the project to nominate a candidate for CDEvents on the TOC board.
    - CDF Outreach Committee
        - CDF Workshops to be set up
            - CDEvents could probably be showcased there eventually
            - [Planning doc](https://docs.google.com/document/d/1-vFw5ypt3wqNUe-tQvcCxwuiBkIA5Fu8PHXLEr_9prg/edit#)
        - Next meeting this Wednesday. Andrea will try to join or follow up with them if/how we should be involved.
    - CNCF Tag App Delivery
        - Andrea and Emil will join there meeting on May 3rd
        - [Meeting link](https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=NGZidjB1aWZhaDNkMHAzdHB2MHI4aW5kODVfMjAyMzA1MDNUMTUwMDAwWiBhbmRyZWEuZnJpdHRvbGlAbQ&tmsrc=andrea.frittoli%40gmail.com&scp=ALL)
        - Anyone else interested can reach out to Emil or Andrea to get the invite as well

- CDEvents Whitepaper for cdCon
    - Not yet finalized
    - Andrea to follow up with Fatih. At least the urls should be updated.

- Project meeting for CDEvents on cdCon

- Interoperability initiative together with VSMI
    - Maybe form a dedicated workgroup under SIG Events for this purpose?
        - [Working charter draft](https://docs.google.com/document/d/1iX3RJ2KH4h-WM-sltfzatNtDzdMpN0T8nm4IylGhPdY/edit)
    - Meet & greet
    - Scope of the meetings to be defined
    - [Proposed charter](https://docs.google.com/document/d/1iX3RJ2KH4h-WM-sltfzatNtDzdMpN0T8nm4IylGhPdY/edit)
    - Input from TOC meeting:
        - Reference architecture
            - To be clarified what level of overlap exists between VSM architecture and CDF reference architecture
            - Can be viewed from a customer value perspective, what kind of customer requirements are shared across between the two architectures
        - Interoperability is a challenge today - wiring data from the different data sources along the value stream
        - Scope is everything that produce metrics along the value stream
        - Very early stages - evaluating the architecture
        - VSMI currently running through OASIS
    - Steve presented the current VSMI Reference Architecture
        - https://docs.google.com/presentation/d/16-c-Kui3LKmiIc54N7f_4tebTt-zcbCnKDypksT5WRo/edit#slide=id.p
        - We ran out of time so could not complete that walkthrough
    - We should discuss on Slack when to meet next to continue this discussion. Emil to initiate that discussion

- Next SIG meeting is conflicting with cdCon
    - It's at 8am local time Vancouver, just before cdCon starts. Emil to bring up on Slack if the meeting should be held anyway or not.


## March 27th

### Participants
* Emil Bäckmark, Ericsson, UTC+1
* Andrea Frittoli, IBM, UTC
* Jalander Ramagiri, Ericsson Software Technology, UTC

### Agenda

- Action items from last meeting

- Updates from other CDF Groups
    - CDEvents Workgroup
        - [Releases](https://github.com/cdevents/spec/milestones)
            - v0.2 released (all admin not finished yet)
                - Focus on incident events and preparations for full DORA metrics PoC
            - v0.3 planned for late April
                - Focus on connected events, supply chain security and improved test events
    - TOC
        - VSMI TC to guest TOC tomorrow
        - SIG sponsors being appointed
        - TOC will probably request from each CDF SIG:
            - Roadmap
            - Quarterly updates
            - End-of-year report
    - CDF Outreach Committee
        - CDF Workshops to be set up
            - CDEvents could probably be showcased there eventually
            - [Planning doc](https://docs.google.com/document/d/1-vFw5ypt3wqNUe-tQvcCxwuiBkIA5Fu8PHXLEr_9prg/edit#)

- Recent/upcoming Conferences
    - [KubeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) April 18-21, Amsterdam - [Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/schedule/)
    - [Cloud-native Rejekts](https://cloud-native.rejekts.io/)
    - [cdCon](https://events.linuxfoundation.org/cdcon-gitopscon/) & [OSS](https://events.linuxfoundation.org/open-source-summit-north-america/) May, Vancouver - [cdCon schedule](https://events.linuxfoundation.org/cdcon-gitopscon/program/schedule/), [OSS schedule](https://events.linuxfoundation.org/open-source-summit-north-america/program/schedule/)
        - [CDF Projects Demo Theater](https://docs.google.com/spreadsheets/d/1QUEIy54lVBYwm2DMK9dIOLowkmXZlFNz9Ivw4y57c88/edit#gid=0)
        - CDF Project meetings at cdCon?
            - See spreadsheet above. Opportunity to meet in the project. Not yet decided if we should request a slot for CDEvents or not.
    - [Raleigh devops](https://devopsdays.org/events/2023-raleigh/welcome/) - April 12th-13th

- CDEvents integration with Spinnaker
    - RFC approved
    - Implementation started

- CDEvents integration with Jenkins
    - Ongoing internally at Fidelity

- JavaSDK for CDEvents v0.2
    - Might be possible for Adam to look into
    - The Java SDK might be able to reuse some parts of the Go SDK for example-event handling and test cases. There is a component that generate the event specifiic parts of the Go SDK from schema files, and it could potentially generate parts of the Java SDK as well.
        - A first version of that code generation tool is available as a PR on the Go SDK: https://github.com/cdevents/sdk-go/pull/43/files
            - The idea is to create an event.<language>.tmpl file per SDK language
    - The Java SDK should also be properly released on GitHub for 0.2


## March 13th

### Participants
 * Emil Bäckmark, Ericsson, UTC+1
 * Jalander Ramagiri, Ericsson Software Technology, UTC
 * Evan Elms, Fidelity Investments, UTC-6

### Agenda

- Action Items From Last meeting

- Updates From other CDF Groups
    - CDEvents Workgroup
        - v0.2
            - [Release planned for March 20th](https://github.com/cdevents/spec/milestones)
            - Focus on incident events and preparations for full DORA metrics PoC
        - v0.3
            - [Release planned for late April](https://github.com/cdevents/spec/milestones) (at least before cdCon)
            - Focus on connected events and supply chain events
    - TOC
        - CDF Projects health monitoring being discussed
            - [LFX Insights for CDEvents](https://insights-v2.lfx.linuxfoundation.org/cdevents/trends)
            - [Workgroup for LFX](https://github.com/cdfoundation/wg-lfx-requirements) initiated
        - OKRs (Objectives and Key Results) [for CDF are being defined](https://github.com/cdfoundation/toc/issues?q=is%3Aissue+label%3Aokr)
    - SIG Supply chain
    - Other SIGs
    - Other

- Recent/upcoming Conferences
    - [KubeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) April 18-21, Amsterdam - [Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/schedule/)
    - [Cloud-native Rejekts](https://cloud-native.rejekts.io/)
    - [cdCon](https://events.linuxfoundation.org/cdcon-gitopscon/) & [OSS](https://events.linuxfoundation.org/open-source-summit-north-america/) May, Vancouver - [cdCon schedule](https://events.linuxfoundation.org/cdcon-gitopscon/program/schedule/), OSS schedule to be announced tomorrow
    - [Raleigh devops](https://devopsdays.org/events/2023-raleigh/welcome/) - April 12th-13th

- CDEvents integration with Spinnaker
    - RFC awaiting approval before implementation can start



## February 27th

### Participants

* Fatih Degirmenci, CDF, UTC+1
* Emil Bäckmark, Ericsson, UTC+1
* Andrea Frittoli, IBM, UTC
* Jalander Ramagiri, Ericsson Software Technology, UTC

### Agenda

- Action Items From Last meeting

- Updates From other CDF Groups
    - CDEvents Workgroup
        - v0.2
            - Release postponed til mid March
            - Focus on incident events
        - v0.3
            - Release planned before cdCon
            - Focus on connected events and supply chain events
    - TOC
        - CDF Project presentations to the TOC are being planned
    - SIG Supply chain
        - Brett to fill us in on discussions from SIG SSC
            - Intregrations of SBOMs
                - The full SBOM is probably not relevant to include in events payload
                - SAS has such references in their system
                    - (Watch the recording for more info on that setup)
            - CVEs
                - Audit of SBOMs
                - Audit of Scan results
            - Signed Provenance
                - Proof of generation
            - Attestations
                - Proof of generation
    - Other SIGs
    - Other
        - (Tekton/Andrea) For supply chain events and properties there is an idea to relate it to how Tekton Chain handles these aspects

- [CDF Awards 2023](https://cd.foundation/cdf-community-awards/)

- Recent/upcoming Conferences
    - [KubeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) April 18-21, Amsterdam - [Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/schedule/)
    - [Cloud-native Rejekts](https://cloud-native.rejekts.io/)
    - [cdCon](https://events.linuxfoundation.org/cdcon-gitopscon/) & [OSS](https://events.linuxfoundation.org/open-source-summit-north-america/) May, Vancouver - CFP closed. Schedule to be announced on March 6th
    - (Brett) [Raleigh devops](https://devopsdays.org/events/2023-raleigh/welcome/) - April 12th-13th

- Spinnaker RFC
    - No updates so far


## February 13th

### Participants

- Emil Bäckmark, Ericsson, UTC+1
- Andrea Frittoli, IBM, UTC
- Jalander Ramagiri, Ericsson Software Technology, UTC
- Bruno Lopes
- Adam Kenihan

### Agenda

- Action Items From Last meeting
    - None

- Presentations
    - Testkube - Bruno Lopes

- Updates From other CDF Groups
    - CDEvents Workgroup
        - [Incident events](https://github.com/cdevents/spec/pull/107)
        - [Connecting Events](https://hackmd.io/-Or6hobHSLWVj4duAWX7nA?view)
        - [Testing events](https://github.com/cdevents/spec/pull/105)
    - TOC
        - Robert wanted us to aim for implementing CDEvents in the other CDF projects. E.g. through GSoC initiatives.
    - SIGs
    - Other
        - User Story from Fidelity published - https://cd.foundation/blog/2023/02/07/launching-cdf-user-stories-first-up-fidelity-investments/

- CDF Awards 2023
    - Each project should decide on their Most Valuable Contributor
    - https://github.com/cdevents/community/issues/20
    - [Guidelines](https://github.com/cdfoundation/foundation/blob/main/CDF%20Awards%20Guidelines.md)
    - Timeline
        - Nominations open: Wednesday, February 8
        - Nominations close: Friday, March 3
        - Voting opens: Wednesday, March 8
        - Voting closes: Tuesday, March 28
        - Winners announced at cdCon: May 8 – 9, 2023

- Recent/upcoming Conferences
    - FOSDEM [CI/CD devroom](https://fosdem.org/2023/schedule/track/continuous_integration_and_continuous_deployment/) Feb 4-5
    - [KubeCon EU](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) April 18-21, Amsterdam - [Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/program/schedule/)
    - [Cloud-native Rejekts](https://cloud-native.rejekts.io/)
    - [cdCon](https://events.linuxfoundation.org/cdcon-gitopscon/) & [OSS](https://events.linuxfoundation.org/open-source-summit-north-america/) May, Vancouver - CFP closed

- Adam: Java SDK for CDEvents
    - Some linting errors remain
    - Will be merged soon regardless, suppressing some of those errors

- Jalander: Proposal for cdCon on general experience on integrating tools using CDEvents

- Jalander: Visualizations of CDEvents being discussed
    - Will look at Eiffel visualizations as input to this


## January 30th

### Participants

- Emil Bäckmark, Ericsson, UTC+1
- Andrea Frittoli, IBM, UTC
- Jalander Ramagiri, Ericsson Software Technology, UTC

### Agenda

- Action Items From Last meeting

- Updates From other CDF Groups
    - CDF TOC
        - [Projects benefits](https://github.com/cdfoundation/toc/issues/168) document about to be written
        - Each CDF project should be invited to have a participant on TOC meetings. Emil is representing CDEvents there.
    - CDEvents Workgroup
        - [Incident events](https://github.com/cdevents/spec/pull/107)
        - [Testing events](https://github.com/cdevents/spec/pull/105)
    - SIG Interoperability
        - [Latest Techstrong TV episode in "The CD Pipeline" series](https://techstrong.tv/videos/cd-pipeline/interoperability-challenges-and-solutions-the-cd-pipeline-ep-4) is about Interoperability
        - Value Stream Mapping
            - Meeting with VSM folks on Jan 24th
            - A working group together with [VSM Interoperability](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=vsmi) is intended to be started

- Google Summer of Code
    - [Mentor organization applications](https://opensource.googleblog.com/2023/01/mentor-organization-applications-are-open-for-google-summer-of-code-2023.html) accepted until Feb 7th
    - The Jenkins community is preparing a GSoC application, based on [these ideas](https://www.jenkins.io/projects/gsoc/2023/project-ideas/)
        - Fidelity is probably interested in a CDEvents project with Jenkins
    - Other CDF projects have similar plans?

- Presentations
    - _No presentation planned today_
    - Testkube planned for next meeting

- Upcoming Conferences
    - [cdCon](https://events.linuxfoundation.org/cdcon-gitopscon/) May 8-9th, Vancouver - CFP closes on Feb 10th
    - [Open Source Summit NA](https://events.linuxfoundation.org/open-source-summit-north-america/) May 10-12th, Vancouver - CFP closes on Feb 5th

- Spinnaker updates
    - Meeeting tomorrow about the PR with CDEvents

- Java SDK for CDEvents v0.1 soon to be done
    - https://github.com/cdevents/sdk-java/pull/12
    - It would be valuable to provide validation in the Java SDK to validate the generated events towards the CDEvents schemas. It's not a requirement for an SDK, but good to have. The Go SDK has something along those lines.

- Visualizations of CDEvents
    - DORA Metrics (with incident events) is probably the only ongoing initiative in this area


## January 16th

### Participants
- Emil Bäckmark, Ericsson, UTC+1
- Ole Lensmar, Kubeshop, UTC+1
- Jalander Ramagiri, Ericsson Software Technology, UTC
- Andrea Frittoli, IBM, UTC

### Agenda

- Action Items From Last meeting

- Updates From other CDF Goups
    - CDF TOC
        - New Chair - Andrea Frittoli
    - CDEvents Workgroup
        - Test events & incident events

- Presentations
    - _No presentation planned_

- Upcoming Conferences
    - [cdCon](https://events.linuxfoundation.org/cdcon-gitopscon/) May 8-9th, Vancouver - CFP Early Bird Jan 17th
        - CFP Closes on Feb 10th

- Events for Software Supply Chain
    - Supply Chain Levels ([SLSA](https://slsa.dev/)), [Supply Chain Maturity](https://github.com/cdfoundation/sig-software-supply-chain/blob/main/workstreams/scmm/README.md)
    - CDEvents Issue: [Secure Software Supply Chain Aspects](https://github.com/cdevents/spec/issues/70)

- cdevents.dev Refresh
    - Refresh is done

- CDEvents Spinnaker integration
    - A meeting to be scheduled by Spinnaker community

- Java SDK
    - Soon to be updated to CDEvents v0.1

- Request for presentations
    - Argo?
    - Flux?
    - Tekton Triggers?
    - Kubeshop/Testkube

- GSoC with CDEvents?
    - Jenkins had a GSoC with Cloudevents summer 2021. Time to revisit/renew that?
    - https://summerofcode.withgoogle.com/


***Previous meeting notes can be found on [GitHub](https://github.com/cdfoundation/sig-events/tree/main/docs)***