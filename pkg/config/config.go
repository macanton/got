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
	CheckoutBranch             OperationType = "CheckoutBranch"
	CheckBranchForNewJiraIssue OperationType = "CheckBranchForNewJiraIssue"
)

// OptionsType is a type for stored app configuration
type OptionsType struct {
	IssueCode int
	Summary   string
	Operation OperationType
	Jira      struct {
		ProjectCode string
		APIEndPoint string
		Email       string
		APIKey      string
	}
}

// Options variable stores app configuration settings
var Options OptionsType

// InitAndRequestAdditionalData function initializes global configuration of the application
func InitAndRequestAdditionalData() error {
	err := readConfigVariables()
	if err != nil {
		return err
	}

	ticketID := flag.Int("b", 0, "Jira ticket number key for new branch")
	createIssue := flag.Bool("cj", false, "Create a new Jira issue and switch to the new branch")
	flag.Parse()

	if *ticketID < 0 {
		return errors.New("Jira ticket number ket should be more than 0")
	}

	if *ticketID > 0 {
		Options.IssueCode = *ticketID
		Options.Operation = CheckoutBranch
		return nil
	}

	if *createIssue {
		Options.Operation = CheckBranchForNewJiraIssue
		fmt.Print("Enter issue summary: ")

		reader := bufio.NewReader(os.Stdin)
		summary, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Failed to read string from buffer reader: %s", err.Error())
		}

		summary = strings.TrimSpace(summary)
		if len(summary) == 0 {
			return errors.New("Summary cannot be empty string")
		}

		Options.Summary = summary
		return nil
	}

	return errors.New("Invalid flags supplied. Cannot determine target operation")
}

// GetIssueKey returns a key of the Jira issue that contains project code and issue code
func GetIssueKey() string {
	return fmt.Sprintf("%s-%d", Options.Jira.ProjectCode, Options.IssueCode)
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
