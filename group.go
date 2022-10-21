package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

// TODO: which name is better. AddGroup or NewGroup ?
// first reaction is NewGroup
func (n *Node) NewGroup(name string) *Node {
	var newT *toolkit.Toolkit
	var gNode *Node
	log.Println("toolkit.NewGroup() START", name)

	if (n.toolkit == nil) {
		log.Println("toolkit.NewGroup() toolkit == nil")
		panic("toolkit should never be nil")
	}

	// make a *Node with a *toolkit.Group
	gNode = n.New(name)
	newT = n.toolkit.NewGroup(name)
	gNode.toolkit = newT
	gNode.Dump()

	return gNode
}

/*
func (n *Node) AddGroup(title string) *Node {
	return n.NewGroup(title + " deprecated AddGroup")
}
*/
