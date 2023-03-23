package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewTextbox(name string) *Node {
	newNode := n.New(name, toolkit.Textbox, func() {
		log(debugGui, "NewTextbox changed =", name)
	})

	var a toolkit.Action
	a.Type = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}
