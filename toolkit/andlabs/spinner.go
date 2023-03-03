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

	if (t.uiBox == nil) {
		log(debugToolkit, "newSpinner() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return nil
	}

	s := ui.NewSpinbox(w.X, w.Y)
	newt.uiSpinbox = s
	newt.uiBox = t.uiBox
	newt.tw = w
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(s *ui.Spinbox) {
		newt.tw.I = newt.uiSpinbox.Value()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func newSpinner(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewTab() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.newSpinner(w)
	mapWidgetsToolkits(w, newt)
}
