package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newTextbox(n *node) {
	newt := new(andlabsT)

	if (n.X == 1) {
		e := ui.NewEntry()
		newt.uiEntry = e
		newt.uiControl = e

		e.OnChanged(func(spin *ui.Entry) {
			n.S = spin.Text()
			n.doUserEvent()
		})
	} else {
		e := ui.NewNonWrappingMultilineEntry()
		newt.uiMultilineEntry = e
		newt.uiControl = e

		e.OnChanged(func(spin *ui.MultilineEntry) {
			n.S = spin.Text()
			n.doUserEvent()
		})
	}
	n.tk = newt
	p.place(n)
}
