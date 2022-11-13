package main

import (
	"log"

	"git.wit.org/wit/gui/toolkit"
)

// This is a map between the widgets in wit/gui and the internal structures of gocui

var mapWidgets map[*andlabsT]*toolkit.Widget
var mapToolkits map[*toolkit.Widget]*andlabsT

// This lists out the know mappings
func listMap() {
	log.Println("listMap() HERE")
	log.Println("listMap() HERE")
	log.Println("listMap() HERE mapWidgets()")
	for t, w := range mapWidgets {
		log.Println("andlabs =", t.Name, "widget =", w.Name)
	}
	log.Println("listMap() HERE mapToolkits()")
	for w, t := range mapToolkits {
		log.Println("andlabs =", t, "widget =", w.Name)
		forceDump(t)
	}
}

func mapWidgetsToolkits(w *toolkit.Widget, t *andlabsT) {
	if (mapToolkits[w] == nil) {
		mapToolkits[w] = t
	} else {
		log.Println("WTF: mapToolkits already installed")
		panic("WTF")
	}

	if (mapWidgets[t] == nil) {
		mapWidgets[t] = w
	} else {
		log.Println("WTF: mapWidgets already installed")
		panic("WTF")
	}
}
