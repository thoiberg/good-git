package commands

import utils "github.com/thoiberg/good-git/internal/utils"

func Status() (string, error) {
	// command := exec.Command(
	// 	"git",
	// 	"-c", "color.status=always", // @see https://stackoverflow.com/a/18304605
	// 	"status",
	// )
	// output, err := command.CombinedOutput()

	output, err := utils.RunGitCommand("git -c color.status=always status")

	if err != nil {
		return "", err
	}

	return output, nil
}
