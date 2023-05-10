package main

import (
	// "fmt"
	// "errors"
	"strconv"
	"strings"

	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// set isCurrent = false everywhere
func UnsetCurrent(w *cuiWidget) {
	w.isCurrent = false

	for _, child := range w.children {
		UnsetCurrent(child)
	}
}

func updateCurrentTabs() {
	me.rootNode.nextW = 0
	me.rootNode.nextH = 0
	me.rootNode.redoTabs(true)
}

// when adding a new widget, this will update the display
// of the current widgets if that widget is supposed
// to be in current display
func (w *cuiWidget) updateCurrent() {
	if w.widgetType == toolkit.Tab {
		if w.IsCurrent() {
			setCurrentTab(w)
		}
		return
	}
	if w.widgetType == toolkit.Window {
		if w.IsCurrent() {
			setCurrentWindow(w)
		}
		return
	}
	if w.widgetType == toolkit.Root {
		return
	}
	w.parent.updateCurrent()
}

// shows the widgets in a window
func setCurrentWindow(w *cuiWidget) {
	if w.widgetType != toolkit.Window {
		return
	}
	UnsetCurrent(me.rootNode)
	me.rootNode.hideWidgets()

	// THIS IS THE BEGINING OF THE LAYOUT
	me.rootNode.nextW = 0
	me.rootNode.nextH = 0

	w.isCurrent = true
	if w.hasTabs {
		// set isCurrent = true on the first tab
		for _, child := range w.children {
			child.isCurrent = true
			break
		}
	}
	me.rootNode.redoTabs(true)

	w.placeWidgets()
	w.showWidgets()
}

// shows the widgets in a tab
func setCurrentTab(w *cuiWidget) {
	if w.widgetType != toolkit.Tab {
		return
	}
	me.current = w
	UnsetCurrent(me.rootNode)
	me.rootNode.hideWidgets()
	w.isCurrent = true
	w.parent.isCurrent = true
	updateCurrentTabs()
	w.placeWidgets()
	w.showWidgets()
}

func (w *cuiWidget) doWidgetClick() {
	switch w.widgetType {
	case toolkit.Root:
		// THIS IS THE BEGINING OF THE LAYOUT
		me.rootNode.nextW = 0
		me.rootNode.nextH = 0
		me.rootNode.redoTabs(true)
	case toolkit.Flag:
		// me.rootNode.redoColor(true)
		me.rootNode.dumpTree(true)
	case toolkit.Window:
		setCurrentWindow(w)
	case toolkit.Tab:
		setCurrentTab(w)
	case toolkit.Group:
		w.placeWidgets()
		w.toggleTree()
	case toolkit.Checkbox:
		if (w.b) {
			w.setCheckbox(false)
		} else {
			w.setCheckbox(true)
		}
		w.doUserEvent()
	case toolkit.Grid:
		me.rootNode.hideWidgets()
		w.placeGrid()
		w.showWidgets()
	case toolkit.Box:
		// w.showWidgetPlacement(logNow, "drawTree()")
		if (w.horizontal) {
			log("BOX IS HORIZONTAL", w.name)
		} else {
			log("BOX IS VERTICAL", w.name)
		}
		w.placeWidgets()
		w.toggleTree()
	case toolkit.Button:
		w.doUserEvent()
	default:
	}
}

// this passes the user event back from the plugin
func (w *cuiWidget) doUserEvent() {
	if (me.callback == nil) {
		log(logError, "doUserEvent() no callback channel was configured")
		return
	}
	var a toolkit.Action
	a.WidgetId = w.id
	a.Name = w.name
	a.Text = w.text
	a.B = w.b
	a.ActionType = toolkit.User
	me.callback <- a
	log(logNow, "END:   sent a button click callback()")
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
		w.showView()
	} else {
		w.deleteView()
	}

	for _, child := range w.children {
		child.drawTree(draw)
	}
}

func click(g *gocui.Gui, v *gocui.View) error {
	// var l string
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

	/*
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
	*/

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
		if ((widget.gocuiSize.w0 <= w) && (w <= widget.gocuiSize.w1) &&
		(widget.gocuiSize.h0 <= h) && (h <= widget.gocuiSize.h1)) {
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
		// log(logNow, "ctrlDown() FOUND widget", widget.id, widget.name)
		t += widget.cuiName + " " + widget.name + "\n"
		widget.showWidgetPlacement(logNow, "ctrlDown() FOUND")
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
	// ? TODO: found.setRealSize()
	me.ctrlDown.gocuiSize.w0 = found.startW
	me.ctrlDown.gocuiSize.h0 = found.startH
	me.ctrlDown.gocuiSize.w1 =  me.ctrlDown.gocuiSize.w0 + found.realWidth
	me.ctrlDown.gocuiSize.h1 =  me.ctrlDown.gocuiSize.h0 + found.realHeight
	if (me.ctrlDown.v == nil) {
		me.ctrlDown.text = found.text
		me.ctrlDown.showWidgetPlacement(logNow, "ctrlDown:")
		me.ctrlDown.showView()
	} else {
		me.ctrlDown.deleteView()
	}

	log(logNow, "ctrlDown()", w, h)
	return nil
}
