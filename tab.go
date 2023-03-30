package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// This function should make a new node with the parent and
// the 'tab' as a child

func (n *Node) NewTab(text string) *Node {
	newNode := n.New(text, toolkit.Tab, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = text
	a.Text = text
	newaction(&a, newNode, n)

	newBox := newNode.NewBox(text, true)
	return newBox
}
