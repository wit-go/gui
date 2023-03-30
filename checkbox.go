package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) Checked() bool {
	return n.widget.B
}

func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.New(name, toolkit.Checkbox, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = name
	a.Text = name
	newaction(&a, newNode, n)

	return newNode
}
