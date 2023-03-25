package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func show(a *toolkit.Action) {
	if (a == nil) {
		log(debugError, "nil is probably already hidden")
		return
	}
	log(debugError, "show()", a.WidgetId)

	t := andlabs[a.WidgetId]
	if (t == nil) {
		log(debugError, "show() toolkit struct == nil. for", a.WidgetId)
		return
	}

	if (a.B) {
		t.uiControl.Show()
	} else {
		t.uiControl.Hide()
	}
}

func enable(a *toolkit.Action) {
	if (a == nil) {
		log(debugError, "nil is probably already hidden")
		return
	}
	log(debugError, "enable() name =", a.WidgetId)

	t := andlabs[a.WidgetId]
	if (t == nil) {
		log(debugToolkit, "enable() toolkit struct == nil. for id =", a.WidgetId)
		return
	}

	if (a.B) {
		t.uiControl.Enable()
	} else {
		t.uiControl.Disable()
	}
}

func pad(a *toolkit.Action) {
	if (a == nil) {
		log(debugError, "pad() ERROR: nil is probably already hidden")
		return
	}
	log(debugError, "pad()")

	t := andlabs[a.WidgetId]
	if (t == nil) {
		log(debugError, "pad() toolkit struct == nil. for", a.WidgetId)
		return
	}

	switch t.Type {
	case toolkit.Group:
		switch a.Type {
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
		switch a.Type {
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
		switch a.Type {
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
		switch a.Type {
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
		switch a.Type {
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
		log(debugError, "TODO: implement expand for", a.Type)
		log(debugError, "TODO: implement expand for", a.Type)
		log(debugError, "TODO: implement expand for", a.Type)
		log(debugError, "TODO: implement expand for", a.Type)
	default:
		log(debugError, "TODO: implement pad() for", a.Type)
	}
}

func move(a *toolkit.Action) {
	log(debugNow, "move()", a.WidgetId, "to", a.WhereId)

	tWidget := andlabs[a.WidgetId]
	if (tWidget == nil) {
		log(debugError, "move() ERROR: toolkit struct == nil. for", a.WidgetId)
		return
	}

	tWhere := andlabs[a.WhereId]
	if (tWhere == nil) {
		log(debugError, "move() ERROR: toolkit struct == nil. for", a.WhereId)
		return
	}

	switch tWhere.Type {
	case toolkit.Group:
		switch a.Type {
		case toolkit.Margin:
			tWhere.uiGroup.SetMargined(true)
		}
	case toolkit.Tab:
		switch a.Type {
		case toolkit.Margin:
			// tabSetMargined(tWhere.uiTab, true)
		}
	case toolkit.Window:
		switch a.Type {
		case toolkit.Pad:
			// t.uiWindow.SetBorderless(false)
		}
	case toolkit.Grid:
		switch a.Type {
		case toolkit.Pad:
			// t.uiGrid.SetPadded(true)
		}
	case toolkit.Box:
		log(debugNow, "TODO: move() for a =", a.Type)
		log(debugNow, "TODO: move() where =", a.WhereId)
		log(debugNow, "TODO: move() for widget =", a.WidgetId)

		stretchy = true
		tWhere.uiBox.Append(tWidget.uiControl, stretchy)
		// log(debugNow, "is there a tWhere parent? =", tWhere.parent)
		// tWhere.uiBox.Delete(0)

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
	default:
		log(debugError, "TODO: need to implement move() for a =", a.Type)
		log(debugError, "TODO: need to implement move() for where =", a.WhereId)
		log(debugError, "TODO: need to implement move() for widget =", a.WidgetId)
	}
}

func uiDelete(a *toolkit.Action) {
	if (andlabs[a.WhereId] == nil) {
		log(debugError, "uiDelete() ERROR: can not uiDelete to nil")
		return
	}
	if (andlabs[a.WidgetId] == nil) {
		log(debugError, "uiDelete() ERROR: can not uiDelete nil")
		return
	}
	log(debugNow, "uiDelete()", a.WidgetId, "to", a.WhereId)

	tWidget := andlabs[a.WidgetId]
	if (tWidget == nil) {
		log(debugError, "uiDelete() ERROR: toolkit struct == nil. for", a.WidgetId)
		return
	}

	tWhere := andlabs[a.WhereId]
	if (tWhere == nil) {
		log(debugError, "uiDelete() ERROR: toolkit struct == nil. for", a.WhereId)
		return
	}

	switch tWhere.Type {
	case toolkit.Group:
		switch a.Type {
		case toolkit.Margin:
			tWhere.uiGroup.SetMargined(true)
		}
	case toolkit.Tab:
		switch a.Type {
		case toolkit.Margin:
			// tabSetMargined(tWhere.uiTab, true)
		}
	case toolkit.Window:
		switch a.Type {
		case toolkit.Pad:
			// t.uiWindow.SetBorderless(false)
		}
	case toolkit.Grid:
		switch a.Type {
		case toolkit.Pad:
			// t.uiGrid.SetPadded(true)
		}
	case toolkit.Box:
		log(debugNow, "tWidget.boxC =", tWhere.Name)
		log(debugNow, "is there a tWhere parent? =", tWhere.parent)
		if (tWidget.boxC < 1) {
			log(debugNow, "Can not delete from Box. already empty. tWidget.boxC =", tWhere.boxC)
			return
		}
		tWidget.uiBox.Delete(0)
		tWidget.boxC -= 1

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
		// tWhere.uiBox.Append(tWidget.uiControl, stretchy)
	default:
		log(debugError, "TODO: need to implement uiDelete() for a =", a.Type)
		log(debugError, "TODO: need to implement uiDelete() for where =", a.WhereId)
		log(debugError, "TODO: need to implement uiDelete() for widget =", a.WidgetId)
	}
}
