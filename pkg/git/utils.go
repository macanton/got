package git

import (
	"errors"
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

// PrependIssueKeysToBranchName prepends issue keys to branch name
func PrependIssueKeysToBranchName(issueKeys []string, branchName string) (string, error) {
	branchNameSubstrings := append(issueKeys, branchName)

	return strings.Join(branchNameSubstrings, config.Options.IssueBranchSeparator), nil
}

// GetIssueKeysFromBranchName returns list of Jira issue keys accosiated with current branch
func GetIssueKeysFromBranchName(branchName string) []string {
	substrings := strings.Split(branchName, config.Options.IssueBranchSeparator)

	issueKeyPrefix := config.GetIssueKeyPrefix()
	filterFunc := func(substring string) bool {
		return strings.HasPrefix(substring, issueKeyPrefix)
	}

	return filter(substrings, filterFunc)
}

func filter(stringsArr []string, filterFunc func(string) bool) (filteredArr []string) {
	for _, str := range stringsArr {
		if filterFunc(str) {
			filteredArr = append(filteredArr, str)
		}
	}
	return
}
