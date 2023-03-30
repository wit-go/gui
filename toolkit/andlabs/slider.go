package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newSlider(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	w := a.Widget
	// log(debugToolkit, w.Name, w.Type, w.X, w.Y)

	s := ui.NewSlider(a.X, a.Y)
	newt.uiSlider = s
	newt.uiControl = s
	newt.tw = w
	newt.Type = toolkit.Slider
	newt.wId = a.WidgetId

	s.OnChanged(func(spin *ui.Slider) {
		newt.tw.I = newt.uiSlider.Value()
		newt.commonChange(newt.tw, a.WidgetId)
	})

	return &newt
}

func newSlider(a *toolkit.Action) {
	var newt *andlabsT
	w := a.Widget
	log(debugToolkit, "newSlider()", w.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		log(debugError, "newSlider() ERROR toolkit struct == nil. name=", w.Name)
		return
	}
	// w.X = a.X
	// w.Y = a.Y
	newt = t.newSlider(a)
	place(a, t, newt)
}
