package editor

import "github.com/rivo/tview"

// Editor implements the methods to manipulate modify editor
// behavior and manipulate ui elements
type Editor struct {
	View   *tview.Flex
	Input  *Input
	Logger *Logger
	Output *tview.TextView

	isLoggerVisible bool
	upperPaneTitle  string
	lowerPaneTitle  string
	app             *tview.Application
}

// SetErrorFunc will set an error function which will be invoked
// when there an error occurs while executing a command
func (editor *Editor) SetErrorFunc(fn func(error)) {
	editor.Input.errorFunc = fn
}

// SetCommandExecFunc will set an exec function which is invoked
// when user presses enter key after entering the command
func (editor *Editor) SetCommandExecFunc(fn func(string) error) {
	editor.Input.commandExecFunc = fn
}

// NewEditor will return a reference to the complete editor ui
// used to run commands in interactive mode
func NewEditor() *Editor {
	editor := &Editor{
		Input:           newInput(),
		Logger:          newLogger(),
		Output:          newOutput(),
		isLoggerVisible: true,
	}

	return editor
}
