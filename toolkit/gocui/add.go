package main

import (
//	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

var fakeStartWidth int = 40
var fakeStartHeight int = 3
func (w *cuiWidget) setFake() {
	w.isFake = true
	t := len(w.name)
	// setup fake labels for non-visable things off screen
	w.realWidth = t + 2
	w.realHeight = me.defaultHeight

	w.gocuiSize.width = t + 2
	w.gocuiSize.height = me.defaultHeight
	w.gocuiSize.w0 = fakeStartWidth
	w.gocuiSize.h0 = fakeStartHeight

	fakeStartHeight += 3
	if (fakeStartHeight > 24) {
		fakeStartHeight = 3
		fakeStartWidth += 20
	}
	w.setWH()
	if (logInfo) {
		w.drawView()
	}
}

// set the widget start width & height
func (w *cuiWidget) addWidget() {
	log(logInfo, "setStartWH() w.id =", w.id, "w.name", w.name)
	switch w.widgetType {
	case toolkit.Root:
		log(logInfo, "setStartWH() rootNode w.id =", w.id, "w.name", w.name)
		w.setFake()
		return
	case toolkit.Flag:
		w.setFake()
		return
	case toolkit.Window:
		w.setTabWH()
		w.drawView()
		// w.frame = false
		return
	case toolkit.Tab:
		w.setTabWH()
		w.drawView()
		return
	case toolkit.Box:
		w.isFake = true
		w.setFake()
		w.startW = w.parent.startW
		w.startH = w.parent.startH
		return
	case toolkit.Grid:
		w.isFake = true
		w.setFake()
		w.startW = w.parent.startW
		w.startH = w.parent.startH
		return
	case toolkit.Group:
		w.startW = w.parent.startW + 4
		w.startH = w.parent.startH + 3

		t := len(w.text)
		w.gocuiSize.width = t + me.buttonPadding
		w.gocuiSize.height = me.defaultHeight
		return
	default:
		w.startW = w.parent.startW
		w.startH = w.parent.startH
	}
	w.showWidgetPlacement(logInfo, "addWidget()")
}
