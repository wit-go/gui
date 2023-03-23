package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSlider(w *toolkit.Widget) *andlabsT {
	// make new node here
	log(debugToolkit, w.Name, w.Type, w.X, w.Y)
	var newt andlabsT

	s := ui.NewSlider(w.X, w.Y)
	newt.uiSlider = s
	newt.uiControl = s
	newt.tw = w

	s.OnChanged(func(spin *ui.Slider) {
		newt.tw.I = newt.uiSlider.Value()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func newSlider(a *toolkit.Action) {
	var newt *andlabsT
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "newSlider()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugError, "newSlider() ERROR toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	w.X = a.X
	w.Y = a.Y
	newt = t.newSlider(w)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
