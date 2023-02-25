package main

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "git.wit.org/wit/gui/toolkit"

func (t *andlabsT) NewDropdown(title string) *andlabsT {
	// make new node here
	var newt andlabsT

	log(debugToolkit, "gui.Toolbox.NewDropdownCombobox()", title)

	if t.broken() {
		return nil
	}

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
			log(debugToolkit, "make map didn't work")
			newt.text = "error"
		}
		newt.text = newt.val[i]
		newt.commonChange("Dropdown")
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
	t.c = t.c + 1
}

func (t andlabsT) SetDropdown(i int) {
	t.uiCombobox.SetSelected(i)
}

func NewDropdown(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "gui.andlabs.NewDropdown()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewDropdown() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap()
	}
	newt := t.NewDropdown(w.Name)
	newt.Custom = w.Custom
	newt.tw = w
	mapWidgetsToolkits(w, newt)
}

func AddDropdownName(w *toolkit.Widget, s string) {
	log(debugToolkit, "gui.andlabs.AddDropdownName()", w.Name, "add:", s)

	t := mapToolkits[w]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.AddDropdownName() toolkit struct == nil. name=", w.Name, s)
		listMap()
	}
	t.AddDropdownName(s)
}

func SetDropdown(w *toolkit.Widget, i int) {
	log(debugToolkit, "gui.andlabs.SetDropdown()", i)

	t := mapToolkits[w]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.SetDropdown() toolkit struct == nil. name=", w.Name, i)
		listMap()
	}
	t.SetDropdown(i)
}
