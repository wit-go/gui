package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (parent *Node) NewLabel(text string) *Node {
	newNode := parent.newNode(text, toolkit.Label)
	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
