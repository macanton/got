# got
**got** is the CLI tool for linking your branches and Jira issues.
It helps to reduce the amount of time spent on creating managing git branches that should be linked to Jira issues

## Available commands
- `got -b XXXX` - creates new git branch with the name generated from Jira issue
- `got -cj` - creates a new Jira issue and if it succeeds creates new git branch for it

## Assumptions
- Your Jira issues have a key in the format of `DD-XXXX`, where `DD` is a project code and `XXXX` is an issue number
- Branch format is `DD-XXXX/ZZZZZ`, where `ZZZZZ` is an issue summary in underscore case
