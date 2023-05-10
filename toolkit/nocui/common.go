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

func addWidget(a *toolkit.Action) *node {
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

	n.W = a.W
	n.H = a.H
	n.AtW = a.AtW
	n.AtH = a.AtH

	// store the internal toolkit information
	n.tk = new(nocuiT)

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
