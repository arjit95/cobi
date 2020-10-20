package cobi

import (
	"github.com/gdamore/tcell"
)

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

// ExecuteInteractive will start the command in interactive mode
func (co *Command) ExecuteInteractive() error {
	co.Editor.Render(co.App)
	co.Editor.Input.SetAutocompleteFunc(co.generateSuggestions)
	co.Editor.SetCommandExecFunc(co.execute)
	co.SetOut(co.Editor.Output)
	co.Editor.SetErrorFunc(co.onError)
	co.App.SetInputCapture(co.handleInputEvents)
	co.interactive = true

	err := co.App.SetRoot(co.Editor.View, true).EnableMouse(true).Run()
	return err
}
