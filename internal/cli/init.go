package cli

import (
	"fmt"
	"os"

	"github.com/hybridz/mere/internal/scaffold"
	cliv2 "github.com/urfave/cli/v2"
)

// InitCommand creates the init command
func InitCommand() *cliv2.Command {
	return &cliv2.Command{
		Name:  "init",
		Usage: "Initialize .claude/ directory with AI-optimized structure",
		Description: `Initialize the current project with a .claude/ directory containing:
- CLAUDE.md: Project context and boundary rules
- commands/: AI command templates for module management`,
		Flags: []cliv2.Flag{
			&cliv2.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Overwrite existing .claude/ directory",
			},
		},
		Action: func(c *cliv2.Context) error {
			force := c.Bool("force")
			return scaffold.InitProject(force)
		},
	}
}

// ShowAppHelp displays the app help
func ShowAppHelp(c *cliv2.Context) error {
	cliv2.ShowAppHelpAndExit(c, 0)
	return nil
}

func init() {
	// Register commands will be called from main
}

// Internal helper to output messages
func outputMessage(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// Internal helper to output errors
func outputError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
