package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

// func NewSlider(b *ui.Box, name string *Toolkit {
func (t Toolkit) NewSpinner(title string, x int, y int) *Toolkit {
	// make new node here
	log.Println("gui.Toolkit.NewSpinner()", x, y)
	var newt Toolkit

	if (t.uiBox == nil) {
		log.Println("gui.ToolkitNode.NewSpinner() node.UiBox == nil. I can't add a range UI element without a place to put it")
		os.Exit(0)
		return nil
	}

	s := ui.NewSpinbox(x, y)
	newt.uiSpinbox = s
	newt.uiBox = t.uiBox
	t.uiBox.Append(s, stretchy)

	s.OnChanged(func(s *ui.Spinbox) {
		i := s.Value()
		if (DebugToolkit) {
			log.Println("gui.Toolkit.ui.OnChanged() val =", i)
			scs := spew.ConfigState{MaxDepth: 1}
			scs.Dump(newt)
		}
		if (t.OnChanged != nil) {
			log.Println("gui.Toolkit.OnChanged() entered val =", i)
			newt.OnChanged(&newt)
		}
	})

	return &newt
}
