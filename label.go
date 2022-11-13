package gui

import "log"
// import "errors"
// import "regexp"

func (n *Node) NewLabel(text string) *Node {
	newNode := n.New(text)

	for _, aplug := range allPlugins {
		log.Println("gui.NewLabel() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewLabel == nil) {
			log.Println("\tgui.NewLabel() aplug.NewLabel = nil", aplug.name)
			continue
		}
		aplug.NewLabel(&n.Widget, &newNode.Widget)
	}

	return newNode
}
