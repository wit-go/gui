package gui

import "go.wit.com/gui/widget"

func (n *Node) Checked() bool {
	return widget.GetBool(n.value)
}

func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.newNode(name, widget.Checkbox)
	newNode.value = name
	newNode.progname = name

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}
	return newNode
}
