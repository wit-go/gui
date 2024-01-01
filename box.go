package gui

import (
	"go.wit.com/gui/gui/toolkit"
)

func (parent *Node) NewBox(name string, b bool) *Node {
	newNode := parent.newNode(name, toolkit.Box)
	newNode.B = b

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
