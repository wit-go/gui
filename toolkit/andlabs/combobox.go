package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newCombobox(w *toolkit.Widget) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "newCombobox() START", w.Name)

	newt.tw = w
	s := ui.NewEditableCombobox()
	newt.uiEditableCombobox = s
	newt.uiControl = s

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	s.OnChanged(func(spin *ui.EditableCombobox) {
		newt.tw.S = spin.Text()
		newt.commonChange(newt.tw)
	})

	return &newt
}

func (t *andlabsT) AddComboboxName(title string) {
	t.uiEditableCombobox.Append(title)
	if (t.val == nil) {
		log(debugToolkit, "make map didn't work")
		return
	}
	t.val[t.c] = title

	// If this is the first menu added, set the dropdown to it
	// if (t.c == 0) {
	// }
	t.c = t.c + 1
}

func newCombobox(a *toolkit.Action) {
	w := a.Widget
	parentW := a.Where
	log(debugToolkit, "newCombobox()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "newCombobox() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newCombobox(w)
	place(a, t, newt)
	mapWidgetsToolkits(a, newt)
}
