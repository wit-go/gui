package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := n.New(name, toolkit.Spinner, func() {
		log(debugGui, "even newer clicker() name in NewSpinner name =", name)
	})
	newNode.widget.X = x
	newNode.widget.Y = y

	send(n, newNode)
	return newNode
}
