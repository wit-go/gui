package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newLabel(a *toolkit.Action) {
	var newt *andlabsT
	w := a.Widget
	log(debugToolkit, "NewLabel()", w.Name)

	t := mapToolkits[a.Where]
	if (t == nil) {
		listMap(debugError)
		log(debugError, "ERROR newLabel() listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		return
	}

	log(debugToolkit, "NewLabel()", w.Name)

	newt = new(andlabsT)

	c := ui.NewLabel(w.Name)
	newt.uiLabel = c
	newt.uiControl = c

	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
