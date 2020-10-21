package editor

import (
	"io"
	"os"

	"github.com/rivo/tview"
)

type writer struct {
	prefix           []byte
	primary          io.Writer
	secondary        io.Writer
	primaryAvailable bool
}

func (w *writer) Write(data []byte) (n int, err error) {
	if w.primaryAvailable {
		return w.primary.Write(data)
	}

	return w.secondary.Write(append(w.prefix, data...))
}

// Logger provides an option to log background logs to ui
type Logger struct {
	Info  *writer
	Warn  *writer
	Error *writer
	View  *tview.TextView
}

// ClearLogger will clear all data from logger pane
func (l *Logger) ClearLogger() {
	l.View.SetText("")
}

// SetPrimary sets the primary writer for logger
func (l *Logger) SetPrimary(w io.Writer) {
	l.Info.primary = w
	l.Warn.primary = w
	l.Error.primary = w
}

// SetSecondary sets the secondary writer for logger
func (l *Logger) SetSecondary(w io.Writer) {
	l.Info.secondary = w
	l.Warn.secondary = w
	l.Error.secondary = w
}

// SetPrimaryAvailable decides if the primary writer can be used or not
func (l *Logger) SetPrimaryAvailable(state bool) {
	l.Info.primaryAvailable = state
	l.Warn.primaryAvailable = state
	l.Error.primaryAvailable = state
}

func newLogger() *Logger {
	view := tview.NewTextView().SetScrollable(true)
	prefix := []string{"[Info] ", "[Warn] ", "[Error] "}
	outs := make([]*writer, len(prefix))

	for i, p := range prefix {
		primary := os.Stdout
		if i == len(prefix)-1 {
			primary = os.Stderr
		}

		outs[i] = &writer{
			prefix:           []byte(p),
			primary:          primary,
			primaryAvailable: true,
			secondary:        view,
		}
	}

	return &Logger{
		View:  view,
		Info:  outs[0],
		Warn:  outs[1],
		Error: outs[2],
	}
}
