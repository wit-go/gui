package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/gui/toolkit"
)

// This routine creates a blank window with a Title and size (W x H)

func (parent *Node) NewWindow(title string) *Node {
	var newNode *Node

	// Windows are created off of the master node of the Binary Tree
	newNode = parent.newNode(title, toolkit.Window)
	newNode.Custom = StandardExit

	log.Info("NewWindow()", title)

	a := newAction(newNode, toolkit.Add)
	sendAction(a)
	return newNode
}
