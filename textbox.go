package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewTextbox(name string) *Node {
	newNode := n.newNode(name, toolkit.Textbox, func() {
		log(debugGui, "NewTextbox changed =", name)
	})

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = name
	a.Text = name
	newaction(&a, newNode, n)

	return newNode
}
