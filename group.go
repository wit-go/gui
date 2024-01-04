package gui

import (
	"go.wit.com/gui/toolkits"
)

// TODO: make a "Group" a "Grid" ?
// probably since right now group is just a 
// pre-canned andlabs/ui gtk,macos,windows thing
func (parent *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = parent.newNode(name, toolkit.Group)

	a := newAction(newNode, toolkit.Add)
	sendAction(a)

	// by default, always pad groups
	newNode.Pad()

	newBox := newNode.NewBox("defaultGroupBox", false)
	return newBox
}
