package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"os"
)

var username string
var password string
var skipVerify bool

func main() {
	app := cli.NewApp()
	app.Name = "jenkins_list_plugins"
	app.Usage = "List Jenkins plugins in shortname:version format"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "username, u",
			Usage:       "Username for Jenkins authentication",
			EnvVar:      "JENKINS_USERNAME",
			Destination: &username,
		},
		cli.StringFlag{
			Name:        "password, p",
			Usage:       "Password for Jenkins authentication",
			EnvVar:      "JENKINS_PASSWORD",
			Destination: &password,
		},
		cli.BoolFlag{
			Name:        "insecure, k",
			Usage:       "Allow connections to SSL sites without certs",
			EnvVar:      "SKIP_TLSVERIFY",
			Destination: &skipVerify,
		},
	}
	app.Action = func(c *cli.Context) {
		if !c.Args().Present() {
			fmt.Println("Missing first argument, Jenkins URL")
			os.Exit(1)
		}
		url := c.Args().First()
		plugins, err := ListPlugins(url, username, password, skipVerify)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		red := color.New(color.FgRed)
		yellow := color.New(color.FgYellow)
		green := color.New(color.FgGreen)

		for _, plugin := range plugins {
			var c *color.Color
			if plugin.HasUpdate {
				c = yellow
			} else if plugin.Active {
				c = green
			} else {
				c = red
			}
			c.Printf("%s:%s\n", plugin.ShortName, plugin.Version)
		}
	}

	app.RunAndExitOnError()
}
