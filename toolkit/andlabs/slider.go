package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func NewSlider(b *ui.Box, name string, x int, y int) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewSpinbox()", x, y)
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewSpinbox() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return &t
	}
	s := ui.NewSlider(x, y)
	t.uiSlider = s
	t.uiBox = b
	t.uiBox.Append(s, false)

	s.OnChanged(func(spin *ui.Slider) {
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
