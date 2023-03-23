package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func show(w *toolkit.Widget) {
	if (w == nil) {
		log(debugError, "nil is probably already hidden")
		return
	}
	log(debugError, "show()", w.Name)

	t := mapToolkits[w]
	if (t == nil) {
		log(debugToolkit, "show() toolkit struct == nil. for", w.Name)
		return
	}

	if (w.B) {
		t.uiControl.Show()
	} else {
		t.uiControl.Hide()
	}
}

func enable(w *toolkit.Widget) {
	if (w == nil) {
		log(debugError, "nil is probably already hidden")
		return
	}
	log(debugError, "enable()", w.Name)

	t := mapToolkits[w]
	if (t == nil) {
		log(debugToolkit, "enable() toolkit struct == nil. for", w.Name)
		return
	}

	if (w.B) {
		t.uiControl.Enable()
	} else {
		t.uiControl.Disable()
	}
}

func pad(a *toolkit.Action) {
	if (a.Widget == nil) {
		log(debugError, "pad() ERROR: nil is probably already hidden")
		return
	}
	log(debugError, "pad()", a.Widget.Name)

	t := mapToolkits[a.Widget]
	if (t == nil) {
		log(debugToolkit, "pad() toolkit struct == nil. for", a.Widget.Name)
		return
	}

	switch a.Widget.Type {
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
	if (a.Where == nil) {
		log(debugError, "move() ERROR: can not move to nil")
		return
	}
	if (a.Widget == nil) {
		log(debugError, "move() ERROR: can not move nil")
		return
	}
	log(debugNow, "move()", a.Widget.Name, "to", a.Where.Name)

	tWidget := mapToolkits[a.Widget]
	if (tWidget == nil) {
		log(debugError, "move() ERROR: toolkit struct == nil. for", a.Widget.Name)
		return
	}

	tWhere := mapToolkits[a.Where]
	if (tWhere == nil) {
		log(debugError, "move() ERROR: toolkit struct == nil. for", a.Where.Name)
		return
	}

	switch a.Where.Type {
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
		log(debugNow, "TODO: move() where =", a.Where.Type)
		log(debugNow, "TODO: move() for widget =", a.Widget.Type)

		stretchy = true
		tWhere.uiBox.Append(tWidget.uiControl, stretchy)
		// log(debugNow, "is there a tWhere parent? =", tWhere.parent)
		// tWhere.uiBox.Delete(0)

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
	default:
		log(debugError, "TODO: need to implement move() for a =", a.Type)
		log(debugError, "TODO: need to implement move() for where =", a.Where.Type)
		log(debugError, "TODO: need to implement move() for widget =", a.Widget.Type)
	}
}

func uiDelete(a *toolkit.Action) {
	if (a.Where == nil) {
		log(debugError, "uiDelete() ERROR: can not uiDelete to nil")
		return
	}
	if (a.Widget == nil) {
		log(debugError, "uiDelete() ERROR: can not uiDelete nil")
		return
	}
	log(debugNow, "uiDelete()", a.Widget.Name, "to", a.Where.Name)

	tWidget := mapToolkits[a.Widget]
	if (tWidget == nil) {
		log(debugError, "uiDelete() ERROR: toolkit struct == nil. for", a.Widget.Name)
		return
	}

	tWhere := mapToolkits[a.Where]
	if (tWhere == nil) {
		log(debugError, "uiDelete() ERROR: toolkit struct == nil. for", a.Where.Name)
		return
	}

	switch a.Where.Type {
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
		log(debugError, "TODO: need to implement uiDelete() for where =", a.Where.Type)
		log(debugError, "TODO: need to implement uiDelete() for widget =", a.Widget.Type)
	}
}
