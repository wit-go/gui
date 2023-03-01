package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) Checked() bool {
	n.Dump()
	return n.widget.B
}

func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.New(name, toolkit.Checkbox, nil)
	send(n, newNode)
	return newNode
}

func (n *Node) NewThing(name string) *Node {
	newNode := n.New(name, toolkit.Button, nil)
	send(n, newNode)
	return newNode
}
