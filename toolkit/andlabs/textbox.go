package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// func newTextbox(a *toolkit.Action) {
func (t *andlabsT) newTextbox(w *toolkit.Widget) *andlabsT {
	var newt andlabsT

	c := ui.NewNonWrappingMultilineEntry()
	newt.uiMultilineEntry = c
	newt.uiControl = c

	newt.Name = w.Name
	newt.tw = w

	c.OnChanged(func(spin *ui.MultilineEntry) {
		t.s = spin.Text()
		// this is still dangerous
		// newt.commonChange(newt.tw)
		log(debugChange, "Not yet safe to trigger on change for ui.MultilineEntry")
	})
	return &newt
}

func newTextbox(a *toolkit.Action) {
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "newCombobox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "newCombobox() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newTextbox(w)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
