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
	t.doAppend(newt, nil)
	/*
	if (defaultBehavior) {
		t.uiBox.Append(c, stretchy)
	}

	button1 := ui.NewButton("a(0,0)")
	c.Append(button1,
                0, 0, 1, 1,
                false, ui.AlignFill, false, ui.AlignFill)

	button2 := ui.NewButton("a(1,0)")
	c.Append(button2,
                1, 0, 1, 1,
                false, ui.AlignFill, false, ui.AlignFill)
	*/

	// Append(child Control, 
	//	left, top int, 
	// xspan, yspan int, 
	// hexpand bool, halign Align, 
	// vexpand bool, valign Align) {

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
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "Grid() ct.broken", ct)
		return
	}
	if (ct.uiGrid == nil) {
	
		log(true, "Grid() uiGrid == nil", ct)
		return
	}
	log(true, "Going to attempt:", c.Action)
	switch c.Action {
	case "Enable":
		ct.uiGrid.Enable()
	case "Disable":
		ct.uiGrid.Disable()
	case "Show":
		ct.uiGrid.Show()
	case "Hide":
		ct.uiGrid.Hide()
	case "Set":
		log(true, "Can I use 'Set' to place a *Node in a Grid?")
	default:
		log(true, "Can't do", c.Action, "to a Grid")
	}
}
