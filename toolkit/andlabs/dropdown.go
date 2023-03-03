package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newDropdown(w *toolkit.Widget) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "gui.Toolbox.newDropdown() START", w.Name)

	if t.broken() {
		return nil
	}

	newt.tw = w
	s := ui.NewCombobox()
	newt.uiCombobox = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	s.OnSelected(func(spin *ui.Combobox) {
		i := spin.Selected()
		if (newt.val == nil) {
			log(debugChange, "make map didn't work")
			newt.text = "error"
		}
		newt.tw.S = newt.val[i]
		newt.commonChange(newt.tw)
	})

	return &newt
}

func (t *andlabsT) AddDropdownName(title string) {
	t.uiCombobox.Append(title)
	if (t.val == nil) {
		log(debugToolkit, "make map didn't work")
		return
	}
	t.val[t.c] = title

	// If this is the first menu added, set the dropdown to it
	if (t.c == 0) {
		log(debugChange, "THIS IS THE FIRST Dropdown", title)
		t.uiCombobox.SetSelected(1)
	}
	t.c = t.c + 1
}

func (t *andlabsT) SetDropdown(i int) {
	t.uiCombobox.SetSelected(i)
}

func AddDropdownName(w *toolkit.Widget, s string) {
	log(debugToolkit, "gui.andlabs.AddDropdownName()", w.Name, "add:", s)

	t := mapToolkits[w]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.AddDropdownName() toolkit struct == nil. name=", w.Name, s)
		listMap(debugToolkit)
		return
	}
	t.AddDropdownName(s)
}

func SetDropdownName(w *toolkit.Widget, s string) {
	log(debugChange, "gui.andlabs.SetDropdown()", w.Name, ",", s)

	t := mapToolkits[w]
	if (t == nil) {
		log(debugError, "ERROR: SetDropdown() FAILED mapToolkits[w] == nil. name=", w.Name, s)
		listMap(debugError)
		return
	}
	t.SetDropdown(1)
	t.tw.S = s
}

func newDropdown(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "gui.andlabs.newDropdown()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.newDropdown() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newDropdown(w)
	mapWidgetsToolkits(w, newt)
}

func doDropdown(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newDropdown(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "Dropdown() ct.broken", ct)
		return
	}
	if (ct.uiCombobox == nil) {
		log(true, "Dropdown() uiCombobox == nil", ct)
		return
	}
	log(true, "Going to attempt:", c.Action)
	switch c.Action {
	case "Add":
		ct.AddDropdownName(c.S)
		// ct.uiCombobox.Enable()
	case "Enable":
		ct.uiCombobox.Enable()
	case "Disable":
		ct.uiCombobox.Disable()
	case "Show":
		ct.uiCombobox.Show()
	case "Hide":
		ct.uiCombobox.Hide()
	case "Set":
		ct.uiCombobox.SetSelected(1)
	default:
		log(true, "Can't do", c.Action, "to a Dropdown")
	}
}
