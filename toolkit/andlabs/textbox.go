package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t Toolkit) NewTextbox(name string) *Toolkit {
	if (DebugToolkit) {
		log.Println("gui.Toolkit.NewTextbox()", name)
	}
	var newt Toolkit

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
