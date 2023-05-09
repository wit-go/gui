package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) Checked() bool {
	return n.B
}

func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.newNode(name, toolkit.Checkbox, nil)

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
