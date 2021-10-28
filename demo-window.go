package gui

import "log"
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

func (n *Node) GetText(name string) string {
	if (n.uiText != nil) {
		return n.uiText.Text()
	}
	return n.Name
}

func (n *Node) AddComboBox(title string, s ...string) *Node {
	box := n.uiBox
	if (box == nil) {
		return n
	}

	ecbox := ui.NewEditableCombobox()

	for id, name := range s {
		log.Println("Adding Combobox Entry:", id, name)
		ecbox.Append(name)
	}

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		test := ecbox.Text()
		log.Println("node.Name = '" + n.Name + "' text for '" + title + "' is now: '" + test + "'")
	})

	box.Append(ecbox, false)

	newNode := n.AddNode(title)
	newNode.uiText = ecbox
	return newNode
}
