package gui

import "log"

func (n *Node) NewButton(name string, custom func()) *Node {
	newNode := n.New(name)

	newNode.Widget.Custom = func() {
		log.Println("even newer clicker() name", newNode.Widget)
		if (custom != nil) {
			custom()
		} else {
			log.Println("wit/gui No callback function is defined for button name =", name)
		}
	}

	for _, aplug := range allPlugins {
		log.Println("gui.NewButton() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewButton == nil) {
			log.Println("\tgui.NewButton() aplug.NewButton = nil", aplug.name)
			continue
		}
		aplug.NewButton(&n.Widget, &newNode.Widget)
	}

	return newNode
}
