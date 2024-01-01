package gui

import (
	"go.wit.com/gui/gui/toolkit"
)

func (parent *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = parent.newNode(name, toolkit.Image)

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
