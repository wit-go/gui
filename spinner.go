package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	var newT *toolkit.Toolkit
	var sNode *Node

	log.Println("toolkit.NewSpinner() START", name)

	if (n.toolkit == nil) {
		log.Println("toolkit.NewSpinner() toolkit == nil")
		panic("toolkit should never be nil")
	}

	// make a *Node with a *toolkit.Group
	sNode = n.New(name + " part1")
	newT = n.toolkit.NewSpinner(name, x, y)
	newT.Name = name
	sNode.toolkit = newT
	sNode.Dump()

	return sNode
}
