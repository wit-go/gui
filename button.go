package gui

import "git.wit.org/wit/gui/toolkit"

func (n *Node) NewButton(name string, custom func()) *Node {
	newNode := n.New(name, toolkit.Button, custom)

	var a toolkit.Action
	a.Title = name
	a.Type = toolkit.Add
	a.Callback = callback
	newaction(&a, newNode, n)

	return newNode
}

func callback(i int) {
	log(debugError, "button callback() i =", i)
}
