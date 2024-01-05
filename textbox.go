package gui

import (
	"go.wit.com/log"

	"go.wit.com/gui/widget"
)

func (parent *Node) NewTextbox(name string) *Node {
	newNode := parent.newNode(name, widget.Textbox)

	newNode.Custom = func() {
		log.Log(GUI, "NewTextbox changed =", name)
	}

	a := newAction(newNode, widget.Add)
	sendAction(a)
	return newNode
}

func (parent *Node) NewEntryLine(name string) *Node {
	newNode := parent.newNode(name, widget.Textbox)

	newNode.X = 1

	newNode.Custom = func() {
		log.Log(GUI, "NewTextbox changed =", name)
	}

	a := newAction(newNode, widget.Add)
	sendAction(a)
	return newNode
}
