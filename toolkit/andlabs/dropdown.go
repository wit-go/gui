package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func (p *node) newDropdown(n *node) {
	newt := new(guiWidget)
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
			log(logError, "make map didn't work")
			n.S = "map did not work. ui.Combobox error"
		} else {
			n.S = newt.val[i]
		}
		n.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}

func (t *guiWidget) addDropdownName(title string) {
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

func (t *guiWidget) SetDropdown(i int) {
	t.uiCombobox.SetSelected(i)
}

func (n *node) AddDropdownName(s string) {
	log(logInfo, "AddDropdownName()", n.WidgetId, "add:", s)

	t := n.tk
	if (t == nil) {
		log(logInfo, "AddDropdownName() toolkit struct == nil. name=", n.Name, s)
		return
	}
	t.addDropdownName(s)
}

func (n *node) SetDropdownName(a *toolkit.Action, s string) {
	log(logInfo, "SetDropdown()", n.WidgetId, ",", s)

	t := n.tk
	if (t == nil) {
		log(debugError, "SetDropdown() FAILED mapToolkits[w] == nil. name=", n.WidgetId, s)
		return
	}
	t.SetDropdown(1)
	// TODO: send back to wit/gui goroutine with the chan
	n.S = s
}
