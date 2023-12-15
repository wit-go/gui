package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newGroup(n *node) {
	log(debugToolkit, "NewGroup()", n.Name)

	newt := new(guiWidget)

	log(debugToolkit, "NewGroup() create", n.Name)

	g := ui.NewGroup(n.Name)
	g.SetMargined(margin)
	newt.uiGroup = g
	newt.uiControl = g

	n.tk = newt
	p.place(n)
}
