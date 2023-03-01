package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewTextbox(name string) *Node {
	newNode := n.New(name, toolkit.Textbox, func() {
		log(debugGui, "wit/gui clicker()NewTextBox BUT IS EMPTY. FIXME name =", name)
	})

	send(n, newNode)
	return newNode
}
