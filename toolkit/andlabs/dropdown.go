package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func (p *node) newDropdown(n *node) {
	newt := new(andlabsT)
	log(debugToolkit, "gui.Toolbox.newDropdown() START", n.Name)

	cb := ui.NewCombobox()
	newt.uiCombobox = cb
	newt.uiControl = cb

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	cb.OnSelected(func(spin *ui.Combobox) {
		i := spin.Selected()
		if (newt.val == nil) {
			log(debugChange, "make map didn't work")
			newt.text = "error"
		}
		newt.s = newt.val[i]
		newt.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}

func (t *andlabsT) addDropdownName(title string) {
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

func (n *node) AddDropdownName(a *toolkit.Action) {
	log(debugToolkit, "gui.andlabs.AddDropdownName()", n.WidgetId, "add:", a.S)

	t := n.tk
	if (t == nil) {
		log(debugToolkit, "go.andlabs.AddDropdownName() toolkit struct == nil. name=", n.Name, a.S)
		listMap(debugToolkit)
		return
	}
	t.addDropdownName(a.S)
}

func (n *node) SetDropdownName(a *toolkit.Action, s string) {
	log(debugChange, "gui.andlabs.SetDropdown()", n.WidgetId, ",", s)

	t := n.tk
	if (t == nil) {
		log(debugError, "ERROR: SetDropdown() FAILED mapToolkits[w] == nil. name=", n.WidgetId, s)
		listMap(debugError)
		return
	}
	t.SetDropdown(1)
	// TODO: send back to wit/gui goroutine with the chan
	t.s = s
}
