package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// This routine creates a blank window with a Title and size (W x H)

func (n *Node) NewWindow(title string) *Node {
	var newNode *Node

	// Windows are created off of the master node of the Binary Tree
	newNode = n.newNode(title, toolkit.Window, StandardExit)

	log(logInfo, "NewWindow()", title)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.X         = n.X
	a.Y         = n.Y
	a.Name      = title
	a.Text      = title
	newaction(&a, newNode, n)

	return newNode
}
