# got
**got** is the CLI tool for linking your branches and Jira issues.
It helps to reduce the amount of time spent on creating/managing git branches that should be linked to Jira issues

## Available flags
`got --help` - list all comands with description

### List of all supported flags
- `got -b XXXX` - creates new git branch with the name generated from Jira issue. If branch is already present then it will switch to it.
- `got -cj` - creates a new Jira issue and if it succeeds creates new git branch for it

Example of created branches:
`PC-1234/jira_issue_summary`, where:
-  `PC` - project code
- `1234` - code of the Jira issue
- `jira_issue_summary`- Jira issue summary in underscore case

## Assumptions
- Your Jira issues have a key in the format of `DD-XXXX`, where `DD` is a project code and `XXXX` is an issue number
- Branch format is `DD-XXXX/ZZZZZ`, where `ZZZZZ` is an issue summary in underscore case

## Tested with
- Mac OS, 10.15.6
- Go 1.15
- Jira ckoud api version 3 (https://YOUR_COMPANY_JIRA_DOMAIN.atlassian.net/rest/api/3)

## Installation
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
