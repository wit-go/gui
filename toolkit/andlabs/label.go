package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func NewLabel(b *ui.Box, name string) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewLabel", name)
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewLabel() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return &t
	}
	l := ui.NewLabel(name)
	t.uiLabel = l
	t.uiBox = b
	t.uiBox.Append(l, false)

	return &t
}
