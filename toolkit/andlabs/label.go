package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newLabel(a *toolkit.Action) {
	var newt *andlabsT
	log(debugToolkit, "NewLabel()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		listMap(debugError)
		log(debugError, "ERROR newLabel() listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		return
	}

	log(debugToolkit, "NewLabel()", a.Name)

	newt = new(andlabsT)

	c := ui.NewLabel(a.Name)
	newt.uiLabel = c
	newt.uiControl = c

	place(a, t, newt)
}
