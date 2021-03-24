# Vocabulary Discussion Meeting Notes

[![hackmd-github-sync-badge](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g/badge)](https://hackmd.io/xjK5ujQbTHSaEZjoY28b8g)

This document contains the notes from the of the Events SIG meetings focused on [vocabulary discussion](https://hackmd.io/lBlDCrL7TvmtNOjxdopJ5g).


## Meeting March 23rd
Meeeting time in your timezone [here](https://time.is/4pm_24_March_2021_in_UTC). You're welcome to join!

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

- [ ] (Andrea) Buckets (or phases) of events
- [ ] (Tracy) GitOps specific events
- [ ] (Mattias) 4keys specific events / mappings
- [ ] (Erik) What is a "change", granularity of events (PR, commit)
- [ ] (Steve) DevSecOps specific events and bucket
- [ ] (Tracy) Continuous Tests as its own bucket
- [ ] (Alois) Deployments and versions (e.g. blue/green)
- [ ] (Steve) Environments
- [ ] (Andrea)Lightweigth vs. heavyweigth events (artifacts, logs)
- [ ] (Erik) New [CD Bucket events](https://github.com/cdfoundation/sig-events/pull/23)
- [ ] (Mauricio) Core / base / orchestration events
- [ ] (Alois) Status changed / heartbeat events

Action Points:

- [ ] (Mauricio) PR to link bucket docs to introduction doc
- [ ] (Steve) Add deployment start/finish events
- [ ] (Andrea) Create issue about removing the bucket from the event type