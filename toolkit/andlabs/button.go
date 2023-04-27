package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newButton(n *node) {
	log(debugToolkit, "newButton()", n.Name)

	t := p.tk
	if (t == nil) {
		log(debugToolkit, "newButton() toolkit struct == nil. name=", n.Name)
		return
	}

	newt := new(andlabsT)

	b := ui.NewButton(n.Text)
	newt.uiButton = b
	newt.uiControl = b
	newt.wId = n.WidgetId
	newt.WidgetType = n.WidgetType
	newt.parent = t

	b.OnClicked(func(*ui.Button) {
		newt.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}
