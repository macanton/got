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
	}
}

func checkoutJiraBranch() {
	issue, err := jira.GetIssue()
	if err != nil {
		printErrorToConsole(err)
		return
	}

	branchName, err := git.FindBranchBySubstring(config.GetIssueKey())
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

	branchName, err = git.GenerateBranchName(issue.Key, issue.Fields.Summary)
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

	branchName, err := git.GenerateBranchName(issueKey, config.Options.Summary)
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
