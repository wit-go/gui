package main

import (
	"git.wit.org/wit/gui/toolkit"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newCheckbox(w *toolkit.Widget) *andlabsT {
	log(debugToolkit, "newCheckbox()", w.Name, w.Type)
	var newt andlabsT
	newt.tw = w

	if t.broken() {
		return nil
	}

	newt.uiCheckbox = ui.NewCheckbox(w.Name)
	newt.uiBox = t.uiBox
	// t.doAppend(&newt, *newt.uiCheckbox)
	t.uiBox.Append(newt.uiCheckbox, stretchy)

	newt.uiCheckbox.OnToggled(func(spin *ui.Checkbox) {
		newt.tw.B = newt.checked()
		log(debugChange, "val =", newt.tw.B)
		newt.commonChange(newt.tw)
	})

	return &newt
}

func (t *andlabsT) checked() bool {
	if t.broken() {
		return false
	}

	return t.uiCheckbox.Checked()
}

func newCheckbox(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "newCheckbox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		return
	}
	newt := t.newCheckbox(w)
	mapWidgetsToolkits(w, newt)
}

func doCheckbox(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newCheckbox(p, c)
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
