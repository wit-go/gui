package main

import (
	"strconv"
	"github.com/andlabs/ui"
	"git.wit.org/wit/gui/toolkit"
)

func (n *node) show(b bool) {
	if n.tk == nil {
		return
	}
	if n.tk.uiControl == nil {
		return
	}
	if (b) {
		n.tk.uiControl.Show()
	} else {
		n.tk.uiControl.Hide()
	}
}

func (n *node) enable(b bool) {
	if n == nil {
		panic("WHAT? enable was passed nil. How does this even happen?")
	}
	if n.tk == nil {
		return
	}
	if n.tk.uiControl == nil {
		return
	}
	if (b) {
		n.tk.uiControl.Enable()
	} else {
		n.tk.uiControl.Disable()
	}
}

func (n *node) pad(at toolkit.ActionType) {
	log(debugError, "pad()")

	t := n.tk
	if (t == nil) {
		log(debugError, "pad() toolkit struct == nil. for", n.WidgetId)
		return
	}

	switch n.WidgetType {
	case toolkit.Group:
		switch at {
		case toolkit.Margin:
			t.uiGroup.SetMargined(true)
		case toolkit.Unmargin:
			t.uiGroup.SetMargined(false)
		case toolkit.Pad:
			t.uiGroup.SetMargined(true)
		case toolkit.Unpad:
			t.uiGroup.SetMargined(false)
		}
	case toolkit.Tab:
		switch at {
		case toolkit.Margin:
			tabSetMargined(t.uiTab, true)
		case toolkit.Unmargin:
			tabSetMargined(t.uiTab, false)
		case toolkit.Pad:
			tabSetMargined(t.uiTab, true)
		case toolkit.Unpad:
			tabSetMargined(t.uiTab, false)
		}
	case toolkit.Window:
		switch at {
		case toolkit.Margin:
			t.uiWindow.SetMargined(true)
		case toolkit.Unmargin:
			t.uiWindow.SetMargined(false)
		case toolkit.Pad:
			t.uiWindow.SetBorderless(false)
		case toolkit.Unpad:
			t.uiWindow.SetBorderless(true)
		}
	case toolkit.Grid:
		switch at {
		case toolkit.Margin:
			t.uiGrid.SetPadded(true)
		case toolkit.Unmargin:
			t.uiGrid.SetPadded(false)
		case toolkit.Pad:
			t.uiGrid.SetPadded(true)
		case toolkit.Unpad:
			t.uiGrid.SetPadded(false)
		}
	case toolkit.Box:
		switch at {
		case toolkit.Margin:
			t.uiBox.SetPadded(true)
		case toolkit.Unmargin:
			t.uiBox.SetPadded(false)
		case toolkit.Pad:
			t.uiBox.SetPadded(true)
		case toolkit.Unpad:
			t.uiBox.SetPadded(false)
		}
	case toolkit.Textbox:
		log(debugError, "TODO: implement ActionType =", at)
	default:
		log(debugError, "TODO: implement pad() for", at)
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

		stretchy = true
		if (p.tk.uiBox != nil) {
			p.tk.uiBox.Append(n.tk.uiControl, stretchy)
		}
		// log(debugNow, "is there a tParent parent? =", tParent.parent)
		// tParent.uiBox.Delete(0)

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
	default:
		log(logError, "TODO: need to implement move() for type =", n.WidgetType)
		log(logError, "TODO: need to implement move() for where =", p.ParentId)
		log(logError, "TODO: need to implement move() for widget =", n.WidgetId)
	}
}

func (n *node) Delete() {
	p := n.parent
	log(debugNow, "uiDelete()", n.WidgetId, "to", p.WidgetId)

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
		log(debugNow, "tWidget.boxC =", p.Name)
		log(debugNow, "is there a tParent parent? =", p.parent)
		if (p.tk.boxC < 1) {
			log(debugNow, "Can not delete from Box. already empty. tWidget.boxC =", p.tk.boxC)
			return
		}
		p.tk.uiBox.Delete(0)
		p.tk.boxC -= 1

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
		// tParent.uiBox.Append(tWidget.uiControl, stretchy)
	default:
		log(debugError, "TODO: need to implement uiDelete() for widget =", n.WidgetId, n.WidgetType)
		log(debugError, "TODO: need to implement uiDelete() for parent =", p.WidgetId, p.WidgetType)
	}
}

func rawAction(a toolkit.Action) {
	log(logInfo, "rawAction() START a.ActionType =", a.ActionType)
	log(logInfo, "rawAction() START a.S =", a.S)

	if (a.ActionType == toolkit.InitToolkit) {
		// TODO: make sure to only do this once
		// go uiMain.Do(func() {
		// 	ui.Main(demoUI)
			// go catchActionChannel()
		// })
		// try doing this on toolkit load in init()
		return
	}

	log(logInfo, "rawAction() START a.WidgetId =", a.WidgetId, "a.ParentId =", a.ParentId)
	switch a.WidgetType {
	case toolkit.Flag:
		flag(&a)
		return
	}

	n := me.rootNode.findWidgetId(a.WidgetId)

	if (a.ActionType == toolkit.Add) {
		ui.QueueMain(func() {
			add(a)
		})
		// TODO: remove this artificial delay
		// sleep(.001)
		return
	}

	if (a.ActionType == toolkit.Dump) {
		log(debugNow, "rawAction() Dump =", a.ActionType, a.WidgetType, n.Name)
		me.rootNode.listChildren(true)
		return
	}

	if (n == nil) {
		me.rootNode.listChildren(true)
		log(true, "rawAction() ERROR findWidgetId found nil", a.ActionType, a.WidgetType)
		log(true, "rawAction() ERROR findWidgetId found nil for id =", a.WidgetId)
		log(true, "rawAction() ERROR findWidgetId found nil", a.ActionType, a.WidgetType)
		log(true, "rawAction() ERROR findWidgetId found nil for id =", a.WidgetId)
		return
		panic("findWidgetId found nil for id = " + strconv.Itoa(a.WidgetId))
	}

	switch a.ActionType {
	case toolkit.Show:
		n.show(true)
	case toolkit.Hide:
		n.show(false)
	case toolkit.Enable:
		n.enable(true)
	case toolkit.Disable:
		n.enable(false)
	case toolkit.Get:
		n.setText(&a)
	case toolkit.GetText:
		switch a.WidgetType {
		case toolkit.Textbox:
			a.S = n.S
		}
	case toolkit.Set:
		n.setText(&a)
	case toolkit.SetText:
		n.setText(&a)
	case toolkit.AddText:
		n.setText(&a)
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
		log(debugNow, "rawAction() attempt to move() =", a.ActionType, a.WidgetType)
		newParent := me.rootNode.findWidgetId(a.ParentId)
		n.move(newParent)
	default:
		log(debugError, "rawAction() Unknown =", a.ActionType, a.WidgetType)
	}
	log(logInfo, "rawAction() END =", a.ActionType, a.WidgetType)
}
