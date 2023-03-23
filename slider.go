package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewSlider(name string, x int, y int) *Node {
	newNode := n.New(name, toolkit.Slider, func() {
		log(debugGui, "even newer clicker() name in NewSlider name =", name)
	})

	var a toolkit.Action
	a.Type = toolkit.Add
	a.X = x
	a.Y = y
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}
