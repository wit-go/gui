package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewSlider(name string, x int, y int) *Node {
	newNode := n.New(name, toolkit.Slider, func() {
		log(debugGui, "even newer clicker() name in NewSlider name =", name)
	})

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.X = x
	a.Y = y
	a.Name = name
	a.Text = name
	newaction(&a, newNode, n)

	return newNode
}
