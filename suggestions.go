package cobi

import (
	"bytes"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

// Adds default flag suggestions to a command
func (co *Command) generateDefaultFlagSuggestions(args []string) {
	child, _, err := co.RootCmd.Find(args)
	if err != nil || child == nil {
		return
	}

	child.InitDefaultHelpFlag()
}

func (co *Command) generateSuggestions(d prompt.Document) []prompt.Suggest {
	text := d.Text

	promptArgs, err := shlex.Split(text)
	if err != nil {
		return nil
	}

	buffer := &bytes.Buffer{}
	bOut := co.RootCmd.OutOrStdout()
	bErr := co.RootCmd.OutOrStderr()

	co.RootCmd.SetOut(buffer)
	co.RootCmd.SetErr(buffer)

	co.generateDefaultFlagSuggestions(promptArgs)
	os.Args = append([]string{os.Args[0], "__complete"}, promptArgs...)
	err = co.Execute()

	// Restore output
	co.RootCmd.SetOut(bOut)
	co.RootCmd.SetErr(bErr)

	if err != nil {
		return nil
	}

	bString := buffer.String()
	commands := trimEmptyLines(strings.Split(bString, "\n"))

	// Trim completions metadata
	commands = commands[:len(commands)-2]

	// No completions present
	if len(commands) == 0 {
		return nil
	}

	var suggestions []prompt.Suggest
	for _, command := range commands {
		cmdMeta := strings.SplitN(command, "\t", 2)
		var suggestion prompt.Suggest

		if len(cmdMeta) == 2 {
			suggestion = prompt.Suggest{Text: cmdMeta[0], Description: cmdMeta[1]}
		} else { // Command without description
			suggestion = prompt.Suggest{Text: cmdMeta[0]}
		}

		suggestions = append(suggestions, suggestion)
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
