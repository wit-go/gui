package main

import (
	"strings"
	"git.wit.org/wit/gui/toolkit"
)

func (n *node) placeBox(startW int, startH int) {
	if (n.WidgetType != toolkit.Box) {
		return
	}
	n.showWidgetPlacement(logNow, "boxS()")

	newW := startW
	newH := startH
	for _, child := range n.children {
		child.placeWidgets(newW, newH)
		// n.showWidgetPlacement(logNow, "boxS()")
		newR := child.realGocuiSize()
		w := newR.w1 - newR.w0
		h := newR.h1 - newR.h0
		if (n.horizontal) {
			log(logNow, "BOX IS HORIZONTAL", n.Name, "newWH()", newW, newH, "child()", w, h, child.Name)
			// expand based on the child width
			newW += w
		} else {
			log(logNow, "BOX IS VERTICAL  ", n.Name, "newWH()", newW, newH, "child()", w, h, child.Name)
			// expand based on the child height
			newH += h
		}
	}

	// just compute this every time?
	// newR := n.realGocuiSize()

	n.showWidgetPlacement(logNow, "boxE()")
}

func (n *node) placeWidgets(startW int, startH int) {
	if (n == nil) {
		return
	}
	if (me.rootNode == nil) {
		return
	}

	switch n.WidgetType {
	case toolkit.Window:
		for _, child := range n.children {
			child.placeWidgets(me.RawW, me.RawH)
			return
		}
	case toolkit.Tab:
		for _, child := range n.children {
			child.placeWidgets(me.RawW, me.RawH)
			return
		}
	case toolkit.Grid:
		n.placeGrid(startW, startH)
	case toolkit.Box:
		n.placeBox(startW, startH)
	case toolkit.Group:
		// move the group to the parent's next location
		n.gocuiSetWH(startW, startH)
		n.showWidgetPlacement(logNow, "group()")

		newW := startW + me.GroupPadW
		newH := startH + 3 // normal hight of the group label
		// now move all the children aka: run place() on them
		for _, child := range n.children {
			child.placeWidgets(newW, newH)
			newR := child.realGocuiSize()
			// w := newR.w1 - newR.w0
			h := newR.h1 - newR.h0

			// increment straight down
			newH += h
		}
	default:
		n.gocuiSetWH(startW, startH)
		// n.moveTo(startW, startH)
	}
}

func (n *node) placeGrid(startW int, startH int) {
	w := n.tk
	n.showWidgetPlacement(logInfo, "grid0:")
	if (n.WidgetType != toolkit.Grid) {
		return
	}

	// first compute the max sizes of the rows and columns
	for _, child := range n.children {
		newR := child.realGocuiSize()
		childW := newR.w1 - newR.w0
		childH := newR.h1 - newR.h0

		// set the child's realWidth, and grid offset
		if (w.widths[child.AtW] < childW) {
			w.widths[child.AtW] = childW
		}
		if (w.heights[child.AtH] < childH) {
			w.heights[child.AtH] = childH
		}
		// child.showWidgetPlacement(logInfo, "grid: ")
		log(logVerbose, "placeGrid:", child.Name, "child()", childW, childH, "At()", child.AtW, child.AtH)
	}

	// find the width and height offset of the grid for AtW,AtH
	for _, child := range n.children {
		child.showWidgetPlacement(logInfo, "grid1:")

		var totalW, totalH int
		for i, w := range w.widths {
			if (i < child.AtW) {
				totalW += w
			}
		}
		for i, h := range w.heights {
			if (i < child.AtH) {
				totalH += h
			}
		}

		// the new corner to move the child to
		newW := startW + totalW
		newH := startH + totalH

		log(logVerbose, "placeGrid:", child.Name, "new()", newW, newH, "At()", child.AtW, child.AtH)
		child.placeWidgets(newW, newH)
		child.showWidgetPlacement(logInfo, "grid2:")
	}
	n.showWidgetPlacement(logInfo, "grid3:")
}

// computes the real, actual size of all the gocli objects in a widget
func (n *node) realGocuiSize() *rectType {
	var f func (n *node, r *rectType)
	newR := new(rectType)
	// initialize the values to opposite
	newR.w0 = 80
	newR.h0 = 24
	if me.baseGui != nil {
		maxW, maxH := me.baseGui.Size()
		newR.w0 = maxW
		newR.h0 = maxH
	}
	newR.w1 = 0
	newR.h1 = 0

	// expand the rectangle to the biggest thing displayed
	f = func(n *node, r *rectType) {
		newR := n.tk.gocuiSize
		if ! n.tk.isFake {
			if r.w0 > newR.w0 {
				r.w0 = newR.w0
			}
			if r.h0 > newR.h0 {
				r.h0 = newR.h0
			}
			if r.w1 < newR.w1 {
				r.w1 = newR.w1
			}
			if r.h1 < newR.h1 {
				r.h1 = newR.h1
			}
		}
		for _, child := range n.children {
			f(child, r)
		}
	}
	f(n, newR)
	return newR
}

func (n *node) textSize() (int, int) {
	var width, height int

	for _, s := range strings.Split(n.Text, "\n") {
		if (width < len(s)) {
			width = len(s)
		}
		height += 1
	}
	return width, height
}
