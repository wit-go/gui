package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newButton(a *toolkit.Action) {
	var t, newt *andlabsT
	var b *ui.Button
	w := a.Widget
	log(debugToolkit, "newButton()", w.Name)

	t = mapToolkits[a.Where]
	if (t == nil) {
		log(debugToolkit, "newButton() toolkit struct == nil. name=", a.Where.Name, w.Name)
		return
	}

	newt = new(andlabsT)

	b = ui.NewButton(w.Name)
	newt.uiButton = b
	newt.uiControl = b
	newt.tw = w
	newt.parent = t

	b.OnClicked(func(*ui.Button) {
		newt.commonChange(newt.tw)
	})

	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
