package utils

import (
	"errors"
	"os/exec"
)

func RunGitCommand(command ...string) (string, error) {
	commandParams := normaliseCommand(command)
	gitCommand := exec.Command(commandParams[0], commandParams[1:]...)
	output, err := gitCommand.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	return bytesToString(output), nil
}

// https://gist.github.com/is73/de4f38e1d8da157fe33e
func bytesToString(data []byte) string {
	return string(data[:])
}

func normaliseCommand(command []string) []string {
	if command[0] != "git" {
		gitArg := []string{"git"}
		return append(gitArg, command...)
	}

	return command
}
