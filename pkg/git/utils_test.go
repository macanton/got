package git

import (
	"testing"
)

func TestGenerateBranchName_WithoutJiraIssueKeys(t *testing.T) {
	branchName, err := GenerateBranchName([]string{}, "Test string  data")

	if err != nil {
		t.Errorf("GenerateBranchName without Jira issue keys returned error %+v", err.Error())
	}

	expectedBranchName := "test_string__data"
	if branchName != expectedBranchName {
		t.Errorf("GenerateBranchName without Jira issue keys returned %+v, want %+v", branchName, expectedBranchName)
	}
}

func TestGenerateBranchName_WithSeveralJiraIssueKeys(t *testing.T) {
	branchName, err := GenerateBranchName([]string{"PC-1234", "PC-345"}, "Test string  data")

	if err != nil {
		t.Errorf("GenerateBranchName without Jira issue keys returned error %+v", err.Error())
	}

	expectedBranchName := "PC-1234/PC-345/test_string__data"
	if branchName != expectedBranchName {
		t.Errorf("GenerateBranchName without Jira issue keys returned %+v, want %+v", branchName, expectedBranchName)
	}
}

func TestPrependIssueKeysToBranchName_WithoutJiraIssueKeys(t *testing.T) {
	branchName, err := PrependIssueKeysToBranchName([]string{}, "test_branch_name")

	if err != nil {
		t.Errorf("GenerateBranchName without Jira issue keys returned error %+v", err.Error())
	}

	expectedBranchName := "test_branch_name"
	if branchName != expectedBranchName {
		t.Errorf("GenerateBranchName without Jira issue keys returned %+v, want %+v", branchName, expectedBranchName)
	}
}

func TestPrependIssueKeysToBranchName_WithJiraIssueKeys(t *testing.T) {
	branchName, err := PrependIssueKeysToBranchName([]string{"PC-123", "PC-345"}, "test_branch_name")

	if err != nil {
		t.Errorf("PrependIssueKeysToBranchName without Jira issue keys returned error %+v", err.Error())
	}

	expectedBranchName := "PC-123/PC-345/test_branch_name"
	if branchName != expectedBranchName {
		t.Errorf(
			"PrependIssueKeysToBranchName without Jira issue keys returned %+v, want %+v",
			branchName,
			expectedBranchName,
		)
	}
}

func TestRemoveIssueKeysFromBranchName_WithoutJiraIssueKeys(t *testing.T) {

}

func TestRemoveIssueKeysFromBranchName_WithJiraIssueKeys(t *testing.T) {

}

func TestRemoveIssueKeysFromBranchName_WithFirstJiraIssueKey(t *testing.T) {

}
func TestRemoveIssueKeysFromBranchName_WithLastJiraIssueKey(t *testing.T) {

}

func TestRemoveIssueKeysFromBranchName_WithDuplicateJiraIssueKey(t *testing.T) {

}
