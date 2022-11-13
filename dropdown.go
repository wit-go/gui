package gui

import "log"

func (n *Node) AddDropdownName(name string) {
	for _, aplug := range allPlugins {
		log.Println("gui.AddDropdownName() aplug =", aplug.name, "name =", name)
		if (aplug.AddDropdownName == nil) {
			log.Println("\tgui.AddDropdownName() aplug.NewDropdown = nil", aplug.name)
			continue
		}
		aplug.AddDropdownName(&n.Widget, name)
	}
}

func (n *Node) SetDropdown(i int) {
}

func (n *Node) NewDropdown(text string) *Node {
	newNode := n.New(text)

	for _, aplug := range allPlugins {
		log.Println("gui.NewDropdown() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewDropdown == nil) {
			log.Println("\tgui.NewDropdown() aplug.NewDropdown = nil", aplug.name)
			continue
		}
		aplug.NewDropdown(&n.Widget, &newNode.Widget)
	}
	return newNode
}
