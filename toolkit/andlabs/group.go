package toolkit

import "log"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// make new Group here
func (t Toolkit) NewGroup(title string) *Toolkit {
	var newt Toolkit

	if (t.uiBox == nil) {
		log.Println("gui.ToolboxNode.NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log.Println("probably could just make a box here?")
		os.Exit(0)
		return nil
	}

	log.Println("gui.Toolbox.NewGroup() create", title)
	g := ui.NewGroup(title)
	g.SetMargined(margin)
	t.uiBox.Append(g, stretchy)

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(padded)
	g.SetChild(hbox)

	newt.uiGroup = g
	newt.uiBox = hbox
	newt.Name = title

	t.Dump()
	newt.Dump()
	// panic("toolkit.NewGroup")
	return &newt
}
