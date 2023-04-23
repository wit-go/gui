package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewBox(name string, b bool) *Node {
	newNode := n.newNode(name, toolkit.Box, nil)
	newNode.B = b

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = name
	a.Text = name
	a.B = b
	newaction(&a, newNode, n)

	return newNode
}
