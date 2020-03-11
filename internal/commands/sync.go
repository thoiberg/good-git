package commands

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/thoiberg/good-git/internal/utils"
)

func Sync() (string, error) {
	output, err := utils.RunGitCommand("git", "branch", "--show-current")

	if err != nil {
		return "", err
	}

	currentBranch := strings.TrimSpace(output)

	// https://stackoverflow.com/a/16879922
	output, err = utils.RunGitCommand("git", "rev-parse", "--abbrev-ref", fmt.Sprintf("%v@{upstream}", currentBranch))

	if err != nil {
		// then there's no upstream set for the branch
		color.Yellow("No remote branch set up for %s", currentBranch)
		fmt.Println("Creating upstream and pushing commits...")

		output, err = utils.RunGitCommand("git", "push", "--set-upstream", "origin", currentBranch)

		if err != nil {
			return "", err
		}

		return output, nil
	}

	fmt.Println("Step 1 of 2: Integrating remote changes into local...")

	output, err = utils.RunGitCommand("git", "pull")

	if err != nil {
		return "", err
	}

	fmt.Println("Step 1 of 2: Complete!")

	fmt.Println("Step 2 of 2: Synchronising local changes to remote server...")

	output, err = utils.RunGitCommand("git", "push")

	if err != nil {
		return "", err
	}

	fmt.Println("Step 2 of 2: Complete!")

	fmt.Println("All changes are now synchronised.")

	return output, nil
}
