package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *node) doWidgetClick() {
	switch n.WidgetType {
	case toolkit.Root:
		// THIS IS THE BEGINING OF THE LAYOUT
		// rootNode.nextW = 0
		// rootNode.nextH = 0
		// rootNode.redoTabs(true)
	case toolkit.Flag:
		// me.rootNode.redoColor(true)
		// rootNode.dumpTree(true)
	case toolkit.Window:
		// setCurrentWindow(w)
	case toolkit.Tab:
		// setCurrentTab(w)
	case toolkit.Group:
		// n.placeWidgets()
		// n.toggleTree()
	case toolkit.Checkbox:
		if (n.B) {
			// n.setCheckbox(false)
		} else {
			// n.setCheckbox(true)
		}
		n.doUserEvent()
	case toolkit.Grid:
		// rootNode.hideWidgets()
		// n.placeGrid()
		// n.showWidgets()
	case toolkit.Box:
		// n.showWidgetPlacement(logNow, "drawTree()")
		if (n.B) {
			log("BOX IS HORIZONTAL", n.Name)
		} else {
			log("BOX IS VERTICAL", n.Name)
		}
	case toolkit.Button:
		n.doUserEvent()
	default:
	}
}

func (n *node) doUserEvent() {
	if (callback == nil) {
		log(logError, "doUserEvent() callback == nil", n.WidgetId)
		return
	}
	var a toolkit.Action
	a.WidgetId = n.WidgetId
	a.Name = n.Name
	a.Text = n.Text
	a.S = n.S
	a.I = n.I
	a.B = n.B
	a.ActionType = toolkit.User
	log(logInfo, "doUserEvent() START: send a user event to the callback channel")
	callback <- a
	log(logInfo, "doUserEvent() END:   sent a user event to the callback channel")
	return
}
