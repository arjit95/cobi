package editor

import (
	"github.com/rivo/tview"
)

type writer struct {
	prefix []byte
	output *tview.TextView
}

func (w *writer) Write(data []byte) (n int, err error) {
	return w.output.Write(append(w.prefix, data...))
}

// Logger provides an option to log background logs to ui
type Logger struct {
	Info   *writer
	Warn   *writer
	Error  *writer
	output *tview.TextView
}

// ClearLogger will clear all data from logger pane
func (l *Logger) ClearLogger() {
	l.output.SetText("")
}

func newLogger() *Logger {
	output := newOutput()
	return &Logger{
		output: output,
		Info:   &writer{prefix: []byte("[Info] "), output: output},
		Warn:   &writer{prefix: []byte("[Warn] "), output: output},
		Error:  &writer{prefix: []byte("[Error] "), output: output},
	}
}
