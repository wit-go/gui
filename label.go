package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewLabel(text string) *Node {
	newNode := n.New(text, toolkit.Label, nil)

	var a toolkit.Action
	a.Type = toolkit.Add
	newaction(&a, newNode, n)

	return newNode
}
