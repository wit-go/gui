package gui

import (
	"go.wit.com/gui/widget"
)

func (parent *Node) NewLabel(text string) *Node {
	newNode := parent.newNode(text, widget.Label)
	a := newAction(newNode, widget.Add)
	a.Text = text
	a.S = text
	sendAction(a)
	return newNode
}
