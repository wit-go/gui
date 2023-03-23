package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewImage(name string) *Node {
	var newNode *Node
	newNode = n.New(name, toolkit.Image, nil)

	var a toolkit.Action
	a.Type = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}
