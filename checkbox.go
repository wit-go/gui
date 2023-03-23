package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) Checked() bool {
	return n.widget.B
}

func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.New(name, toolkit.Checkbox, nil)

	var a toolkit.Action
	a.Type = toolkit.Add
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a, newNode, n)
	newaction(&a, newNode, n)

	return newNode
}
