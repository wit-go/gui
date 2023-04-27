package main

import (
	"git.wit.org/wit/gui/toolkit"
)

// searches the binary tree for a WidgetId
func (n *node) findWidgetId(id int) *node {
	if (n == nil) {
		return nil
	}

	if n.WidgetId == id {
		return n
	}

	for _, child := range n.children {
		newN := child.findWidgetId(id)
		if (newN != nil) {
			return newN
		}
	}
	return nil
}

func addWidget(a *toolkit.Action, tk *andlabsT) *node {
	n := new(node)
	n.WidgetType = a.WidgetType
	n.WidgetId = a.WidgetId
	n.ParentId = a.ParentId

	// copy the data from the action message
	n.Name = a.Name
	n.Text = a.Text
	n.I = a.I
	n.S = a.S
	n.B = a.B
	n.X = a.X
	n.Y = a.Y

	// store the internal toolkit information
	n.tk = tk

	if (a.WidgetType == toolkit.Root) {
		log(logInfo, "addWidget() Root")
		return n
	}

	if (rootNode.findWidgetId(a.WidgetId) != nil) {
		log(logError, "addWidget() WidgetId already exists", a.WidgetId)
		return rootNode.findWidgetId(a.WidgetId)
	}

	// add this new widget on the binary tree
	n.parent = rootNode.findWidgetId(a.ParentId)
	if n.parent != nil {
		n.parent.children = append(n.parent.children, n)
	}
	return n
}

func (n *node) doUserEvent() {
	if (callback == nil) {
		log(debugError, "doUserEvent() callback == nil", n.WidgetId)
		return
	}
	var a toolkit.Action
	a.WidgetId = n.WidgetId
	a.Name = n.Name
	a.Text = n.Text
	a.S = n.S
	a.I = n.I
	a.B = n.B
	a.ActionType = toolkit.User
	log(logInfo, "doUserEvent() START: send a user event to the callback channel")
	callback <- a
	log(logInfo, "doUserEvent() END:   sent a user event to the callback channel")
	return
}
