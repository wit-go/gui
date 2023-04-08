package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = n.newNode(name, toolkit.Image, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}
