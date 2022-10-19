package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// create a new box
func (t *Toolkit) GetBox() *ui.Box {
	return t.uiBox
}

// create a new box
func (t *Toolkit) NewBox() *Toolkit {
	log.Println("gui.Toolbox.NewBox() START create default")
	t.Dump()
	if (t.uiGroup != nil) {
		log.Println("gui.Toolbox.NewBox() is a Group")
		var newTK Toolkit

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		t.uiGroup.SetChild(vbox)
		newTK.uiBox = vbox

		return &newTK
	}
	if (t.uiBox != nil) {
		log.Println("gui.Toolbox.NewBox() is a Box")
		// return t
	}
	log.Println("gui.Toolbox.NewBox() FAILED. Couldn't figure out where to make a box")
	t.Dump()
	return nil
}

// Make a new box
func MakeBox(name string) *Toolkit {
	var newt Toolkit

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(border)
	newt.uiBox = vbox
	newt.Name = name

	log.Println("gui.Toolbox.MakeBox() name =", name)
	newt.Dump()
	return &newt
}
