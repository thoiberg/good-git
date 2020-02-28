package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "branch")
	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", stdoutStderr)
}
