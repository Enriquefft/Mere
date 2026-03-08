package cli

import (
	"fmt"

	cliv2 "github.com/urfave/cli/v2"
)

// VersionCommand creates the version command
func VersionCommand() *cliv2.Command {
	return &cliv2.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Display the current version of Mere",
		Action: func(c *cliv2.Context) error {
			fmt.Printf("Mere version %s\n", c.App.Version)
			fmt.Println("The minimal, module-first architecture toolkit for Claude Code")
			return nil
		},
	}
}

// GetVersion returns the current version string
func GetVersion() string {
	// This will be injected at build time via ldflags
	// For now, return a default value
	return "dev"
}

// VersionAction is the action handler for the version command
func VersionAction(c *cliv2.Context) error {
	version := GetVersion()
	fmt.Printf("Mere version %s\n", version)
	fmt.Println("The minimal, module-first architecture toolkit for Claude Code")
	return nil
}

func init() {
	// Version command initialization if needed
}
