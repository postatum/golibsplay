package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/postatum/golibsplay/core"
	"path/filepath"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:  "jsonfirst",
		Usage: "get first json element",
		Action: func(c *cli.Context) error {
			path, err := filepath.Abs(c.Args().First())
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			el, err := core.GetFirstFileElement(path)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println(el)
			return nil
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
