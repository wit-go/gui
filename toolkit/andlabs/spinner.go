package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newSpinner(n *node) {
	newt := new(andlabsT)

	s := ui.NewSpinbox(n.X, n.Y)
	newt.uiSpinbox = s
	newt.uiControl = s

	s.OnChanged(func(s *ui.Spinbox) {
		newt.i = newt.uiSpinbox.Value()
		newt.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}
