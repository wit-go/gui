package main

import (
	"strconv"
	"git.wit.org/wit/gui/toolkit"
)

func makeWidget(n *node) *cuiWidget {
	var w *cuiWidget
	w = new(cuiWidget)
	// Set(w, "default")

	w.frame = true

	// set the name used by gocui to the id
	w.cuiName = strconv.Itoa(n.WidgetId)

	if n.WidgetType == toolkit.Root {
		log(logInfo, "setupWidget() FOUND ROOT w.id =", n.WidgetId)
		n.WidgetId = 0
		me.rootNode = n
		return w
	}

	if (n.WidgetType == toolkit.Box) {
		if (n.B) {
			n.horizontal = true
		} else {
			n.horizontal = false
		}
	}

	if (n.WidgetType == toolkit.Grid) {
		w.widths = make(map[int]int) // how tall each row in the grid is
		w.heights = make(map[int]int) // how wide each column in the grid is
	}

	return w
}

func setupCtrlDownWidget() {
	a := new(toolkit.Action)
	a.Name = "ctrlDown"
	a.WidgetType = toolkit.Dialog
	a.WidgetId = -1
	a.ParentId = 0
	n := addNode(a)

	me.ctrlDown = n
}

func (n *node) deleteView() {
	w := n.tk
	if (w.v != nil) {
		w.v.Visible = false
		return
	}
	// make sure the view isn't really there
	me.baseGui.DeleteView(w.cuiName)
	w.v = nil
}

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

// searches the binary tree for a WidgetId
func (n *node) findWidgetName(name string) *node {
	if (n == nil) {
		return nil
	}

	if n.tk.cuiName == name {
		return n
	}

	for _, child := range n.children {
		newN := child.findWidgetName(name)
		if (newN != nil) {
			return newN
		}
	}
	return nil
}

func addNode(a *toolkit.Action) *node {
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
	n.tk = makeWidget(n)

	if (a.WidgetType == toolkit.Root) {
		log(logInfo, "addNode() Root")
		return n
	}

	if (me.rootNode.findWidgetId(a.WidgetId) != nil) {
		log(logError, "addNode() WidgetId already exists", a.WidgetId)
		return me.rootNode.findWidgetId(a.WidgetId)
	}

	// add this new widget on the binary tree
	n.parent = me.rootNode.findWidgetId(a.ParentId)
	if n.parent != nil {
		n.parent.children = append(n.parent.children, n)
		//w := n.tk
		//w.parent = n.parent.tk
		//w.parent.children = append(w.parent.children, w)
	}
	return n
}

func (n *node) IsCurrent() bool {
	w := n.tk
	if (n.WidgetType == toolkit.Tab) {
		return w.isCurrent
	}
	if (n.WidgetType == toolkit.Window) {
		return w.isCurrent
	}
	if (n.WidgetType == toolkit.Root) {
		return false
	}
	return n.parent.IsCurrent()
}

func (n *node) Visible() bool {
	if (n == nil) {
		return false
	}
	if (n.tk == nil) {
		return false
	}
	if (n.tk.v == nil) {
		return false
	}
	return n.tk.v.Visible
}

func (n *node) SetVisible(b bool) {
	if (n == nil) {
		return
	}
	if (n.tk == nil) {
		return
	}
	if (n.tk.v == nil) {
		return
	}
	n.tk.v.Visible = b
}
