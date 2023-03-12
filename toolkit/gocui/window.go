package main

import (
	"git.wit.org/wit/gui/toolkit"
)

func NewWindow(w *toolkit.Widget) {
	if (w == nil) {
		log("wit/gui plugin error. widget == nil")
		return
	}
	if (w.Name == "") {
		w.Name = "nil newWindow"
	}
	log("gui.gocui.AddWindow", w.Name)
}
