# Events SIG Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g/badge)](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g)

This SIG meets bi-weekly on Mondays at [3pm UTC](https://time.is/3pm_in_UTC) during summer time and at [4pm UTC](https://time.is/4pm_in_UTC) during winter time). The meeting can be accessed through [this zoom link](https://zoom.us/j/97660712600?pwd=Z3BqYTE5YzNsbEhmck16cjdZNEFIUT09). For older meetings please see our [playlist on YouTube](https://www.youtube.com/playlist?list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC)

The SIG was initiated as a [workstream under SIG Interoperability ]( https://github.com/cdfoundation/sig-interoperability/blob/master/docs/meetings_2020.md#May-28-2020) and its first meeting was held on June 8th 2020.

Meeting notes for the workstream are managed on HackMD [here](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g), and published to GitHub [here](https://github.com/cdfoundation/sig-events/blob/main/docs/meetings.md).


## Agenda Template

- Action Items From Last meeting
- Updates From Other Groups
    - CDF/TOC
    - SIG Interoperability
    - SIG Best Practices
    - [Vocabulary workstream](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)
- CD Events Project set up
- Meetups / Conferences
- Blog Posts / Links
- Presentations
- PR / Design Discussions

## Next Meeting (Feb 28th)
- Presentation from Fidelity (Jamie Plower) on their use of Cloudevents in CI/CD pipelines?

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