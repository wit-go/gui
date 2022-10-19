package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

// TODO: which name is better. AddGroup or NewGroup ?
// first reaction is NewGroup
func (n *Node) NewGroup(name string) *Node {
	var newT *toolkit.Toolkit
	var gNode *Node
	log.Println("toolkit.NewGroup() START", name)

	if (n.Toolkit == nil) {
		log.Println("toolkit.NewGroup() Toolkit == nil")
		panic("Toolkit should never be nil")
	}

	// make a *Node with a *toolkit.Group
	gNode = n.New(name + " part1")
	newT = n.Toolkit.NewGroup(name)
	gNode.Toolkit = newT
	log.Println("################## gNode #######   ", name)
	gNode.Dump()

	return gNode
}

func (n *Node) AddGroup(title string) *Node {
	return n.NewGroup(title + "deprecated AddGroup")
}
