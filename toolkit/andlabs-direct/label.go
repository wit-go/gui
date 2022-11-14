package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t *Toolkit) NewLabel(name string) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewLabel", name)

	if t.broken() {
		return nil
	}
	var newt Toolkit
	newt.uiLabel = ui.NewLabel(name)
	newt.uiBox = t.uiBox
	t.uiBox.Append(newt.uiLabel, false)

	return &newt
}
