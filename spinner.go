package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	var newT *toolkit.Toolkit
	var sNode *Node

	log.Println("toolkit.NewSpinner() START", name)

	n.verify()

	// make a *Node with a *toolkit.Group
	sNode = n.New(name + " part1")
	newT = n.toolkit.NewSpinner(name, x, y)
	newT.Name = name
	sNode.toolkit = newT
	// sNode.Dump()

	newT.Custom = func () {
		commonCallback(sNode)
	}

	return sNode
}
