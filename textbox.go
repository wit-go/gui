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
		if (Config.Options.DebugChange) {
			log.Println("AM IN CALLBACK. SETTING NODE.checked START")
			c.Dump()
			c.toolkit.Dump()
		}
		c.text = c.toolkit.GetText()
		if (c.OnChanged == nil) {
			if (Config.Options.DebugChange) {
				log.Println("this is println?")
			}
		} else {
			if (Config.Options.DebugChange) {
				log.Println("this is println? running c.OnChanged() here")
			}
			c.OnChanged(n)
		}
		if (Config.Options.DebugChange) {
			log.Println("n.toolkit.GetText() =", c.text)
			log.Println("AM IN CALLBACK. SETTING NODE.checked END")
		}
	}

	return c
}
