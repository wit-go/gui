package main

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t andlabsT) NewCheckbox(name string) *andlabsT {
	log.Println("gui.Toolkit.NewCheckbox()", name)
	var newt andlabsT

	if t.broken() {
		return nil
	}

	c := ui.NewCheckbox(name)
	newt.uiCheckbox = c
	newt.uiBox = t.uiBox
	t.uiBox.Append(c, stretchy)

	c.OnToggled(func(spin *ui.Checkbox) {
		newt.commonChange("Checkbox")
	})

	return &newt
}

func (t andlabsT) Checked() bool {
	if t.broken() {
		return false
	}

	return t.uiCheckbox.Checked()
}
