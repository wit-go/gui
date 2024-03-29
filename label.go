package gui

import (
	"go.wit.com/gui/widget"
)

func (parent *Node) NewLabel(text string) *Node {
	newNode := parent.newNode(text, widget.Label)
	newNode.value = text
	newNode.progname = text

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}
	return newNode
}
