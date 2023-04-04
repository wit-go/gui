package main

import (
	"fmt"
	"errors"
	"strconv"
	"strings"

	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

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

func (w *cuiWidget) doWidgetClick() {
	switch w.widgetType {
	case toolkit.Root:
		me.rootNode.redoTabs(true)
		// me.rootNode.redoFake(true)
	case toolkit.Flag:
		me.rootNode.redoColor(true)
	case toolkit.Window:
		w.redoBox(true)
		w.toggleTree()
	case toolkit.Tab:
		w.redoBox(true)
		w.toggleTree()
	case toolkit.Grid:
		w.gridBounds()
		for _, child := range w.children {
			child.showWidgetPlacement(logNow, "gridBounds:")
			if (child.v == nil) {
				child.drawView()
			} else {
				child.deleteView()
			}
		}
		// w.toggleTree()
		// w.redoBox(true)
	case toolkit.Box:
		// w.showWidgetPlacement(logNow, "drawTree()")
		if (w.horizontal) {
			log("BOX IS HORIZONTAL", w.nextW, w.nextH, w.name)
		} else {
			log("BOX IS VERTICAL", w.nextW, w.nextH, w.name)
		}
		w.redoBox(true)
		w.toggleTree()
	default:
	}
}

var toggle bool = true
func (w *cuiWidget) toggleTree() {
	if (toggle) {
		w.drawTree(toggle)
		toggle = false
	} else {
		w.hideWidgets()
		toggle = true
	}
}


// display the widgets in the binary tree
func (w *cuiWidget) drawTree(draw bool) {
	if (w == nil) {
		return
	}
	w.showWidgetPlacement(logNow, "drawTree()")
	if (draw) {
		// w.textResize()
		w.drawView()
	} else {
		w.deleteView()
	}

	for _, child := range w.children {
		child.drawTree(draw)
	}
}

func click(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	log(logNow, "click() START", v.Name())
	i, err := strconv.Atoi(v.Name())
	if (err != nil) {
		log(logNow, "click() Can't find widget. error =", err)
	} else {
		log(logNow, "click() ok v.Name() =", v.Name())
		w := findWidget(i, me.rootNode)
		if (w == nil) {
			log(logError, "click() CANT FIND VIEW in binary tree. v.Name =", v.Name())
			return nil
		}
		log(logNow, "click() Found widget =", w.id, w.name, ",", w.text)
		w.doWidgetClick()
		return nil
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2, 0); err == nil || errors.Is(err, gocui.ErrUnknownView) {
		v.Clear()
		v.SelBgColor = gocui.ColorCyan
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, l)
	}

	// this seems to delete the button(?)
	// g.SetViewOnBottom(v.Name())
	log(logNow, "click() END")
	return nil
}

// display the widgets in the binary tree

func ctrlDown(g *gocui.Gui, v *gocui.View) error {
	var found *cuiWidget
	var widgets []*cuiWidget
	var f func (widget *cuiWidget)
	w, h := g.MousePosition()

	// find buttons that are below where the mouse button click
	f = func(widget *cuiWidget) {
		// if ((widget.logicalSize.w0 < w) && (w < widget.logicalSize.w1)) {
		if ((widget.realSize.w0 <= w) && (w <= widget.realSize.w1) &&
		(widget.realSize.h0 <= h) && (h <= widget.realSize.h1)) {
			widgets = append(widgets, widget)
			found = widget
		}

		for _, child := range widget.children {
			f(child)
		}
	}
	f(me.rootNode)
	var t string
	for _, widget := range widgets {
		log(logNow, "ctrlDown() FOUND widget", widget.id, widget.name)
		t += widget.cuiName + " " + widget.name + "\n"
		// widget.showWidgetPlacement(logNow, "drawTree()")
	}
	t = strings.TrimSpace(t)
	if (me.ctrlDown == nil) {
		setupCtrlDownWidget()
		me.ctrlDown.text = "ctrlDown" // t
		me.ctrlDown.cuiName = "ctrlDown"
		me.ctrlDown.parent = me.rootNode
	}
	if (found == nil) {
		found = me.rootNode
	}
	found.updateLogicalSizes()
	me.ctrlDown.realSize.w0 = found.logicalSize.w0
	me.ctrlDown.realSize.w1 = found.logicalSize.w1
	me.ctrlDown.realSize.h0 = found.logicalSize.h0
	me.ctrlDown.realSize.h1 = found.logicalSize.h1

	if (me.ctrlDown.v == nil) {
		me.ctrlDown.text = found.text
		me.ctrlDown.showWidgetPlacement(logNow, "drawTree()")
		me.ctrlDown.drawView()
	} else {
		me.ctrlDown.deleteView()
	}

	/*
	v, err := g.SetView("ctrlDown", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2, 0)
	if (err != nil) {
		log(logError, "ctrlDown() g.SetView() error:", err)
		return
	}
	v.Clear()
	v.SelBgColor = gocui.ColorCyan
	v.SelFgColor = gocui.ColorBlack
	fmt.Fprintln(v, l)
	*/

	log(logNow, "ctrlDown()", w, h)
	return nil
}
