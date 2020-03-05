package main

import (
	"log"
	"os"
	"strings"

	commands "github.com/thoiberg/good-git/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "good-git",
		Usage:                "gg",
		UsageText:            "gg command [command options] [arguments...]",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:      "show",
				Usage:     "shows the local and remote branches. Can switch to the branch by entering the number",
				UsageText: "gg show",
				Action: func(c *cli.Context) error {
					message, err := commands.Show()

					if err != nil {
						return cli.Exit(err, 1)
					}

					return cli.Exit(message, 0)
				},
			},
			{
				Name:      "grab",
				Usage:     "[INCOMPLETE] Checkout a remote branch and switch to it locally, tracking the remote with the same branch name.",
				UsageText: "gg grab <branch name>",
				Action: func(c *cli.Context) error {
					return cli.Exit("Still being built!", 1)
				},
			},
			{
				Name:      "save",
				Usage:     "Commit all working files with a given commit message. Usually followed by gg sync.",
				UsageText: "gg sync <commit message>",
				Action: func(c *cli.Context) error {
					// combines all args into a single string so we don't need to use quotation marks
					commitMessage := strings.Join(c.Args().Slice(), " ")

					message, err := commands.Save(commitMessage)

					if err != nil {
						return cli.Exit(err, 1)
					}

					return cli.Exit(message, 0)
				},
			},
			{
				Name:      "sync",
				Usage:     "Pull any commits from your remote branch into your local, and push any local commits to the remote.",
				UsageText: "gg sync",
				Action: func(c *cli.Context) error {
					message, err := commands.Sync()

					if err != nil {
						return cli.Exit(err, 1)
					}

					return cli.Exit(message, 0)
				},
			},
			{
				Name:      "status",
				Usage:     "[INCOMPLETE] Show the status of the git repo right now.",
				UsageText: "gg status",
				Aliases:   []string{"huh"},
				Action: func(c *cli.Context) error {
					message, err := commands.Status()

					if err != nil {
						return cli.Exit(err, 1)
					}

					return cli.Exit(message, 0)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
