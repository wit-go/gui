package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = n.newNode(name, toolkit.Image, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	newaction(&a, newNode, n)

	return newNode
}
