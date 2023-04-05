package main

import (
	// if you include more than just this import
	// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
	"git.wit.org/wit/gui/toolkit"
)

func Quit() {
	me.baseGui.Close()
}

// set the widget start width & height
// re-run this when things change to recalibrate the position of the gocui view rect
func (w *cuiWidget) setStartWH() {
	log(logInfo, "setStartWH() w.id =", w.id, "w.name", w.name)
	switch w.widgetType {
	case toolkit.Root:
		log(logInfo, "setStartWH() rootNode w.id =", w.id, "w.name", w.name)
		w.startW = 1
		w.startH = 1
		w.id = 0
		w.isFake = true
		w.setFake()
		w.showWidgetPlacement(logNow, "StartWH:")
		return
	case toolkit.Flag:
		w.startW = 1
		w.startH = 1
		w.isFake = true
		w.setFake()
		w.showWidgetPlacement(logNow, "StartWH:")
		return
	case toolkit.Window:
		w.startW = 1
		w.startH = 3
		w.setTabWH()
		w.showWidgetPlacement(logNow, "StartWH:")
		return
	case toolkit.Tab:
		w.startW = 1
		w.startH = 3
		w.setTabWH()
		w.showWidgetPlacement(logNow, "StartWH:")
		return
	}
	p := w.parent
	switch p.widgetType {
	case toolkit.Box:
		w.getBoxWH()
		return
	case toolkit.Grid:
		w.getGridWH()
		return
	case toolkit.Group:
		w.getGroupWH()
		return
	}
}

func Action(a *toolkit.Action) {
	log(logInfo, "Action() START", a.WidgetId, a.ActionType, a.WidgetType, a.Name)
	w := findWidget(a.WidgetId, me.rootNode)
	switch a.ActionType {
	case toolkit.Add:
		w = setupWidget(a)
		w.setStartWH()
		w.moveTo(w.startW, w.startH)
		w.drawView()

		// findPlace(w)
	case toolkit.Show:
		if (a.B) {
			w.drawView()
		} else {
			w.hideWidgets()
		}
	case toolkit.Set:
		w.Set(a.A)
	case toolkit.SetText:
		w.SetText(a.S)
	case toolkit.AddText:
		w.AddText(a.S)
	case toolkit.Move:
		log(logNow, "attempt to move() =", a.ActionType, a.WidgetType, a.Name)
	default:
		log(logError, "Action() Unknown =", a.ActionType, a.WidgetType, a.Name)
	}
	log(logInfo, "Action() END")
}

func (w *cuiWidget) AddText(text string) {
	if (w == nil) {
		log(logNow, "widget is nil")
		return
	}
	w.vals = append(w.vals, text)
	for i, s := range w.vals {
		log(logNow, "AddText()", w.name, i, s)
	}
	w.SetText(text)
}

func (w *cuiWidget) SetText(text string) {
	if (w == nil) {
		log(logNow, "widget is nil")
		return
	}
	w.text = text
	w.s = text
	w.textResize()
	w.deleteView()
	w.drawView()
}

func (w *cuiWidget) Set(val any) {
	log(logInfo, "Set() value =", val)
	var a toolkit.Action
	a.ActionType = toolkit.Set

	switch v := val.(type) {
	case bool:
		w.b = val.(bool)
	case string:
		w.SetText(val.(string))
	case int:
		w.i = val.(int)
	default:
		log(logError, "Set() unknown type =", v, "a =", a)
	}
}
