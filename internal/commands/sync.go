package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Sync() (string, error) {
	command := exec.Command("git", "branch", "--show-current")
	output, err := command.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	currentBranch := strings.TrimSpace(bytesToString(output))

	// https://stackoverflow.com/a/16879922
	command = exec.Command("git", "rev-parse", "--abbrev-ref", fmt.Sprintf("%v@{upstream}", currentBranch))
	output, err = command.CombinedOutput()

	if err != nil {
		// then there's no upstream set for the branch
		color.Yellow("No remote branch set up for: ", currentBranch)
		fmt.Println("Creating upstream and pushing commits...")

		command = exec.Command("git", "push", "--set-upstream", "origin", currentBranch)
		output, err = command.CombinedOutput()

		if err != nil {
			return "", errors.New(bytesToString(output))
		}

		return bytesToString(output), nil
	}

	fmt.Println("Step 1 of 2: Integrating remote changes into local...")

	command = exec.Command("git", "pull")
	output, err = command.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	fmt.Println("Step 1 of 2: Complete!")

	fmt.Println("Step 2 of 2: Synchronising local changes to remote server...")

	command = exec.Command("git", "push")
	output, err = command.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	fmt.Println("Step 2 of 2: Complete!")

	fmt.Println("All changes are now synchronised.")

	return bytesToString(output), nil
}
