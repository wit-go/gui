package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// This recreates the whole GUI for a plugin

func Redraw(s string) {
	var p *aplug
	log(logNow, "attempt to feed the binary tree to", s)
	for _, aplug := range allPlugins {
		log("Loaded plugin:", aplug.name, aplug.filename)
		if (aplug.name == s) {
			log("Found plugin:", aplug.name, aplug.filename)
			p = aplug
		}
	}
	if (p == nil) {
		log("Plugin", s, "is not loaded")
		return
	}
	Config.rootNode.Redraw(p)
}

// func (n *Node) ListChildren(dump bool, dropdown *Node, mapNodes map[string]*Node) {
func (n *Node) Redraw(p *aplug) {
	if (n == nil) {
		return
	}

	n.redo(p)
	for _, child := range n.children {
		child.Redraw(p)
	}
	return
}

func (n *Node) redo(p *aplug) {
	log(logNow, "redo() on n.Widget")

	var a *toolkit.Action
	a = new(toolkit.Action)
	a.Name = n.Name
	a.Text = n.Text
	a.ActionType = toolkit.Add
	a.WidgetType = n.WidgetType
	a.WidgetId = n.id
	a.Width = n.Width
	a.Height = n.Height

	if (n.parent == nil) {
		a.ParentId = 0
	} else {
		a.ParentId = n.parent.id
	}

	p.Action(a)
}
