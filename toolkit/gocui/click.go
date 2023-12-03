package main

import (
	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// set isCurrent = false everywhere
func UnsetCurrent(n *node) {
	w := n.tk
	w.isCurrent = false

	for _, child := range n.children {
		UnsetCurrent(child)
	}
}

// when adding a new widget, this will update the display
// of the current widgets if that widget is supposed
// to be in current display
func (n *node) updateCurrent() {
	log("updateCurrent()", n.Name)
	if n.WidgetType == toolkit.Tab {
		if n.IsCurrent() {
			setCurrentTab(n)
		}
		return
	}
	if n.WidgetType == toolkit.Window {
		if n.IsCurrent() {
			// setCurrentWindow(n)
		}
		return
	}
	if n.WidgetType == toolkit.Root {
		return
	}
	n.parent.updateCurrent()
}

// shows the widgets in a window
func setCurrentWindow(n *node) {
	if n.IsCurrent() {
		return
	}
	w := n.tk
	if n.WidgetType != toolkit.Window {
		return
	}
	UnsetCurrent(me.rootNode)

	if n.hasTabs {
		// set isCurrent = true on the first tab
		for _, child := range n.children {
			child.tk.isCurrent = true
			break
		}
	} else {
		w.isCurrent = true
	}
}

// shows the widgets in a tab
func setCurrentTab(n *node) {
	w := n.tk
	if n.WidgetType != toolkit.Tab {
		return
	}
	UnsetCurrent(me.rootNode)
	w.isCurrent = true
	p := n.parent.tk
	p.isCurrent = true
	log("setCurrent()", n.Name)
}

func (n *node) doWidgetClick() {
	switch n.WidgetType {
	case toolkit.Root:
		// THIS IS THE BEGINING OF THE LAYOUT
		log("doWidgetClick()", n.Name)
		redoWindows(0,0)
	case toolkit.Flag:
		// me.rootNode.redoColor(true)
		me.rootNode.dumpTree(true)
	case toolkit.Window:
		me.rootNode.hideWidgets()
		n.redoTabs(me.TabW, me.TabH)
		if ! n.hasTabs {
			setCurrentWindow(n)
			n.placeWidgets(me.RawW, me.RawH)
			n.showWidgets()
		}
	case toolkit.Tab:
		setCurrentTab(n)
		n.placeWidgets(me.RawW, me.RawH)
		n.showWidgets()
	case toolkit.Group:
		// n.placeWidgets(p.tk.startH, newH)
		n.toggleTree()
	case toolkit.Checkbox:
		if (n.B) {
			n.setCheckbox(false)
		} else {
			n.setCheckbox(true)
		}
		n.doUserEvent()
	case toolkit.Grid:
		n.placeGrid(n.tk.size.w0, n.tk.size.h0)
		n.showWidgets()
	case toolkit.Box:
		// w.showWidgetPlacement(logNow, "drawTree()")
		if (n.horizontal) {
			log("BOX IS HORIZONTAL", n.Name)
		} else {
			log("BOX IS VERTICAL", n.Name)
		}
		// n.placeWidgets()
		n.toggleTree()
	case toolkit.Button:
		n.doUserEvent()
	default:
	}
}

var toggle bool = true
func (n *node) toggleTree() {
	if (toggle) {
		n.drawTree(toggle)
		toggle = false
	} else {
		n.hideWidgets()
		toggle = true
	}
}


// display the widgets in the binary tree
func (n *node) drawTree(draw bool) {
	w := n.tk
	if (w == nil) {
		return
	}
	n.showWidgetPlacement(logNow, "drawTree()")
	if (draw) {
		// w.textResize()
		n.showView()
	} else {
		n.deleteView()
	}

	for _, child := range n.children {
		child.drawTree(draw)
	}
}

func click(g *gocui.Gui, v *gocui.View) error {
	// var l string
	// var err error

	log(logVerbose, "click() START", v.Name())
	// n := me.rootNode.findWidgetName(v.Name())
	n := findUnderMouse()
	if (n != nil) {
		log(logNow, "click() Found widget =", n.WidgetId, n.Name, ",", n.Text)
		n.doWidgetClick()
	} else {
		log(logNow, "click() could not find node name =", v.Name())
	}
	/*
	i, err := strconv.Atoi(v.Name())
	if (err != nil) {
		log(logError, "click() Can't find widget. error =", err)
	} else {
		log(logVerbose, "click() ok v.Name() =", v.Name())
		n := me.rootNode.findWidgetId(i)
		if (n == nil) {
			log(logError, "click() CANT FIND VIEW in binary tree. v.Name =", v.Name())
			return nil
		}
		log(logNow, "click() Found widget =", n.WidgetId, n.Name, ",", n.Text)
		n.doWidgetClick()
		return nil
	}
	*/

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	log(logVerbose, "click() END")
	return nil
}
func findUnderMouse() *node {
	var found *node
	var widgets []*node
	var f func (n *node)
	w, h := me.baseGui.MousePosition()

	// find buttons that are below where the mouse button click
	f = func(n *node) {
		widget := n.tk
		// ignore widgets that are not visible
		if n.Visible() {
			if ((widget.gocuiSize.w0 <= w) && (w <= widget.gocuiSize.w1) &&
			(widget.gocuiSize.h0 <= h) && (h <= widget.gocuiSize.h1)) {
				widgets = append(widgets, n)
				found = n
			}
		}

		for _, child := range n.children {
			f(child)
		}
	}
	f(me.rootNode)
	// widgets has everything that matches
	// TODO: pop up menu with a list of them
	for _, n := range widgets {
		//log(logNow, "ctrlDown() FOUND widget", widget.id, widget.name)
		n.showWidgetPlacement(logNow, "ctrlDown() FOUND")
	}
	return found
}

// find the widget under the mouse click
func ctrlDown(g *gocui.Gui, v *gocui.View) error {
	var found *node
	// var widgets []*node
	// var f func (n *node)
	found = findUnderMouse()
	/*
	w, h := g.MousePosition()

	// find buttons that are below where the mouse button click
	f = func(n *node) {
		widget := n.tk
		// ignore widgets that are not visible
		if n.Visible() {
			if ((widget.gocuiSize.w0 <= w) && (w <= widget.gocuiSize.w1) &&
			(widget.gocuiSize.h0 <= h) && (h <= widget.gocuiSize.h1)) {
				widgets = append(widgets, n)
				found = n
			}
		}

		for _, child := range n.children {
			f(child)
		}
	}
	f(me.rootNode)
	*/
	if (me.ctrlDown == nil) {
		setupCtrlDownWidget()
		me.ctrlDown.Text = found.Name
		me.ctrlDown.tk.cuiName = "ctrlDown"
		// me.ctrlDown.parent = me.rootNode
	}
	cd := me.ctrlDown.tk
	if (found == nil) {
		found = me.rootNode
	}
	me.ctrlDown.Text = found.Name
	newR := found.realGocuiSize()
	cd.gocuiSize.w0 = newR.w0
	cd.gocuiSize.h0 = newR.h0
	cd.gocuiSize.w1 = newR.w1
	cd.gocuiSize.h1 = newR.h1
	if me.ctrlDown.Visible() {
		me.ctrlDown.deleteView()
	} else {
		me.ctrlDown.updateView()
	}
	me.ctrlDown.showWidgetPlacement(logNow, "ctrlDown:")
	return nil
}
