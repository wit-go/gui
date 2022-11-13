package gui

import (
	"log"
)

// This function should make a new node with the parent and
// the 'tab' as a child

func (n *Node) NewTab(text string) *Node {
	newNode := n.New(text)

	for _, aplug := range allPlugins {
		log.Println("gui.NewTab() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewTab == nil) {
			log.Println("\tgui.NewTab() aplug.NewTab = nil", aplug.name)
			continue
		}
		aplug.NewTab(&n.Widget, &newNode.Widget)
	}

	return newNode
}
