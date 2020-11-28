package config

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// OperationType is a type for enum values of requested by user operation
type OperationType string

// CheckoutBranch is a holder of operation name
const (
	CheckoutBranch               OperationType = "CheckoutBranch"
	ModifyBranch                 OperationType = "ModifyBranch"
	CheckBranchForNewJiraIssue   OperationType = "CheckBranchForNewJiraIssue"
	PrintInfo                    OperationType = "PrintInfo"
	LinkJiraIssueToCurrentBranch OperationType = "LinkJiraIssueToBranch"
)

// OptionsType is a type for stored app configuration
type OptionsType struct {
	IssueCode            int
	Summary              string
	Operation            OperationType
	IssueBranchSeparator string
	Jira                 struct {
		ProjectCode string
		APIEndPoint string
		Email       string
		APIKey      string
	}
}

// Options variable stores app configuration settings
var Options OptionsType = OptionsType{
	IssueBranchSeparator: "/",
}

// InitAndRequestAdditionalData function initializes global configuration of the application
func InitAndRequestAdditionalData() error {
	err := readConfigVariables()
	if err != nil {
		return err
	}

	ticketID := flag.Int("b", 0, "Jira ticket number key for new branch")
	modifyBranch := flag.Bool("m", false, "Update branch name with Jira issue summary")
	createIssue := flag.Bool("cj", false, "Create a new Jira issue and switch to the new branch")
	printIssuesInfo := flag.Bool("info", false, "Print current branch Jira issues information")
	issueCodeForLinking := flag.Int("lj", 0, "Links Jira Issue to current branch")
	flag.Parse()

	if *ticketID < 0 {
		return errors.New("Jira ticket number ket should be more than 0")
	}

	if *ticketID > 0 {
		Options.IssueCode = *ticketID
		Options.Operation = CheckoutBranch
		return nil
	}

	if *issueCodeForLinking > 0 {
		Options.Operation = LinkJiraIssueToCurrentBranch
		Options.IssueCode = *issueCodeForLinking
		return nil
	}

	if *createIssue {
		Options.Operation = CheckBranchForNewJiraIssue
		err := readJiraIssueSummary()
		return err
	}

	if *modifyBranch {
		Options.Operation = ModifyBranch
		err := readJiraIssueSummary()
		return err
	}

	if *printIssuesInfo {
		Options.Operation = PrintInfo
		return nil
	}

	return errors.New("Invalid flags supplied. Cannot determine target operation, use --help")
}

// GetIssueKey returns a key of the Jira issue that contains project code and issue code
func GetIssueKey() string {
	return fmt.Sprintf("%s-%d", Options.Jira.ProjectCode, Options.IssueCode)
}

// GetIssueKeyPrefix returns Jira issue key prefix
func GetIssueKeyPrefix() string {
	return fmt.Sprintf("%s-", Options.Jira.ProjectCode)
}

func readJiraIssueSummary() error {
	fmt.Print("Enter Jira issue summary: ")

	reader := bufio.NewReader(os.Stdin)
	summary, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("Failed to read string from buffer reader: %s", err.Error())
	}

	summary = strings.TrimSpace(summary)
	if len(summary) == 0 {
		return errors.New("Summary cannot be an empty string")
	}

	Options.Summary = summary

	return nil
}

func readConfigVariables() error {
	var err error
	Options.Jira.APIEndPoint, err = getJiraAPIEnpoint()
	if err != nil {
		return err
	}

	Options.Jira.ProjectCode, err = getJiraProjectCode()
	if err != nil {
		return err
	}

	Options.Jira.Email, err = getJiraEmail()
	if err != nil {
		return err
	}

	Options.Jira.APIKey, err = getJiraAPIKey()
	if err != nil {
		return err
	}

	return nil
}

func getJiraAPIKey() (string, error) {
	apiKey := os.Getenv("JIRA_API_KEY")

	if apiKey == "" {
		return "", errors.New("env variable JIRA_API_KEY not specified")
	}
	return apiKey, nil
}

func getJiraEmail() (string, error) {
	email := os.Getenv("JIRA_EMAIL")

	if email == "" {
		return "", errors.New("env variable JIRA_EMAIL not specified")
	}
	return email, nil
}

func getJiraProjectCode() (string, error) {
	projectCode := os.Getenv("JIRA_PROJECT_CODE")

	if projectCode == "" {
		return "", errors.New("env variable JIRA_PROJECT_CODE not specified")
	}
	return projectCode, nil
}

func getJiraAPIEnpoint() (string, error) {
	apiEndpoint := os.Getenv("JIRA_API_ENDPOINT")

	if apiEndpoint == "" {
		return "", errors.New("env variable JIRA_API_ENDPOINT not specified")
	}
	return apiEndpoint, nil

}
