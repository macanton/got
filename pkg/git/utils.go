package git

import (
	"errors"
	"fmt"
	"got/pkg/config"
	"regexp"
	"strings"
)

// GenerateBranchName generates branch name for issue keys and summary
func GenerateBranchName(issueKeys []string, summary string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return "", errors.New("Failed to create regexp for branch name")
	}

	branchName := reg.ReplaceAllString(summary, "")
	reg, err = regexp.Compile("[ ]")
	if err != nil {
		return "", errors.New("Failed to create regexp for branch name with underscores")
	}

	filteredSummary := strings.ToLower(reg.ReplaceAllString(strings.TrimSpace(branchName), "_"))
	branchNameSubstrings := append(issueKeys, filteredSummary)

	return strings.Join(branchNameSubstrings, config.Options.IssueBranchSeparator), nil
}

// AddIssueKeysToBranchName add issue keys to branch name
func AddIssueKeysToBranchName(issueKeys []string, branchName string) (string, error) {
	branchNameSubstrings := append(issueKeys, branchName)

	return strings.Join(branchNameSubstrings, config.Options.IssueBranchSeparator), nil
}

// GetIssueKeysFromBranchName returns list of Jira issue keys accosiated with current branch
func GetIssueKeysFromBranchName(branchName string) ([]string, error) {
	substrings := strings.Split(branchName, config.Options.IssueBranchSeparator)
	if len(substrings) == 0 {
		return nil, fmt.Errorf(
			"Branch name '%s' is in wrong format. Is should contain '%s'", branchName, config.Options.IssueBranchSeparator,
		)
	}

	issueKeyPrefix := config.GetIssueKeyPrefix()
	filterFunc := func(substring string) bool {
		return strings.HasPrefix(substring, issueKeyPrefix)
	}

	issueKeys := filter(substrings, filterFunc)
	if len(issueKeys) == 0 {
		return nil, fmt.Errorf(
			"Branch name '%s' does not contain issue keys with prefix '%s'", branchName, issueKeyPrefix,
		)
	}

	return issueKeys, nil
}

func filter(stringsArr []string, filterFunc func(string) bool) (filteredArr []string) {
	for _, str := range stringsArr {
		if filterFunc(str) {
			filteredArr = append(filteredArr, str)
		}
	}
	return
}
