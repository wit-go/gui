package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

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
		newt.commonChange("Slider")
	})

	return &newt
}
