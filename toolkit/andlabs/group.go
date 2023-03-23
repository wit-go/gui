package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func newGroup(a *toolkit.Action) {
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "NewGroup()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "NewGroup() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
	}
	newt := t.rawGroup(w.Name)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}

// make new Group here
func (t *andlabsT) rawGroup(title string) *andlabsT {
	var newt andlabsT
	newt.Name = title

	log(debugToolkit, "NewGroup() create", newt.Name)

	g := ui.NewGroup(newt.Name)
	g.SetMargined(margin)
	newt.uiGroup = g
	newt.uiControl = g

//	hbox := ui.NewVerticalBox()
//	hbox.SetPadded(padded)
//	g.SetChild(hbox)

//	newt.uiBox = hbox
//	newt.uiWindow = t.uiWindow
//	newt.uiTab = t.uiTab

	return &newt
}
