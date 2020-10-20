package cobi

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var root *Command

type scenario struct {
	command            string
	expectedSuggestion []string
}

func init() {
	root = NewCommand(&cobra.Command{
		Use:   "cobi-test",
		Short: "Test cases for cobi",
	})

	testCommands := []*cobra.Command{
		{
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
		{
			Use: "test2",
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
	}

	testCommands[1].Flags().BoolP("debug", "d", false, "Testing debug flag")
	testCommands[1].Flags().String("namespace", "", "Select namespace")
	testCommands[1].RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ns1", "ns2"}, cobra.ShellCompDirectiveNoFileComp
	})

	testCommands[1].AddCommand(&cobra.Command{
		Use:   "deep",
		Short: "Nested command for test2",
		Run:   func(cmd *cobra.Command, args []string) {},
	})

	for _, cmd := range testCommands {
		root.AddCommand(cmd)
	}
}

func runScenarios(t *testing.T, scenarios []scenario) {
	for _, s := range scenarios {
		suggestions := root.generateSuggestions(s.command)
		assert.Equal(t, len(s.expectedSuggestion), len(suggestions))
		assert.EqualValues(t, s.expectedSuggestion, suggestions)
	}
}

func execCommand(cmd *Command, in string) string {
	buffer := &bytes.Buffer{}
	bOut := root.OutOrStdout()
	bErr := root.OutOrStderr()

	root.SetOut(buffer)
	root.SetErr(buffer)

	err := root.execute(in)

	root.SetOut(bOut)
	root.SetErr(bErr)

	if err != nil {
		return err.Error()
	}

	return buffer.String()
}

func TestNumCommands(t *testing.T) {
	var visibleCmds []string
	for _, cmd := range root.Commands() {
		if strings.Index(cmd.Use, "__") != 0 { // Hide complete commands
			visibleCmds = append(visibleCmds, cmd.Use)
		}
	}

	assert.Equal(t, 2, len(visibleCmds))
}

func TestInvalidShellCommand(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            `test "S`,
			expectedSuggestion: nil,
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
