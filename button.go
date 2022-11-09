package gui

import "log"

func (n *Node) NewButton(name string, custom func()) *Node {
	if (n.toolkit == nil) {
		log.Println("gui.Node.NewButton() filed node.toolkit == nil")
		panic("gui.Node.NewButton() filed node.toolkit == nil")
		return n
	}
	newNode := n.New(name)
	newNode.toolkit = n.toolkit.NewButton(name)

	log.Println("gui.Node.NewButton()", name)
	if (PlugGocliOk) {
		log.Println("wit/gui gocui is loaded", PlugGocliOk)
		greeter.AddButton(name)
		log.Println("GOT HERE PlugGocliOk TRUE")
	} else {
		log.Println("GOT HERE PlugGocliOk FALSE")
	}

	// TODO: this is still confusing and probably wrong. This needs to communicate through a channel
	newNode.toolkit.Custom = func() {
		if (Config.Options.Debug) {
			log.Println("gui.Newutton() Button Clicked. Running custom() from outside toolkit START")
		}
		if (custom != nil) {
			custom()
		} else {
			log.Println("wit/gui No callback function is defined for button name =", name)
		}
		if (Config.Options.Debug) {
			log.Println("gui.NewButton() Button Clicked. Running custom() from outside toolkit END")
		}
	}
	newNode.custom = custom

	return newNode
}
