package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = n.New(name, toolkit.Image, nil)
	send(n, newNode)
	return newNode
}
