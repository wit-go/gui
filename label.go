package gui

import (
	"go.wit.com/gui/gui/toolkit"
)

func (parent *Node) NewLabel(text string) *Node {
	newNode := parent.newNode(text, toolkit.Label)
	a := newAction(newNode, toolkit.Add)
	a.Text = text
	a.S = text
	sendAction(a)
	return newNode
}
