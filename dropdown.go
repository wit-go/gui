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
	n.AddText(name)
}

// Set the dropdown menu to 'name'
func (n *Node) SetDropdownName(name string) {
	n.SetText(name)
}

func (n *Node) NewDropdown(name string) *Node {
	newNode := n.New(name, toolkit.Dropdown, nil) 

	var a toolkit.Action
	a.ActionType = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}

func (n *Node) NewCombobox(name string) *Node {
	newNode := n.New(name, toolkit.Combobox, nil) 

	var a toolkit.Action
	a.ActionType = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}
