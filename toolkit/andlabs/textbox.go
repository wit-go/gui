package main

import "git.wit.org/wit/gui/toolkit"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t andlabsT) NewTextbox(name string) *andlabsT {
	var newt andlabsT

	log(debugToolkit, "gui.Toolkit.NewTextbox()", name)
	if t.broken() {
		return nil
	}

	c := ui.NewNonWrappingMultilineEntry()
	newt.uiMultilineEntry = c

	newt.uiBox = t.uiBox
	newt.Name = name
	if (defaultBehavior) {
		t.uiBox.Append(c, true)
	} else {
		t.uiBox.Append(c, stretchy)
	}

	c.OnChanged(func(spin *ui.MultilineEntry) {
		newt.commonChange("Textbox")
	})

	return &newt
}

func NewTextbox(parentW *toolkit.Widget, w *toolkit.Widget) {
	var t, newt *andlabsT
	log(debugToolkit, "gui.andlabs.NewTextbox()", w.Name)

	t = mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewTextbox() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}

	if t.broken() {
		return
	}
	newt = new(andlabsT)

	newt.uiLabel = ui.NewLabel(w.Name)
	newt.uiBox = t.uiBox

	log(debugToolkit, "gui.Toolbox.NewTextbox() about to append to Box parent t:", w.Name)
	t.Dump()
	log(debugToolkit, "gui.Toolbox.NewTextbox() about to append to Box new t:", w.Name)
	newt.Dump()

	if (t.uiBox != nil) {
		t.uiBox.Append(newt.uiLabel, false)
	} else {
		log(debugToolkit, "ERROR: wit/gui andlabs couldn't place this Textbox in a box")
		return
	}

	mapWidgetsToolkits(w, newt)
}
