package main

import (
	"git.wit.org/wit/gui/toolkit"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t andlabsT) NewCheckbox(w *toolkit.Widget) *andlabsT {
	log(debugToolkit, "NewCheckbox()", w.Name, w.Type)
	var newt andlabsT
	newt.tw = w

	if t.broken() {
		return nil
	}

	c := ui.NewCheckbox(w.Name)
	newt.uiCheckbox = c
	newt.uiBox = t.uiBox
	t.uiBox.Append(c, stretchy)

	c.OnToggled(func(spin *ui.Checkbox) {
		newt.tw.B = newt.Checked()
		log(debugChange, "val =", newt.tw.B)
		newt.commonChange(newt.tw)
	})

	return &newt
}

func (t andlabsT) Checked() bool {
	if t.broken() {
		return false
	}

	return t.uiCheckbox.Checked()
}

func NewCheckbox(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "NewCheckbox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		return
	}
	newt := t.NewCheckbox(w)
	mapWidgetsToolkits(w, newt)
}

func doCheckbox(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		NewCheckbox(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "checkbox() ct.broken", ct)
		return
	}
	if (ct.uiCheckbox == nil) {
		log(true, "checkbox() uiCheckbox == nil", ct)
		return
	}
	log(true, "Going to attempt:", c.Action)
	switch c.Action {
	case "Enable":
		ct.uiCheckbox.Enable()
	case "Disable":
		ct.uiCheckbox.Disable()
	case "Show":
		ct.uiCheckbox.Show()
	case "Hide":
		ct.uiCheckbox.Hide()
	case "Set":
		ct.uiCheckbox.SetText(c.S)
		ct.uiCheckbox.SetChecked(c.B)
	default:
		log(true, "Can't do", c.Action, "to a checkbox")
	}
}
