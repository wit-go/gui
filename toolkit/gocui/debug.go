package main

import (
	"fmt"

	"git.wit.org/wit/gui/toolkit"
//	"github.com/awesome-gocui/gocui"
)

func actionDump(b bool, a *toolkit.Action) {
	if (a == nil) {
		log(b, "action = nil")
		return
	}

	log(b, "a.Name             =", a.Name)
	log(b, "a.Text             =", a.Text)
	log(b, "a.WidgetId         =", a.WidgetId)
	log(b, "a.ParentId         =", a.ParentId)
	log(b, "a.B                =", a.B)
	log(b, "a.S                =", a.S)
}

func (w *cuiWidget) dumpTree(draw bool) {
	if (w == nil) {
		return
	}
	w.showWidgetPlacement(logNow, "Tree:")

	for _, child := range w.children {
		child.dumpTree(draw)
	}
}

func (w *cuiWidget) showWidgetPlacement(b bool, s string) {
	if (w == nil) {
		log(logError, "WTF w == nil")
		return
	}
	if (w.id == 0) {
		log(logVerbose, "showWidgetPlacement() parent == nil ok. This is the rootNode", w.id, w.cuiName)
		return
	}
	if (w.parent == nil) {
		log(logError, "showWidgetPlacement() ERROR parent == nil", w.id, w.cuiName)
	}
	log(b, "dump()", s,
		fmt.Sprintf("(wId,pId)=(%3d,%3d)", w.id, w.parent.id),
		fmt.Sprintf("real()=(%3d,%3d,%3d,%3d)", w.realSize.w0, w.realSize.h0, w.realSize.w1, w.realSize.h1),
		"next()=(", w.nextW, ",", w.nextH, ")",
		"logical()=(", w.logicalSize.w0, ",", w.logicalSize.h0, ",", w.logicalSize.w1, ",", w.logicalSize.h1, ")",
		w.widgetType, ",", w.name, "text=", w.text)

	if (w.realWidth != (w.realSize.w1 - w.realSize.w0)) {
		log(b, "dump()", s,
			"badsize()=(", w.realWidth, ",", w.realHeight, ")",
			"badreal()=(", w.realSize.w0, ",", w.realSize.h0, ",", w.realSize.w1, ",", w.realSize.h1, ")",
			w.widgetType, ",", w.name)
	}
	if (w.realHeight != (w.realSize.h1 - w.realSize.h0)) {
		log(b, "dump()", s,
			"badsize()=(", w.realWidth, ",", w.realHeight, ")",
			"badreal()=(", w.realSize.w0, ",", w.realSize.h0, ",", w.realSize.w1, ",", w.realSize.h1, ")",
			w.widgetType, ",", w.name)
	}
}
