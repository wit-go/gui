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
func newGrid(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT
	log(debugToolkit, "NewGrid()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		log(debugError, "ERROR newGrid() listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		return
	}

	log(debugToolkit, "NewGrid()", w.Name)
	if t.broken() {
		return
	}

	newt = new(andlabsT)

	c := ui.NewGrid()
	newt.uiGrid = c
	newt.uiBox = t.uiBox
	newt.tw = w
	t.doAppend(toolkit.Grid, newt, nil)
	mapWidgetsToolkits(w, newt)
}

func doGrid(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newGrid(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Grid() ct.broken", ct)
		return
	}
	if (ct.uiGrid == nil) {
		log(debugError, "Grid() uiGrid == nil", ct)
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	switch c.Action {
	case "Enable":
		ct.uiGrid.Enable()
	case "Disable":
		ct.uiGrid.Disable()
	case "Show":
		ct.uiGrid.Show()
	case "Hide":
		log(debugError, "trying Hide on grid")
		ct.uiGrid.Hide()
	case "SetMargin":
		log(debugError, "trying SetMargin on grid")
		ct.uiGrid.SetPadded(c.B)
	case "Set":
		log(debugError, "Can I use 'Set' to place a *Node in a Grid?")
	/*
	case "AddGrid":
		log(true, "how do I add a thing to a grid?")
		dump(p, c, true)
		newt.uiGrid.Append(button1,
			0, 2, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
	*/
	default:
		log(debugError, "Can't do", c.Action, "to a Grid")
	}
}
