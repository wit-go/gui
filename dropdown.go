package gui

// functions to create 'Dropdown' and 'Combobox'
// Combobox is a Dropdown you can edit
// Thererfore, AddDropdownName() is used on both combobox and dropdown nodes
// since it is the same. confusing names? maybe...

import (
	"git.wit.org/wit/gui/toolkit"
)

// add a new entry to the dropdown name
func (n *Node) AddDropdownName(name string) {
	n.Add(name)
}

// Set the dropdown menu to 'name'
func (n *Node) SetDropdownName(name string) {
	n.SetText(name)
}

func (n *Node) NewDropdown(name string) *Node {
	newNode := n.New(name, toolkit.Dropdown, nil) 
	send(n, newNode)
	return newNode
}

func (n *Node) NewCombobox(name string) *Node {
	newNode := n.New(name, toolkit.Combobox, nil) 
	send(n, newNode)
	return newNode
}
