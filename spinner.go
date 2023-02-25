package gui

func (n *Node) NewSpinner(name string, x int, y int) *Node {
	newNode := n.New(name, "Spinner")
	newNode.Widget.X = x
	newNode.Widget.Y = y

	newNode.Widget.Custom = func() {
		log(debugGui, "even newer clicker() name in NewSpinner", newNode.Widget)
	}

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewSpinner() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewSpinner == nil) {
			log(debugGui, "\tgui.NewSpinner() aplug.NewSpinner = nil", aplug.name)
			continue
		}
		aplug.NewSpinner(&n.Widget, &newNode.Widget)
	}

	return newNode
}
