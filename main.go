package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	gitBranchesCmd := exec.Command("git", "branch", "--list")
	gitBranchesStdoutStderr, gitBranchesErr := gitBranchesCmd.CombinedOutput()

	if gitBranchesErr != nil {
		log.Fatal(gitBranchesErr)
	}

	gitCurrentBranchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	gitCurrentBranchStdoutStderr, gitCurrentBranchErr := gitCurrentBranchCmd.CombinedOutput()

	if gitCurrentBranchErr != nil {
		log.Fatal(gitCurrentBranchErr)
	}

	fmt.Printf("all branches: %q\n", gitBranchesStdoutStderr)
	fmt.Printf("current branch: %q\n", gitCurrentBranchStdoutStderr)

	// split all branches on \n
	branches := strings.Split(bytesToString(gitBranchesStdoutStderr), "\n")
	// remove the * from the current branch

	fmt.Printf("total amount of branches is: %d\n", len(branches))
	// map each branch with it's index
	for index, element := range branches {
		fmt.Println("mapping the branches:")
		if len(element) > 0 {
			fmt.Printf("%d) %v\n", index, element)
		}
	}

	// listen for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("input: %v", text)
	// if input matches a branch number:
	inputInt, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))

	if err != nil {
		log.Fatal(err)
	}

	branchToCheckout := branches[inputInt]

	fmt.Println("branch to checkout:")
	fmt.Println(branchToCheckout)
	fmt.Println(len(branchToCheckout))
	fmt.Println(len(strings.TrimSpace(branchToCheckout)))
	// checkout that branch
	gitCheckoutCmd := exec.Command("git", "checkout", strings.TrimSpace(branchToCheckout))
	gitCheckoutStdStderr, gitCheckoutErr := gitCheckoutCmd.CombinedOutput()

	if gitCheckoutErr != nil {
		log.Fatal(gitCheckoutErr)
	}

	fmt.Printf("blah: %q", gitCheckoutStdStderr)
}

// https://gist.github.com/is73/de4f38e1d8da157fe33e
func bytesToString(data []byte) string {
	return string(data[:])
}
