package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// TODO: make a "Group" a "Grid" ?
// probably since right now group is just a 
// pre-canned andlabs/ui gtk,macos,windows thing
func (parent *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = parent.newNode(name, toolkit.Group, nil)

	a := newAction(newNode, toolkit.Add)
	sendAction(a, newNode, parent)

	newBox := newNode.NewBox("defaultGroupBox", false)
	return newBox
}
