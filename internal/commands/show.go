package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/thoiberg/good-git/internal/utils"
)

func Show() (string, error) {
	gitBranchesStdoutStderr, gitBranchesErr := utils.RunGitCommand("git", "branch", "--list")

	if gitBranchesErr != nil {
		return "", gitBranchesErr
	}

	// @see https://stackoverflow.com/a/6245587
	gitCurrentBranchStdoutStderr, gitCurrentBranchErr := utils.RunGitCommand("git", "rev-parse", "--abbrev-ref", "HEAD")

	if gitCurrentBranchErr != nil {
		return "", gitCurrentBranchErr
	}

	branches := strings.Split(gitBranchesStdoutStderr, "\n")
	currentBranch := strings.Trim(gitCurrentBranchStdoutStderr, "\n")

	var normalisedBranchNames = normaliseGitBranchOutput(branches)

	numberOfBranches := len(normalisedBranchNames)

	cols := 1

	if numberOfBranches > 9 {
		cols = 2
	}

	fmt.Print("Your available local branches are:\n\n")

	for index, branch := range normalisedBranchNames {
		if branch == currentBranch {
			color.Green(" * %s  %v\n", alignRight(index, cols), branch)
		} else {
			fmt.Printf("   %s  %v\n", alignRight(index, cols), branch)
		}
	}

	fmt.Print("\n")

	branchNumberToCheckout := -1
	reader := bufio.NewReader(os.Stdin)

	// loop until we get a valid branch number
	for branchNumberToCheckout == -1 {
		fmt.Print("Enter the number next to the branch name to switch: ")
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

	// checkout the branch
	branchToCheckout := normalisedBranchNames[branchNumberToCheckout]
	gitCheckoutStdStderr, gitCheckoutErr := utils.RunGitCommand("git", "checkout", branchToCheckout)

	if gitCheckoutErr != nil {
		return "", gitCheckoutErr
	}

	return gitCheckoutStdStderr, nil
}

// https://gist.github.com/is73/de4f38e1d8da157fe33e
func bytesToString(data []byte) string {
	return string(data[:])
}

func normaliseGitBranchOutput(branches []string) []string {
	var normalisedBranchNames []string

	for _, branch := range branches {
		if len(branch) > 0 {
			trimmedBranch := strings.TrimSpace(branch)
			onlyBranchName := strings.Replace(trimmedBranch, "* ", "", 1)
			normalisedBranchNames = append(normalisedBranchNames, onlyBranchName)
		}
	}

	return normalisedBranchNames
}

func alignRight(input int, columns int) string {
	str := strconv.Itoa(input)
	if input < 10 && columns == 2 {
		return " " + str
	} else {
		return str
	}
}
