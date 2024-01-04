package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/toolkits"
)

func (parent *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := parent.newNode(name, toolkit.Spinner)

	newNode.Custom = func() {
		log.Info("default NewSpinner() change", name)
	}

	newNode.X = x
	newNode.Y = y

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
