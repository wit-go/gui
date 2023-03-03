package main

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// create a new box
func (t *andlabsT) getBox() *ui.Box {
	return t.uiBox
}

// create a new box
func (t *andlabsT) newBox() *andlabsT {
	log(debugToolkit, "newBox() START create default")
	t.Dump(debugToolkit)
	if (t.uiGroup != nil) {
		log(debugToolkit, "\tnewBox() is a Group")
		var newTK andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiGroup.SetChild(vbox)
		newTK.uiBox = vbox

		return &newTK
	}
	if (t.uiBox != nil) {
		log(debugToolkit, "\tnewBox() is a Box")
		var newTK andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiBox.Append(vbox, stretchy)
		newTK.uiBox = vbox
		newTK.Name = t.Name

		return &newTK
	}
	if (t.uiWindow != nil) {
		log(debugToolkit, "\tnewBox() is a Window")
		var newT andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiWindow.SetChild(vbox)
		newT.uiBox = vbox
		newT.Name = t.Name

		// panic("WTF")
		return &newT
	}
	log(debugToolkit, "\tnewBox() FAILED. Couldn't figure out where to make a box")
	t.Dump(debugToolkit)
	return nil
}
