package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newCheckbox(n *node) {
	newt := new(andlabsT)
	log(debugToolkit, "newCheckbox()", n.Name, n.WidgetType)
	newt.WidgetType = n.WidgetType
	newt.wId = n.WidgetId
	newt.Name = n.Name
	newt.Text = n.Text

	newt.uiCheckbox = ui.NewCheckbox(n.Text)
	newt.uiControl = newt.uiCheckbox

	newt.uiCheckbox.OnToggled(func(spin *ui.Checkbox) {
		newt.b = newt.checked()
		log(debugChange, "val =", newt.b)
		newt.doUserEvent()
	})

	n.tk = newt
	p.place(n)
}

func (t *andlabsT) checked() bool {
	return t.uiCheckbox.Checked()
}

/*
func newCheckbox(a *toolkit.Action) {
	log(debugToolkit, "newCheckbox()", a.Name)

	t := andlabs[a.ParentId]
	if (t == nil) {
		listMap(debugError)
		return
	}
	newt := t.newCheckbox(a)
}
*/
