package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) verify() {
	if (n.toolkit == nil) {
		log.Println("gui/wit node.Verify(): toolkit == nil", n.Name)
		panic("gui/wit node.Verify(): toolkit == nil")
	}
}

func (n *Node) Checked() bool {
	n.Dump()
	return n.checked
}

func (n *Node) NewCheckbox(name string) *Node {
	var newt *toolkit.Toolkit
	var c *Node

	log.Println("toolkit.NewCheckbox() START", name)

	n.verify()

	// make a *Node with a *toolkit.Group
	c = n.New(name + " part1")
	newt = n.toolkit.NewCheckbox(name)
	newt.Name = name
	c.toolkit = newt
	c.custom = n.custom
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
	c.Dump()

	return c
}
