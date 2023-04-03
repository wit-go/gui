package main

import (
	"strconv"
	"git.wit.org/wit/gui/toolkit"
//	"github.com/awesome-gocui/gocui"
)

func setupWidget(a *toolkit.Action) *cuiWidget {
	var w *cuiWidget
	w = new(cuiWidget)

	w.name = a.Name
	w.text = a.Text
	w.b = a.B
	w.i = a.I
	w.s = a.S
	w.x = a.X
	w.y = a.Y
	w.width = a.Width
	w.height = a.Height

	w.widgetType = a.WidgetType
	w.id = a.WidgetId
	// set the name used by gocui to the id
	w.cuiName = strconv.Itoa(w.id)

	w.parent = findWidget(a.ParentId, me.rootNode)
	log(logInfo, "setupWidget() w.id =", w.id, "w.parent", w.parent, "ParentId =", a.ParentId)
	if (w.id == 0) {
		me.rootNode = w
		// this is the rootNode
		return w
	}
	if (w.parent == nil) {
		log(logError, "setupWidget() ERROR: PARENT = NIL w.id =", w.id, "w.parent", w.parent, "ParentId =", a.ParentId)
		// just use the rootNode (hopefully it's not nil)
		w.parent = me.rootNode
		// return w
	}
	if (w.parent == nil) {
		log(logError, "setupWidget() ERROR: PARENT = NIL w.id =", w.id, "w.parent", w.parent, "ParentId =", a.ParentId)
		me.rootNode = w
		return w
	}
	// add this widget as a child for the parent
	w.parent.Append(w)

	if (a.WidgetType == toolkit.Box) {
		if (a.B) {
			w.horizontal = true
		} else {
			w.horizontal = false
		}
	}

	// w.showWidgetPlacement(logNow)
	return w
}

func setupCtrlDownWidget() {
	var w *cuiWidget
	w = new(cuiWidget)

	w.name = "ctrlDown"

	w.widgetType = toolkit.Flag
	w.id = -1
	me.ctrlDown = w
	// me.rootNode.Append(w)
}

func (w *cuiWidget) deleteView() {
	if (w.v != nil) {
		me.baseGui.DeleteView(w.cuiName)
	}
	w.v = nil
}

func (n *cuiWidget) Append(child *cuiWidget) {
	n.children = append(n.children, child)
	// child.parent = n
}

// find widget by number
func findWidget(i int, w *cuiWidget) (*cuiWidget) {
	if (w == nil) {
		log(logVerbose, "findWidget() Trying to find i =", i, "currently checking against w.id = nil")
		return nil
	}
	log(logVerbose, "findWidget() Trying to find i =", i, "currently checking against w.id =", w.id)

	if (w.id == i) {
		log(logInfo, "findWidget() FOUND w.id ==", i, w.widgetType, w.name)
		return w
	}

	for _, child := range w.children {
		newW := findWidget(i, child)
		log(logVerbose, "findWidget() Trying to find i =", i, "currently checking against child.id =", child.id)
		if (newW != nil) {
			return newW
		}
	}
	return nil
}

