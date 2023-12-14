package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// set isCurrent = false everywhere
func unsetCurrent(n *node) {
	w := n.tk
	w.isCurrent = false

	if n.WidgetType == toolkit.Tab {
		// n.tk.color = &colorTab
		// n.setColor()
	}

	for _, child := range n.children {
		unsetCurrent(child)
	}
}

// when adding a new widget, this will update the display
// of the current widgets if that widget is supposed
// to be in current display
func (n *node) updateCurrent() {
	log("updateCurrent()", n.Name)
	if n.WidgetType == toolkit.Tab {
		if n.IsCurrent() {
			// n.tk.color = &colorActiveT
			n.setColor(&colorActiveT)
			n.hideView()
			n.showView()
			setCurrentTab(n)
		} else {
			// n.tk.color = &colorTab
			// n.setColor()
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
	unsetCurrent(me.rootNode)

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
	unsetCurrent(me.rootNode)
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
		if (me.currentWindow == n) {
			return
		}
		if (me.currentWindow != nil) {
			unsetCurrent(me.currentWindow)
			me.currentWindow.setColor(&colorWindow)
			me.currentWindow.hideWidgets()
		}
		n.hideWidgets()
		me.currentWindow = n
		// setCurrentWindow(n) // probably delete this
		n.setColor(&colorActiveW)
		n.redoTabs(me.TabW, me.TabH)
		for _, child := range n.children {
			if (child.currentTab == true) {
				log(true, "FOUND CURRENT TAB", child.Name)
				setCurrentTab(child)
				child.placeWidgets(me.RawW, me.RawH)
				child.showWidgets()
				return
			}
		}
		/* FIXME: redo this
		if ! n.hasTabs {
		}
		*/
	case toolkit.Tab:
		if (n.IsCurrent()) {
			return // do nothing if you reclick on the already selected tab
		}
		// find the window and disable the active tab
		p := n.parent
		if (p != nil) {
			p.hideWidgets()
			p.redoTabs(me.TabW, me.TabH)
			unsetCurrent(p)
			for _, child := range p.children {
				if child.WidgetType == toolkit.Tab {
					child.setColor(&colorTab)
					n.currentTab = false
				}
			}
		}
		n.currentTab = true
		n.setColor(&colorActiveT)
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
	case toolkit.Dropdown:
		log(true, "do the dropdown here")
		if (me.ddview == nil) {
			me.ddview = addDropdown()
			tk := me.ddview.tk
			tk.gocuiSize.w0 = 20
			tk.gocuiSize.w1 = 40
			tk.gocuiSize.h0 = 10
			tk.gocuiSize.h1 = 25
			tk.v, _ = me.baseGui.SetView("ddview",
				tk.gocuiSize.w0,
				tk.gocuiSize.h0,
				tk.gocuiSize.w1,
				tk.gocuiSize.h1, 0)
			if (tk.v == nil) {
				return
			}
			tk.v.Wrap = true
			tk.v.Frame = true
			tk.v.Clear()
			fmt.Fprint(tk.v, "example.com\nwit.org\nwit.com")
			me.ddview.SetVisible(true)
			return
		}
		log(true, "doWidgetClick() visible =", me.ddview.Visible())
		if (me.ddview.Visible()) {
			me.ddview.SetVisible(false)
			me.baseGui.DeleteView("ddview")
			me.ddview.tk.v = nil
		} else {
			var dnsList string
			for i, s := range n.vals {
				log(logNow, "AddText()", n.Name, i, s)
				dnsList += s + "\n"
			}
			me.ddNode = n
			log(logNow, "new dns list should be set to:", dnsList)
			me.ddview.Text = dnsList
			me.ddview.SetText(dnsList)
			me.ddview.SetVisible(true)
		}
		for i, s := range n.vals {
			log(logNow, "AddText()", n.Name, i, s)
		}
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
		if (n.Name == "DropBox") {
			log(logNow, "click() this is the dropdown menu. set a flag here what did I click? where is the mouse?")
			log(logNow, "click() set a global dropdown clicked flag=true here")
			me.ddClicked = true
		}
		n.doWidgetClick()
	} else {
		log(logNow, "click() could not find node name =", v.Name())
	}

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
		if (n == me.ddview) {
			log(true, "findUnderMouse() found ddview")
			if n.Visible() {
				log(true, "findUnderMouse() and ddview is visable. hide it here. TODO: find highlighted row")
				found = n
				// find the actual value here and set the dropdown widget
				me.baseGui.DeleteView("ddview")
			} else {
				log(true, "findUnderMouse() I was lying, actually it's not found")
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
		n.showWidgetPlacement(logNow, "findUnderMouse() FOUND")
	}
	return found
}

// find the widget under the mouse click
func ctrlDown(g *gocui.Gui, v *gocui.View) error {
	var found *node
	// var widgets []*node
	// var f func (n *node)
	found = findUnderMouse()
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
		me.ctrlDown.hideView()
	} else {
		me.ctrlDown.showView()
	}
	me.ctrlDown.showWidgetPlacement(logNow, "ctrlDown:")
	return nil
}
