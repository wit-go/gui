package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (parent *Node) NewSlider(name string, x int, y int) *Node {
	newNode := parent.newNode(name, toolkit.Slider)

	newNode.Custom = func() {
		log(debugGui, "even newer clicker() name in NewSlider name =", name)
	}

	newNode.X = x
	newNode.Y = y
	a := newAction(newNode, toolkit.Add)
	a.X = x
	a.Y = y
	sendAction(a)

	return newNode
}
