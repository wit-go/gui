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

	w.gocuiSize.width = t + 2
	w.gocuiSize.height = me.defaultHeight
	w.gocuiSize.startW = fakeStartWidth
	w.gocuiSize.startH = fakeStartHeight

	fakeStartHeight += 3
	if (fakeStartHeight > 24) {
		fakeStartHeight = 3
		fakeStartWidth += 20
	}
	w.setWH()
	w.showWidgetPlacement(logNow, "setFake()")
}

// find the start (w,h) for child a inside a box widget
func (w *cuiWidget) getBoxWH() {
	p := w.parent // the parent must be a box widget

	// update parent gocuiSize
	p.realWidth = 0
	p.realHeight = 0
	for _, child := range p.children {
		p.realWidth += child.realWidth
		p.realHeight += child.realHeight
	}

	// compute child offset
	w.startW = p.startW
	w.startH = p.startH
	for _, child := range p.children {
		if (p.horizontal) {
			log("BOX IS HORIZONTAL (w,h)", w.startW, w.startH)
			log("BOX IS HORIZONTAL (w,h)", w.startW, w.startH)
			log("BOX IS HORIZONTAL (w,h)", w.startW, w.startH)
			w.startW += child.realWidth
		} else {
			log("BOX IS VERTICAL (w,h)", w.startW, w.startH)
			log("BOX IS VERTICAL (w,h)", w.startW, w.startH)
			log("BOX IS VERTICAL (w,h)", w.startW, w.startH)
			w.startH += child.realHeight
		}
		if child == w {
			return
		}
	}
	return
}

// find the start (w,h) for child a inside a Group widget
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

// find the start (w,h) for child a inside a Grid widget
func (w *cuiWidget) getGridWH() {
	p := w.parent
	w.startW = p.startW
	w.startH = p.startH
	w.nextW = p.startW
	w.nextH = p.startH
	w.drawGrid()
}

func (w *cuiWidget) drawBox() {
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

	switch w.widgetType {
	case toolkit.Window:
		// draw only one thing
		for _, child := range w.children {
			child.drawBox()
			return
		}
	case toolkit.Tab:
		// draw only one thing
		for _, child := range w.children {
			child.drawBox()
			return
		}
	case toolkit.Grid:
		w.startW = p.startW
		w.startH = p.startH
		w.getGridWH()
		w.showWidgetPlacement(logNow, "drawBox:")
	case toolkit.Box:
		w.startW = p.startW
		w.startH = p.startH
		var maxW int
		var maxH int
		for _, child := range w.children {
			child.drawBox()
			if (w.horizontal) {
				log("BOX IS HORIZONTAL")
				// expand based on the child width
				w.startW += child.realWidth
			} else {
				log("BOX IS VERTICAL")
				// expand based on the child height
				w.startH += child.realHeight
			}
			if (maxW < child.realWidth) {
				maxW = child.realWidth
			}
			if (maxH < child.realHeight) {
				maxH = child.realHeight
			}
		}
		w.realWidth = maxW
		w.realHeight = maxH
		w.showWidgetPlacement(logNow, "drawBox:")
	case toolkit.Group:
		w.startW = p.startW
		w.startH = p.startH
		w.gocuiSize.startW = w.startW
		w.gocuiSize.startH = w.startH
		w.realWidth = w.gocuiSize.width
		w.realHeight = w.gocuiSize.height
		w.setWH()

		w.startW = p.startW + 4
		w.startH = p.startH + 3
		var maxW int
		var maxH int
		for _, child := range w.children {
			child.drawBox()
			// reset nextW to straight down
			w.startH += child.realHeight
			if (maxW < child.realWidth) {
				maxW = child.realWidth
			}
			if (maxH < child.realHeight) {
				maxH = child.realHeight
			}
		}
		w.realWidth += maxW
		w.realHeight += maxH
		w.showWidgetPlacement(logNow, "drawBox:")
	default:
		w.startW = p.startW
		w.startH = p.startH
		w.gocuiSize.startW = w.startW
		w.gocuiSize.startH = w.startH
		w.setWH()
		w.showWidgetPlacement(logNow, "drawBox:")
	}
}

func (w *cuiWidget) setWH() {
	w.gocuiSize.w0 = w.gocuiSize.startW
	w.gocuiSize.h0 = w.gocuiSize.startH
	w.gocuiSize.w1 = w.gocuiSize.w0 + w.gocuiSize.width
	w.gocuiSize.h1 = w.gocuiSize.h0 + w.gocuiSize.height
}

func (w *cuiWidget) moveTo(leftW int, topH int) {
	if (w.isFake) {
		return
	}
	w.gocuiSize.startW = leftW
	w.gocuiSize.startH = topH

	w.setWH()
	w.showWidgetPlacement(logNow, "moveTo()")
}

/*
func (w *cuiWidget) updateLogicalSizes() {
	for _, child := range w.children {
		// if (w.isReal)
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
*/

func (w *cuiWidget) drawGrid() {
	w.showWidgetPlacement(logNow, "gridBounds:")

	var wCount int = 0
	var hCount int = 0
	for _, child := range w.children {
		// increment for the next child
		w.nextW = w.startW + wCount * 20
		w.nextH = w.startH + hCount * 2
		// child.drawBox()

		// set the child's realWidth, and grid offset
		child.parentH = hCount
		child.parentW = wCount
		if (w.widths[wCount] < child.realWidth) {
			w.widths[wCount] = child.realWidth
		}
		if (w.heights[hCount] < child.realHeight) {
			w.heights[hCount] = child.realHeight
		}
		log(logNow, "redoBox(GRID) (w,h count)", wCount, hCount, "(X,Y)", w.x, w.y, w.name)
		child.showWidgetPlacement(logNow, "grid:")

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
		child.showWidgetPlacement(logVerbose, "gridBounds:")
		var totalW, totalH int
		for i, val := range w.widths {
			if (i < child.parentW) {
				log(logVerbose, "gridBounds() (w, widths[])", i, val)
				totalW += w.widths[i]
			}
		}
		for i, h := range w.heights {
			if (i < child.parentH) {
				totalH += h
			}
		}

		// the new corner to move the child to
		realW := w.nextW + totalW
		realH := w.nextH + totalH

		log(logNow, "gridBounds()", child.id, "parent (W,H) =", child.parentW, child.parentH,
			"total (W,H) =", totalW, totalH,
			"real (W,H) =", realW, realH)
		w.startW = realW
		w.startH = realH
		child.drawBox()
		child.showWidgetPlacement(logInfo, "gridBounds:")
		log(logInfo)
	}
	// w.updateLogicalSizes()
	w.showWidgetPlacement(logNow, "gridBounds:")
}
