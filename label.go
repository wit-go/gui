package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewLabel(text string) *Node {
	newNode := n.New(text, toolkit.Label, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = text
	a.Text = text
	newaction(&a, newNode, n)

	return newNode
}
