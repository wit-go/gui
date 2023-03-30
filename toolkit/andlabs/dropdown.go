package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newDropdown(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	w := a.Widget
	log(debugToolkit, "gui.Toolbox.newDropdown() START", a.Name)

	newt.tw = w
	newt.Type = w.Type
	newt.wId = a.WidgetId
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
		newt.commonChange(newt.tw, a.WidgetId)
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
	log(debugToolkit, "gui.andlabs.AddDropdownName()", a.WidgetId, "add:", a.S)

	t := andlabs[a.WidgetId]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.AddDropdownName() toolkit struct == nil. name=", a.Widget.Name, a.S)
		listMap(debugToolkit)
		return
	}
	t.AddDropdownName(a.S)
}

func SetDropdownName(a *toolkit.Action, s string) {
	log(debugChange, "gui.andlabs.SetDropdown()", a.WidgetId, ",", s)

	t := andlabs[a.WidgetId]
	if (t == nil) {
		log(debugError, "ERROR: SetDropdown() FAILED mapToolkits[w] == nil. name=", a.WidgetId, s)
		listMap(debugError)
		return
	}
	t.SetDropdown(1)
	// TODO: send back to wit/gui goroutine with the chan
	t.tw.S = s
}

func newDropdown(a *toolkit.Action) {
	log(debugToolkit, "gui.andlabs.newDropdown()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.newDropdown() toolkit struct == nil. name=", a.WidgetId)
		listMap(debugToolkit)
		return
	}
	newt := t.newDropdown(a)
	place(a, t, newt)
	// mapWidgetsToolkits(a, newt)
}
