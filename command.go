package cobi

import (
	"fmt"
	"os"
	"strings"

	"github.com/arjit95/cobi/editor"
	"github.com/google/shlex"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// Command is a modified struct to handle interactive sessions
// combined with tview & cobra.
type Command struct {
	*cobra.Command
	App         *tview.Application
	Editor      *editor.Editor
	interactive bool
	pipes       *iopipes
}

func (co *Command) onError(err error) {
	if err != nil {
		fmt.Fprintf(co.Editor.Logger.Error, "%s\n", err)
	}
}

func (co *Command) execute(in string) error {
	promptArgs, err := shlex.Split(in)
	if err != nil {
		return err
	}

	os.Args = append([]string{os.Args[0]}, promptArgs...)
	return co.Execute()
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

// NewCommand returns an instance of cobi command
func NewCommand(cmd *cobra.Command) *Command {
	instance := &Command{
		App:     tview.NewApplication(),
		Editor:  editor.NewEditor(),
		Command: cmd,
	}

	return instance
}
