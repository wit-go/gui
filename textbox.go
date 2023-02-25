package gui

func (n *Node) NewTextbox(name string) *Node {
	newNode := n.New(name, "Textbox")

	newNode.Widget.Custom = func() {
		log(debugGui, "wit/gui clicker()NewTextBox BUT IS EMPTY. FIXME", newNode.Widget)
	}

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewTextbox() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewTextbox == nil) {
			log(debugGui, "\tgui.NewTextbox() aplug.NewTextbox = nil", aplug.name)
			continue
		}
		aplug.NewTextbox(&n.Widget, &newNode.Widget)
	}

	return newNode
}
