package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newButton(a *toolkit.Action) {
	var t, newt *andlabsT
	var b *ui.Button
	log(debugToolkit, "newButton()", a.Title)

	t = andlabs[a.WhereId]
	if (t == nil) {
		log(debugToolkit, "newButton() toolkit struct == nil. name=", a.Title)
		return
	}

	newt = new(andlabsT)

	b = ui.NewButton(a.Title)
	newt.uiButton = b
	newt.uiControl = b
	newt.tw = a.Widget
	newt.Type = a.WidgetT
	newt.parent = t

	b.OnClicked(func(*ui.Button) {
		newt.commonChange(newt.tw, a.WidgetId)
	})

	place(a, t, newt)
	// mapWidgetsToolkits(a, newt)
}
