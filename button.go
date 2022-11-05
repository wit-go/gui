package gui

import "log"

func (n *Node) NewButton(name string, custom func()) *Node {
	if (n.toolkit == nil) {
		log.Println("gui.Node.AppendButton() filed node.toolkit == nil")
		panic("gui.Node.AppendButton() filed node.toolkit == nil")
		return n
	}
	newNode := n.New(name)
	newNode.toolkit = n.toolkit.NewButton(name)

	// TODO: this is still confusing and probably wrong. This needs to communicate through a channel
	newNode.toolkit.Custom = func() {
		if (Config.Options.Debug) {
			log.Println("gui.AppendButton() Button Clicked. Running custom() from outside toolkit START")
		}
		custom()
		if (Config.Options.Debug) {
			log.Println("gui.AppendButton() Button Clicked. Running custom() from outside toolkit END")
		}
	}
	newNode.custom = custom

	return newNode
}
