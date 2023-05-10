package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// Grid numbering by (X,Y)
// -----------------------------
// -- (1,1) -- (2,1) -- (3,1) --
// -- (1,2) -- (2,1) -- (3,1) --
// -----------------------------
func (p *node) newGrid(n *node) {
	var newt *andlabsT
	log(debugToolkit, "newGrid()", n.WidgetId, "to", n.ParentId)

	newt = new(andlabsT)

	c := ui.NewGrid()
	newt.uiGrid = c
	newt.uiControl = c

	n.tk = newt
	p.place(n)
}
