package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Show() {
	gitBranchesCmd := exec.Command("git", "branch", "--list")
	gitBranchesStdoutStderr, gitBranchesErr := gitBranchesCmd.CombinedOutput()

	if gitBranchesErr != nil {
		log.Fatal(bytesToString(gitBranchesStdoutStderr))
	}

	gitCurrentBranchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	gitCurrentBranchStdoutStderr, gitCurrentBranchErr := gitCurrentBranchCmd.CombinedOutput()

	if gitCurrentBranchErr != nil {
		log.Fatal(gitCurrentBranchErr)
	}

	branches := strings.Split(bytesToString(gitBranchesStdoutStderr), "\n")
	// TODO: remove the * from the current branch

	for index, element := range branches {
		if len(element) > 0 {
			fmt.Printf("%d) %v\n", index, element)
		}
	}

	// listen for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	inputInt, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	// TODO: Need to handle when the input is not within the valid range

	if err != nil {
		log.Fatal(err)
	}

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
