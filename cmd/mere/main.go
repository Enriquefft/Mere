package main

import (
	"fmt"
	"os"

	"github.com/Enriquefft/Mere/internal/cli"
	cliv2 "github.com/urfave/cli/v2"
)

// version is injected at build time via ldflags
var version = "dev"

func main() {
	app := &cliv2.App{
		Name:     "mere",
		Version:  version,
		Usage:    "Minimal, module-first architecture toolkit for AI-native development",
		UsageText: "mere [command] [options]",
		Commands: []*cliv2.Command{
			cli.InitCommand(),
			cli.VersionCommand(),
		},
		Flags: []cliv2.Flag{
			&cliv2.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose output",
			},
		},
		Before: func(c *cliv2.Context) error {
			if c.Bool("verbose") {
				fmt.Printf("Mere v%s\n", version)
			}
			return nil
		},
		Action: func(c *cliv2.Context) error {
			if c.Args().Len() == 0 {
				return cli.ShowAppHelp(c)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
