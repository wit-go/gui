package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// This recreates the whole GUI for a plugin

// func (n *Node) ListChildren(dump bool, dropdown *Node, mapNodes map[string]*Node) {
func (n *Node) redraw(p *aplug) {
	if (n == nil) {
		return
	}

	n.redo(p)
	for _, child := range n.children {
		child.redraw(p)
	}
	return
}

func (n *Node) redo(plug *aplug) {
	log(logNow, "redo()", plug.name, n.id, n.WidgetType, n.Name)

	var a *toolkit.Action
	a = new(toolkit.Action)
	a.Name = n.Name
	a.Text = n.Text

	a.ActionType = toolkit.Add
	a.WidgetType = n.WidgetType
	a.WidgetId = n.id

	// used for new Windows
	a.Width = n.Width
	a.Height = n.Height

	// used for anything that needs a range
	a.X = n.X
	a.Y = n.Y

	// implement here for grids and tables ?
//	a.NextX = n.NextX
//	a.NextY = n.NextY

	// used for values
	a.I = n.I
	a.S = n.S
	a.B = n.B

	if (n.parent == nil) {
		a.ParentId = 0
	} else {
		a.ParentId = n.parent.id
	}

	plug.pluginChan <- *a
	sleep(.5)
}
