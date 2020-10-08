package cobi

import (
	"io"
	"os"

	"github.com/gdamore/tcell"
)

type iopipes struct {
	stdoutW *os.File
	stdoutR *os.File
	stderrR *os.File
	stderrW *os.File

	oStdout *os.File
	oStderr *os.File

	stdoutChan chan error
	stderrChan chan error
}

func (co *Command) handleInputEvents(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlO:
		co.Editor.Output.SetText("")
		break
	case tcell.KeyCtrlL:
		co.Editor.Logger.ClearLogger()
		break
	}

	return event
}

// InteractiveMode returns true if the commands are running
// in interactive mode
func (co *Command) InteractiveMode() bool {
	return co.interactive
}

func (co *Command) pipeStdio() *iopipes {
	pipes := &iopipes{}

	pipes.stdoutR, pipes.stdoutW, _ = os.Pipe()
	pipes.stderrR, pipes.stderrW, _ = os.Pipe()

	pipes.oStdout = os.Stdout
	pipes.oStderr = os.Stderr

	os.Stdout = pipes.stdoutW
	os.Stderr = pipes.stderrW

	go func() {
		io.Copy(co.Editor.Output, pipes.stdoutR)
	}()

	go func() {
		io.Copy(co.Editor.Logger.Error, pipes.stderrR)
	}()

	return pipes
}

// ExecuteInteractive will start the execute the command in interactive mode
func (co *Command) ExecuteInteractive() error {
	if co.Editor.GetUpperPaneTitle() == "" {
		co.Editor.SetUpperPaneTitle("Commands")
	}

	if co.Editor.GetLowerPaneTitle() == "" {
		co.Editor.SetLowerPaneTitle("Logs")
	}

	co.Editor.Render(co.App)
	co.Editor.Input.SetFieldBackgroundColor(tcell.ColorBlack)
	co.Editor.Input.SetAutocompleteFunc(co.generateSuggestions)
	co.Editor.SetCommandExecFunc(co.execute)
	co.SetOut(co.Editor.Output)
	co.Editor.SetErrorFunc(co.onError)
	co.App.SetInputCapture(co.handleInputEvents)
	co.interactive = true

	co.pipes = co.pipeStdio()
	err := co.App.SetRoot(co.Editor.View, true).EnableMouse(true).Run()
	co.pipes.stderrW.Close()
	co.pipes.stdoutW.Close()

	os.Stdout = co.pipes.oStdout
	os.Stderr = co.pipes.oStderr

	return err
}
