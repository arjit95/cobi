package cobi

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/stretchr/testify/assert"
)

func TestInteractiveSupport(t *testing.T) {
	root.BuildInteractiveSession(false)
	app, editor := root.app, root.editor

	wg := &sync.WaitGroup{}

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1", str)
		defer wg.Done()
		return nil
	})

	editor.SetErrorFunc(func(err error) {
		fmt.Fprint(editor.Logger.Error, err)
	})

	app.SetAfterDrawFunc(func(screen tcell.Screen) {
		wg.Add(1)
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
