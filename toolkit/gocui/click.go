package main

import (
	"fmt"
	"errors"
	"strconv"
	"strings"

	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

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

		// w.toggleTree()
		// me.rootNode.redoColor(true)
	case toolkit.Box:
		w.showWidgetPlacement(logNow, "drawTree()")
		if (w.horizontal) {
			log("BOX IS HORIZONTAL", w.nextW, w.nextH, w.name)
		} else {
			log("BOX IS VERTICAL", w.nextW, w.nextH, w.name)
		}
		// w.redoBox(true)
	default:
		// w.textResize()
		// something
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
		w.textResize()
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
	var widgets []*cuiWidget
	var f func (widget *cuiWidget)
	w, h := g.MousePosition()

	f = func(widget *cuiWidget) {
		if ((widget.logicalSize.w0 < w) && (w < widget.logicalSize.w1)) {
			widgets = append(widgets, widget)
		}

		for _, child := range widget.children {
			f(child)
		}
	}
	f(me.rootNode)
	var t string
	for i, widget := range widgets {
		log(logNow, "ctrlDown() FOUND widget", i, widget.name)
		t += widget.cuiName + " " + widget.name + "\n"
		// widget.showWidgetPlacement(logNow, "drawTree()")
	}
	t = strings.TrimSpace(t)
	if (me.ctrlDown == nil) {
		setupCtrlDownWidget()
	}
	me.ctrlDown.text = t
	me.ctrlDown.realSize.w0 = w
	me.ctrlDown.realSize.h0 = h
	me.ctrlDown.textResize()
	me.ctrlDown.drawView()

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
