package main

import (
	"git.wit.org/wit/gui/toolkit"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newCheckbox(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	w := a.Widget
	log(debugToolkit, "newCheckbox()", w.Name, w.Type)
	newt.tw = w
	newt.Type = w.Type
	newt.wId = a.WidgetId

	newt.uiCheckbox = ui.NewCheckbox(w.Name)
	newt.uiControl = newt.uiCheckbox

	newt.uiCheckbox.OnToggled(func(spin *ui.Checkbox) {
		newt.tw.B = newt.checked()
		log(debugChange, "val =", newt.tw.B)
		newt.commonChange(newt.tw, a.WidgetId)
	})

	return &newt
}

func (t *andlabsT) checked() bool {
	return t.uiCheckbox.Checked()
}

func newCheckbox(a *toolkit.Action) {
	log(debugToolkit, "newCheckbox()", a.Title)

	t := andlabs[a.WhereId]
	if (t == nil) {
		listMap(debugError)
		return
	}
	newt := t.newCheckbox(a)
	place(a, t, newt)
}
