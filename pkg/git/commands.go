package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// FindBranchBySubstring finds branch
func FindBranchBySubstring(substring string) (string, error) {
	listAllBranchesCmd := exec.Command("git", "branch", "-a")
	grepCmd := exec.Command("grep", substring)
	var err error
	pipe, err := listAllBranchesCmd.StdoutPipe()
	defer pipe.Close()
	if err != nil {
		formattedError := fmt.Errorf(fmt.Sprintf("Failed to create pipe for grep '%s'. Error: '%s'", substring, err.Error()))
		return "", formattedError
	}

	grepCmd.Stdin = pipe

	err = listAllBranchesCmd.Start()
	if err != nil {
		formattedError := fmt.Errorf(fmt.Sprintf("Failed to start grep command for '%s'. Error: '%s'", substring, err.Error()))
		return "", formattedError
	}

	output, _ := grepCmd.Output()
	if err != nil {
		formattedError := fmt.Errorf(fmt.Sprintf("Failed to get result from grep for '%s'. Error: '%s'", substring, err.Error()))
		return "", formattedError
	}

	branches := strings.Split(string(output), "\n")
	if len(branches) == 0 {
		return "", nil
	}

	branch := strings.ReplaceAll(branches[0], "* ", "")
	branch = strings.ReplaceAll(branch, "remotes/origin/", "")
	branch = strings.ReplaceAll(branch, " ", "")

	return strings.TrimSpace(branch), nil
}

// CheckoutBranch checkouts git branch by name
func CheckoutBranch(branchName string) (string, error) {
	cmd := exec.Command("git", "checkout", branchName)
	output, err := cmd.Output()
	if err != nil {
		formattedError := fmt.Errorf(fmt.Sprintf("Failed to switch to branch '%s'. Error: '%s'", branchName, err.Error()))
		return string(output), formattedError
	}

	return string(output), nil
}

// CheckoutNewBranch creates and checks out new branch
func CheckoutNewBranch(branchName string) ([]byte, error) {
	cmd := exec.Command("git", "checkout", "-b", branchName)
	output, err := cmd.Output()
	if err != nil {
		return output, fmt.Errorf(
			fmt.Sprintf("Failed to create a new branch '%s'. Error: '%s'", branchName, err.Error()),
		)
	}

	return output, nil
}

// UpdateBranchName updates current branch name
func UpdateBranchName(branchName string) ([]byte, error) {
	cmd := exec.Command("git", "branch", "-m", branchName)
	output, err := cmd.Output()
	if err != nil {
		return output, fmt.Errorf(
			fmt.Sprintf("Failed to update branch name to '%s'. Error: '%s'", branchName, err.Error()),
		)
	}

	return output, nil

}

// GetCurrentBranch returns current git branch
func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return string(output), fmt.Errorf(
			fmt.Sprintf("Failed get current git branch with. Error: '%s'", err.Error()),
		)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
