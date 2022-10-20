package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSlider(name string, x int, y int) *Node {
	var newT *toolkit.Toolkit
	var sNode *Node

	log.Println("toolkit.NewSlider() START", name)

	if (n.toolkit == nil) {
		log.Println("toolkit.NewSlider() toolkit == nil")
		panic("Toolkit should never be nil")
	}

	// make a *Node with a *toolkit.Group
	sNode = n.New(name + " part1")
	newT = n.toolkit.NewSlider(name, x, y)
	newT.Name = name
	sNode.custom = n.custom
	newT.Custom = func () {
		// TODO: make all of this common code to all the widgets
		if (n.custom == nil) {
			log.Println("Not Running n.custom(n) == nil")
		} else {
			log.Println("Running n.custom(n)")
			sNode.custom(sNode)
		}
		if (sNode.OnChanged == nil) {
			log.Println("Not Running n.OnChanged(n) == nil")
		} else {
			log.Println("Running n.OnChanged(n)")
			sNode.OnChanged(sNode)
		}
	}
	sNode.toolkit = newT
	sNode.Dump()
	// panic("checking Custom()")

	return sNode
}
