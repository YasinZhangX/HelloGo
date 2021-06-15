package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cliDemo",
		Usage: "Cli Demo",
	}

	app.Commands = []*cli.Command{
		&testCommand,
	}

	app.EnableBashCompletion = true
	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var testCommand cli.Command = cli.Command{
	Name:  "test",
	Usage: "test demo",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "test config",
		},
		&cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
			// Required: true,
		},
		&cli.IntFlag{
			Name:        "port",
			Usage:       "Use a randomized port",
			Value:       0,
			DefaultText: "random",
		},
	},
	Subcommands: []*cli.Command{
		{
			Name:     "add",
			Usage:    "add a new template",
			Category: "template",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "remove",
			Usage:    "remove an existing template",
			Category: "template",
			Action: func(c *cli.Context) error {
				if len(c.String("lang")) != 0 {
					fmt.Println("lang is ", c.String("lang"))
				} else {
					fmt.Println("no lang")
				}
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
	Action: func(c *cli.Context) error {
		var output string
		if c.String("lang") == "spanish" {
			output = "Hola"
		} else {
			output = "Hello"
		}
		fmt.Println(output)
		return nil
	},
}
