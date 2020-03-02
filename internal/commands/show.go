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

	for index, branch := range branches {
		if len(branch) > 0 {
			if isCurrentBranch(branch) {
				normalisedBranchName := strings.Replace(branch, "* ", "", 1)
				color.Green("%d)   %v\n", index, normalisedBranchName)
			} else {
				fmt.Printf("%d) %v\n", index, branch)
			}
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
	// TODO: Need to handle when the input is not within the valid range

	branchToCheckout := branches[inputInt]
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
