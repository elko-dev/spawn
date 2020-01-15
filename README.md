# spawn

Spawn creates project scaffolding, integrating GitLab with Heroku.

## Requirements

**GitLab:** [Generate an API token](#https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) and copy it to the clipboard.

**Heroku:** Setup token access by first running `heroku login.` 



### Usage
``` bash
export GOPATH=$HOME/go && export PATH=$GOPATH/bin:$PATH
./spawn application
```