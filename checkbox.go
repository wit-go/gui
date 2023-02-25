package gui

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
	newNode := n.New(name, "Checkbox")
	newNode.custom = n.custom

	newNode.Widget.Custom = func() {
		log(debugChange, "wit go gui checkbox", newNode.Widget)
		if (n.custom != nil) {
			log(debugChange, "trying parent.custom() callback() name =", name)
			n.custom()
		} else {
			log(debugChange, "wit/gui No parent.custom() function is defined for button name =", name)
		}
		if (newNode.custom != nil) {
			log(debugChange, "trying newNode.custom() callback name =", name)
			newNode.custom()
		} else {
			log(debugChange, "wit/gui No newNode.custom() function is defined for button name =", name)
		}
	}

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewCheckbox() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewCheckbox == nil) {
			log(debugGui, "\tgui.NewCheckbox() aplug.NewCheckbox = nil", aplug.name)
			continue
		}
		aplug.NewCheckbox(&n.Widget, &newNode.Widget)
	}

	return newNode
}
