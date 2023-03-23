package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) NewButton(name string, custom func()) *Node {
	newNode := n.New(name, toolkit.Button, custom)

	var a toolkit.Action
	a.Type = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, newNode, n)

	return newNode
}
