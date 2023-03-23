package main

import (
	"git.wit.org/wit/gui/toolkit"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newCheckbox(w *toolkit.Widget) *andlabsT {
	log(debugToolkit, "newCheckbox()", w.Name, w.Type)
	var newt andlabsT
	newt.tw = w

	newt.uiCheckbox = ui.NewCheckbox(w.Name)
	newt.uiControl = newt.uiCheckbox

	newt.uiCheckbox.OnToggled(func(spin *ui.Checkbox) {
		newt.tw.B = newt.checked()
		log(debugChange, "val =", newt.tw.B)
		newt.commonChange(newt.tw)
	})

	return &newt
}

func (t *andlabsT) checked() bool {
	return t.uiCheckbox.Checked()
}

func newCheckbox(a *toolkit.Action) {
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "newCheckbox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		listMap(debugError)
		return
	}
	newt := t.newCheckbox(w)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
