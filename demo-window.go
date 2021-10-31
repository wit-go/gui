package gui

// import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var mybox *ui.Box

func (n *Node) AddGroup(title string) *Node {
	hbox := n.uiBox
	if (hbox == nil) {
		return n
	}
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	newNode := n.AddNode(title)
	newNode.uiBox = vbox
	return newNode
}
