package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (t Toolkit) NewCheckbox(name string) *Toolkit {
	log.Println("gui.Toolkit.NewCheckbox()", name)
	var newt Toolkit

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

func (t Toolkit) Checked() bool {
	if t.broken() {
		return false
	}

	return t.uiCheckbox.Checked()
}
