package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func newGroup(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "NewGroup()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "NewGroup() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
	}
	newt := t.rawGroup(w.Name)
	mapWidgetsToolkits(w, newt)
}

// make new Group here
func (t *andlabsT) rawGroup(title string) *andlabsT {
	var newt andlabsT
	newt.Name = title

	log(debugToolkit, "NewGroup() create", newt.Name)

	g := ui.NewGroup(newt.Name)
	g.SetMargined(margin)
	newt.uiGroup = g

	t.doAppend(toolkit.Group, &newt, nil)

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(padded)
	g.SetChild(hbox)

	newt.uiBox = hbox
	newt.uiWindow = t.uiWindow
	newt.uiTab = t.uiTab

	return &newt
}

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

func doGroup(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	if (c.Action == "New") {
		newGroup(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Group() ct.broken", ct)
		return
	}
	if (ct.uiGroup == nil) {
		log(debugError, "Label() uiGroup == nil", ct)
		return
	}
	switch c.Action {
	case "Enable":
		ct.uiGroup.Enable()
	case "Disable":
		ct.uiGroup.Disable()
	case "Show":
		ct.uiGroup.Show()
	case "Hide":
		ct.uiGroup.Hide()
	case "Get":
		c.S = ct.uiGroup.Title()
	case "Set":
		ct.uiGroup.SetTitle(c.S)
	case "SetText":
		ct.uiGroup.SetTitle(c.S)
	case "SetMargin":
		ct.uiGroup.SetMargined(c.B)
	case "Destroy":
		ct.uiGroup.Destroy()
	default:
		log(debugError, "Can't do", c.Action, "to a Group")
	}
}
