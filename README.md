# got
**got** is the CLI tool for linking your branches and Jira issues.
It helps to reduce the amount of time spent on creating/managing git branches that should be linked to Jira issues

## Available flags
`got --help` - list all comands with description

### List of all supported flags
- `got -b XXXX` - creates new git branch with the name generated from Jira issue. If the branch already exists (locally or remotely) then it will switch to it.
- `got -lj XXXX` - links Jira issue to the current branch if not linked already
- `got -uj XXXX` - unlinks Jira issue from the current branch
- `got -cj` - creates a new Jira issue and if it succeeds creates new git branch for it
- `got -m` - modifies Jira issue summary and current branch name
- `got -info` - prints current branch Jira issues info

Example of created branches:
`PC-1234/jira_issue_summary`, where:
- `PC` - project code
- `1234` - code of the Jira issue
- `jira_issue_summary`- Jira issue summary in underscore case

## Tested with
- Git 2.29.1
- Mac OS, 10.15.6
- Go 1.15
- Jira cloud api version 3 (https://YOUR_COMPANY_JIRA_DOMAIN.atlassian.net/rest/api/3)

## Installation
### Home brew
```
brew tap macanton/got
brew install got
```

### From source code
```
brew install go

git clone git@github.com:macanton/got.git
cd got
go build
mv got /usr/local/bin
```

## Setup environment
In order to allow app to connect to Jira you need to setup environment variables (later part of the settings will be moved to config file in order to allow to use app with several git repos/ jira projects)
```
export JIRA_API_KEY=
export JIRA_EMAIL=
export JIRA_API_ENDPOINT=
export JIRA_PROJECT_CODE=
```
