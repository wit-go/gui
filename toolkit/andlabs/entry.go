package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func NewEntry(b *ui.Box, name string) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewEntry", name)
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewEntry() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return &t
	}
	l := ui.NewEntry()
	t.uiEntry = l
	t.uiBox = b
	t.uiBox.Append(l, false)

	return &t
}
