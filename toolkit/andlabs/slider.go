package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t andlabsT) NewSlider(title string, x int, y int) *andlabsT {
	// make new node here
	log(debugToolkit, "gui.Toolkit.NewSpinbox()", x, y)
	var newt andlabsT

	if (t.uiBox == nil) {
		log(debugToolkit, "gui.ToolkitNode.NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugToolkit, "probably could just make a box here?")
		exit("internal golang wit/gui/toolkit error")
		return nil
	}

	s := ui.NewSlider(x, y)
	newt.uiSlider = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(spin *ui.Slider) {
		newt.commonChange("Slider")
	})

	return &newt
}

func NewSlider(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT
	log(debugToolkit, "gui.andlabs.NewTab()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewTab() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.NewSlider(w.Name, w.X, w.Y)
	mapWidgetsToolkits(w, newt)
}
