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
		w.nextW = p.nextW
		w.nextH = p.nextH
		w.gridBounds()
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
		w.moveTo(p.nextW, p.nextH)

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
		w.moveTo(p.nextW, p.nextH)
		w.nextW = w.realSize.w1
		w.nextH = w.realSize.h1
	}
}

func (w *cuiWidget) moveTo(leftW int, topH int) {
	w.realSize.w0 = leftW
	w.realSize.h0 = topH 
	w.realSize.w1 = leftW + w.realWidth
	w.realSize.h1 = topH + w.realHeight

	w.logicalSize.w0 = w.realSize.w0
	w.logicalSize.h0 = w.realSize.h0
	w.logicalSize.w1 = w.realSize.w1
	w.logicalSize.h1 = w.realSize.h1

	w.showWidgetPlacement(logNow, "moveTo()")
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

func (w *cuiWidget) gridBounds() {
	w.showWidgetPlacement(logNow, "gridBounds:")
	p := w.parent

	/*
	for a := 0; a < w.x; a++ {
		for b := 0; b < w.y; b++ {
			log(logNow, "gridBounds() (w,h)", a, b,
				"logical(W,H)", w.logicalW[a], w.logicalH[b],
				"p.next(W,H)", p.nextW, p.nextH)
		}
		log("\n")
	}
	*/
	var wCount int = 0
	var hCount int = 0
	for _, child := range w.children {
		// increment for the next child
		w.nextW = p.nextW + wCount * 20
		w.nextH = p.nextH + hCount * 2
		child.redoBox(true)

		// set the child's realWidth, and grid offset
		child.parentH = hCount
		child.parentW = wCount
		if (w.logicalW[wCount] < child.realWidth) {
			w.logicalW[wCount] = child.realWidth
		}
		if (w.logicalH[hCount] < child.realHeight) {
			w.logicalH[hCount] = child.realHeight
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


	for _, child := range w.children {
		child.showWidgetPlacement(logVerbose, "gridBounds:")
		var totalW, totalH int
		for i, val := range w.logicalW {
			if (i < child.parentW) {
				log(logVerbose, "gridBounds() (w, logicalW[])", i, val)
				totalW += w.logicalW[i]
			}
		}
		for i, h := range w.logicalH {
			if (i < child.parentH) {
				totalH += h
			}
		}

		// the new corner to move the child to
		realW := w.nextW + totalW
		realH := w.nextH + totalH


		log(logInfo, "gridBounds()", child.id, "parent (W,H) =", child.parentW, child.parentH,
			"total (W,H) =", totalW, totalH,
			"real (W,H) =", realW, realH)
		child.moveTo(realW, realH)
		child.showWidgetPlacement(logInfo, "gridBounds:")
		log(logInfo)
	}
	w.updateLogicalSizes()
	w.showWidgetPlacement(logNow, "gridBounds:")
}
