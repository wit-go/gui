package main

import (
	// "git.wit.org/wit/gui/toolkit"
)

// This is a map between the widgets in wit/gui and the internal structures of gocui

// var mapWidgets map[*andlabsT]*toolkit.Widget
// var mapToolkits map[*toolkit.Widget]*andlabsT

// This lists out the known mappings
// deprecate and use instead the GUI interface
func listMap(b bool) {
	log(b, "listMap() disabled HERE. output too big")
	return
	log(b, "listMap() HERE mapToolkits()")
	for i, t := range andlabs {
		log(b, "andlabs =", t, "widgetId =", i)
		t.Dump(b)
	}
	log(b, "listMap() HERE mapWidgets()")
	log(b, "listMap() HERE FIXME. output too big")
}
