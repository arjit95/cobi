package cobi

import (
	"errors"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/stretchr/testify/assert"
)

func TestInteractiveValidCmd(t *testing.T) {
	editor := root.Editor

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1", str)
		assert.Equal(t, root.InteractiveMode(), true)
		return nil
	})

	root.App.SetBeforeDrawFunc(func(_ tcell.Screen) bool {
		editor.Input.SetText("test1")
		editor.Input.Done()
		root.App.QueueEvent(tcell.NewEventKey(tcell.KeyCtrlO, 0, tcell.ModNone))
		root.App.QueueEvent(tcell.NewEventKey(tcell.KeyCtrlE, 0, tcell.ModNone))
		defer root.App.QueueEvent(tcell.NewEventKey(tcell.KeyCtrlC, 0, tcell.ModNone))
		return true
	})

	err := root.ExecuteInteractive()
	assert.NoError(t, err)
}

func TestInteractiveInvalidCmd(t *testing.T) {
	root = NewCommand(root.Command)

	editor := root.Editor

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1 ", str)
		return errors.New("Invalid command")
	})

	root.App.SetBeforeDrawFunc(func(_ tcell.Screen) bool {
		editor.Input.SetText("test1 invalid")
		editor.Input.Done()
		root.App.QueueEvent(tcell.NewEventKey(tcell.KeyCtrlC, 0, tcell.ModNone))
		return true
	})

	err := root.ExecuteInteractive()
	assert.NoError(t, err)
}
