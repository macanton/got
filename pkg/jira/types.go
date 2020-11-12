package jira

import (
	"regexp"
)

// Issue is a struct for Jira issue
type Issue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
	RenderedFields struct {
		Description string `json:"description"`
	} `json:"renderedFields"`
}

// IssueForm is a struct for Jira Issue form values
type IssueForm struct {
	Fields IssueFormFields `json:"fields"`
}

// IssueFormFields is a type for IssueForm nested structure
type IssueFormFields struct {
	Project   IssueFormProject   `json:"project"`
	Summary   string             `json:"summary"`
	IssueType IssueFormIssueType `json:"issuetype"`
}

// IssueFormProject is a type for IssueForm nested structure
type IssueFormProject struct {
	Key string `json:"key"`
}

// IssueFormIssueType is a type for IssueForm nested structure
type IssueFormIssueType struct {
	Name string `json:"name"`
}

// IssueCreationResponse type for response on Jira issue creation request
type IssueCreationResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

// GetStrippedDescription returns issues descriptions without html tags
func (issue Issue) GetStrippedDescription() string {
	reg := regexp.MustCompile("<.*?>")
	return reg.ReplaceAllString(issue.RenderedFields.Description, "")
}
