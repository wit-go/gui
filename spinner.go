package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := n.newNode(name, toolkit.Spinner, func() {
		log(debugChange, "default NewSpinner() change", name)
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
