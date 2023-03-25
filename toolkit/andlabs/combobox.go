package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newCombobox(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	w := a.Widget
	log(debugToolkit, "newCombobox() START", w.Name)

	newt.tw = w
	newt.wId = a.WidgetId
	newt.Type = w.Type
	s := ui.NewEditableCombobox()
	newt.uiEditableCombobox = s
	newt.uiControl = s

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	s.OnChanged(func(spin *ui.EditableCombobox) {
		newt.tw.S = spin.Text()
		newt.commonChange(newt.tw, a.WidgetId)
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
	log(debugToolkit, "newCombobox()", a.Title)

	t := andlabs[a.WhereId]
	if (t == nil) {
		log(debugToolkit, "newCombobox() toolkit struct == nil. name=", a.Title)
		listMap(debugToolkit)
		return
	}
	newt := t.newCombobox(a)
	place(a, t, newt)
}
