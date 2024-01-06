package gui

// functions to create 'Dropdown' and 'Combobox'
// Combobox is a Dropdown you can edit
// Thererfore, AddDropdownName() is used on both combobox and dropdown nodes
// since it is the same. confusing names? maybe...

import (
	"go.wit.com/gui/widget"
)

// add a new entry to the dropdown name
func (n *Node) AddDropdownName(name string) {
	n.AddText(name)
}

// Set the dropdown menu to 'name'
func (n *Node) SetDropdownName(name string) {
	n.SetText(name)
}

func (n *Node) NewDropdown(name string) *Node {
	newNode := n.newNode(name, widget.Dropdown) 

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}

	return newNode
}

func (n *Node) NewCombobox(name string) *Node {
	newNode := n.newNode(name, widget.Combobox) 

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}

	return newNode
}
