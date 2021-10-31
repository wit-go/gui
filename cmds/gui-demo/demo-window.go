package main

import "log"

import "git.wit.org/wit/gui"

func demoClick (n *gui.Node) {
		log.Println("demoClick() Dumping node:")
		n.Dump()
}

func addDemoTab(n *gui.Node, title string) {
	newNode := n.AddTab(title, nil)
	if (gui.Config.Debug) {
		newNode.Dump()
	}
	newNode.ListChildren(false)

	groupNode1 := newNode.AddGroup("group 1")
	groupNode1.AddComboBox("demoCombo1", "foo", "bar", "stuff")
	groupNode1.AddComboBox("demoCombo3", "foo 3", "bar", "stuff")

	groupNode1.Dump()

	butNode1 := groupNode1.AddButton("button1", demoClick)
	butNode1.Dump()

	butNode2 := groupNode1.AddButton("button2", demoClick)
	butNode2.Dump()

	groupNode2 := newNode.AddGroup("group 2")
	groupNode2.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
}
