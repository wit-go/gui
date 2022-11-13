package main

import (
	"log"

	"git.wit.org/wit/gui/toolkit"

	"github.com/awesome-gocui/gocui"
)

// This is a map between the widgets in wit/gui and the internal structures of gocui
var viewWidget map[*gocui.View]*toolkit.Widget
var stringWidget map[string]*toolkit.Widget

func Quit() {
	baseGui.Close()
}

// This lists out the know mappings
func listMap() {
	for v, w := range viewWidget {
		log.Println("view =", v.Name, "widget name =", w.Name)
	}
	for s, w := range stringWidget {
		log.Println("string =", s, "widget =", w)
	}
}

