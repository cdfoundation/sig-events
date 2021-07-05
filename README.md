# CDF Events Special Interest Group

## Quick links

- [Meeting Information](#meetings)
- [Members](#members)
- [New Members](#new-members)
- [Governance](#governance)
- [Communication](#communication)
- [Meetings](#meetings)
- [Specification](vocabulary-draft/README.md)
- [Proof of Concept](poc/README.md)
- [SDK](cde/sdk/go/README.md)

## Group Charter

### Introduction

Today’s CI/CD systems do not talk to each other in a standardized way. This leads to problems related to interoperability, notification of failure issues, and poor automation.

This group is looking at how events can help to create CI/CD systems with a decoupled architecture that is easy to scale and makes it resilient to failures. Using events could also increase automation when connecting workflows from different systems to each other, and as a result empowering tracing/visualizing/auditing of the connected workflows through these events.

### Purpose

The group will focus on the use of events to provide interoperability through topics like:

- When are events suited? For triggers, audits, monitoring, management
  - Common guidelines for at-least-once, at-most-once, exactly once, ordering… When to apply what strategy?
- Best practices for event-driven CI/CD systems
- Events to be used by tools for orchestration/workflows
- Pipeline to pipeline communication via events
- Tracing/auditing/graphing/visualizing of the entire process, e.g., through events. What truly occurs?
- Metrics, e.g., how many versions have been deployed, how many PRs (Pull Requests) have been raised, how many events have been issued?
- How are events related and how are they ordered (links vs trace context)?

### Outcome

The group is working on a standardized event protocol that caters for technology agnostic machine-to-machine communication in CI/CD systems. This specification will be published, reviewed and agreed upon between relevant Linux Foundation projects/members.

The group aims to provide reference implementations such as event listeners and event senders on top of CloudEvents.

### History

This is group started as a work-stream within the CDF SIG Interoperability.
The forming of the workstream was suggested on a [SIG Interoperability meeting](https://github.com/cdfoundation/sig-interoperability/blob/master/docs/meetings_2020.md#May-28-2020) and its first meeting was held on June 8th 2020.
The SIG was formed in February 2021.

## What is an event?

We currently stick with the [definition used by CloudEvents](https://github.com/cloudevents/spec/blob/v1.0/spec.md#terminology):
An *event* is a data record expressing an *occurrence* and its context, where *occurrence* is the capture of a statement of fact during the operation of a software system.

## Members

Current members:

- Ravi Lachhman, [@ravilach](https://github.com/ravilach), Harness
- Andreas Grimmer, [@agrimmer](https://github.com/agrimmer), Dynatrace
- Emil Bäckmark, [@e-backmark-ericsson](https://github.com/e-backmark-ericsson), Ericsson
- Ramin Akhbari, ([@rakhbari](https://github.com/rakhbari)), eBay
- Mattias Linnér, [@m-linner-ericsson](https://github.com/m-linner-ericsson), Ericsson
- Andrea Frittoli, [@afrittoli](https://github.com/afrittoli), IBM
- Mauricio Salatino [@salaboy](https://github.com/salaboy), [Camunda](https://camunda.com) / [LearnK8s](https://learnk8s.io)
- Steve Taylor [@sbtaylor15](https://github.com/sbtaylor15), [DeployHub](https://www.deployhub.com) / [Ortelius OS](https://ortelius.io)
- Tracy Ragan [@tracyragan](https://github.com/tracyragan), [DeployHub](https://www.deployhub.com) / [Ortelius OS](https://ortelius.io)
- Erik Sternerson [@erkist](https://github.com/erkist), doWhile
- Cameron Motevasselani ([@link108](https://github.com/link108)), [Armory](https://www.armory.io/)
- Alois Reitbauer ([@aloisreitbauer](https://github.com/aloisreitbauer)), [Dynatrace](https://www.dynatrace.com/)
- Fredrik Fristedt [@fredjn](https://github.com/fredjn), Axis Communications

### New Members

Membership to this SIG is open to public and self-declared.

New members are advised to:

- Join the [SIG](https://lists.cd.foundation/g/sig-events) and
[CDF TOC](https://lists.cd.foundation/g/cdf-toc) mailing lists.
- Join the [CDF Slack](https://join.slack.com/t/cdeliveryfdn/shared_invite/zt-ao8y4qhd-BQcTUg5l7m0HxXyBvJrT4w%20) and jump into the [#sig-events channel](https://cdeliveryfdn.slack.com/archives/C0151BTKEJX) and introduce yourself.
- Go through the [README.md](README.md) document.
- Regularly join the SIG meetings.
- Submit a PR to add yourself to the [members list](#members).
- Here are various ways to get involved:
  - Share your thoughts by joining the meetings, posting to the mailing list, discussions forum or Slack channel.
  - Add a topic you would like to discuss to [the agenda](docs/meetings.md) of upcoming meeting.
  - Create a [new issue](https://github.com/cdfoundation/sig-events/issues) or start a [discussion](https://github.com/cdfoundation/sig-events/discussions) to start gathering feedback and collaborating.
  - Choose an issue where [help is needed](https://github.com/cdfoundation/sig-events/issues/labels/help%20wanted) and comment on it expressing interest.

## Communication

- Mailing list: [lists.cd.foundation/g/sig-events](https://lists.cd.foundation/g/sig-events).
- GitHub Discussions: [cdfoundation/sig-events/discussions](https://github.com/cdfoundation/sig-events/discussions).
- Slack: #sig-events [slack channel](https://cdeliveryfdn.slack.com/archives/C0151BTKEJX) on the CDF slack.

## Governance

SIG Events is a [CDF Special Interest Group](https://github.com/cdfoundation/toc/tree/master/sigs).

The process SIG Events follows can be seen from [here](https://github.com/cdfoundation/toc/blob/master/GROUPS.md#sigs).

Chairs and the TOC Sponsor of the SIG are

* Emil Bäckmark ([@e-backmark-ericsson](https://github.com/e-backmark-ericsson)), Ericsson - Co-Chair
* Andrea Frittoli ([@afrittoli](https://github.com/afrittoli)), IBM - Co-chair
* Isaac Mosquera ([@imosquera](https://github.com/imosquera)), Armory - TOC Sponsor

## Meetings

SIG Events meets bi-weekly on Mondays at 3pm UTC in the summer time and at 4pm UTC in winter time. (*See your timezone [here](https://time.is/3pm_in_UTC)*).

- [Meeting agenda and minutes](./docs/meetings.md).
- [Zoom Bridge](https://zoom.us/j/97660712600?pwd=Z3BqYTE5YzNsbEhmck16cjdZNEFIUT09).
- [Zoom International dial-in numbers](https://zoom.us/zoomconference).
- [CDF Public Calendar (UTC)](https://calendar.google.com/calendar/u/0/embed?src=linuxfoundation.org_mhf0kmgedn67ihni8r129avp24@group.calendar.google.com&ctz=UTC).
- [Recorded meetings on Youtube](https://www.youtube.com/playlist?list=PL2KXbZ9-EY9RlxWAnAjxs8Azuz11XVhkC)
- [Archived meeting notes (2020)](docs/meetings_2020.md).
