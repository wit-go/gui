package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// TODO: make a "Group" a "Grid" ?
// probably since right now group is just a 
// pre-canned andlabs/ui gtk,macos,windows thing
func (n *Node) NewGroup(name string) *Node {
	var newNode *Node
	newNode = n.New(name, toolkit.Group, nil)

	var a toolkit.Action
	a.Type = toolkit.Add
	newaction(&a, newNode, n)

	newBox := newNode.NewBox("group vBox", false)
	return newBox
}
