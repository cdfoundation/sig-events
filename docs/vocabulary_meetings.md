# Vocabulary Discussion Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA/badge)](https://hackmd.io/2FRGlw9fTMmKN1OQUVvguA)

This document contains the notes from the of the Events SIG meetings focused on [vocabulary discussion](https://hackmd.io/lBlDCrL7TvmtNOjxdopJ5g).

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
- [ ] (Mattias) 4keys specific events / mappings
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