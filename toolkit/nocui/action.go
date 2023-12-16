package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func (n *node) show(b bool) {
}

func (n *node) enable(b bool) {
}

func (n *node) pad(at toolkit.ActionType) {
	switch n.WidgetType {
	case toolkit.Group:
		switch at {
		case toolkit.Margin:
			// SetMargined(true)
		case toolkit.Unmargin:
			// SetMargined(false)
		case toolkit.Pad:
			// SetMargined(true)
		case toolkit.Unpad:
			// SetMargined(false)
		}
	case toolkit.Tab:
	case toolkit.Window:
	case toolkit.Grid:
	case toolkit.Box:
	case toolkit.Textbox:
		log(logError, "TODO: implement ActionType =", at)
	default:
		log(logError, "TODO: implement pad() for", at)
	}
}

func (n *node) move(newParent *node) {
	p := n.parent

	switch p.WidgetType {
	case toolkit.Group:
	case toolkit.Tab:
		// tabSetMargined(tParent.uiTab, true)
	case toolkit.Window:
		// t.uiWindow.SetBorderless(false)
	case toolkit.Grid:
		// t.uiGrid.SetPadded(true)
	case toolkit.Box:
		log(logInfo, "TODO: move() where =", p.ParentId)
		log(logInfo, "TODO: move() for widget =", n.WidgetId)
	default:
		log(logError, "TODO: need to implement move() for type =", n.WidgetType)
		log(logError, "TODO: need to implement move() for where =", p.ParentId)
		log(logError, "TODO: need to implement move() for widget =", n.WidgetId)
	}
}

func (n *node) Delete() {
	p := n.parent
	log(logNow, "uiDelete()", n.WidgetId, "to", p.WidgetId)

	switch p.WidgetType {
	case toolkit.Group:
		// tParent.uiGroup.SetMargined(true)
	case toolkit.Tab:
		// tabSetMargined(tParent.uiTab, true)
	case toolkit.Window:
		// t.uiWindow.SetBorderless(false)
	case toolkit.Grid:
		// t.uiGrid.SetPadded(true)
	case toolkit.Box:
		log(logNow, "tWidget.boxC =", p.Name)
		log(logNow, "is there a tParent parent? =", p.parent)
		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
		// tParent.uiBox.Append(tWidget.uiControl, stretchy)
	default:
		log(logError, "TODO: need to implement uiDelete() for widget =", n.WidgetId, n.WidgetType)
		log(logError, "TODO: need to implement uiDelete() for parent =", p.WidgetId, p.WidgetType)
	}
}

func doAction(a *toolkit.Action) {
	log(logNow, "doAction() START a.ActionType =", a.ActionType)
	log(logNow, "doAction() START a.S =", a.S)

	if (a.ActionType == toolkit.InitToolkit) {
		// TODO: make sure to only do this once
		// go uiMain.Do(func() {
		// 	ui.Main(demoUI)
			// go catchActionChannel()
		// })
		// try doing this on toolkit load in init()
		return
	}

	log(logNow, "doAction() START a.WidgetId =", a.WidgetId, "a.ParentId =", a.ParentId)
	switch a.WidgetType {
	case toolkit.Root:
		me.rootNode = addNode(a)
		log(logNow, "doAction() found rootNode")
		return
	case toolkit.Flag:
		// flag(&a)
		return
	}

	n := me.rootNode.findWidgetId(a.WidgetId)

	switch a.ActionType {
	case toolkit.Add:
		addNode(a)
	case toolkit.Show:
		n.show(true)
	case toolkit.Hide:
		n.show(false)
	case toolkit.Enable:
		n.enable(true)
	case toolkit.Disable:
		n.enable(false)
	case toolkit.Get:
		// n.setText(a.S)
	case toolkit.GetText:
		switch a.WidgetType {
		case toolkit.Textbox:
			a.S = n.S
		}
	case toolkit.Set:
		// n.setText(a.S)
	case toolkit.SetText:
		// n.setText(a.S)
	case toolkit.AddText:
		// n.setText(a.S)
	case toolkit.Margin:
		n.pad(toolkit.Unmargin)
	case toolkit.Unmargin:
		n.pad(toolkit.Margin)
	case toolkit.Pad:
		n.pad(toolkit.Pad)
	case toolkit.Unpad:
		n.pad(toolkit.Unpad)
	case toolkit.Delete:
		n.Delete()
	case toolkit.Move:
		log(logNow, "doAction() attempt to move() =", a.ActionType, a.WidgetType)
		newParent := me.rootNode.findWidgetId(a.ParentId)
		n.move(newParent)
	default:
		log(logError, "doAction() Unknown =", a.ActionType, a.WidgetType)
	}
	log(logInfo, "doAction() END =", a.ActionType, a.WidgetType)
}
