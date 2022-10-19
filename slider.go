package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSlider(name string, x int, y int) *Node {
	var newT *toolkit.Toolkit
	var sNode *Node

	log.Println("toolkit.NewSlider() START", name)

	if (n.Toolkit == nil) {
		log.Println("toolkit.NewSlider() Toolkit == nil")
		panic("Toolkit should never be nil")
	}

	// make a *Node with a *toolkit.Group
	sNode = n.New(name + " part1")
	newT = n.Toolkit.NewSlider(name, x, y)
	newT.Name = name
	sNode.Toolkit = newT
	sNode.Dump()

	return sNode
}
