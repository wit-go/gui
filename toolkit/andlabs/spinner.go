package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSpinner(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	w := a.Widget
	log(debugToolkit, "newSpinner()", w.X, w.Y)

	s := ui.NewSpinbox(w.X, w.Y)
	newt.uiSpinbox = s
	newt.uiControl = s
	newt.tw = w
	newt.wId = a.WidgetId
	newt.Type = toolkit.Spinner

	s.OnChanged(func(s *ui.Spinbox) {
		newt.tw.I = newt.uiSpinbox.Value()
		newt.commonChange(newt.tw, a.WidgetId)
	})

	return &newt
}

func newSpinner(a *toolkit.Action) {
	var newt *andlabsT
	w := a.Widget

	t := andlabs[a.WhereId]
	if (t == nil) {
		log(debugError, "NewSpinner() toolkit struct == nil. name=", w.Name)
		return
	}
	w.X = a.X
	w.Y = a.Y
	newt = t.newSpinner(a)
	place(a, t, newt)
}
