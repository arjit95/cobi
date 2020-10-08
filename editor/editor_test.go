package editor

import (
	"fmt"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestEditorSupport(t *testing.T) {
	editor := NewEditor()
	app := tview.NewApplication()

	editor.SetUpperPaneTitle("Commands")
	editor.SetLowerPaneTitle("Logs")

	editor.Render(app)
	editor.Input.SetFieldBackgroundColor(tcell.ColorBlack)

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1", str)
		assert.Equal(t, editor.Logger.Info.output.GetText(true), "[Info] Hello World")

		return nil
	})

	editor.SetErrorFunc(func(err error) {
		fmt.Fprint(editor.Logger.Error, err)
	})

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		defer app.QueueEvent(tcell.NewEventKey(tcell.KeyCtrlC, 0, tcell.ModNone))
		fmt.Fprint(editor.Logger.Info, "Hello World")
		editor.Input.SetText("test1")
		editor.Input.Done()

		return true
	})

	err := app.SetRoot(editor.View, true).Run()
	assert.NoError(t, err)
}
