package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// make new Box here
func (p *node) newBox(n *node) {
	log(debugToolkit, "newBox()", n.Name)

	t := p.tk
	if (t == nil) {
		log(debugToolkit, "newBox() toolkit struct == nil. name=", n.Name)
		listMap(debugToolkit)
	}
	newt := t.rawBox(n.Text, n.B)
	newt.boxC = 0
	n.tk = newt
	p.place(n)
}

// make new Box using andlabs/ui
func (t *andlabsT) rawBox(title string, b bool) *andlabsT {
	var newt andlabsT
	var box *ui.Box
	newt.Name = title

	log(debugToolkit, "rawBox() create", newt.Name)

	if (b) {
		box = ui.NewHorizontalBox()
	} else {
		box = ui.NewVerticalBox()
	}
	box.SetPadded(padded)

	newt.uiBox = box
	newt.uiControl = box

	return &newt
}
