package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func newGroup(a *toolkit.Action) {
	// w := a.Widget
	log(debugToolkit, "NewGroup()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		log(debugToolkit, "NewGroup() toolkit struct == nil. name=", a.Name)
		listMap(debugToolkit)
	}
	newt := t.rawGroup(a.Name)
	place(a, t, newt)
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

	return &newt
}
