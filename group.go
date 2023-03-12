package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// TODO: treat a "Group" like a "Grid"
func (n *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = n.New(name, toolkit.Group, nil)
	send(n, newNode)
	return newNode
}
