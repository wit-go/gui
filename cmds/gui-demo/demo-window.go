package main

import "git.wit.org/wit/gui"

func addDemoTab(n *gui.Node, title string) {
	newNode := n.AddTab(title, nil)
	if (gui.Config.Debug) {
		newNode.Dump()
	}
	newNode.ListChildren(false)

	groupNode1 := newNode.AddGroup("group 1")
	groupNode1.AddComboBox("demoCombo1", "foo", "bar", "stuff")

	groupNode2 := newNode.AddGroup("group 2")
	groupNode2.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
}
