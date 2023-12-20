package main

import (
	"go.wit.com/gui/toolkit"
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
		n.doUserEvent()
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
