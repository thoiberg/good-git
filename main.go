package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cmd := exec.Command("git", "branch")
	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}
}
