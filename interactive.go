package cobi

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

func (co *Command) onError(err error) {
	if err != nil {
		co.RootCmd.PrintErr(err)
	}
}

func (co *Command) execute(in string) {
	promptArgs, err := shlex.Split(in)
	if err != nil {
		co.onError(err)
	}

	os.Args = append([]string{os.Args[0]}, promptArgs...)
	co.onError(co.Execute())
}

// RunInteractive will start the shell in interactive mode
func (co *Command) RunInteractive() {
	p := prompt.New(
		co.execute,
		co.generateSuggestions,
		co.GoPromptOptions...,
	)

	p.Run()
}
