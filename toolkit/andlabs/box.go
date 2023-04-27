package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// make new Box here
func (p *node) newBox(n *node) {
	log(debugToolkit, "newBox()", n.Name)

	newt := new(andlabsT)
	var box *ui.Box

	log(debugToolkit, "rawBox() create", n.Name)

	if (n.B) {
		box = ui.NewHorizontalBox()
	} else {
		box = ui.NewVerticalBox()
	}
	box.SetPadded(padded)

	newt.uiBox = box
	newt.uiControl = box
	newt.boxC = 0
	n.tk = newt
	p.place(n)
}
