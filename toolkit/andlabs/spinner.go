package main

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t andlabsT) NewSpinner(title string, x int, y int) *andlabsT {
	// make new node here
	log(debugToolkit, "gui.Toolkit.NewSpinner()", x, y)
	var newt andlabsT

	if (t.uiBox == nil) {
		log(debugToolkit, "gui.ToolkitNode.NewSpinner() node.UiBox == nil. I can't add a range UI element without a place to put it")
		exit("internal golang wit/gui error")
		return nil
	}

	s := ui.NewSpinbox(x, y)
	newt.uiSpinbox = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(s *ui.Spinbox) {
		newt.commonChange("Spinner")
	})

	return &newt
}
