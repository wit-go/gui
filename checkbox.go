package gui

import "log"

func (n *Node) Checked() bool {
	n.Dump()
	return n.checked
}

/*
This was the old code
	newt.Custom = func () {
		println("AM IN CALLBACK. SETTING NODE.checked START")
		if newt.Checked() {
			println("is checked")
			c.checked = true
		} else {
			println("is not checked")
			c.checked = false
		}
		commonCallback(c)
		println("AM IN CALLBACK. SETTING NODE.checked END")
	}
*/


func (n *Node) NewCheckbox(name string) *Node {
	newNode := n.New(name)
	newNode.custom = n.custom

	newNode.Widget.Custom = func() {
		log.Println("even newer clicker() name", newNode.Widget)
		if (n.custom != nil) {
			n.custom()
		} else {
			log.Println("wit/gui No callback function is defined for button name =", name)
		}
	}

	for _, aplug := range allPlugins {
		log.Println("gui.NewCheckbox() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewCheckbox == nil) {
			log.Println("\tgui.NewCheckbox() aplug.NewCheckbox = nil", aplug.name)
			continue
		}
		aplug.NewCheckbox(&n.Widget, &newNode.Widget)
	}

	return newNode
}
