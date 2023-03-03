package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newCombobox(w *toolkit.Widget) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "newCombobox() START", w.Name)

	if t.broken() {
		return nil
	}

	newt.tw = w
	s := ui.NewEditableCombobox()
	newt.uiEditableCombobox = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	s.OnChanged(func(spin *ui.EditableCombobox) {
		newt.tw.S = spin.Text()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func (t *andlabsT) AddComboboxName(title string) {
	t.uiEditableCombobox.Append(title)
	if (t.val == nil) {
		log(debugToolkit, "make map didn't work")
		return
	}
	t.val[t.c] = title

	// If this is the first menu added, set the dropdown to it
	// if (t.c == 0) {
	// }
	t.c = t.c + 1
}

func newCombobox(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "newCombobox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "newCombobox() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newCombobox(w)
	mapWidgetsToolkits(w, newt)
}

func doCombobox(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newCombobox(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "Combobox() ct.broken", ct)
		return
	}
	if (ct.uiEditableCombobox == nil) {
		log(true, "Combobox() uiEditableCombobox == nil", ct)
		return
	}
	log(true, "Going to attempt:", c.Action)
	switch c.Action {
	case "Add":
		ct.AddComboboxName(c.S)
	case "Enable":
		ct.uiEditableCombobox.Enable()
	case "Disable":
		ct.uiEditableCombobox.Disable()
	case "Show":
		ct.uiEditableCombobox.Show()
	case "Hide":
		ct.uiEditableCombobox.Hide()
	case "Set":
		ct.uiEditableCombobox.SetText(c.S)
	default:
		log(true, "Can't do", c.Action, "to a Combobox")
	}
}
