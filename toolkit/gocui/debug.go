package main

import (
	"fmt"

	"git.wit.org/wit/gui/toolkit"
//	"github.com/awesome-gocui/gocui"
)

var debugAction	bool = false

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
	s1 += fmt.Sprintf("start()=(%2d,%2d) ", w.startW, w.startH)
	s1 += fmt.Sprintf("size()=(%2d,%2d) ", w.realWidth, w.realHeight)
	s1 += fmt.Sprintf("gocui()=(%2d,%2d,%2d,%2d,%2d,%2d) ",
		w.gocuiSize.width, w.gocuiSize.height,
		w.gocuiSize.w0, w.gocuiSize.h0, w.gocuiSize.w1, w.gocuiSize.h1)
	switch w.widgetType {
	case toolkit.Grid:
		s1 += fmt.Sprintf("next()=(%2d,%2d)", w.nextW, w.nextH)
	default:
	}
	log(b, s1, s, w.widgetType, ",", w.name) // , "text=", w.text)
}
