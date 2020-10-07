package editor

import (
	"github.com/rivo/tview"
)

func newOutput() *tview.TextView {
	return tview.NewTextView().SetScrollable(true)
}
