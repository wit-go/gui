package gui

import (
	"go.wit.com/gui/widget"
)

func (parent *Node) NewBox(name string, b bool) *Node {
	newNode := parent.newNode(name, widget.Box)

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		if b {
			a.Direction = widget.Horizontal
		} else {
			a.Direction = widget.Vertical
		}
		sendAction(a)
	}
	return newNode
}

func (parent *Node) NewHorizontalBox(name string) *Node {
	newNode := parent.newNode(name, widget.Box)

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		a.Direction = widget.Horizontal
		sendAction(a)
	}
	return newNode
}

func (parent *Node) NewVerticalBox(name string) *Node {
	newNode := parent.newNode(name, widget.Box)

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		a.Direction = widget.Vertical
		sendAction(a)
	}
	return newNode
}
