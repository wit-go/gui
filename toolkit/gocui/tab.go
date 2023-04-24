package main

// implements widgets 'Window' and 'Tab'

import (
	"git.wit.org/wit/gui/toolkit"
//	"github.com/awesome-gocui/gocui"
)

func (w *cuiWidget) hideWidgets() {
	switch w.widgetType {
	case toolkit.Root:
	case toolkit.Flag:
	case toolkit.Window:
	case toolkit.Tab:
	case toolkit.Box:
	case toolkit.Grid:
	default:
		w.deleteView()
	}
	for _, child := range w.children {
		child.hideWidgets()
	}
}

func (w *cuiWidget) hideFake() {
	if (w.isFake) {
		w.deleteView()
	}
	for _, child := range w.children {
		child.hideFake()
	}
}

func (w *cuiWidget) showFake() {
	if (w.isFake) {
		w.drawView()
	}
	for _, child := range w.children {
		child.showFake()
	}
}

func (w *cuiWidget) showWidgets() {
	if (w.isFake) {
		// don't display by default
	} else {
		w.drawView()
	}
	for _, child := range w.children {
		child.showWidgets()
	}
}

func (w *cuiWidget) setTabWH() {
	// set the start and size of the tab gocui button
	t := len(w.text)
	w.gocuiSize.width = t + me.buttonPadding
	w.gocuiSize.height = me.defaultHeight
	w.gocuiSize.w0 = me.rootNode.nextW
	w.gocuiSize.h0 = me.rootNode.nextH

	// move the rootNode width over for the next window or tab
	me.rootNode.nextW += w.gocuiSize.width + me.padW

	w.startW = me.rawW
	w.startH = me.rawH
	w.nextW = me.rawW
	w.nextH = me.rawH

	w.setWH()
	w.showWidgetPlacement(logNow, "setTabWH:")
}

func (w *cuiWidget) setLabel() {
	// set the start and size of the tab gocui button
	t := len(w.text)
	w.gocuiSize.width = t + me.buttonPadding
	w.gocuiSize.height = 2
	w.gocuiSize.w0 = me.rootNode.nextW
	w.gocuiSize.h0 = me.rootNode.nextH

	// move the rootNode width over for the next window or tab
	me.rootNode.nextW += w.gocuiSize.width + me.padW

	w.startW = me.rawW
	w.startH = me.rawH
	w.nextW = me.rawW
	w.nextH = me.rawH

	w.setWH()
	w.showWidgetPlacement(logNow, "setLabel:")
}

func (w *cuiWidget) redoTabs(draw bool) {
	if (w.widgetType == toolkit.Window) {
		var tabs bool = false
		// figure out if the window is just a bunch of tabs
		for _, child := range w.children {
			if (child.widgetType == toolkit.Tab) {
				tabs = true
			}
		}
		if (tabs) {
			// window is tabs. Don't show it as a standard button
			w.frame = false
			w.setLabel()
		} else {
			w.frame = true
			w.setTabWH()
		}

		w.deleteView()
		w.drawView()
	}
	if (w.widgetType == toolkit.Tab) {
		w.deleteView()
		w.drawView()
	}

	for _, child := range w.children {
		child.redoTabs(draw)
	}
}
