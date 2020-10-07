package editor

import (
	"github.com/rivo/tview"
)

// SetUpperPaneTitle will set the title for upper pane.
// If no title is provided the pane will be rendered
// without title
func (editor *Editor) SetUpperPaneTitle(title string) {
	editor.upperPaneTitle = title
}

// GetUpperPaneTitle will return the title for upper pane
func (editor *Editor) GetUpperPaneTitle() string {
	return editor.upperPaneTitle
}

// SetLowerPaneTitle will set the title for lower pane.
// If no title is provided the pane will be rendered
// without title
func (editor *Editor) SetLowerPaneTitle(title string) {
	editor.lowerPaneTitle = title
}

// GetLowerPaneTitle will return the title for lower pane
func (editor *Editor) GetLowerPaneTitle() string {
	return editor.lowerPaneTitle
}

// Render initializes all the ui elements and adds
// them to the parent container
func (editor *Editor) Render(app *tview.Application) {
	view := tview.NewFlex().SetDirection(tview.FlexRow)
	child := tview.NewFlex().SetDirection(tview.FlexRow)

	if editor.GetUpperPaneTitle() != "" {
		editor.Output.SetTitle(editor.upperPaneTitle).SetBorder(true)
	}

	editor.Output.SetBorder(true)
	child = child.AddItem(editor.Output, 0, 2, false)
	child.AddItem(editor.Input, 1, 1, true)
	view = view.AddItem(child, 0, 2, true)

	if editor.GetLowerPaneTitle() != "" {
		editor.Logger.output.SetTitle(editor.lowerPaneTitle)
	}

	editor.Logger.output.SetBorder(true)
	view = child.AddItem(editor.Logger.output, 0, 1, false)
	editor.View = view
	editor.app = app
}
