package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

// func NewSlider(b *ui.Box, name string *Toolkit {
func (t Toolkit) NewSlider(title string, x int, y int) *Toolkit {
	// make new node here
	log.Println("gui.Toolkit.NewSpinbox()", x, y)
	var newt Toolkit

	if (t.uiBox == nil) {
		log.Println("gui.ToolkitNode.NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log.Println("probably could just make a box here?")
		os.Exit(0)
		return nil
	}

	s := ui.NewSlider(x, y)
	newt.uiSlider = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(spin *ui.Slider) {
		i := spin.Value()
		log.Println("gui.Toolkit.ui.Slider.OnChanged() val =", i)
		if (DebugToolkit) {
			log.Println("gui.Toolkit.ui.OnChanged() val =", i)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(newt)
		}
		if (newt.OnChanged != nil) {
			log.Println("gui.Toolkit.OnChanged() trying to run toolkit.OnChanged() entered val =", i)
			newt.OnChanged(&newt)
			return
		}
		if (newt.Custom != nil) {
			log.Println("gui.Toolkit.OnChanged() Running toolkit.Custom()")
			newt.Custom()
			return
		}
		log.Println("gui.Toolkit.OnChanged() ENDED without finding any callback")
	})

	return &newt
}
