package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newButton(a *toolkit.Action) {
	var t, newt *andlabsT
	var b *ui.Button
	log(debugToolkit, "newButton()", a.Name)

	t = andlabs[a.ParentId]
	if (t == nil) {
		log(debugToolkit, "newButton() toolkit struct == nil. name=", a.Name)
		return
	}

	newt = new(andlabsT)

	b = ui.NewButton(a.Text)
	newt.uiButton = b
	newt.uiControl = b
	newt.tw = a.Widget
	newt.Type = a.WidgetType
	newt.parent = t

	b.OnClicked(func(*ui.Button) {
		newt.commonChange(newt.tw, a.WidgetId)
	})

	place(a, t, newt)
	// mapWidgetsToolkits(a, newt)
}
