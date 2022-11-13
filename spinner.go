package gui

import "log"

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := n.New(name)
	newNode.Widget.Name = name
	newNode.Widget.X = x
	newNode.Widget.Y = y

	newNode.Widget.Custom = func() {
		log.Println("even newer clicker() name in NewSpinner", newNode.Widget)
	}

	for _, aplug := range allPlugins {
		log.Println("gui.NewSpinner() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewSpinner == nil) {
			log.Println("\tgui.NewSpinner() aplug.NewSpinner = nil", aplug.name)
			continue
		}
		aplug.NewSpinner(&n.Widget, &newNode.Widget)
	}

	return newNode
}
