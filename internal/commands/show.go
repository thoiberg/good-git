package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Show() (string, error) {
	gitBranchesCmd := exec.Command("git", "branch", "--list")
	gitBranchesStdoutStderr, gitBranchesErr := gitBranchesCmd.CombinedOutput()

	if gitBranchesErr != nil {
		return "", errors.New(bytesToString(gitBranchesStdoutStderr))
	}

	branches := strings.Split(bytesToString(gitBranchesStdoutStderr), "\n")

	var normalisedBranchNames []string

	for _, branch := range branches {
		if len(branch) > 0 {
			trimmedBranch := strings.TrimSpace(branch)
			if isCurrentBranch(trimmedBranch) {
				onlyBranchName := strings.Replace(trimmedBranch, "* ", "", 1)
				coloredBranchName := color.GreenString("%v", onlyBranchName)
				normalisedBranchNames = append(normalisedBranchNames, coloredBranchName)
			} else {
				normalisedBranchNames = append(normalisedBranchNames, trimmedBranch)
			}
		}
	}

	for index, branch := range normalisedBranchNames {
		fmt.Printf("%d)\t%v\n", index, branch)
	}

	branchNumberToCheckout := -1
	// listen for input
	reader := bufio.NewReader(os.Stdin)

	for branchNumberToCheckout == -1 {
		fmt.Println("Enter the number next to the branch name to switch: ")
		text, _ := reader.ReadString('\n')

		inputInt, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
		if err != nil {
			color.Red("input must be a number")
		} else if inputInt < 0 || inputInt > len(normalisedBranchNames)-1 {
			color.Red("Branch number is invalid. Please enter one of the numbers next to the branch name")
		} else {
			branchNumberToCheckout = inputInt
		}
	}

	branchToCheckout := normalisedBranchNames[branchNumberToCheckout]
	gitCheckoutCmd := exec.Command("git", "checkout", strings.TrimSpace(branchToCheckout))
	gitCheckoutStdStderr, gitCheckoutErr := gitCheckoutCmd.CombinedOutput()

	if gitCheckoutErr != nil {
		return "", errors.New(bytesToString(gitCheckoutStdStderr))
	}

	return bytesToString(gitCheckoutStdStderr), nil
}

// https://gist.github.com/is73/de4f38e1d8da157fe33e
func bytesToString(data []byte) string {
	return string(data[:])
}

func isCurrentBranch(s string) bool {
	return strings.HasPrefix(s, "*")
}
