package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) NewButton(name string, custom func()) *Node {
	newNode := n.New(name, toolkit.Button, custom)
	send(n, newNode)
	return newNode
}
