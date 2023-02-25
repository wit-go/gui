package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

var pwLabel *toolkit.Widget
var wLabel *toolkit.Widget
var tmpNewt *andlabsT

func NewLabel(parentW *toolkit.Widget, w *toolkit.Widget) {
	pwLabel = parentW
	wLabel = w
	tmpNewt = new(andlabsT)
	tmpNewt.Width = 10
	log(debugToolkit, "mapWidgets in ui.QueueMain() START newt =", tmpNewt.Width, tmpNewt)
	if (tmpNewt == nil) {
		log(debugToolkit, "mapWidgets WHY THE HELL IS THIS NIL?", tmpNewt.Width, tmpNewt)
	}
	ui.QueueMain(newLabel)

	log(true, "sleep(.2) HACK. TODO: wrap spinlock around andlabs.ui goroutine")
	// log(true, "sleep(.2) HACK. TODO: verify newLabel() is running inside andlabs/ui goroutine") // verified that worked
	sleep(.2)
	log(debugToolkit, "mapWidgets parentW/wLabel =", pwLabel)
	log(debugToolkit, "mapWidgets new/newt =", tmpNewt)
	mapWidgetsToolkits(w, tmpNewt)

	/*
	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewDropdown() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap()
	}
	newt := t.NewDropdown(w.Name)
	mapWidgetsToolkits(w, newt)
	*/
}

func newLabel() {
	var t, newt *andlabsT
	log(debugToolkit, "gui.andlabs.NewButton()", wLabel.Name)

	t = mapToolkits[pwLabel]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewButton() toolkit struct == nil. name=", pwLabel.Name, wLabel.Name)
		return
	}

	if t.broken() {
		return
	}
	newt = tmpNewt

	newt.uiLabel = ui.NewLabel(wLabel.Name)
	newt.uiBox = t.uiBox

	log(debugToolkit, "gui.Toolbox.NewButton() about to append to Box parent t:", wLabel.Name)
	t.Dump()
	log(debugToolkit, "gui.Toolbox.NewButton() about to append to Box new t:", wLabel.Name)
	newt.Dump()

	if (t.uiBox != nil) {
		t.uiBox.Append(newt.uiLabel, false)
	} else {
		log(debugToolkit, "ERROR: wit/gui andlabs couldn't place this label in a box")
		return
	}
}
