package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (parent *Node) NewTextbox(name string) *Node {
	newNode := parent.newNode(name, toolkit.Textbox, func() {
		log(debugGui, "NewTextbox changed =", name)
	})

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
