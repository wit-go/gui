package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewTextbox(name string) *Node {
	var newt *toolkit.Toolkit
	var c *Node

	log.Println("toolkit.NewTextbox() START", name)

	n.verify()

	// make a new Node and a new toolbox struct
	c = n.New(name)
	newt = n.toolkit.NewTextbox(name)

	c.toolkit = newt
	c.custom = n.custom

	newt.Name = name
	// newt.Custom = func () {
	newt.OnChanged = func (*toolkit.Toolkit) {
		println("AM IN CALLBACK. SETTING NODE.checked START")
		c.text = c.toolkit.GetText()
		c.Dump()
		c.toolkit.Dump()
		c.OnChanged(n)
		println("n.toolkit.GetText() =", c.text)
		println("AM IN CALLBACK. SETTING NODE.checked END")
	}

	return c
}
