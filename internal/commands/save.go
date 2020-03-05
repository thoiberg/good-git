package commands

import (
	"errors"
	"os/exec"
)

func Save(commitMessage string) (string, error) {
	gitCommitCmd := exec.Command("git", "commit", "-am", commitMessage)
	output, err := gitCommitCmd.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	return bytesToString(output), nil
}
