package main

import (
	"git.wit.org/wit/gui/toolkit"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (t *andlabsT) newCheckbox(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "newCheckbox()", a.Name, a.WidgetType)
	newt.WidgetType = a.WidgetType
	newt.wId = a.WidgetId
	newt.Name = a.Name
	newt.Text = a.Text

	newt.uiCheckbox = ui.NewCheckbox(a.Text)
	newt.uiControl = newt.uiCheckbox

	newt.uiCheckbox.OnToggled(func(spin *ui.Checkbox) {
		newt.b = newt.checked()
		log(debugChange, "val =", newt.b)
		newt.doUserEvent()
	})

	return &newt
}

func (t *andlabsT) checked() bool {
	return t.uiCheckbox.Checked()
}

func newCheckbox(a *toolkit.Action) {
	log(debugToolkit, "newCheckbox()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		listMap(debugError)
		return
	}
	newt := t.newCheckbox(a)
	place(a, t, newt)
}
