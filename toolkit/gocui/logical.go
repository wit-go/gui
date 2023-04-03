package main

import (
	// "git.wit.org/wit/gui/toolkit"
)

var adjusted bool = false

// expands the logical size of the parents
func (w *cuiWidget) setParentLogical(p *cuiWidget) {
	if (w.visable) {
		// expand the parent logicalsize to include the widget realSize
		if (p.logicalSize.w0 > w.realSize.w0) {
			p.logicalSize.w0 = w.realSize.w0
			adjusted = true
		}
		if (p.logicalSize.h0 > w.realSize.h0) {
			p.logicalSize.h0 = w.realSize.h0
			adjusted = true
		}
		if (p.logicalSize.w1 < w.realSize.w1) {
			p.logicalSize.w1 = w.realSize.w1
			adjusted = true
		}
		if (p.logicalSize.h1 < w.realSize.h1) {
			p.logicalSize.h1 = w.realSize.h1
			adjusted = true
		}
	} else {
		// expand the parent logicalsize to include the widget logicalsize
		if (p.logicalSize.w0 > w.logicalSize.w0) {
			p.logicalSize.w0 = w.logicalSize.w0
			adjusted = true
		}
		if (p.logicalSize.h0 > w.logicalSize.h0) {
			p.logicalSize.h0 = w.logicalSize.h0
			adjusted = true
		}
		if (p.logicalSize.w1 < w.logicalSize.w1) {
			p.logicalSize.w1 = w.logicalSize.w1
			adjusted = true
		}
		if (p.logicalSize.h1 < w.logicalSize.h1) {
			p.logicalSize.h1 = w.logicalSize.h1
			adjusted = true
		}
	}
	if (w.visable) {
		// adjust the widget realSize to the top left corner of the logicalsize
		if (w.logicalSize.w0 > w.realSize.w0) {
			w.realSize.w0 = w.logicalSize.w0
			w.realSize.w1 = w.realSize.w0 + w.realWidth
			adjusted = true
		}
		if (w.logicalSize.h0 > w.realSize.h0) {
			w.realSize.h0 = w.logicalSize.h0
			w.realSize.h1 = w.realSize.h0 + w.realHeight
			adjusted = true
		}
	}
	w.showWidgetPlacement(logNow, "setParentLogical() widget")
	p.showWidgetPlacement(logNow, "setParentLogical() parent")
	if (w.id == 0) || (p.id == 0) {
		// stop resizing when you hit the root widget
		return
	}
	// pass the logical resizing up
	pP := w.parent
	if (pP != nil) {
		p.setParentLogical(pP)
	}
}
