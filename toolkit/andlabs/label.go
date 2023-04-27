package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newLabel(n *node) {
	log(logInfo, "NewLabel()", n.Name)

	newt := new(andlabsT)
	c := ui.NewLabel(n.Name)
	newt.uiLabel = c
	newt.uiControl = c

	n.tk = newt
	p.place(n)
}
