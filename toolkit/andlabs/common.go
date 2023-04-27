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

	// deprecate this when this toolkit uses the binary tree instead
	if (andlabs[a.WidgetId] == nil) {
		andlabs[a.WidgetId] = tk
	}

	return n
}

func (t *andlabsT) doUserEvent() {
	if (callback == nil) {
		log(debugError, "doUserEvent() callback == nil", t.wId)
		return
	}
	var a toolkit.Action
	a.WidgetId = t.wId
	a.Name = t.Name
	a.S = t.s
	a.I = t.i
	a.B = t.b
	a.ActionType = toolkit.User
	log(logInfo, "doUserEvent() START: send a user event to the callback channel")
	callback <- a
	log(logInfo, "doUserEvent() END:   sent a user event to the callback channel")
	return
}
