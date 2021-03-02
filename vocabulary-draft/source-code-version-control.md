# Source Code Version Control Events

This phase includes events related to changes in Source Code repositories that are relevant from a Continuous Delivery perspective.


## Repository Events

These events are related to Source Code repositories
- **Repository Created Event**: a new Source Code Repository was created to host source code for a project
- **Repository Modified Event**: a Source Code Repository modified some of its attributes, like location, or owner
- **Repository Deleted Event**: a Source Code Repository was deleted and it is not longer available
- **Branch Created Event**: a Branch inside the Repository was created 
- **Branch Deleted Event**: a Branch inside the Repository was deleted
- **Change Committed**: a change is commited to a branch
- **Change Approved**: a change request is created by a user, that requires review and approval before being applied to an existing branch inside the repository
- **Change Rejected**: a change was reviewed and rejected by a user
- **Change Merged**: a change request has been merged to a branch in an existing repository


Repository Events MUST include the following attributes:
- **Event Type**: the type is restricted to include `CD.Repository**` prefix. For example `CD.Repository.Created` or `CD.Repository.ChangeApproved`
- **Repository URL**: indicates the location of the source code repository for API operations, this URL needs to include the protocol used to connect to the repository. For example `git://` , `ssh://`, `https://`
- **Repository Name**: friendly name to list this repository to users

Optional attributes: 
- **Repository Owner**: indicates who is the owner of the repository, usually a `user` or an `organization`
- **Repository URL View**: URL to access the repository from a user web browser