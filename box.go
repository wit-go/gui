package gui

import (
	"go.wit.com/gui/widget"
)

func (parent *Node) NewBox(name string, b bool) *Node {
	newNode := parent.newNode(name, widget.Box)
	newNode.B = b

	a := newAction(newNode, widget.Add)
	sendAction(a)
	return newNode
}
