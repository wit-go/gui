package main

import (
	"strconv"
	"git.wit.org/wit/gui/toolkit"
)

func initWidget(n *node) *guiWidget {
	var w *guiWidget
	w = new(guiWidget)
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

func addDropdown() *node {
	n := new(node)
	n.WidgetType = toolkit.Flag
	n.WidgetId = -2
	n.ParentId = 0

	// copy the data from the action message
	n.Name = "DropBox"
	n.Text = "DropBox text"

	// store the internal toolkit information
	n.tk = new(guiWidget)
	n.tk.frame = true

	// set the name used by gocui to the id
	n.tk.cuiName = "-1 dropbox"

	n.tk.color = &colorFlag

	// add this new widget on the binary tree
	n.parent = me.rootNode
	if n.parent != nil {
		n.parent.children = append(n.parent.children, n)
	}
	return n
}
