package main

import (
	"fmt"
	"got/pkg/config"
	"got/pkg/git"
	"got/pkg/jira"
)

func main() {
	err := config.InitAndRequestAdditionalData()
	if err != nil {
		printErrorToConsole(err)
		return
	}

	switch config.Options.Operation {
	case config.CheckoutBranch:
		checkoutJiraBranch()
	case config.CheckBranchForNewJiraIssue:
		createJiraTicketAndCheckBranch()
	case config.ModifyBranch:
		modifyBranch()
	case config.PrintInfo:
		printInfo()
	case config.LinkJiraIssueToCurrentBranch:
		linkJiraIssueToCurrentBranch()
	case config.UnlinkJiraIssueFromCurrentBranch:
		unlinkJiraIssueFromCurrentBranch()
	}
}

func checkoutJiraBranch() {
	issue, err := jira.GetIssue(config.GetIssueKey())
	if err != nil {
		printErrorToConsole(err)
		return
	}

	branchName, err := git.FindBranchBySubstring(config.GetIssueKey() + config.Options.IssueBranchSeparator)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	if branchName != "" {
		output, err := git.CheckoutBranch(branchName)
		if err != nil {
			printErrorToConsole(err)
			return
		}

		printInfoToConsole(string(output))
		return
	}

	branchName, err = git.GenerateBranchName([]string{issue.Key}, issue.Fields.Summary)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	output, err := git.CheckoutNewBranch(branchName)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	printInfoToConsole(string(output))
	printJiraIssueData(issue)
}

func createJiraTicketAndCheckBranch() {
	issueKey, err := jira.CreateIssue(config.Options.Summary)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	branchName, err := git.GenerateBranchName([]string{issueKey}, config.Options.Summary)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	output, err := git.CheckoutNewBranch(branchName)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	printInfoToConsole(string(output))
}

func modifyBranch() {
	currentBranchName, err := git.GetCurrentBranchName()
	if err != nil {
		printErrorToConsole(err)
		return
	}

	issueKeys := git.GetIssueKeysFromBranchName(currentBranchName)
	if len(issueKeys) == 0 {
		printErrorToConsole(fmt.Errorf(
			"Branch name '%s' does not contain issue keys with prefix '%s'", currentBranchName, config.GetIssueKeyPrefix(),
		))
		return
	}

	issueKey := issueKeys[0]

	summary, err := jira.UpdateIssueSummary(issueKey, config.Options.Summary)
	if err != nil {
		printErrorToConsole(err)
		return
	}
	printInfoToConsole(fmt.Sprintf("Jira issue summary updated to '%s'", summary))

	newBranchName, err := git.GenerateBranchName(issueKeys, summary)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	output, err := git.UpdateCurrentBranchName(newBranchName)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	printInfoToConsole(string(output))
}

func linkJiraIssueToCurrentBranch() {
	currentBranchName, err := git.GetCurrentBranchName()
	if err != nil {
		printErrorToConsole(err)
		return
	}

	issueKeys := git.GetIssueKeysFromBranchName(currentBranchName)
	issueKey := config.GetIssueKey()
	if stringInSlice(issueKey, issueKeys) {
		printInfoToConsole(fmt.Sprintf("Jira issue %s already linked to the current branch", issueKey))
		return
	}

	updatedBranchName, err := git.PrependIssueKeysToBranchName([]string{issueKey}, currentBranchName)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	output, err := git.UpdateCurrentBranchName(updatedBranchName)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	printInfoToConsole(string(output))
}

func unlinkJiraIssueFromCurrentBranch() {
	currentBranchName, err := git.GetCurrentBranchName()
	if err != nil {
		printErrorToConsole(err)
		return
	}

	issueKey := config.GetIssueKey()
	updatedBranchName := git.RemoveIssueKeysFromBranchName([]string{issueKey}, currentBranchName)

	output, err := git.UpdateCurrentBranchName(updatedBranchName)
	if err != nil {
		printErrorToConsole(err)
		return
	}

	printInfoToConsole(fmt.Sprintf("Jira issue '%s' was unlinked from the current branch", issueKey))
	printInfoToConsole(string(output))
}

func printInfo() {
	currentBranchName, err := git.GetCurrentBranchName()
	if err != nil {
		printErrorToConsole(err)
		return
	}

	issueKeys := git.GetIssueKeysFromBranchName(currentBranchName)
	if len(issueKeys) == 0 {
		printErrorToConsole(fmt.Errorf(
			"Branch name '%s' does not contain issue keys with prefix '%s'", currentBranchName, config.GetIssueKeyPrefix(),
		))
		return
	}

	for i := 0; i < len(issueKeys); i++ {
		issue, err := jira.GetIssue(issueKeys[i])
		if err != nil {
			printErrorToConsole(err)
			continue
		}
		printJiraIssueData(issue)
	}
}

func printInfoToConsole(data string) {
	if len(data) > 0 {
		fmt.Println(data)
	}
}

func printErrorToConsole(err error) {
	fmt.Println(fmt.Sprintf("[ERROR] %s", err.Error()))
}

func printJiraIssueData(issue jira.Issue) {
	fmt.Println(fmt.Sprintf("---------%s---------", issue.Key))
	fmt.Println(fmt.Sprintf("Summary: %s", issue.Fields.Summary))
	fmt.Println("Description:")
	fmt.Println(issue.GetStrippedDescription())
	fmt.Println("--------------------")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
