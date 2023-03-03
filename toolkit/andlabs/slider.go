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

	if (t.uiBox == nil) {
		log(debugToolkit, "node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugToolkit, "probably could just make a box here?")
		exit("internal golang wit/gui/toolkit error")
		return nil
	}

	s := ui.NewSlider(w.X, w.Y)
	newt.uiSlider = s
	newt.uiBox = t.uiBox
	newt.tw = w
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(spin *ui.Slider) {
		newt.tw.I = newt.uiSlider.Value()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func newSlider(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT
	log(debugToolkit, "gui.andlabs.NewTab()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewTab() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.newSlider(w)
	mapWidgetsToolkits(w, newt)
}
