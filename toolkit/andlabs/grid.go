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
	log(debugToolkit, "newGrid()", a.WidgetId, "to", a.ParentId)

	newt = new(andlabsT)

	c := ui.NewGrid()
	newt.uiGrid = c
	newt.uiControl = c
	newt.tw = a.Widget
	newt.WidgetType = toolkit.Grid
	newt.gridX = 0
	newt.gridY = 0

	t := andlabs[a.ParentId]
	place(a, t, newt)
}
