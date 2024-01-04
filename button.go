package gui

import "go.wit.com/gui/toolkits"

func (parent *Node) NewButton(name string, custom func()) *Node {
	newNode := parent.newNode(name, toolkit.Button)
	newNode.Custom = custom

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}

// find widget by number
func (n *Node) FindId(i int) (*Node) {
	if (n == nil) {
		return nil
	}

	if (n.id == i) {
		return n
	}

	for _, child := range n.children {
		newN := child.FindId(i)
		if (newN != nil) {
			return newN
		}
	}
	return nil
}
