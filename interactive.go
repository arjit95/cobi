package cobi

import (
	Editor "github.com/arjit95/cobi/editor"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
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
	return co.interactve
}

// BuildInteractiveSession will start the shell in interactive mode
// start bool is for testing purposes, it should always be set to true.
func (co *Command) BuildInteractiveSession(start bool) {
	co.app = tview.NewApplication()
	co.Editor = Editor.NewEditor()

	co.Editor.SetUpperPaneTitle("Commands")
	co.Editor.SetLowerPaneTitle("Logs")

	co.Editor.Render(co.app)
	co.Editor.Input.SetFieldBackgroundColor(tcell.ColorBlack)
	co.Editor.Input.SetAutocompleteFunc(co.generateSuggestions)
	co.Editor.SetCommandExecFunc(co.execute)
	co.SetOut(co.Editor.Output)
	co.Editor.SetErrorFunc(co.onError)
	co.app.SetInputCapture(co.handleInputEvents)
	co.interactve = true

	if !start {
		return
	}

	if err := co.app.SetRoot(co.Editor.View, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
