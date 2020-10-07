package cobi

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteractiveValidCmd(t *testing.T) {
	if !root.InteractiveMode() {
		root.BuildInteractiveSession(false)
	}

	editor := root.Editor

	wg := &sync.WaitGroup{}
	wg.Add(1)

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1", str)
		assert.Equal(t, root.InteractiveMode(), true)
		defer wg.Done()
		return nil
	})

	editor.Input.SetText("test1")
	editor.Input.Done()
	wg.Wait()
}

func TestInteractiveInvalidCmd(t *testing.T) {
	if !root.InteractiveMode() {
		root.BuildInteractiveSession(false)
	}

	editor := root.Editor

	wg := &sync.WaitGroup{}
	wg.Add(2)

	invalidCommand := false

	editor.SetCommandExecFunc(func(str string) error {
		assert.Equal(t, "test1 asdasdasd", str)
		defer wg.Done()
		return errors.New("Invalid command")
	})

	editor.SetErrorFunc(func(err error) {
		invalidCommand = true
		root.onError(err)
		defer wg.Done()
	})

	t.Logf("setting txt")
	editor.Input.SetText("test1 asdasdasd")
	t.Logf("text txt")
	editor.Input.Done()

	wg.Wait()
	assert.Equal(t, true, invalidCommand)
}
