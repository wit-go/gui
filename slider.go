package gui

func (n *Node) NewSlider(name string, x int, y int) *Node {
	newNode := n.New(name, "Slider")
	newNode.Widget.X = x
	newNode.Widget.Y = y

	newNode.Widget.Custom = func() {
		log(debugGui, "even newer clicker() name in NewSlider", newNode.Widget)
	}

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewSlider() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewSlider == nil) {
			log(debugGui, "\tgui.NewSlider() aplug.NewSlider = nil", aplug.name)
			continue
		}
		aplug.NewSlider(&n.Widget, &newNode.Widget)
	}

	return newNode
}
