package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// func NewLabel(b *ui.Box, name string) *Toolkit {

func (t *Toolkit) NewLabel(name string) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewLabel", name)

	if (t.uiBox == nil) {
		log.Println("gui.ToolboxNode.NewLabel() node.UiBox == nil. I can't add a range UI element without a place to put it")
		os.Exit(0)
		return nil
	}
	var newt Toolkit
	newt.uiLabel = ui.NewLabel(name)
	newt.uiBox = t.uiBox
	t.uiBox.Append(newt.uiLabel, false)
	log.Println("parent toolkit")
	t.Dump()
	log.Println("newt toolkit")
	newt.Dump()
	// panic("got here")

	return &newt
}
