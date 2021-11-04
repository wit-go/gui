package gui

import "log"
// import "os"
// import "reflect"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

func (n *Node) AddComboBox(title string, s ...string) *Node {
	newNode := n.AddNode(title)
	box := n.uiBox
	if (box == nil) {
		return n
	}

	ecbox := ui.NewEditableCombobox()
	newNode.uiText = ecbox
	// newNode.Dump()
	// log.Println("ecbox", ecbox)

	for id, name := range s {
		log.Println("Adding Combobox Entry:", id, name)
		ecbox.Append(name)
	}

	ecbox.OnChanged(func(*ui.EditableCombobox) {
		test := ecbox.Text()
		log.Println("node.Name = '" + newNode.Name + "' text for '" + title + "' is now: '" + test + "'")
		if (newNode.OnChanged == nil) {
			log.Println("Not doing custom OnChanged since OnChanged == nil")
			newNode.Dump()
		} else {
			newNode.OnChanged()
		}
	})

	box.Append(ecbox, false)
	return newNode
}

func (n *Node) GetText() string {
	if (n.uiText == nil) {
		return ""
	}
	ecbox := n.uiText
	return ecbox.Text()
}
