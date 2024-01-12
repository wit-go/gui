package gui

import (
	"go.wit.com/gui/widget"
)

// TODO: make a "Group" a "Grid" ?
// probably since right now group is just a 
// pre-canned andlabs/ui gtk,macos,windows thing
func (parent *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = parent.newNode(name, widget.Group)
	newNode.progname = name
	newNode.value = name

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}

	// by default, always pad groups
	newNode.Pad()

	// newBox := newNode.NewBox("defaultGroupBox", false)
	return newNode
}
