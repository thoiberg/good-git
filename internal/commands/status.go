package commands

import utils "github.com/thoiberg/good-git/internal/utils"

func Status() (string, error) {
	output, err := utils.RunGitCommand("git", "-c", "color.status=always status")

	if err != nil {
		return "", err
	}

	return output, nil
}
