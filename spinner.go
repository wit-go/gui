package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := n.New(name, toolkit.Spinner, func() {
		log(debugChange, "default NewSpinner() change", name)
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
