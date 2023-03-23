package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// make new Box here
func newBox(a *toolkit.Action) {
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "newBox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "newBox() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
	}
	newt := t.rawBox(w.Name, a.B)
	newt.boxC = 0
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}

// make new Box using andlabs/ui
func (t *andlabsT) rawBox(title string, b bool) *andlabsT {
	var newt andlabsT
	var box *ui.Box
	newt.Name = title

	log(debugToolkit, "rawBox() create", newt.Name)

	if (b) {
		box = ui.NewHorizontalBox()
	} else {
		box = ui.NewVerticalBox()
	}
	box.SetPadded(padded)

	newt.uiBox = box
	newt.uiControl = box

	return &newt
}
