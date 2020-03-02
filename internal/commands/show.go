package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Show() {
	gitBranchesCmd := exec.Command("git", "branch", "--list")
	gitBranchesStdoutStderr, gitBranchesErr := gitBranchesCmd.CombinedOutput()

	if gitBranchesErr != nil {
		log.Fatal(bytesToString(gitBranchesStdoutStderr))
	}

	branches := strings.Split(bytesToString(gitBranchesStdoutStderr), "\n")

	var onlyBranchNames []string

	for _, branchName := range branches {
		if len(branchName) > 0 {
			onlyBranchNames = append(onlyBranchNames, branchName)
		}
	}

	for index, branch := range onlyBranchNames {
		if isCurrentBranch(branch) {
			normalisedBranchName := strings.Replace(branch, "* ", "", 1)
			color.Green("%d)   %v\n", index, normalisedBranchName)
		} else {
			fmt.Printf("%d) %v\n", index, branch)
		}
	}

	// listen for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number next to the branch name to switch: ")
	text, _ := reader.ReadString('\n')

	inputInt, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	if inputInt < 0 || inputInt > len(onlyBranchNames)-1 {
		log.Fatal("Branch number is invalid. Please enter one of the numbers next to the branch number")
	}

	branchToCheckout := onlyBranchNames[inputInt]
	gitCheckoutCmd := exec.Command("git", "checkout", strings.TrimSpace(branchToCheckout))
	gitCheckoutStdStderr, gitCheckoutErr := gitCheckoutCmd.CombinedOutput()

	if gitCheckoutErr != nil {
		log.Fatal(bytesToString(gitCheckoutStdStderr))
	}

	fmt.Print(bytesToString(gitCheckoutStdStderr))
}

// https://gist.github.com/is73/de4f38e1d8da157fe33e
func bytesToString(data []byte) string {
	return string(data[:])
}

func isCurrentBranch(s string) bool {
	return strings.HasPrefix(s, "*")
}
