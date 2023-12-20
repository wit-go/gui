package main

// if you include more than just this import
// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
import "go.wit.com/gui/toolkit"

// delete the child widget from the parent
// p = parent, c = child
func (n *node) destroy() {
	pId := n.parent.WidgetId
	cId := n.WidgetId
	log(logNow, "delete()", pId, cId)

	pt := n.parent.tk
	ct := n.tk
	if (ct == nil) {
		log(true, "delete FAILED (ct = mapToolkit[c] == nil) for c", pId, cId)
		// this pukes out a whole universe of shit
		// listMap()
		return
	}

	switch n.WidgetType {
	case toolkit.Button:
		log(true, "Should delete Button here:", n.Name)
		log(true, "Parent:")
		pt.Dump(true)
		log(true, "Child:")
		ct.Dump(true)
		if (pt.uiBox == nil) {
			log(true, "Don't know how to destroy this")
		} else {
			log(true, "Fuck it, destroy the whole box", n.parent.Name)
			// pt.uiBox.Destroy() // You have a bug: You cannot destroy a uiControl while it still has a parent.
			pt.uiBox.SetPadded(false)
			pt.uiBox.Delete(4)
			ct.uiButton.Disable()
			// ct.uiButton.Hide()
			ct.uiButton.Destroy()
		}

	case toolkit.Window:
		log(true, "Should delete Window here:", n.Name)
	default:
		log(true, "Fuckit, let's destroy a button")
		if (ct.uiButton != nil) {
			pt.uiBox.Delete(4)
			ct.uiButton.Destroy()
		}
	}
}
