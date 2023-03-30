package main

// implements widgets 'Window' and 'Tab'

import (
//	"fmt"
	"strconv"

//	"git.wit.org/wit/gui/toolkit"
//	"github.com/awesome-gocui/gocui"
)

func adjustWidgets() {
	for i := 0; i <= me.highest; i++ {
		w := me.widgets[i]
		if (w == nil) {
			continue
		}
		p := me.widgets[w.parentId]
		if (p != nil) {
			w.setParentLogical(p)
		}
	}
}

func hideWidgets() {
	for i := 0; i <= me.highest; i++ {
		w := me.widgets[i]
		if (w == nil) {
			continue
		}
		if (w.visable) {
			if (w.v != nil) {
				cuiName := strconv.Itoa(i)
				log(logNow, "about to delete", cuiName, w.name)
				me.baseGui.DeleteView(cuiName)
			}
		}
	}
}

func showWidgets() {
	for i := 0; i <= me.highest; i++ {
		w := me.widgets[i]
		if (w == nil) {
			continue
		}
		if (w.visable) {
			w.drawView()
		}
	}
}
