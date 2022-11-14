package main

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// create a new box
func (t *andlabsT) GetBox() *ui.Box {
	return t.uiBox
}

// create a new box
func (t *andlabsT) NewBox() *andlabsT {
	if (DebugToolkit) {
		log.Println("gui.Toolbox.NewBox() START create default")
	}
	t.Dump()
	if (t.uiGroup != nil) {
		if (DebugToolkit) {
			log.Println("\tgui.Toolbox.NewBox() is a Group")
		}
		var newTK andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiGroup.SetChild(vbox)
		newTK.uiBox = vbox

		return &newTK
	}
	if (t.uiBox != nil) {
		if (DebugToolkit) {
			log.Println("\tgui.Toolbox.NewBox() is a Box")
		}
		var newTK andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiBox.Append(vbox, stretchy)
		newTK.uiBox = vbox
		newTK.Name = t.Name

		return &newTK
	}
	if (t.uiWindow != nil) {
		if (DebugToolkit) {
			log.Println("\tgui.Toolbox.NewBox() is a Window")
		}
		var newT andlabsT

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiWindow.SetChild(vbox)
		newT.uiBox = vbox
		newT.Name = t.Name

		// panic("WTF")
		return &newT
	}
	if (DebugToolkit) {
		log.Println("\tgui.Toolbox.NewBox() FAILED. Couldn't figure out where to make a box")
	}
	t.Dump()
	return nil
}
