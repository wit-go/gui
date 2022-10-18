package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// make new Group here
func NewGroup(b *ui.Box, title string) *Toolkit {
	var t Toolkit

	if (b == nil) {
		log.Println("gui.ToolboxNode.NewSpinbox() node.UiBox == nil. I can't add a range UI element without a place to put it")
		return &t
	}

	log.Println("gui.Toolbox.NewGroup() create", title)
	g := ui.NewGroup(title)
	g.SetMargined(true)
	t.uiGroup = g
	t.uiBox = b
	t.uiBox.Append(g, false)

	return &t
}

// create a new box
func (t *Toolkit) GetBox() *ui.Box {
	return t.uiBox
}

// create a new box
func (t *Toolkit) NewBox() *Toolkit {
	log.Println("gui.Toolbox.NewBox() START create default")
	if (t.uiGroup != nil) {
		log.Println("gui.Toolbox.NewBox() is a Group")
		var newTK Toolkit

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		t.uiGroup.SetChild(vbox)
		newTK.uiBox = vbox

		return &newTK
	}
	log.Println("gui.Toolbox.NewBox() FAILED")
	return nil
}

/*
func (n *Node) NewGroup(title string) *Node {
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)
*/
