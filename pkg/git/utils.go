package git

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// GenerateBranchName generates branch name from issue key and summary
func GenerateBranchName(issueKey string, summary string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return "", errors.New("Failed to create regexp for branch name")
	}

	branchName := reg.ReplaceAllString(summary, "")
	reg, err = regexp.Compile("[ ]")
	if err != nil {
		return "", errors.New("Failed to create regexp for branch name with underscores")
	}

	branchName = strings.ToLower(reg.ReplaceAllString(strings.TrimSpace(branchName), "_"))

	return fmt.Sprintf("%s/%s", issueKey, branchName), nil
}
