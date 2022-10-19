package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func NewSpinbox(b *ui.Box, name string, x int, y int) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewSpinbox()", x, y)
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewSpinbox() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return nil
	}
	spin := ui.NewSpinbox(x, y)
	t.uiSpinbox = spin
	t.uiBox = b
	t.uiBox.Append(spin, false)

	spin.OnChanged(func(spin *ui.Spinbox) {
		i := spin.Value()
		if (DebugToolkit) {
			log.Println("gui.Toolbox.ui.OnChanged() val =", i)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(t)
		}
		if (t.OnChanged != nil) {
			log.Println("gui.Toolbox.OnChanged() entered val =", i)
			t.OnChanged(&t)
		}
	})

	return &t
}
