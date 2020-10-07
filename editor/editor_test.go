package editor

import (
	"fmt"
	"sync"
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

	wg := &sync.WaitGroup{}

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1", str)
		assert.Equal(t, editor.Logger.Info.output.GetText(true), "[Info] Hello World")

		defer wg.Done()
		return nil
	})

	editor.SetErrorFunc(func(err error) {
		fmt.Fprint(editor.Logger.Error, err)
	})

	app.SetAfterDrawFunc(func(screen tcell.Screen) {
		wg.Add(1)
		fmt.Fprint(editor.Logger.Info, "Hello World")
		editor.Input.SetText("test1")
		editor.Input.Done()
	})

	go func() {
		wg.Wait()
		app.Stop()
	}()

	if err := app.SetRoot(editor.View, true).Run(); err != nil {
		assert.Error(t, err)
	}
}
