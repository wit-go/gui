package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSlider(w *toolkit.Widget) *andlabsT {
	// make new node here
	log(debugToolkit, w.Name, w.Type, w.X, w.Y)
	var newt andlabsT

	if (t.uiBox == nil) {
		log(debugToolkit, "node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugToolkit, "probably could just make a box here?")
		exit("internal golang wit/gui/toolkit error")
		return nil
	}

	s := ui.NewSlider(w.X, w.Y)
	newt.uiSlider = s
	newt.uiBox = t.uiBox
	newt.tw = w
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(spin *ui.Slider) {
		newt.tw.I = newt.uiSlider.Value()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func newSlider(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT
	log(debugToolkit, "newSlider()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugError, "newSlider() ERROR toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.newSlider(w)
	mapWidgetsToolkits(w, newt)
}

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

func doSlider(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	if (c.Action == "New") {
		newSlider(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Slider() ct.broken", ct)
		return
	}
	if (ct.uiSlider == nil) {
		log(debugError, "Label() uiSlider == nil", ct)
		return
	}
	switch c.Action {
	case "Enable":
		ct.uiSlider.Enable()
	case "Disable":
		ct.uiSlider.Disable()
	case "Show":
		ct.uiSlider.Show()
	case "Hide":
		ct.uiSlider.Hide()
	case "Set":
		ct.uiSlider.SetValue(c.I)
	case "Get":
		c.I = ct.uiSlider.Value()
	default:
		log(debugError, "Can't do", c.Action, "to a Slider")
	}
}
