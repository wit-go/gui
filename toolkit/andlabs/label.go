package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newLabel(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT
	log(debugToolkit, "NewLabel()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		log(debugError, "ERROR newLabel() listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		log(debugError, "ERROR FFFFFFFFFFFF listMap()")
		return
	}

	log(debugToolkit, "NewLabel()", w.Name)
	if t.broken() {
		return
	}

	newt = new(andlabsT)

	c := ui.NewLabel(w.Name)
	newt.uiLabel = c

	newt.uiBox = t.uiBox
	newt.tw = w
	if (defaultBehavior) {
		t.uiBox.Append(c, stretchy)
	}

	mapWidgetsToolkits(w, newt)
}

func doLabel(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newLabel(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "Label() ct.broken", ct)
		return
	}
	if (ct.uiLabel == nil) {
	
		log(true, "Label() uiLabel == nil", ct)
		return
	}
	log(true, "Going to attempt:", c.Action)
	switch c.Action {
	case "Enable":
		ct.uiLabel.Enable()
	case "Disable":
		ct.uiLabel.Disable()
	case "Show":
		ct.uiLabel.Show()
	case "Hide":
		ct.uiLabel.Hide()
	case "Set":
		ct.uiLabel.SetText(c.S)
	default:
		log(true, "Can't do", c.Action, "to a Label")
	}
}
