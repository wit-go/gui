package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newCombobox(n *node) {
	newt := new(andlabsT)
	log(debugToolkit, "newCombobox() START", n.Name)

	cb := ui.NewEditableCombobox()
	newt.uiEditableCombobox = cb
	newt.uiControl = cb

	// initialize the index
	newt.c = 0
	newt.val = make(map[int]string)

	cb.OnChanged(func(spin *ui.EditableCombobox) {
		newt.s = spin.Text()
		newt.doUserEvent()
	})

	n.tk = newt
	p.place(n)
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
