package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewLabel(text string) *Node {
	newNode := n.newNode(text, toolkit.Label, nil)

	n.Name = text
	n.Text = text
	a := newAction(n, toolkit.Add)
	sendAction(a, newNode, n)

	return newNode
}
