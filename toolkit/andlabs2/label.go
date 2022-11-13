package main

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "git.wit.org/wit/gui/toolkit"

func NewLabel(parentW *toolkit.Widget, w *toolkit.Widget) {
	var t, newt *andlabsT
	log.Println("gui.andlabs.NewButton()", w.Name)

	t = mapToolkits[parentW]
	if (t == nil) {
		log.Println("go.andlabs.NewButton() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}

	if t.broken() {
		return
	}
	newt = new(andlabsT)

	newt.uiLabel = ui.NewLabel(w.Name)
	newt.uiBox = t.uiBox

	if (DebugToolkit) {
		log.Println("gui.Toolbox.NewButton() about to append to Box parent t:", w.Name)
		t.Dump()
		log.Println("gui.Toolbox.NewButton() about to append to Box new t:", w.Name)
		newt.Dump()
	}
	if (t.uiBox != nil) {
		t.uiBox.Append(newt.uiLabel, false)
	} else {
		log.Println("ERROR: wit/gui andlabs couldn't place this label in a box")
		return
	}

	mapWidgetsToolkits(w, newt)
}
