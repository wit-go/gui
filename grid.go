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
