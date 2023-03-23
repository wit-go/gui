package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSpinner(w *toolkit.Widget) *andlabsT {
	// make new node here
	log(debugToolkit, "newSpinner()", w.X, w.Y)
	var newt andlabsT

	s := ui.NewSpinbox(w.X, w.Y)
	newt.uiSpinbox = s
	newt.uiControl = s
	newt.tw = w

	s.OnChanged(func(s *ui.Spinbox) {
		newt.tw.I = newt.uiSpinbox.Value()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func newSpinner(a *toolkit.Action) {
	var newt *andlabsT
	w := a.Widget
	parentW := a.Where

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugError, "NewSpinner() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	w.X = a.X
	w.Y = a.Y
	newt = t.newSpinner(w)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
