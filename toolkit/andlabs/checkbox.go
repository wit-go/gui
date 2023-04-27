package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newCheckbox(n *node) {
	newt := new(andlabsT)
	log(debugToolkit, "newCheckbox()", n.Name, n.WidgetType)

	newt.uiCheckbox = ui.NewCheckbox(n.Text)
	newt.uiControl = newt.uiCheckbox

	newt.uiCheckbox.OnToggled(func(spin *ui.Checkbox) {
		n.B = newt.checked()
		log(debugChange, "val =", n.B)
		n.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}

func (t *andlabsT) checked() bool {
	return t.uiCheckbox.Checked()
}
