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
	t := len(w.text)

	w.gocuiSize.width = t + me.buttonPadding
	w.gocuiSize.height = me.defaultHeight

	w.gocuiSize.startW = me.rootNode.startW
	w.gocuiSize.startH = me.rootNode.startH

	w.startW = me.rawW
	w.startH = me.rawH

	var f func (widget *cuiWidget)

	// find buttons that are below where the mouse button click
	f = func(widget *cuiWidget) {
		if (widget == w) {
			return
		}
		if ((widget.widgetType == toolkit.Window) || (widget.widgetType == toolkit.Tab)) {
			w.gocuiSize.startW += widget.gocuiSize.width + me.padW
		}

		for _, child := range widget.children {
			f(child)
		}
	}
	f(me.rootNode)

	w.setWH()
	w.showWidgetPlacement(logNow, "setTabWH:")
}

func (w *cuiWidget) redoTabs(draw bool) {
	if ((w.widgetType == toolkit.Window) || (w.widgetType == toolkit.Tab)) {
		w.deleteView()
		w.drawView()
	}

	for _, child := range w.children {
		child.redoTabs(draw)
	}
}
