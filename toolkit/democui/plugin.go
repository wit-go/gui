package main

import (
	// if you include more than just this import
	// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
	"git.wit.org/wit/gui/toolkit"
)

func Quit() {
	me.baseGui.Close()
}

func Action(a *toolkit.Action) {
	log(logNow, "Action()", a)
	w := setupWidgetT(a)
	place(w, a)
	log(logInfo, "Action() END")
}
