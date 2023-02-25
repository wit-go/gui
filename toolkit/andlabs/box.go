package main

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// create a new box
func (t *andlabsT) GetBox() *ui.Box {
	return t.uiBox
}

// create a new box
func (t *andlabsT) NewBox() *andlabsT {
	log(debugToolkit, "gui.Toolbox.NewBox() START create default")
	t.Dump()
	if (t.uiGroup != nil) {
		log(debugToolkit, "\tgui.Toolbox.NewBox() is a Group")
		var newTK andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiGroup.SetChild(vbox)
		newTK.uiBox = vbox

		return &newTK
	}
	if (t.uiBox != nil) {
		log(debugToolkit, "\tgui.Toolbox.NewBox() is a Box")
		var newTK andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiBox.Append(vbox, stretchy)
		newTK.uiBox = vbox
		newTK.Name = t.Name

		return &newTK
	}
	if (t.uiWindow != nil) {
		log(debugToolkit, "\tgui.Toolbox.NewBox() is a Window")
		var newT andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiWindow.SetChild(vbox)
		newT.uiBox = vbox
		newT.Name = t.Name

		// panic("WTF")
		return &newT
	}
	log(debugToolkit, "\tgui.Toolbox.NewBox() FAILED. Couldn't figure out where to make a box")
	t.Dump()
	return nil
}
