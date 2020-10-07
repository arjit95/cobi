package cobi

import (
	Editor "github.com/arjit95/cobi/editor"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (co *Command) handleInputEvents(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlL:
		co.editor.Output.SetText("")
		break
	case tcell.KeyCtrlE:
		co.editor.Logger.ClearLogger()
		break
	}

	return event
}

// BuildInteractiveSession will start the shell in interactive mode
// start bool is for testing purposes, the user will always send it
// as true
func (co *Command) BuildInteractiveSession(start bool) {
	co.app = tview.NewApplication()
	co.editor = Editor.NewEditor()

	co.editor.SetUpperPaneTitle("Commands")
	co.editor.SetLowerPaneTitle("Logs")

	co.editor.Render(co.app)
	co.editor.Input.SetFieldBackgroundColor(tcell.ColorBlack)
	co.editor.Input.SetAutocompleteFunc(co.generateSuggestions)
	co.editor.SetCommandExecFunc(co.execute)
	co.SetOut(co.editor.Output)
	co.editor.SetErrorFunc(co.onError)
	co.app.SetInputCapture(co.handleInputEvents)

	if !start {
		return
	}

	if err := co.app.SetRoot(co.editor.View, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
