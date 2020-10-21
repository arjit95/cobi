package cobi

import (
	"bytes"
	"os"
	"strings"

	"github.com/google/shlex"
)

// Adds default flag suggestions to a command
func (co *Command) generateDefaultFlagSuggestions(args []string) {
	child, _, err := co.Find(args)
	if err != nil || child == nil {
		return
	}

	child.InitDefaultHelpFlag()
}

// Generate command suggestions by invoking __complete api for cobra commands
func (co *Command) generateSuggestions(text string) []string {
	promptArgs, err := shlex.Split(text)
	if err != nil {
		return nil
	}

	buffer := &bytes.Buffer{}
	bOut := co.OutOrStdout()
	bErr := co.OutOrStderr()

	co.SetOut(buffer)
	co.SetErr(buffer)

	if len(promptArgs) == 0 {
		promptArgs = append(promptArgs, "")
	}

	co.generateDefaultFlagSuggestions(promptArgs)
	os.Args = append([]string{co.Use, "__complete"}, promptArgs...)
	err = co.Execute()

	// Restore output
	co.SetOut(bOut)
	co.SetErr(bErr)

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

	var suggestions []string
	for _, command := range commands {
		cmdMeta := strings.SplitN(command, "\t", 2)
		paLen := len(promptArgs)
		if paLen == 1 {
			suggestions = append(suggestions, cmdMeta[0])
			continue
		}

		last := promptArgs[paLen-1]
		isFlag := strings.Index(last, "-") == 0
		suggestion := promptArgs[:paLen-1]
		suggestion = append(suggestion, cmdMeta[0])

		// Complete normal suggestions or flags having type
		// --flag value
		if !isFlag || strings.Index(last, "=") == -1 {
			suggestions = append(suggestions, strings.Join(suggestion, " "))
			continue
		}

		// complete -flag=value
		parts := strings.Split(last, "=")
		parts[1] = cmdMeta[0]
		suggestion[len(suggestion)-1] = strings.Join(parts, "=")
		suggestions = append(suggestions, strings.Join(suggestion, " "))
	}

	return suggestions
}
