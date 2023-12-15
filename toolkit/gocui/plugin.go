package main

import (
	// if you include more than just this import
	// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
	"git.wit.org/wit/gui/toolkit"
)

func action(a *toolkit.Action) {
	log(logInfo, "action() START", a.WidgetId, a.ActionType, a.WidgetType, a.Name)
	n := me.rootNode.findWidgetId(a.WidgetId)
	var w *guiWidget
	if (n != nil) {
		w = n.tk
	}
	switch a.ActionType {
	case toolkit.Add:
		if (w == nil) {
			n := addNode(a)
			// w = n.tk
			n.addWidget()
		} else {
			// this is done to protect the plugin being 'refreshed' with the
			// widget binary tree. TODO: find a way to keep them in sync
			log(logError, "action() Add ignored for already defined widget",
				a.WidgetId, a.ActionType, a.WidgetType, a.Name)
		}
	case toolkit.Show:
		if (a.B) {
			n.showView()
		} else {
			n.hideWidgets()
		}
	case toolkit.Set:
		if a.WidgetType == toolkit.Flag {
			log(logNow, "TODO: set flag here", a.ActionType, a.WidgetType, a.Name)
			log(logNow, "TODO: n.WidgetType =", n.WidgetType, "n.Name =", a.Name)
		} else {
			if (a.A == nil) {
				log(logError, "TODO: Set here. a == nil", a.ActionType, "WidgetType =", a.WidgetType, "Name =", a.Name)
			} else {
				n.Set(a.A)
			}
		}
	case toolkit.SetText:
		n.SetText(a.S)
	case toolkit.AddText:
		n.AddText(a.S)
	case toolkit.Move:
		log(logNow, "attempt to move() =", a.ActionType, a.WidgetType, a.Name)
	case toolkit.CloseToolkit:
		log(logNow, "attempting to close the plugin and release stdout and stderr")
		standardExit()
	default:
		log(logError, "action() Unknown =", a.ActionType, a.WidgetType, a.Name)
	}
	log(logInfo, "action() END")
}

func (n *node) AddText(text string) {
	if (n == nil) {
		log(logNow, "widget is nil")
		return
	}
	n.vals = append(n.vals, text)
	for i, s := range n.vals {
		log(logNow, "AddText()", n.Name, i, s)
	}
	n.SetText(text)
}

func (n *node) SetText(text string) {
	if (n == nil) {
		log(logNow, "widget is nil")
		return
	}
	n.S = text
	n.Text = text

	n.textResize()
	n.deleteView()
	n.showView()
}

func (n *node) Set(val any) {
	// w := n.tk
	log(logInfo, "Set() value =", val)

	switch v := val.(type) {
	case bool:
		n.B = val.(bool)
		n.setCheckbox(val.(bool))
	case string:
		n.SetText(val.(string))
	case int:
		n.I = val.(int)
	default:
		log(logError, "Set() unknown type =", val, v)
	}
}
