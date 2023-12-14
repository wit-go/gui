package main

import (
	"git.wit.org/wit/gui/toolkit"
)

var fakeStartWidth int = me.FakeW
var fakeStartHeight int = me.TabH + me.FramePadH
// setup fake labels for non-visible things off screen
func (n *node) setFake() {
	w := n.tk
	w.isFake = true

	n.gocuiSetWH(fakeStartWidth, fakeStartHeight)

	fakeStartHeight += w.gocuiSize.Height()
	// TODO: use the actual max hight of the terminal window
	if (fakeStartHeight > 24) {
		fakeStartHeight = me.TabH
		fakeStartWidth += me.FakeW
	}
	if (logInfo) {
		n.showView()
	}
}

// set the widget start width & height
func (n *node) addWidget() {
	nw := n.tk
	log(logInfo, "setStartWH() w.id =", n.WidgetId, "n.name", n.Name)
	switch n.WidgetType {
	case toolkit.Root:
		log(logInfo, "setStartWH() rootNode w.id =", n.WidgetId, "w.name", n.Name)
		nw.color = &colorRoot
		n.setFake()
		return
	case toolkit.Flag:
		nw.color = &colorFlag
		n.setFake()
		return
	case toolkit.Window:
		nw.frame = false
		nw.color = &colorWindow
		// redoWindows(0,0)
		return
	case toolkit.Tab:
		nw.color = &colorTab
		// redoWindows(0,0)
		return
	case toolkit.Button:
		nw.color = &colorButton
	case toolkit.Box:
		nw.color = &colorBox
		nw.isFake = true
		n.setFake()
		return
	case toolkit.Grid:
		nw.color = &colorGrid
		nw.isFake = true
		n.setFake()
		return
	case toolkit.Group:
		nw.color = &colorGroup
		nw.frame = false
		return
	case toolkit.Label:
		nw.color = &colorLabel
		nw.frame = false
		return
	default:
		/*
		if n.IsCurrent() {
			n.updateCurrent()
		}
		*/
	}
	n.showWidgetPlacement(logInfo, "addWidget()")
}
