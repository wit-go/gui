package gui

import "log"

func (n *Node) AddButton(name string, custom func(*Node)) *Node {
	if (n.toolkit == nil) {
		log.Println("gui.Node.AppendButton() filed node.toolkit == nil")
		panic("gui.Node.AppendButton() filed node.toolkit == nil")
		return n
	}
	newNode := n.New(name)
	newNode.toolkit = n.toolkit.NewButton(name)
	newNode.toolkit.Custom = func() {
		log.Println("gui.AppendButton() Button Clicked. Running custom()")
		custom(newNode)
	}
	newNode.custom = custom

	return newNode
}
