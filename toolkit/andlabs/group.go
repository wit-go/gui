package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newGroup(n *node) {
	log(debugToolkit, "NewGroup()", n.Name)

	t := p.tk
	if (t == nil) {
		log(debugToolkit, "NewGroup() toolkit struct == nil. name=", n.Name)
		listMap(debugToolkit)
	}
	newt := t.rawGroup(n.Name)
	n.tk = newt
	p.place(n)
}

// make new Group here
func (t *andlabsT) rawGroup(title string) *andlabsT {
	var newt andlabsT
	newt.Name = title

	log(debugToolkit, "NewGroup() create", newt.Name)

	g := ui.NewGroup(newt.Name)
	g.SetMargined(margin)
	newt.uiGroup = g
	newt.uiControl = g

	return &newt
}
