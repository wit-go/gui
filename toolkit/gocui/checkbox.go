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
	w.gocuiSize.w1 = w.gocuiSize.w0 + t

	w.realWidth = w.gocuiSize.Width() + me.PadW
	w.realHeight = w.gocuiSize.Height() + me.PadH

	if w.frame {
		w.realWidth += me.FramePadW
		w.realHeight += me.FramePadH
	}

	w.deleteView()
	w.showView()
}
