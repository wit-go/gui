package gui

import (
	"go.wit.com/gui/toolkit"
)

// This function should make a new node with the parent and
// the 'tab' as a child

func (n *Node) NewTab(text string) *Node {
	// check to make sure n is actually a window

	if (n.WidgetType != toolkit.Window) {
		// figure out what the actual window is
		log(logError, "NewTab() is being requested on something that isn't a Window. node =", n)
		if (n.parent == nil) {
			// TODO: find a window. any window. never give up. never die.
			log(logError, "NewTab() TODO: make a window here", n)
			panic("NewTab did not get passed a window")
		}
		log(logError, "NewTab() parent =", n.parent)
		if (n.parent.WidgetType == toolkit.Root) {
			// also broken
			log(logError, "NewTab() TODO: make or find a window here", n)
			panic("NewTab() did not get passed a window")
		}
		// go up the binary tree until we find a window widget to add a tab too
		return n.parent.NewTab(text)
	}
	newNode := n.newNode(text, toolkit.Tab)

	a := newAction(newNode, toolkit.Add)
	sendAction(a)

	// by default, create a box inside the tab
	// TODO: allow this to be configurable
	newBox := newNode.NewBox(text + " box", true)

	return newBox
}
