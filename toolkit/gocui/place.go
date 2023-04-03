package main

import (
//	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

var fakeStartWidth int = 40
var fakeStartHeight int = 3
func (w *cuiWidget) setFake() {
	if (w.isFake == false) {
		return
	}
	t := len(w.name)
	// setup fake labels for non-visable things off screen
	w.realWidth = t + 2
	w.realHeight = me.defaultHeight
	w.realSize.w0 = fakeStartWidth
	w.realSize.h0 = fakeStartHeight
	w.realSize.w1 =  w.realSize.w0 + w.realWidth
	w.realSize.h1 =  w.realSize.h0 + w.realHeight
	fakeStartHeight += 3
	if (fakeStartHeight > 24) {
		fakeStartHeight = 3
		fakeStartWidth += 20
	}
	w.showWidgetPlacement(logNow, "setFake()")
}

func findPlace(w *cuiWidget) {
	w.isFake = false
	w.visable = true
	switch w.widgetType {
	case toolkit.Root:
		w.isFake = true
		w.setFake()
	case toolkit.Flag:
		w.isFake = true
		w.setFake()
	case toolkit.Grid:
		w.isFake = true
		w.setFake()
	case toolkit.Box:
		w.isFake = true
		w.setFake()
	default:
		// w.redoBox(true)
	}
}

func (w *cuiWidget) redoBox(draw bool) {
	if (w == nil) {
		return
	}
	if (me.rootNode == nil) {
		return
	}
	p := w.parent
	if (p == nil) {
		log(logInfo, "redoBox()", w.id, "parent == nil")
		return
	}

	t := len(w.text)
	w.visable = true
	switch w.widgetType {
	case toolkit.Window:
		for _, child := range w.children {
			child.redoBox(draw)
		}
	case toolkit.Tab:
		for _, child := range w.children {
			child.redoBox(draw)
		}
	case toolkit.Grid:
		log("redoBox GRID", p.nextW, p.nextH, p.name)
		log("redoBox GRID", p.nextW, p.nextH, p.name)
		log("redoBox GRID", w.nextW, w.nextH, w.name, w.text)
		// hmm

		w.nextW = p.nextW
		w.nextH = p.nextH

		var wCount int = 0
		var hCount int = 0
		for _, child := range w.children {
			// increment for the next child
			w.nextW = p.nextW + wCount * 20
			w.nextH = p.nextH + hCount * 4
			child.redoBox(draw)

			log(logNow, "redoBox(GRID) (w,h count)", wCount, hCount, "(X,Y)", w.x, w.y, w.name)
			child.showWidgetPlacement(logNow, "grid:")

			if ((wCount + 1) < w.y) {
				wCount += 1
			} else {
				wCount = 0
				hCount += 1
			}
		}
		w.showWidgetPlacement(logNow, "grid:")
	case toolkit.Box:
		w.logicalSize.w0 = p.nextW
		w.logicalSize.h0 = p.nextH
		w.logicalSize.w1 = p.nextW
		w.logicalSize.h1 = p.nextH

		w.nextW = p.nextW
		w.nextH = p.nextH
		for _, child := range w.children {
			child.redoBox(draw)
			if (w.horizontal) {
				log("BOX IS HORIZONTAL", p.nextW, p.nextW, p.name)
				log("BOX IS HORIZONTAL", w.nextW, w.nextH, w.name)
				log("BOX IS HORIZONTAL")
				// expand based on the child width
				w.nextW = child.nextW + me.horizontalPadding
				// reset height to parent
				w.nextH = p.nextH
			} else {
				log("BOX IS VERTICAL", p.nextW, p.nextW, p.name)
				log("BOX IS VERTICAL", w.nextW, w.nextH, w.name)
				log("BOX IS VERTICAL")
				// go straight down
				w.nextW = p.nextW
				// expand based on the child height
				w.nextH = child.nextH
			}
		}
		w.showWidgetPlacement(logNow, "box:")
	case toolkit.Group:
		w.realWidth = t + me.buttonPadding
		w.realHeight = me.defaultHeight

		w.realSize.w0 = p.nextW
		w.realSize.h0 = p.nextH
		w.realSize.w1 = w.realSize.w0 + w.realWidth
		w.realSize.h1 = w.realHeight

		w.logicalSize.w0 = w.realSize.w0
		w.logicalSize.h0 = w.realSize.h0
		w.logicalSize.w1 = w.realSize.w1
		w.logicalSize.h1 = w.realSize.h1

		w.nextW = p.nextW + me.groupPadding
		w.nextH = p.nextH + me.buttonPadding
		for _, child := range w.children {
			child.redoBox(draw)
			// reset nextW to straight down
			w.nextW = p.nextW + 4
			w.nextH = child.nextH
		}
		// expand the height of the parent now that the group is done
		// p.nextW = w.nextW
		// p.nextH = w.nextH
		w.showWidgetPlacement(logNow, "group:")
	default:
		w.realWidth = t + 3
		w.realHeight = me.defaultHeight
		w.realSize.w0 = p.nextW
		w.realSize.h0 = p.nextH
		w.realSize.w1 = p.nextW + w.realWidth
		w.realSize.h1 = p.nextH + w.realHeight

		w.logicalSize.w0 = p.nextW
		w.logicalSize.h0 = p.nextH
		w.logicalSize.w1 = p.nextW + w.realWidth
		w.logicalSize.h1 = p.nextH + w.realHeight

		w.nextW = w.realSize.w1
		w.nextH = w.realSize.h1
	}
}

func (w *cuiWidget) boxedPlace() {
	t := len(w.name)
	if (w.id == 0) {
		w.realWidth = 0
		w.realHeight = 0
		return
	}
	p := w.parent
	if (p == nil) {
		log(logError, "boxedPlace() parentId widget == nil")
		return
	}

	w.realWidth = t + 3
	w.realHeight = me.defaultHeight
	w.realSize.w0 = p.nextW
	w.realSize.h0 = p.nextH
	w.realSize.w1 = p.nextW + w.realWidth
	w.realSize.h1 = p.nextH + w.realHeight

	w.logicalSize.w0 = p.nextW
	w.logicalSize.h0 = p.nextH
	w.logicalSize.w1 = p.nextW + w.realWidth
	w.logicalSize.h1 = p.nextH + w.realHeight

	w.nextW = w.realSize.w1
	w.nextH = w.realSize.h1

	w.showWidgetPlacement(logNow, "bP widget")
}

func (w *cuiWidget) updateLogicalSizes() {
	for _, child := range w.children {
		child.updateLogicalSizes()
		if (w.logicalSize.w0 > child.logicalSize.w0) {
			w.logicalSize.w0 = child.logicalSize.w0
		}
		if (w.logicalSize.w1 < child.logicalSize.w1) {
			w.logicalSize.w1 = child.logicalSize.w1
		}
		if (w.logicalSize.h0 > child.logicalSize.h0) {
			w.logicalSize.h0 = child.logicalSize.h0
		}
		if (w.logicalSize.h1 < child.logicalSize.h1) {
			w.logicalSize.h1 = child.logicalSize.h1
		}
	}
}
