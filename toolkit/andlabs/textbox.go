package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// func newTextbox(a *toolkit.Action) {
func (t *andlabsT) newTextbox() *andlabsT {
	var newt andlabsT

	c := ui.NewNonWrappingMultilineEntry()
	newt.uiMultilineEntry = c
	newt.uiControl = c

	newt.WidgetType = toolkit.Textbox

	c.OnChanged(func(spin *ui.MultilineEntry) {
		newt.s = spin.Text()
		// this is still dangerous
		log(debugChange, "Not yet safe to trigger on change for ui.MultilineEntry")
		newt.s = spin.Text()
		newt.doUserEvent()
	})
	return &newt
}

func newTextbox(a *toolkit.Action) {
	log(debugToolkit, "newCombobox()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		log(debugToolkit, "newCombobox() toolkit struct == nil. name=", a.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newTextbox()
	newt.Name = a.Name
	newt.wId = a.WidgetId
	place(a, t, newt)
}
