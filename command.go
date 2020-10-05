package cobi

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// Command is a modified struct to handle interactive sessions
// combined with go-prompt & cobra.
type Command struct {
	RootCmd         *cobra.Command
	GoPromptOptions []prompt.Option
}

// AddCommand adds a new child to an existing command
func (co *Command) AddCommand(nCo *Command) {
	co.RootCmd.AddCommand(nCo.RootCmd)
}

// Execute runs the command
func (co *Command) Execute() error {
	return co.RootCmd.Execute()
}

func trimEmptyLines(args []string) []string {
	var lines []string
	for _, arg := range args {
		trimmed := strings.TrimSpace(arg)
		if trimmed != "" {
			lines = append(lines, trimmed)
		}
	}

	return lines
}

// InitDefaultExitCmd adds an exit command. Useful while running in interactive shell
func (co *Command) InitDefaultExitCmd() {
	co.AddCommand(&Command{
		RootCmd: &cobra.Command{
			Use:   "exit",
			Short: "Exits the app",
			Run: func(cmd *cobra.Command, args []string) {
				os.Exit(0)
			},
		},
	})
}
