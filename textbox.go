package gui

import (
	"go.wit.com/log"

	"go.wit.com/gui/gui/toolkit"
)

func (parent *Node) NewTextbox(name string) *Node {
	newNode := parent.newNode(name, toolkit.Textbox)

	newNode.Custom = func() {
		log.Log(GUI, "NewTextbox changed =", name)
	}

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}

func (parent *Node) NewEntryLine(name string) *Node {
	newNode := parent.newNode(name, toolkit.Textbox)

	newNode.X = 1

	newNode.Custom = func() {
		log.Log(GUI, "NewTextbox changed =", name)
	}

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
