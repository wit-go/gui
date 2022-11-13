package gui

import "log"

func (n *Node) NewSlider(name string, x int, y int) *Node {
	newNode := n.New(name)
	newNode.Widget.Name = name
	newNode.Widget.X = x
	newNode.Widget.Y = y

	newNode.Widget.Custom = func() {
		log.Println("even newer clicker() name in NewSlider", newNode.Widget)
	}

	for _, aplug := range allPlugins {
		log.Println("gui.NewSlider() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewSlider == nil) {
			log.Println("\tgui.NewSlider() aplug.NewSlider = nil", aplug.name)
			continue
		}
		aplug.NewSlider(&n.Widget, &newNode.Widget)
	}

	return newNode
}
