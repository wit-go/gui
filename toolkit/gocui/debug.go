package main

import (
	"fmt"
	"git.wit.org/wit/gui/toolkit"
)

func (n *node) dumpTree(draw bool) {
	w := n.tk
	if (w == nil) {
		return
	}
	n.showWidgetPlacement(logNow, "Tree:")

	for _, child := range n.children {
		child.dumpTree(draw)
	}
}

func (n *node) showWidgetPlacement(b bool, s string) {
	if (n == nil) {
		log(logError, "WTF w == nil")
		return
	}
	w := n.tk

	var s1 string
	var pId int
	if (n.parent == nil) {
		log(logVerbose, "showWidgetPlacement() parent == nil", n.WidgetId, w.cuiName)
		pId = 0
	} else {
		pId = n.parent.WidgetId
	}
	s1 = fmt.Sprintf("(wId,pId)=(%2d,%2d) ", n.WidgetId, pId)
	s1 += fmt.Sprintf("size=(%2d,%2d)(%2d,%2d,%2d,%2d)",
		w.size.Width(), w.size.Height(),
		w.size.w0, w.size.h0, w.size.w1, w.size.h1)
	if n.Visible() {
		s1 += fmt.Sprintf("gocui=(%2d,%2d)(%2d,%2d,%2d,%2d)",
			w.gocuiSize.Width(), w.gocuiSize.Height(),
			w.gocuiSize.w0, w.gocuiSize.h0, w.gocuiSize.w1, w.gocuiSize.h1)
	}
	if (n.parent != nil) {
		if (n.parent.WidgetType == toolkit.Grid) {
			s1 += fmt.Sprintf("At(%2d,%2d) ", n.AtW, n.AtH)
		}
	}
	log(b, s1, s, n.WidgetType, ",", n.Name) // , "text=", w.text)
}

func (n *node) dumpWidget(pad string) {
	log(true, "node:", pad, n.WidgetId, "At(", n.AtW, n.AtH, ") ,", n.WidgetType, ",", n.Name)
}

func (n *node) listWidgets() {
	if (n == nil) {
		return
	}

	var pad string
	for i := 0; i < me.depth; i++ {
		pad = pad + "    "
	}
	n.dumpWidget(pad)

	for _, child := range n.children {
		me.depth += 1
		child.listWidgets()
		me.depth -= 1
	}
	return
}
