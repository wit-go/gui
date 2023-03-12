package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSpinner(w *toolkit.Widget) *andlabsT {
	// make new node here
	log(debugToolkit, "newSpinner()", w.X, w.Y)
	var newt andlabsT

	if (t.uiBox == nil) {
		log(debugToolkit, "newSpinner() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return nil
	}

	s := ui.NewSpinbox(w.X, w.Y)
	newt.uiSpinbox = s
	newt.uiBox = t.uiBox
	newt.tw = w
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(s *ui.Spinbox) {
		newt.tw.I = newt.uiSpinbox.Value()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func newSpinner(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugError, "NewSpinner() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.newSpinner(w)
	mapWidgetsToolkits(w, newt)
}

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

func doSpinner(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	if (c.Action == "New") {
		newSpinner(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Spinner() ct.broken", ct)
		return
	}
	if (ct.uiSpinbox == nil) {
		log(debugError, "Label() uiSpinbox == nil", ct)
		return
	}
	switch c.Action {
	case "Enable":
		ct.uiSpinbox.Enable()
	case "Disable":
		ct.uiSpinbox.Disable()
	case "Show":
		ct.uiSpinbox.Show()
	case "Hide":
		ct.uiSpinbox.Hide()
	case "Set":
		ct.uiSpinbox.SetValue(c.I)
	case "Get":
		c.I = ct.uiSpinbox.Value()
	default:
		log(debugError, "Can't do", c.Action, "to a Spinbox")
	}
}
