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

// CreateIssueData is a struct for Jira Issue form values
type CreateIssueData struct {
	Fields CreateIssueDataFields `json:"fields"`
}

// CreateIssueDataFields is a type for IssueForm nested structure
type CreateIssueDataFields struct {
	Project   CreateIssueDataProject   `json:"project"`
	Summary   string                   `json:"summary"`
	IssueType CreateIssueDataIssueType `json:"issuetype"`
}

// CreateIssueDataProject is a type for IssueForm nested structure
type CreateIssueDataProject struct {
	Key string `json:"key"`
}

// CreateIssueDataIssueType is a type for IssueForm nested structure
type CreateIssueDataIssueType struct {
	Name string `json:"name"`
}

// CreateIssueResponse type for response on Jira issue creation request
type CreateIssueResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

// UpdateIssueData is a type for issue update request data
type UpdateIssueData struct {
	Update UpdateIssueDataFields `json:"update"`
}

// UpdateIssueDataFields is a type for issue update data fields
type UpdateIssueDataFields struct {
	Summary []UpdateIssueSummaryFieldOperationData `json:"summary,omitempty"`
	Labels []UpdateIssueLabels `json:"labels,omitempty"`
}

type UpdateIssueLabels struct {
	Add string `json:"add"`
}

// UpdateIssueSummaryFieldOperationData is a type for issue summary set operation
type UpdateIssueSummaryFieldOperationData struct {
	Set string `json:"set"`
}

// GetStrippedDescription returns issues descriptions without html tags
func (issue Issue) GetStrippedDescription() string {
	reg := regexp.MustCompile("<.*?>")
	return reg.ReplaceAllString(issue.RenderedFields.Description, "")
}
