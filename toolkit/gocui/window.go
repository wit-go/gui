package main

import (
	"log"

	"git.wit.org/wit/gui/toolkit"
)

func NewWindow(w *toolkit.Widget) {
	if (w == nil) {
		log.Println("wit/gui plugin error. widget == nil")
		return
	}
	if (w.Name == "") {
		w.Name = "nil newWindow"
	}
	log.Println("gui.gocui.AddWindow", w.Name)
}
