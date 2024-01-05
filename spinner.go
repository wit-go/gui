package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/widget"
)

func (parent *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := parent.newNode(name, widget.Spinner)

	newNode.Custom = func() {
		log.Info("default NewSpinner() change", name)
	}

	newNode.X = x
	newNode.Y = y

	a := newAction(newNode, widget.Add)
	sendAction(a)
	return newNode
}
