package gui

// TODO: which name is better. AddGroup or NewGroup ?
// first reaction is NewGroup
func (n *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = n.New(name, "Group")

	log(debugGui, "toolkit.NewGroup() START", name)


	log(debugGui, "gui.Node.NewGroup()", name)
	for _, aplug := range allPlugins {
		log(debugGui, "gui.Node.NewGroup() toolkit plugin =", aplug.name)
		if (aplug.NewGroup == nil) {
			continue
		}
		aplug.NewGroup(&n.Widget, &newNode.Widget)
	}

	return newNode
}
