package main

import (
	"os"
	// if you include more than just this import
	// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
	"git.wit.org/wit/gui/toolkit"
)

func Quit() {
	me.baseGui.Close()
}

func action(a *toolkit.Action) {
	log(logInfo, "action() START", a.WidgetId, a.ActionType, a.WidgetType, a.Name)
	w := findWidget(a.WidgetId, me.rootNode)
	switch a.ActionType {
	case toolkit.Add:
		if (w == nil) {
			w = makeWidget(a)
			w.addWidget()
		} else {
			// this is done to protect the plugin being 'refreshed' with the
			// widget binary tree. TODO: find a way to keep them in sync
			log(logError, "action() Add ignored for already defined widget",
				a.WidgetId, a.ActionType, a.WidgetType, a.Name)
		}
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
	case toolkit.CloseToolkit:
		log(logNow, "attempting to close the plugin and release stdout and stderr")
		me.baseGui.Close()
		// defer outf.Close()
		setOutput(os.Stdout)
	default:
		log(logError, "action() Unknown =", a.ActionType, a.WidgetType, a.Name)
	}
	log(logInfo, "action() END")
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

	switch v := val.(type) {
	case bool:
		w.b = val.(bool)
		w.setCheckbox(val.(bool))
	case string:
		w.SetText(val.(string))
	case int:
		w.i = val.(int)
	default:
		log(logError, "Set() unknown type =", val, v)
	}
}
