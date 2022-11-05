package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

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
		newt.commonChange("Spinner")
	})

	return &newt
}
