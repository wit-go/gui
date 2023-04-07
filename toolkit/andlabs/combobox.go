package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"git.wit.org/wit/gui/toolkit"
)

func (t *andlabsT) newCombobox(a *toolkit.Action) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "newCombobox() START", a.Name)

	newt.wId = a.WidgetId
	newt.WidgetType = a.WidgetType
	s := ui.NewEditableCombobox()
	newt.uiEditableCombobox = s
	newt.uiControl = s

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	s.OnChanged(func(spin *ui.EditableCombobox) {
		newt.s = spin.Text()
		newt.doUserEvent()
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
	log(debugToolkit, "newCombobox()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		log(debugToolkit, "newCombobox() toolkit struct == nil. name=", a.Name)
		listMap(debugToolkit)
		return
	}
	newt := t.newCombobox(a)
	place(a, t, newt)
}
