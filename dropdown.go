package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func commonCallback(n *Node) {
	// TODO: make all of this common code to all the widgets
	if (n.OnChanged == nil) {
		log.Println("Not Running n.OnChanged(n) == nil")
	} else {
		log.Println("Running n.OnChanged(n)")
		n.OnChanged(n)
	}

	if (n.custom == nil) {
		log.Println("Not Running n.custom(n) == nil")
	} else {
		log.Println("Running n.custom()")
		n.custom()
	}
}

func (n *Node) NewDropdown(name string) *Node {
	var newT *toolkit.Toolkit
	var sNode *Node

	log.Println("toolkit.NewDropdown() START", name)

	n.verify()

	sNode = n.New(name + " part1")
	newT = n.toolkit.NewDropdown(name)
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

func (n *Node) AddDropdown(name string) {
	n.toolkit.AddDropdown(name)
}

func (n *Node) SetDropdown(i int) {
	n.toolkit.SetDropdown(i)
}
