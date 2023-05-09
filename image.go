package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (parent *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = parent.newNode(name, toolkit.Image, nil)

	a := newAction(newNode, toolkit.Add)
	sendAction(a, newNode, parent)
	return newNode
}
