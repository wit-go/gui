package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

// func NewSlider(b *ui.Box, name string *Toolkit {
func (t Toolkit) NewSlider(title string, x int, y int) *Toolkit {
	// make new node here
	log.Println("gui.Toolbox.NewSpinbox()", x, y)
	var newt Toolkit

	if (t.uiBox == nil) {
		log.Println("gui.ToolboxNode.NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log.Println("probably could just make a box here?")
		os.Exit(0)
		return nil
	}

	s := ui.NewSlider(x, y)
	newt.uiSlider = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, false)

	s.OnChanged(func(spin *ui.Slider) {
		i := spin.Value()
		log.Println("gui.Toolbox.ui.Slider.OnChanged() val =", i)
		if (DebugToolkit) {
			log.Println("gui.Toolbox.ui.OnChanged() val =", i)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(newt)
		}
		if (t.OnChanged != nil) {
			log.Println("gui.Toolbox.OnChanged() entered val =", i)
			newt.OnChanged(&newt)
		}
	})

	return &newt
}
