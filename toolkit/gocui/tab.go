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
	w.drawView()
	for _, child := range w.children {
		child.showWidgets()
	}
}

func (w *cuiWidget) redoTabs(draw bool) {
	log(logNow, "redoTabs() START", w.name)
	if (w == nil) {
		return
	}
	if (w.widgetType == toolkit.Root) {
		w.logicalSize.w0 = 0
		w.logicalSize.h0 = 0
		w.logicalSize.w1 = 0
		w.logicalSize.h1 = 0

		w.nextW = 2
		w.nextH = 2
	}

	log(logNow, "redoTabs() about to check for window and tab ", w.name)
	w.text = w.name
	t := len(w.text)
	if ((w.widgetType == toolkit.Window) || (w.widgetType == toolkit.Tab)) {
		log(logNow, "redoTabs() in Window and Tab", w.name)
		w.realWidth = t + 2
		w.realHeight = me.defaultHeight

		w.realSize.w0 = me.rootNode.logicalSize.w1
		w.realSize.h0 = 0
		w.realSize.w1 = w.realSize.w0 + w.realWidth
		w.realSize.h1 = w.realHeight

		w.logicalSize.w0 = 0
		w.logicalSize.h0 = 0
		w.logicalSize.w1 = 0
		w.logicalSize.h1 = w.realHeight

		// spaces right 1 space to next tab widget
		// spaces down 1 line to the next widget
		w.nextW = 2
		w.nextH = w.realHeight + 1

		me.rootNode.logicalSize.w1 = w.realSize.w1 + 1
		me.rootNode.logicalSize.h1 = 0

		w.deleteView()
		w.v = nil
		w.drawView()
		w.showWidgetPlacement(logNow, "redoTabs()")
	}

	log(logNow, "redoTabs() about to for loop children", w.name)
	for _, child := range w.children {
		log(logNow, "redoTabs() got to child", child.name)
		child.redoTabs(draw)
	}
}
