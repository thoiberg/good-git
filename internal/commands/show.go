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

	// @see https://stackoverflow.com/a/6245587
	gitCurrentBranch := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	gitCurrentBranchStdoutStderr, gitCurrentBranchErr := gitCurrentBranch.CombinedOutput()

	if gitCurrentBranchErr != nil {
		return "", errors.New(bytesToString(gitBranchesStdoutStderr))
	}

	branches := strings.Split(bytesToString(gitBranchesStdoutStderr), "\n")
	currentBranch := strings.Trim(bytesToString(gitCurrentBranchStdoutStderr), "\n")

	var normalisedBranchNames = normaliseGitBranchOutput(branches)

	numberOfBranches := len(normalisedBranchNames)

	cols := 1

	if numberOfBranches > 9 {
		cols = 2
	}

	for index, branch := range normalisedBranchNames {
		if branch == currentBranch {
			color.Green("%s)\t%v\n", alignRight(index, cols), branch)
		} else {
			fmt.Printf("%s)\t%v\n", alignRight(index, cols), branch)
		}
	}

	branchNumberToCheckout := -1
	reader := bufio.NewReader(os.Stdin)

	// loop until we get a valid branch number
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

	// checkout the branch
	branchToCheckout := normalisedBranchNames[branchNumberToCheckout]
	gitCheckoutCmd := exec.Command("git", "checkout", branchToCheckout)
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