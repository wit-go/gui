package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// add a new entry to the dropdown name
func (n *Node) AddDropdownName(name string) {
	for _, aplug := range allPlugins {
		log(debugPlugin, "AddDropdownName() aplug =", aplug.name, "name =", name)
		if (aplug.AddDropdownName == nil) {
			log(debugPlugin, "\taplug.AddDropdownName() = nil")
			continue
		}
		aplug.AddDropdownName(&n.widget, name)
	}
}

// Set the dropdown menu to 'name'
func (n *Node) SetDropdownName(name string) {
	log(debugGui, "SetDropdownName() work. name =", name)
	for _, aplug := range allPlugins {
		log(debugPlugin, "SetDropdownName() aplug =", aplug.name, "name =", name)
		if (aplug.SetDropdownName == nil) {
			log(true, "\taplug.SetDropdownName() aplug = nil")
			continue
		}
		aplug.SetDropdownName(&n.widget, name)
	}
}

func (n *Node) NewDropdown(name string) *Node {
	newNode := n.New(name, toolkit.Dropdown, nil) 

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewDropdown() aplug =", aplug.name, "name =", newNode.widget.Name)
		if (aplug.NewDropdown == nil) {
			log(debugGui, "\tgui.NewDropdown() aplug.NewDropdown = nil", aplug.name)
			continue
		}
		aplug.NewDropdown(&n.widget, &newNode.widget)
	}

	// TODO, this doesn't work for some reason (over-written by plugin?)
	return newNode
}
