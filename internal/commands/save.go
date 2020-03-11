package commands

import (
	"github.com/thoiberg/good-git/internal/utils"
)

func Save(addNewFiles bool, commitMessage string) (string, error) {
	if addNewFiles {
		_, err := utils.RunGitCommand("git", "add", "--all")

		if err != nil {
			return "", err
		}
	}

	output, err := utils.RunGitCommand("git", "commit", "-am", commitMessage)

	if err != nil {
		return "", err
	}

	return output, nil
}
