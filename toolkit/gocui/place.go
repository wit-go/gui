package main

import (
	"fmt"
//	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

func (w *cuiWidget) placeBox() {
	if (w.widgetType != toolkit.Box) {
		return
	}
	w.startW = w.parent.nextW
	w.startH = w.parent.nextH
	w.nextW = w.parent.nextW
	w.nextH = w.parent.nextH
	w.realWidth = 0
	w.realHeight = 0

	var maxW int
	var maxH int
	for _, child := range w.children {
		w.showWidgetPlacement(logNow, "boxS()")
		child.placeWidgets()
		if (w.horizontal) {
			log(logVerbose, "BOX IS HORIZONTAL")
			// expand based on the child width
			w.nextW += child.realWidth
			w.realWidth += child.realWidth
		} else {
			log(logVerbose, "BOX IS VERTICAL")
			// expand based on the child height
			w.nextH += child.realHeight
			w.realHeight += child.realHeight
		}
		if (maxW < child.realWidth) {
			maxW = child.realWidth
		}
		if (maxH < child.realHeight) {
			maxH = child.realHeight
		}
	}
	if (w.horizontal) {
		w.realHeight = maxH
	} else {
		w.realWidth = maxW
	}
	w.showWidgetPlacement(logNow, "boxE()")
}

func (w *cuiWidget) getGroupWH() {
	p := w.parent // the parent must be a group widget

	// update parent gocuiSize
	p.realWidth = 0
	p.realHeight = 0
	p.realHeight += me.buttonPadding // pad height for the group label
	for _, child := range p.children {
		p.realWidth += child.realWidth
		p.realHeight += child.realHeight
	}

	// compute child offset
	w.startW = p.startW
	w.startH = p.startH
	for _, child := range p.children {
		w.startH += child.realHeight
		if child == w {
			return
		}
	}
	return
}

func (w *cuiWidget) placeWidgets() {
	if (w == nil) {
		return
	}
	if (me.rootNode == nil) {
		return
	}
	p := w.parent
	if (p == nil) {
		log(logInfo, "place()", w.id, "parent == nil")
		return
	}

	switch w.widgetType {
	case toolkit.Window:
		for _, child := range w.children {
			w.startW = me.rawW
			w.startH = me.rawH
			w.nextW = me.rawW
			w.nextH = me.rawH
			w.showWidgetPlacement(logNow, "place()")
			child.placeWidgets()
			if (w.realWidth < child.realWidth) {
				w.realWidth = child.realWidth
			}
			if (w.realHeight < child.realHeight) {
				w.realHeight = child.realHeight
			}
			w.showWidgetPlacement(logNow, "place()")
		}
	case toolkit.Tab:
		for _, child := range w.children {
			w.startW = me.rawW
			w.startH = me.rawH
			w.nextW = me.rawW
			w.nextH = me.rawH
			w.showWidgetPlacement(logNow, "place()")
			child.placeWidgets()
			if (w.realWidth < child.realWidth) {
				w.realWidth = child.realWidth
			}
			if (w.realHeight < child.realHeight) {
				w.realHeight = child.realHeight
			}
			w.showWidgetPlacement(logNow, "place()")
		}
	case toolkit.Grid:
		w.showWidgetPlacement(logNow, "place()")
		w.placeGrid()
		w.showWidgetPlacement(logNow, "place()")
	case toolkit.Box:
		w.showWidgetPlacement(logNow, "place()")
		w.placeBox()
		w.showWidgetPlacement(logNow, "place()")
	case toolkit.Group:
		w.startW = p.nextW
		w.startH = p.nextH
		w.nextW = p.nextW
		w.nextH = p.nextH

		w.moveTo(p.nextW, p.nextH)

		// set real width at the beginning
		w.realWidth = w.gocuiSize.width
		w.realHeight = w.gocuiSize.height

		// indent the widgets for a group
		w.nextW = p.nextW + 4
		w.nextH = p.nextH + 3
		w.showWidgetPlacement(logNow, "place()")
		var maxW int
		for _, child := range w.children {
			child.showWidgetPlacement(logNow, "place()")
			child.placeWidgets()
			child.showWidgetPlacement(logNow, "place()")

			// increment straight down
			w.nextH += child.realHeight
			w.realHeight += child.realHeight

			// track largest width
			if (maxW < child.realWidth) {
				maxW = child.realWidth
			}

		}
		// add real width of largest child
		w.realWidth += maxW
		w.showWidgetPlacement(logNow, "place()")
	default:
		w.startW = p.nextW
		w.startH = p.nextH
		w.nextW = p.nextW
		w.nextH = p.nextH
		w.gocuiSize.w0 = w.startW
		w.gocuiSize.h0 = w.startH
		w.setWH()
		w.showWidgetPlacement(logNow, "place()")
	}
}

func (w *cuiWidget) setWH() {
	w.gocuiSize.w1 = w.gocuiSize.w0 + w.gocuiSize.width
	w.gocuiSize.h1 = w.gocuiSize.h0 + w.gocuiSize.height
}

func (w *cuiWidget) moveTo(leftW int, topH int) {
	if (w.isFake) {
		return
	}
	w.gocuiSize.w0 = leftW
	w.gocuiSize.h0 = topH
	w.gocuiSize.w1 = w.gocuiSize.w0 + w.gocuiSize.width
	w.gocuiSize.h1 = w.gocuiSize.h0 + w.gocuiSize.height
	w.showWidgetPlacement(logInfo, "moveTo()")
}

func (w *cuiWidget) placeGrid() {
	w.showWidgetPlacement(logNow, "grid0:")
	if (w.widgetType != toolkit.Grid) {
		return
	}
	w.startW = w.parent.nextW
	w.startH = w.parent.nextH
	w.nextW = w.parent.nextW
	w.nextH = w.parent.nextH

	var wCount int = 0
	var hCount int = 0
	for _, child := range w.children {
		// increment for the next child
		w.nextW = w.startW + wCount * 20
		w.nextH = w.startH + hCount * 2

		// set the child's realWidth, and grid offset
		child.parentH = hCount
		child.parentW = wCount
		if (w.widths[wCount] < child.realWidth) {
			w.widths[wCount] = child.realWidth
		}
		if (w.heights[hCount] < child.realHeight) {
			w.heights[hCount] = child.realHeight
		}
		log(logVerbose, "grid1: (w,h count)", wCount, hCount, "(X,Y)", w.x, w.y, w.name)
		child.showWidgetPlacement(logNow, "grid1: " + fmt.Sprintf("next()=(%2d,%2d)", w.nextW, w.nextH))

		if ((wCount + 1) < w.y) {
			wCount += 1
		} else {
			wCount = 0
		hCount += 1
		}
	}

	// reset the size of the whole grid
	w.realWidth = 0
	w.realHeight = 0
	for _, val := range w.widths {
		w.realWidth += val
	}
	for _, val := range w.heights {
		w.realHeight += val
	}

	for _, child := range w.children {
		child.showWidgetPlacement(logVerbose, "grid2:")
		var totalW, totalH int
		for i, val := range w.widths {
			if (i < child.parentW) {
				log(logVerbose, "grid2: (w, widths[])", i, val)
				totalW += w.widths[i]
			}
		}
		for i, h := range w.heights {
			if (i < child.parentH) {
				totalH += h
			}
		}

		// the new corner to move the child to
		w.nextW = w.startW + totalW
		w.nextH = w.startH + totalH

		child.placeWidgets()
		child.showWidgetPlacement(logInfo, "grid2:")
		log(logInfo)
	}
	// w.updateLogicalSizes()
	w.showWidgetPlacement(logNow, "grid3:")
}

func (w *cuiWidget) setRealSize() {
}
