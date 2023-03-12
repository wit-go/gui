package main

import (
	"git.wit.org/wit/gui/toolkit"
)

// This is a map between the widgets in wit/gui and the internal structures of gocui

var mapWidgets map[*andlabsT]*toolkit.Widget
var mapToolkits map[*toolkit.Widget]*andlabsT

// This lists out the know mappings
func listMap(b bool) {
	log(b, "listMap() HERE")
	log(b, "listMap() HERE")
	for t, w := range mapWidgets {
		log(b, "andlabs =", t.Name, "widget =", w.Name)
	}
	log(b, "listMap() HERE mapToolkits()")
	for w, t := range mapToolkits {
		log(b, "andlabs =", t, "widget =", w.Name)
		t.Dump(b)
	}
	log(b, "listMap() HERE mapWidgets()")
	log(b, "listMap() HERE FIXME. output too big")
}

func mapWidgetsToolkits(w *toolkit.Widget, t *andlabsT) {
	if ((mapToolkits[w] == nil) && (mapWidgets[t] == nil)) {
		log(debugToolkit, "map a new widget", w, t)
		mapToolkits[w] = t
		mapWidgets[t] = w
		return
	}

	if (mapToolkits[w] != nil) {
		tw := mapToolkits[w]
		if (tw == nil) {
			// logic corruption somewhere? Have you been deleting things recently?
			log(true, "mapToolkits[w] is set, but mapWidgets[t] is nil")
			panic("WTF mapWidgets[t] == nil")
		}
		log(debugToolkit, "mapToolkits[w] is", tw)
		if (tw == nil) {
			log(debugError, "BAD map? mapWidgets[w] tw == nil")
		} else {
			log(debugError, "BAD map? mapWidgets[w] is", tw)
			tw.Dump(debugError)
		}
	}

	if (mapWidgets[t] != nil) {
		wt := mapWidgets[t]
		if (mapToolkits[w] == nil) {
			// logic corruption somewhere? Have you been deleting things recently?
			log(true, "mapWidgets[t] is set, but mapToolkits[w] is nil")
			panic("WTF mapToolkits[w] == nil")
		}
		if (wt == nil) {
			log(debugError, "BAD map? mapWidgets[t] wt == nil")
		} else {
			log(debugError, "BAD map? mapWidgets[t] is", wt)
			widgetDump(debugError, wt)
		}
	}
	log(debugToolkit, "map of widget worked", w.Type, ",", w.Name, ",", w.Action)
}
