package main

import (
	"fmt"
)

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
	var s1 string
	var pId int
	if (w == nil) {
		log(logError, "WTF w == nil")
		return
	}
	if (w.parent == nil) {
		log(logVerbose, "showWidgetPlacement() parent == nil", w.id, w.cuiName)
		pId = 0
	} else {
		pId = w.parent.id
	}
	s1 = fmt.Sprintf("(wId,pId)=(%2d,%2d) ", w.id, pId)
	s1 += fmt.Sprintf("s/n (%2d,%2d) (%2d,%2d) ", w.startW, w.startH, w.nextW, w.nextH)
	s1 += fmt.Sprintf("size (%2d,%2d) ", w.realWidth, w.realHeight)
	s1 += fmt.Sprintf("gocui=(%2d,%2d)(%2d,%2d,%2d,%2d)",
		w.gocuiSize.Width(), w.gocuiSize.Height(),
		w.gocuiSize.w0, w.gocuiSize.h0, w.gocuiSize.w1, w.gocuiSize.h1)
	log(b, s1, s, w.widgetType, ",", w.name) // , "text=", w.text)
}
