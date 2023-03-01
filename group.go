package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// TODO: which name is better. AddGroup or NewGroup ?
// first reaction is NewGroup
func (n *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = n.New(name, toolkit.Group, nil)
	send(n, newNode)
	return newNode
}
