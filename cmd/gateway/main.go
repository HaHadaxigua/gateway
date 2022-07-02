package main

import (
	"fmt"
	"os"

	"github.com/HaHadaxigua/gateway"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	BuildTime   = ""
	BuildNumber = ""
	GitCommit   = ""
	Version     = "1.0.0"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		logrus.Fatal(err.Error())
	}
}

func newApp() *cli.App {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer,
			"version:    %s\n"+
				"Git Commit: %s\n"+
				"Build Time: %s\n"+
				"Build:      %s\n",
			c.App.Version, GitCommit, BuildTime, BuildNumber)
	}
	return &cli.App{
		Name:                 "gateway",
		Version:              Version,
		Usage:                "get gateway",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			newSubnetCommand(),
		},
	}
}

func newSubnetCommand() *cli.Command {
	return &cli.Command{
		Name:  "subnet",
		Usage: "gateway subnet",
		Action: func(c *cli.Context) error {
			discoverGateway, err := gateway.DiscoverInterface()
			if err != nil {
				logrus.Errorf("failed to get subnet")
			}
			os.Stdout.Write([]byte(discoverGateway.String()))
			return nil
		},
	}
}
