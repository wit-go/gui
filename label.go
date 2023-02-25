package gui

// import "errors"
// import "regexp"

func (n *Node) NewLabel(text string) *Node {
	newNode := n.New(text, "Label")

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewLabel() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewLabel == nil) {
			log(debugGui, "\tgui.NewLabel() aplug.NewLabel = nil", aplug.name)
			continue
		}
		aplug.NewLabel(&n.Widget, &newNode.Widget)
	}

	return newNode
}
