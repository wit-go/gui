package gui

import "go.wit.com/gui/widget"

func (n *Node) Checked() bool {
	return n.B
}

func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.newNode(name, widget.Checkbox)

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}
	return newNode
}
