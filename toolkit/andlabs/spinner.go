package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSpinner(a *toolkit.Action) *andlabsT {
	var newt andlabsT

	s := ui.NewSpinbox(a.X, a.Y)
	newt.uiSpinbox = s
	newt.uiControl = s
	newt.wId = a.WidgetId
	newt.WidgetType = toolkit.Spinner

	s.OnChanged(func(s *ui.Spinbox) {
		newt.i = newt.uiSpinbox.Value()
		newt.doUserEvent()
	})

	return &newt
}

func newSpinner(a *toolkit.Action) {
	var newt *andlabsT

	t := andlabs[a.ParentId]
	if (t == nil) {
		log(debugError, "NewSpinner() toolkit struct == nil. name=", a.Name)
		return
	}
	newt = t.newSpinner(a)
	place(a, t, newt)
}
