package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newSlider(n *node) {
	newt := new(guiWidget)

	s := ui.NewSlider(n.X, n.Y)
	newt.uiSlider = s
	newt.uiControl = s

	s.OnChanged(func(spin *ui.Slider) {
		n.I = newt.uiSlider.Value()
		n.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}
