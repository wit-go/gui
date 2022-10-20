package gui

import (
	"log"
)

// This function should make a new node with the parent and
// the 'tab' as a child

func (n *Node) NewTab(title string) *Node {
	log.Println("gui.Node.NewTab() START name =", title)

	// TODO: standardize these checks somewhere
	if (n.toolkit == nil) {
		n.Dump()
		panic("NewTab() failed. toolkit == nil")
	}
	log.Println("Make new node")
	newN := n.New(title)
	log.Println("New tab to window")
	t := n.toolkit.AddTab(title)
	newN.toolkit = t

	n.Append(newN)
	return newN
}
