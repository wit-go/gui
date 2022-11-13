package gui

import "log"

// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"
// import newtoolkit	"git.wit.org/wit/gui/toolkit"

// TODO: which name is better. AddGroup or NewGroup ?
// first reaction is NewGroup
func (n *Node) NewGroup(name string) *Node {
// 	var newT *toolkit.Toolkit
	var newNode *Node

	if (GetDebug()) {
		log.Println("toolkit.NewGroup() START", name)
	}

// 	if (n.toolkit == nil) {
// 		log.Println("toolkit.NewGroup() toolkit == nil")
// 		panic("toolkit should never be nil")
// 	}

	newNode = n.New(name)

	log.Println("gui.Node.NewGroup()", name)
	for _, aplug := range allPlugins {
		log.Println("gui.Node.NewGroup() toolkit plugin =", aplug.name)
		if (aplug.NewGroup == nil) {
			continue
		}
		aplug.NewGroup(&n.Widget, &newNode.Widget)
	}

	// make a *Node with a *toolkit.Group
	// newT = n.toolkit.NewGroup(name)
	// newNode.toolkit = newT
	// newNode.Dump()

	return newNode
}

/*
func (n *Node) AddGroup(title string) *Node {
	return n.NewGroup(title + " deprecated AddGroup")
}
*/
