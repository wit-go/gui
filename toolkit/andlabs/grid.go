package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

// Grid numbering by (X,Y)
// -----------------------------
// -- (1,1) -- (2,1) -- (3,1) --
// -- (1,2) -- (2,1) -- (3,1) --
// -----------------------------
func newGrid(a *toolkit.Action) {
	var newt *andlabsT
	log(debugToolkit, "newGrid()", a.Widget.Name, "to", a.Where.Type)

	t := mapToolkits[a.Where]

	newt = new(andlabsT)

	c := ui.NewGrid()
	newt.uiGrid = c
	newt.uiControl = c
	newt.tw = a.Widget
	newt.gridX = 0
	newt.gridY = 0

	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
