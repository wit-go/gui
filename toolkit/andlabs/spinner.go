package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newSpinner(n *node) {
	newt := new(guiWidget)

	s := ui.NewSpinbox(n.X, n.Y)
	newt.uiSpinbox = s
	newt.uiControl = s

	s.OnChanged(func(s *ui.Spinbox) {
		n.I = newt.uiSpinbox.Value()
		n.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}
