package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/widget"
)

func (parent *Node) NewSlider(name string, x int, y int) *Node {
	newNode := parent.newNode(name, widget.Slider)

	newNode.Custom = func() {
		log.Log(GUI, "even newer clicker() name in NewSlider name =", name)
	}

	newNode.X = x
	newNode.Y = y
	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		a.X = x
		a.Y = y
		sendAction(a)
	}

	return newNode
}
