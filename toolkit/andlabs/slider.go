package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newSlider(n *node) {
	newt := new(andlabsT)

	s := ui.NewSlider(n.X, n.Y)
	newt.uiSlider = s
	newt.uiControl = s

	s.OnChanged(func(spin *ui.Slider) {
		newt.i = newt.uiSlider.Value()
		newt.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}
