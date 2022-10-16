package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSpinbox(name string, x int, y int) *Node {
	// make new node here
	log.Println("toolkit.NewSpinbox", x, y)

	newNode := n.makeNode(name, 767, 676 + Config.counter)
	newNode.Name = name

	t := toolkit.NewSpinbox(n.uiBox, name, x, y)
	t.OnChanged = func(t *toolkit.Toolkit) {
		log.Println("toolkit.NewSpinbox() value =", t.Value())
		if (newNode.OnChanged != nil) {
			newNode.OnChanged(newNode)
		}
	}
	newNode.Toolkit = t

	return newNode
}
