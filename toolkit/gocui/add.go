package main

import (
//	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// TODO: make these defaults in config struct definition
var fakeStartWidth int = me.DevelOffsetW
var fakeStartHeight int = me.TabH + me.FramePadH
func (w *cuiWidget) setFake() {
	w.isFake = true
	t := len(w.name)
	// setup fake labels for non-visable things off screen

	w.gocuiSize.w0 = fakeStartWidth
	w.gocuiSize.h0 = fakeStartHeight
	w.gocuiSize.w1 = w.gocuiSize.w0 + t + me.PadW
	w.gocuiSize.h1 = w.gocuiSize.h0 + me.DefaultHeight + me.PadH

	w.realWidth = w.gocuiSize.Width() + me.FramePadW
	w.realHeight = w.gocuiSize.Height() + me.FramePadH

	fakeStartHeight += w.realHeight
	// TODO: use the actual max hight of the terminal window
	if (fakeStartHeight > 24) {
		fakeStartHeight = me.TabH + me.FramePadH
		fakeStartWidth += me.DevelOffsetW
	}
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
		w.startH = w.parent.startH + me.DefaultHeight + me.FramePadH

		t := len(w.text)
		w.gocuiSize.w1 = w.gocuiSize.w0 + t + me.FramePadW
		w.gocuiSize.h1 = w.gocuiSize.h0 + me.DefaultHeight + me.FramePadH
		return
	default:
		w.startW = w.parent.startW
		w.startH = w.parent.startH
	}
	w.showWidgetPlacement(logInfo, "addWidget()")
}
