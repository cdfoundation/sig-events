# Source Code Version Control Events

__Note:__ This is a work-in-progress draft version and is being worked on by members of the Events SIG. You are very welcome to join the work and the discussions!

This phase includes events related to changes in Source Code repositories that are relevant from a Continuous Delivery perspective.

## Repository Events

These events are related to Source Code repositories
- **Repository Created Event**: a new Source Code Repository was created to host source code for a project
- **Repository Modified Event**: a Source Code Repository modified some of its attributes, like location, or owner
- **Repository Deleted Event**: a Source Code Repository was deleted and it is not longer available
- **Branch Created Event**: a Branch inside the Repository was created 
- **Branch Deleted Event**: a Branch inside the Repository was deleted



Repository Events MUST include the following attributes:
- **Event Type**: the type is restricted to include `cd.**` prefix. For example `cd.repository.created` or `cd.repository.changeproposalreviewed`
- **Repository URL**: indicates the location of the source code repository for API operations, this URL needs to include the protocol used to connect to the repository. For example `git://` , `ssh://`, `https://`
- **Repository Name**: friendly name to list this repository to users

Optional attributes: 
- **Repository Owner**: indicates who is the owner of the repository, usually a `user` or an `organization`
- **Repository View URL**: URL to access the repository from a user web browser

## Change proposal events

From each repository you can emit events related with proposed source code changes. Each change can include a single or multiple commits that can also be tracked.

We use the term "change proposal" to represent one or more commits that form a proposal for a change to the source (e.g. [a Pull Request in GitHub, a Merge Request in GitLab or a Change in Gerrit](https://github.com/cdfoundation/sig-interoperability/blob/master/docs/vocabulary.md#scm-tools-and-technologies).), which becomes a change through approval and merge.

- **Change Proposal Created Event**: a source code change proposal was created and submitted to a repository specific branch.
- **Change Proposal Reviewed Event**: someone (user) or an automated system submitted a review to the source code change proposal. A user or an automated system needs to be in charge of understanding how many approvals/rejections are needed for this change proposal to be merged or rejected. The review event needs to include if the change proposal is approved by the reviewer, more changes are needed or if the change proposal is rejected.
- **Change Proposal Abandoned Event**: a tool or a user decides that the change proposal should be abandoned, for instance because it has been inactive for a while.
- **Change Proposal Updated Event**: the change proposal has been updated, for example a new commit is added or removed from an existing change proposal.
- **Change Proposal Merged Event**: the change proposal is merged to its target branch.


Optional attributes for **Change Proposal** Events: 
- (TBD)
