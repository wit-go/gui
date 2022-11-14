package gui

import "log"

// TODO: which name is better. AddGroup or NewGroup ?
// first reaction is NewGroup
func (n *Node) NewGroup(name string) *Node {
	var newNode *Node

	if (GetDebug()) {
		log.Println("toolkit.NewGroup() START", name)
	}

	newNode = n.New(name)

	log.Println("gui.Node.NewGroup()", name)
	for _, aplug := range allPlugins {
		log.Println("gui.Node.NewGroup() toolkit plugin =", aplug.name)
		if (aplug.NewGroup == nil) {
			continue
		}
		aplug.NewGroup(&n.Widget, &newNode.Widget)
	}

	return newNode
}
