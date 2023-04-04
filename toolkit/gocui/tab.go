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
	log(logVerbose, "redoTabs() START", w.name)
	if (w == nil) {
		return
	}
	if (w.widgetType == toolkit.Root) {
		w.logicalSize.w0 = 0
		w.logicalSize.h0 = 0
		w.logicalSize.w1 = 0
		w.logicalSize.h1 = 0

		w.nextW = 1
		w.nextH = 1
	}

	log(logVerbose, "redoTabs() about to check for window and tab ", w.name)
	w.text = w.name
	t := len(w.text)
	if ((w.widgetType == toolkit.Window) || (w.widgetType == toolkit.Tab)) {
		log(logVerbose, "redoTabs() in Window and Tab", w.name)
		w.realWidth = t + me.buttonPadding
		w.realHeight = me.defaultHeight

		w.realSize.w0 = me.rootNode.logicalSize.w1
		w.realSize.w1 = w.realSize.w0 + w.realWidth
		w.realSize.h0 = 0
		w.realSize.h1 = w.realHeight

		// start logical sizes windows and in the top left corner
		w.logicalSize.w0 = 2
		w.logicalSize.w1 = 2
		w.logicalSize.h0 = w.realHeight
		w.logicalSize.h1 = w.realHeight

		// start all windows and in the top left corner
		w.nextW = w.logicalSize.w0
		w.nextH = w.logicalSize.h0

		me.rootNode.logicalSize.w1 = w.realSize.w1
		me.rootNode.logicalSize.h1 = w.realSize.h1

		w.deleteView()
		w.drawView()
		w.showWidgetPlacement(logNow, "redoTabs()")
	}

	log(logVerbose, "redoTabs() about to for loop children", w.name)
	for _, child := range w.children {
		log(logVerbose, "redoTabs() got to child", child.name)
		child.redoTabs(draw)
	}
}
