package gui

import (
	"go.wit.com/gui/widget"
)

func (parent *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = parent.newNode(name, widget.Image)

	a := newAction(newNode, widget.Add)
	sendAction(a)
	return newNode
}
