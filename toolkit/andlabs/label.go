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

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

func doLabel(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	if (c.Action == "New") {
		newLabel(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Label() ct.broken", ct)
		return
	}
	if (ct.uiLabel == nil) {
	
		log(debugError, "Label() uiLabel == nil", ct)
		return
	}
	switch c.Action {
	case "Enable":
		ct.uiLabel.Enable()
	case "Disable":
		ct.uiLabel.Disable()
	case "Show":
		ct.uiLabel.Show()
	case "Hide":
		ct.uiLabel.Hide()
	case "SetText":
		ct.uiLabel.SetText(c.S)
	case "Set":
		ct.uiLabel.SetText(c.S)
	default:
		log(debugError, "Can't do", c.Action, "to a Label")
	}
}
