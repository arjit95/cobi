package editor

import (
	"os"

	"github.com/rivo/tview"
)

type Output struct {
	*writer
	View *tview.TextView
}

func newOutput() *Output {
	txtview := tview.NewTextView().SetScrollable(true)
	output := &Output{
		View: txtview,
		writer: &writer{
			prefix:           []byte{},
			primary:          os.Stdout,
			secondary:        txtview,
			primaryAvailable: true,
		},
	}

	return output
}
