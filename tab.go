package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// This function should make a new node with the parent and
// the 'tab' as a child

func (n *Node) NewTab(text string) *Node {
	// check to make sure n is actually a window

	if (n.WidgetType != toolkit.Window) {
		// figure out what the actual window is
		log(logError, "NewTab() is being requested on something that isn't a Window. node =", n)
		log(logError, "NewTab() parent", n.parent)
		return n.parent.NewTab(text)
		/*
		if (n.parent.WidgetType == toolkit.Window) {
		} else {
			if (n.parent.WidgetType == toolkit.Window) {
				return n.parent.NewTab(text)
			// TODO: find a window. any window. never give up. never die.
			panic("NewTab did not get passed a window")
		}
		*/
	}
	newNode := n.newNode(text, toolkit.Tab, nil)

	var a toolkit.Action
	a.ActionType = toolkit.Add
	a.Name = text
	a.Text = text
	newaction(&a, newNode, n)

	newBox := newNode.NewBox(text, true)
	return newBox
}
