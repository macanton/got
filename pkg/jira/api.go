package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"got/pkg/config"
	"io/ioutil"
	"net/http"
)

type jiraAPIEndpoint string

const (
	jiraRequestPathGetIssue    jiraAPIEndpoint = "issue/%s?expand=renderedFields"
	jiraRequestPathCreateIssue jiraAPIEndpoint = "issue/"
)

type jiraOperation string

const (
	jiraOperationGetIssue    jiraOperation = "getIssue"
	jiraOperationCreateIssue jiraOperation = "createIssue"
)

type issueTypes string

const (
	storyIssueType issueTypes = "Story"
)

// GetIssue tries to find issue by project code from configuration and by code
func GetIssue() (Issue, error) {
	client := &http.Client{}

	requestURL, err := getRequestURL(jiraOperationGetIssue)
	if err != nil {
		return Issue{}, err
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return Issue{}, err
	}
	setJiraRequestHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return Issue{}, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Issue{}, fmt.Errorf("Failed to parse reponse body: %s", err.Error())
	}

	var issue Issue
	err = json.Unmarshal([]byte(bodyText), &issue)
	if err != nil || issue.Key != config.GetIssueKey() {
		return issue, fmt.Errorf("Jira ticket with key %s not found: %s", config.GetIssueKey(), err)
	}

	return issue, nil
}

// CreateIssue creates Jira issue with specified summary
func CreateIssue(summary string) (string, error) {
	client := &http.Client{}

	requestURL, err := getRequestURL(jiraOperationCreateIssue)
	if err != nil {
		return "", err
	}

	formValues := IssueForm{
		Fields: IssueFormFields{
			Project:   IssueFormProject{Key: config.Options.Jira.ProjectCode},
			IssueType: IssueFormIssueType{Name: string(storyIssueType)},
			Summary:   summary,
		},
	}
	formValuesByte, err := json.Marshal(formValues)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewReader(formValuesByte))
	if err != nil {
		return "", fmt.Errorf("Failed to convert form values to json: %s", err)
	}

	setJiraRequestHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("Failed to create Jira ticket. Status code: %d", resp.StatusCode)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to parse reponse body: %s", err.Error())
	}

	var response IssueCreationResponse
	err = json.Unmarshal([]byte(bodyText), &response)
	if err != nil {
		return "", fmt.Errorf("Jira ticket '%s' was not created: %s\n%s", summary, err, string(bodyText))
	}

	fmt.Println(response)

	return response.Key, nil
}

func getRequestURL(operation jiraOperation) (string, error) {
	switch operation {
	case jiraOperationGetIssue:
		formattedPath := fmt.Sprintf(string(jiraRequestPathGetIssue), config.GetIssueKey())
		return fmt.Sprintf("%s/%s", config.Options.Jira.APIEndPoint, formattedPath), nil
	case jiraOperationCreateIssue:
		return fmt.Sprintf("%s/%s", config.Options.Jira.APIEndPoint, jiraRequestPathCreateIssue), nil
	default:
		return "", fmt.Errorf("Invalid jira operation: %s", operation)
	}
}

func setJiraRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.Options.Jira.Email, config.Options.Jira.APIKey)
}
