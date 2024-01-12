package gui

import (
	"go.wit.com/log"
	"go.wit.com/gui/widget"
)

// This routine creates a blank window with a Title and size (W x H)

func (parent *Node) NewWindow(title string) *Node {
	var newNode *Node

	// Windows are created off of the master node of the Binary Tree
	newNode = parent.newNode(title, widget.Window)
	newNode.Custom = StandardExit

	log.Info("NewWindow()", title)
	newNode.progname = title
	newNode.value = title

	if ! newNode.hidden {
		a := newAction(newNode, widget.Add)
		sendAction(a)
	}
	return newNode
}

// allow window create without actually sending it to the toolkit
func (parent *Node) RawWindow(title string) *Node {
	var newNode *Node

	// Windows are created off of the master node of the Binary Tree
	newNode = parent.newNode(title, widget.Window)
	newNode.Custom = StandardExit
	newNode.hidden = true

	log.Info("RawWindow()", title)
	return newNode
}

// TODO: should do this recursively
func (n *Node) UnDraw() *Node {
	if ! n.hidden {
		n.Hide()
	}
	n.hidden = true
	return n
}

// TODO: should do this recursively
func (n *Node) Draw() *Node {
	n.hidden = false

	a := newAction(n, widget.Add)
	sendAction(a)
	return n
}

// if the toolkit supports a gui with pixels, it might honor this. no promises
// consider this a 'recommendation' or developer 'preference' to the toolkit
func (n *Node) PixelSize(w, h int) *Node {
	n.width = w
	n.height = w
	return n
}
