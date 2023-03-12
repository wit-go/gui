package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func newGroup(parentW *toolkit.Widget, w *toolkit.Widget) {
	if (parentW == nil) {
		log(debugError, "plugin error. parent widget == nil")
		return
	}
	if (w == nil) {
		log(debugPlugin, "plugin error. widget == nil")
		return
	}
	if (w.Name == "") {
		w.Name = parentW.Name
	}
	if (w.Name == "") {
		w.Name = "nil newGroup"
	}
	log("AddGroup", w.Name)
	addGroup(w.Name)
	stringWidget[w.Name] = w
}

func addGroup(name string) {
	log("addGroup() START name =", name)
	log("addGroup() START groupSize =", groupSize, "currentY =", currentY, "currentX =", currentX)

	currentY = 2
	currentX += groupSize + 5
	groupSize = 0

	log("addGroup() START, RESET Y = 3, RESET X = ", currentX)
}
