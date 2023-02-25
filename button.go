package gui

func (n *Node) NewButton(name string, custom func()) *Node {
	newNode := n.New(name, "Button")

	newNode.Widget.Custom = func() {
		log(debugGui, "even newer clicker() name", newNode.Widget)
		if (custom != nil) {
			custom()
		} else {
			log(debugGui, "wit/gui No callback function is defined for button name =", name)
		}
	}

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewButton() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewButton == nil) {
			log(debugGui, "\tgui.NewButton() aplug.NewButton = nil", aplug.name)
			continue
		}
		aplug.NewButton(&n.Widget, &newNode.Widget)
	}

	return newNode
}
