package gui

func (n *Node) AddDropdownName(name string) {
	for _, aplug := range allPlugins {
		log(debugGui, "gui.AddDropdownName() aplug =", aplug.name, "name =", name)
		if (aplug.AddDropdownName == nil) {
			log(debugGui, "\tgui.AddDropdownName() aplug.NewDropdown = nil", aplug.name)
			continue
		}
		aplug.AddDropdownName(&n.Widget, name)
	}

	if (n.Widget.Custom == nil) {
		n.SetDropdownChange( func() {
			log(debugChange, "gui.Dropdown change() REAL Custom() name =", name)
			log(debugChange, "gui.Dropdown change() REAL n.Widget.S =", n.Widget.S)
		})
	}
	// TODO, this makes functions over and over for each dropdown menu
	/*
	n.Widget.Custom = func() {
		log(debugChange, "gui.Dropdown change() START Custom() name =", name)
		log(debugChange, "gui.Dropdown change() START n.Widget.S =", n.Widget.S)
	}
	*/
}

func (n *Node) SetDropdown(s any) {
	log(debugGui, "gui.SetDropdown() TODO: make this work. s =", s)
}

func (n *Node) SetDropdownChange(f func()) {
	n.Widget.Custom = f
}

func (n *Node) NewDropdown(name string) *Node {
	newNode := n.New(name, "Dropdown")

	for _, aplug := range allPlugins {
		log(debugGui, "gui.NewDropdown() aplug =", aplug.name, "name =", newNode.Widget.Name)
		if (aplug.NewDropdown == nil) {
			log(debugGui, "\tgui.NewDropdown() aplug.NewDropdown = nil", aplug.name)
			continue
		}
		aplug.NewDropdown(&n.Widget, &newNode.Widget)
	}

	// TODO, this doesn't work for some reason (over-written by plugin?)
	newNode.Widget.Custom = func() {
		log(true, "gui.NewDropdown() START Custom()")
	}
	return newNode
}
