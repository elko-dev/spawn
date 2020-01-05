# spawn

Spawn creates project scaffolding, integrating GitLab with Heroku.

## Requirements

- macOS
  - [Homebrew](#https://brew.sh/)
  - [Heroku CLI](#https://devcenter.heroku.com/articles/heroku-cli) - `brew tap heroku/brew && brew install heroku`

## Setup

Configure user, GitLab and Heroku information before using Spawn:

**User:** `spawn setup --user` sets user name, email and organization information.

**GitLab:** [Generate an API token](#https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) and copy it to the clipboard. Run `spawn setup --gitlab` and paste the API token when prompted.

**Heroku:** Setup token access by first running `heroku login.` Run `spawn setup --heroku` and enter your
user details.

## Workflow

Ex. `spawn init newproject --node --gitlab --heroku`

### PreFlight

- check for name = `checkUserName`
- check for email = `checkUserEmail`
- check for GitLab token = `checkTokenGitLab`
- check for Heroku token = `checkTokenHeroku`

### User Input

- prompt for project name = `validateProjName`

### Flow

- initialize Git repo = `initLocal`
- create README = `createReadme`
- create .gitignore = `createGitgnore`
- initialize new Node project = `initNode`
- initialize GitLab repo = `initGitLab`
- initialize Herok = `initHeroku`


### Usage
``` bash
export GOPATH=$HOME/go && export PATH=$GOPATH/bin:$PATH
./spawn cr --projectname=test
```