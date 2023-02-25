package main

import (
	"git.wit.org/wit/gui/toolkit"
)

// This is a map between the widgets in wit/gui and the internal structures of gocui

var mapWidgets map[*andlabsT]*toolkit.Widget
var mapToolkits map[*toolkit.Widget]*andlabsT

// This lists out the know mappings
func listMap() {
	log(debugToolkit, "listMap() HERE")
	log(debugToolkit, "listMap() HERE")
	log(debugToolkit, "listMap() HERE mapWidgets()")
	for t, w := range mapWidgets {
		log(debugToolkit, "andlabs =", t.Name, "widget =", w.Name)
	}
	log(debugToolkit, "listMap() HERE mapToolkits()")
	for w, t := range mapToolkits {
		log(debugToolkit, "andlabs =", t, "widget =", w.Name)
		forceDump(t)
	}
}

func mapWidgetsToolkits(w *toolkit.Widget, t *andlabsT) {
	if (mapToolkits[w] == nil) {
		mapToolkits[w] = t
	} else {
		log(debugToolkit, "WTF: mapToolkits was sent nil. this should not happen w =", w)
		log(debugToolkit, "WTF: mapToolkits was sent nil. this should not happen t =", t.Width)
		log(debugToolkit, "WTF: mapToolkits map already set to ", mapToolkits[w])
		panic("WTF mapWidgetsToolkits() w == nil")
	}

	if (mapWidgets[t] == nil) {
		mapWidgets[t] = w
	} else {
		log(debugToolkit, "WTF: mapWidgets already installed. w =", w)
		log(debugToolkit, "WTF: mapWidgets already installed. t =", t.Width, t)
		log(SPEW, &t)
		log(SPEW, t)
		log(SPEW, *t)
		log(debugToolkit, "WTF: mapWidgets already mapped to", mapWidgets[t])
		log(SPEW, mapWidgets[t])
		panic("WTF. mapWidget andlabs toolkit already mapped to gui toolkit")
	}
}
