package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "show",
				Usage: "[INCOMPLETE] shows the local and remote branches. Can switch to the branch by entering the number",
				Action: func(c *cli.Context) error {
					return cli.Exit("Still being built!", 1)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
