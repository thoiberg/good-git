package commands

import (
	"errors"
	"fmt"
	"os/exec"
)

func Sync() (string, error) {
	fmt.Println("Step 1 of 2: Integrating remote changes into local...")

	command := exec.Command("git", "pull")
	output, err := command.CombinedOutput()

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
