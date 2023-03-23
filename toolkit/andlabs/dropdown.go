package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newDropdown(w *toolkit.Widget) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "gui.Toolbox.newDropdown() START", w.Name)

	newt.tw = w
	s := ui.NewCombobox()
	newt.uiCombobox = s
	newt.uiControl = s

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
		t.uiCombobox.SetSelected(0)
	}
	t.c = t.c + 1
}

func (t *andlabsT) SetDropdown(i int) {
	t.uiCombobox.SetSelected(i)
}

func AddDropdownName(a *toolkit.Action) {
	log(debugToolkit, "gui.andlabs.AddDropdownName()", a.Widget.Name, "add:", a.S)

	t := mapToolkits[a.Widget]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.AddDropdownName() toolkit struct == nil. name=", a.Widget.Name, a.S)
		listMap(debugToolkit)
		return
	}
	t.AddDropdownName(a.S)
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

func newDropdown(a *toolkit.Action) {
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "gui.andlabs.newDropdown()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.newDropdown() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newDropdown(w)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
