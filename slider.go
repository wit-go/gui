package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) NewSlider(name string, x int, y int) *Node {
	// make new node here
	log.Println("toolkit.NewSlider", x, y)

	newNode := n.makeNode(name, 767, 676 + Config.counter)
	newNode.Name = name

	t := toolkit.NewSlider(n.uiBox, name, x, y)
	t.OnChanged = func(t *toolkit.Toolkit) {
		log.Println("toolkit.NewSlider() value =", t.Value())
		if (newNode.OnChanged != nil) {
			newNode.OnChanged(newNode)
		}
	}
	newNode.Toolkit = t

	return newNode
}
