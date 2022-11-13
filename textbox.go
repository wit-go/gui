package gui

import "log"

func (n *Node) NewTextbox(name string) *Node {
	newNode := n.New(name)

	newNode.Widget.Custom = func() {
		log.Println("even newer clicker() name in NewTextBox", newNode.Widget)
	}

	for _, aplug := range allPlugins {
		log.Println("gui.NewTextbox() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewTextbox == nil) {
			log.Println("\tgui.NewTextbox() aplug.NewTextbox = nil", aplug.name)
			continue
		}
		aplug.NewTextbox(&n.Widget, &newNode.Widget)
	}

	return newNode
}
