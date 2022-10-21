package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSlider(name string, x int, y int) *Node {
	var newT *toolkit.Toolkit
	var sNode *Node

	log.Println("toolkit.NewSlider() START", name)

	n.verify()

	// make a *Node with a *toolkit.Group
	sNode = n.New(name + " part1")
	newT = n.toolkit.NewSlider(name, x, y)
	newT.Name = name
	sNode.custom = n.custom
	newT.Custom = func () {
		commonCallback(sNode)
	}
	sNode.toolkit = newT
	sNode.Dump()
	// panic("checking Custom()")

	return sNode
}
