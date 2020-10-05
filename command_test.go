package cobi

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var root *Command

type scenario struct {
	command            string
	expectedSuggestion []string
}

func init() {
	root = &Command{
		RootCmd: &cobra.Command{
			Use:   "copt-test",
			Short: "Test cases for copt",
		},
	}

	testCommands := []*Command{
		&Command{
			RootCmd: &cobra.Command{
				Use:   "test1",
				Short: "Description for test1",
				Args:  cobra.ExactValidArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					fmt.Printf("Test1 args %v\n", args)
				},
				ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
					return []string{"Suggestion1", "DiffSuggestion"}, cobra.ShellCompDirectiveNoFileComp
				},
			},
		},
		&Command{
			RootCmd: &cobra.Command{
				Use: "test2",
				Run: func(cmd *cobra.Command, args []string) {

				},
			},
		},
	}

	testCommands[1].RootCmd.LocalFlags().BoolP("debug", "d", false, "Testing debug flag")
	testCommands[1].AddCommand(&Command{
		RootCmd: &cobra.Command{
			Use:   "deep",
			Short: "Nested command for test2",
			Run:   func(cmd *cobra.Command, args []string) {},
		},
	})

	for _, cmd := range testCommands {
		root.AddCommand(cmd)
	}

	root.InitDefaultExitCmd()
}

func runScenarios(t *testing.T, scenarios []scenario) {
	for _, s := range scenarios {
		buf := prompt.NewBuffer()
		buf.InsertText(s.command, false, true)
		suggestions := root.generateSuggestions(*buf.Document())

		assert.Equal(t, len(s.expectedSuggestion), len(suggestions))

		for idx, suggestion := range suggestions {
			assert.Equal(t, suggestion.Text, s.expectedSuggestion[idx])
		}
	}
}

func execCommand(cmd *Command, in string) string {
	buffer := &bytes.Buffer{}
	bOut := root.RootCmd.OutOrStdout()
	bErr := root.RootCmd.OutOrStderr()

	root.RootCmd.SetOut(buffer)
	root.RootCmd.SetErr(buffer)

	root.execute(in)

	root.RootCmd.SetOut(bOut)
	root.RootCmd.SetErr(bErr)

	return buffer.String()
}

func TestNumCommands(t *testing.T) {
	var visibleCmds []string
	for _, cmd := range root.RootCmd.Commands() {
		if strings.Index(cmd.Use, "__") != 0 { // Hide complete commands
			visibleCmds = append(visibleCmds, cmd.Use)
		}
	}

	assert.Equal(t, 3, len(visibleCmds))
}

func TestInvalidShellCommand(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            `test "S`,
			expectedSuggestion: []string{},
		},
	}

	runScenarios(t, scenarioTable)
}

func TestSubCommand(t *testing.T) {
	usage := execCommand(root, "test1 --help")
	test1Desc := "Description for test1"
	assert.Equal(t, true, strings.Index(usage, test1Desc) >= 0)

	invalid := execCommand(root, `test1 "asdasd --asdgdg`)
	assert.Equal(t, true, strings.Index(invalid, "EOF") >= 0)
}
