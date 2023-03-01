package main

import "git.wit.org/wit/gui/toolkit"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func newTextbox(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "NewTexbox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		log(debugError, "newTextbox() listMap()")
		log(debugError, "FFFFFFFFFFFF listMap()")
		log(debugError, "FFFFFFFFFFFF listMap()")
	}
	// t.NewTextbox(w)
// func (t andlabsT) NewTextbox(w *toolkit.Widget) *andlabsT {
	var newt *andlabsT
	newt = new(andlabsT)

	log(debugToolkit, "gui.Toolkit.NewTextbox()", w.Name)
	if t.broken() {
		return
	}

	c := ui.NewNonWrappingMultilineEntry()
	newt.uiMultilineEntry = c

	newt.uiBox = t.uiBox
	newt.Name = w.Name
	newt.tw = w
	if (defaultBehavior) {
		t.uiBox.Append(c, true)
	} else {
		t.uiBox.Append(c, stretchy)
	}

	c.OnChanged(func(spin *ui.MultilineEntry) {
		w.S = newt.uiMultilineEntry.Text()
		// this is still dangerous
		// newt.commonChange(newt.tw)
		log(debugChange, "Not yet safe to trigger on change for ui.MultilineEntry")
	})
	mapWidgetsToolkits(w, newt)
}


func doTextbox(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newTextbox(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "Textbox() ct.broken", ct)
		return
	}
	if (ct.uiMultilineEntry == nil) {
		log(debugError, "Textbox() uiMultilineEntry == nil", ct)
		return
	}
	// the dns control panel isn't crashing anymore (?)
	Queue(ct.doSimpleAction)
}

func (t *andlabsT) doSimpleAction() {
	if (t.tw == nil) {
		log(true, "doSimpleAction() got an empty widget")
		log(true, "THIS SHOULD NEVER HAPPEN")
		panic("crap. panic. widget == nil")
	}
	log(debugChange, "Going to attempt:", t.tw.Action)
	switch  t.tw.Action {
	case "Enable":
		t.uiMultilineEntry.Enable()
	case "Disable":
		t.uiMultilineEntry.Disable()
	case "Show":
		t.uiMultilineEntry.Show()
	case "Hide":
		t.uiMultilineEntry.Hide()
	case "Set":
		t.uiMultilineEntry.SetText(t.tw.S)
	default:
		log(debugError, "Can't do", t.tw.Action, "to a Textbox")
	}
}
