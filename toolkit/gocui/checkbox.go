package main

import (
//	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

func (w *cuiWidget) setCheckbox(b bool) {
	if (w.widgetType != toolkit.Checkbox) {
		return
	}
	if (b) {
		w.b = b
		w.text = "X " + w.name
	} else {
		w.b = b
		w.text = "  " + w.name
	}
	t := len(w.text) + 1
	w.realWidth = t
	w.gocuiSize.width = t

	w.setWH()
	w.deleteView()
	w.drawView()
}
