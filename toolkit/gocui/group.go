package main

import (
	"log"

	"git.wit.org/wit/gui/toolkit"
)

func NewGroup(parentW *toolkit.Widget, w *toolkit.Widget) {
	if (parentW == nil) {
		log.Println("wit/gui plugin error. parent widget == nil")
		return
	}
	if (w == nil) {
		log.Println("wit/gui plugin error. widget == nil")
		return
	}
	if (w.Name == "") {
		w.Name = parentW.Name
	}
	if (w.Name == "") {
		w.Name = "nil newGroup"
	}
	log.Println("gui.gocui.AddGroup", w.Name)
	addGroup(w.Name)
	stringWidget[w.Name] = w
}

func addGroup(name string) {
	log.Println("addGroup() START name =", name)
	log.Println("addGroup() START groupSize =", groupSize, "currentY =", currentY, "currentX =", currentX)

	currentY = 2
	currentX += groupSize + 5
	groupSize = 0

	log.Println("addGroup() START, RESET Y = 3, RESET X = ", currentX)
}
