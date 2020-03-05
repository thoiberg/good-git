package commands

import (
	"errors"
	"os/exec"
)

func Status() (string, error) {
	command := exec.Command(
		"git",
		"-c", "color.status=always", // @see https://stackoverflow.com/a/18304605
		"status",
	)
	output, err := command.CombinedOutput()

	if err != nil {
		return "", errors.New(bytesToString(output))
	}

	return bytesToString(output), nil
}