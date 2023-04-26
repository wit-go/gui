package main

// implements widgets 'Window' and 'Tab'

import (
	"git.wit.org/wit/gui/toolkit"
//	"github.com/awesome-gocui/gocui"
)

func (w *cuiWidget) hideWidgets() {
	w.isCurrent = false
	switch w.widgetType {
	case toolkit.Root:
	case toolkit.Flag:
	case toolkit.Window:
	// case toolkit.Tab:
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
		w.setFake()
		w.showWidgetPlacement(logNow, "showFake:")
		w.showView()
	}
	for _, child := range w.children {
		child.showFake()
	}
}

func (w *cuiWidget) showWidgets() {
	if (w.isFake) {
		// don't display by default
	} else {
		if w.IsCurrent() {
			w.showWidgetPlacement(logNow, "current:")
			w.showView()
		} else {
			w.showWidgetPlacement(logNow, "not:")
			// w.drawView()
		}
	}
	for _, child := range w.children {
		child.showWidgets()
	}
}

func (w *cuiWidget) setWindowWH() {
	w.gocuiSize.w0 = me.rootNode.nextW
	w.gocuiSize.h0 = me.WindowH

	t := len(w.text)
	w.gocuiSize.w1 = w.gocuiSize.w0 + t + me.PadW
	w.gocuiSize.h1 = w.gocuiSize.h0 + me.DefaultHeight + me.PadH

	w.realWidth = w.gocuiSize.Width()
	w.realHeight = w.gocuiSize.Height()

	// move the rootNode width over for the next window
	me.rootNode.nextW += w.realWidth + me.WindowPadW

	w.nextW = 4
	w.nextH = 2

	w.showWidgetPlacement(logNow, "setWindowWH:")
}

func (w *cuiWidget) setTabWH() {
	// set the start and size of the tab gocui button

	w.gocuiSize.w0 = w.parent.nextW
	w.gocuiSize.h0 = me.TabH

	t := len(w.text)
	w.gocuiSize.w1 = w.gocuiSize.w0 + t + me.PadW
	w.gocuiSize.h1 = w.gocuiSize.h0 + me.DefaultHeight + me.PadH

	w.realWidth = w.gocuiSize.Width()
	w.realHeight = w.gocuiSize.Height()

	w.realWidth += me.FramePadW
	w.realHeight += me.FramePadH

	w.parent.nextW += w.realWidth + me.TabPadW

	w.showWidgetPlacement(logNow, "setTabWH:")
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
			w.hasTabs = true
		} else {
			w.frame = true
			w.hasTabs = false
		}
		w.setWindowWH()
		w.deleteView()
		w.showView()
	}
	if (w.widgetType == toolkit.Tab) {
		w.setTabWH()
		w.deleteView()
		// show all the tabs for the current window
		if w.parent.isCurrent {
			w.showView()
		}
	}

	for _, child := range w.children {
		child.redoTabs(draw)
	}
}
