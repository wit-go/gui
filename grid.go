package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewGrid(name string, x int, y int) *Node {
	newNode := n.New(name, toolkit.Grid, func() {
		log(debugChange, "click() NewGrid not defined =", name)
	})
	newNode.widget.X = x
	newNode.widget.Y = y

	send(n, newNode)
	return newNode
}

// a box is just a grid with a single set of widgets that are either horizontal or vertical
func (n *Node) NewBox(name string, horizontal bool) *Node {
	var newNode *Node
	newNode = n.New(name, toolkit.Box, nil)

	newNode.widget.X = 3
	newNode.widget.Y = 1
	newNode.widget.B = horizontal

	send(n, newNode)
	return newNode
}

func (n *Node) AddGrid(a *Node, x int, y int) {
	n.widget.X = x
	n.widget.Y = y

	a.widget.Action = "AddGrid"
	send(n, a)
}
