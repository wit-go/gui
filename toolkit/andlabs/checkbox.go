package main

import "git.wit.org/wit/gui/toolkit"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t andlabsT) NewCheckbox(name string, f func()) *andlabsT {
	log(debugToolkit, "gui.Toolkit.NewCheckbox()", name)
	var newt andlabsT

	if t.broken() {
		return nil
	}

	c := ui.NewCheckbox(name)
	newt.uiCheckbox = c
	newt.uiBox = t.uiBox
	t.uiBox.Append(c, stretchy)
	// newt.Custom = f

	c.OnToggled(func(spin *ui.Checkbox) {
		// log(debugToolkit, "gui.Toolkit.NewCheckbox() clicked", name)
		newt.commonChange("Checkbox")
		/*
		if (f != nil) {
			log(debugToolkit, "Run custom() here", f)
			log(SPEW, f)
			f()
		} else {
			log(debugToolkit, "No custom() function here")
		}
		*/
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
	log(debugToolkit, "gui.andlabs.NewCheckbox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap()
	}
	newt := t.NewCheckbox(w.Name, w.Custom)
	newt.Custom = w.Custom
	/*
	if (w.Custom != nil) {
		log(true, "go.andlabs.NewCheckbox() toolkit struct == nil. name=", parentW.Name, w.Name)
		log(true, "Run custom() START here", w.Custom)
		w.Custom()
		log(true, "Run custom() END")
		// exit("ran it here")
	} else {
		log(true, "No custom() function here")
		// exit("nothing here")
	}
	*/
	mapWidgetsToolkits(w, newt)
}
