package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *Node) NewLabel(text string) *Node {
	newNode := n.New(text, toolkit.Label, func() {
		log(debugChange, "TextBox changed", text)
	})

	send(n, newNode)
	return newNode
}
