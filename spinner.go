package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (parent *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := parent.newNode(name, toolkit.Spinner, func() {
		log(debugChange, "default NewSpinner() change", name)
	})

	newNode.X = x
	newNode.Y = y

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
