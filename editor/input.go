package editor

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Input handles all the commands entered during interactive mode
type Input struct {
	*tview.InputField
	commandExecFunc func(string) error
	errorFunc       func(error)
	history         []string
	historyIdx      int
	text            string
}

// doneFunc is invoked when the user is done entering the command
// i.e after user presses enter
func (i *Input) doneFunc(key tcell.Key) {
	if i.commandExecFunc == nil || strings.TrimSpace(i.GetText()) == "" {
		return
	}

	var err error

	text := i.GetText()
	switch key {
	case tcell.KeyEnter:
		i.history = append(i.history, text)
		i.historyIdx = len(i.history)

		err = i.commandExecFunc(text)
		i.SetText("")
		break
	}

	if err != nil && i.errorFunc != nil {
		i.errorFunc(err)
	}
}

func (i *Input) handleInput(event *tcell.EventKey) *tcell.EventKey {
	historyLen := len(i.history)
	i.text = i.GetText()

	switch event.Key() {
	case tcell.KeyUp:
		if i.historyIdx > 0 {
			i.historyIdx--
			i.SetText(i.history[i.historyIdx])
			return nil
		}

		break
	case tcell.KeyDown:
		if i.historyIdx < historyLen-1 {
			i.historyIdx++
			i.SetText(i.history[i.historyIdx])

			return nil
		}

		i.historyIdx = len(i.history)
		i.SetText(i.text)
		break
	}

	return event
}

// Done function emulates enter key behavior and invokes onDone callback
func (i *Input) Done() {
	i.doneFunc(tcell.KeyEnter)
}

func newInput() *Input {
	view := tview.NewInputField()
	instance := &Input{InputField: view, historyIdx: -1}
	instance.SetDoneFunc(instance.doneFunc)
	instance.SetInputCapture(instance.handleInput)

	return instance
}
