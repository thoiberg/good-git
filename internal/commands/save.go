package commands

import (
	"errors"
	"os/exec"
)

func Save(addNewFiles bool, commitMessage string) (string, error) {
	if addNewFiles {
		gitAddCmd := exec.Command("git", "add", "--all")
		output, err := gitAddCmd.CombinedOutput()

		if err != nil {
			return "", errors.New(bytesToString(output))
		}
	}

	gitCommitCmd := exec.Command("git", "commit", "-am", commitMessage)
	output, err := gitCommitCmd.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	return bytesToString(output), nil
}
