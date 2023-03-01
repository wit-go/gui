package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newButton(parentW *toolkit.Widget, w *toolkit.Widget) {
	var t, newt *andlabsT
	var b *ui.Button
	log(debugToolkit, "gui.andlabs.NewButton()", w.Name)

	t = mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewButton() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}

	if t.broken() {
		return
	}
	newt = new(andlabsT)

	b = ui.NewButton(w.Name)
	newt.uiButton = b
	newt.tw = w
	newt.parent = t

	b.OnClicked(func(*ui.Button) {
		newt.commonChange(newt.tw)
	})

	log(debugToolkit, "gui.Toolbox.NewButton() about to append to Box parent t:", w.Name)
	log(debugToolkit, "gui.Toolbox.NewButton() about to append to Box new t:", w.Name)

	if (t.uiBox != nil) {
		t.uiBox.Append(b, stretchy)
	} else if (t.uiWindow != nil) {
		t.uiWindow.SetChild(b)
	} else {
		log(debugToolkit, "ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		log(debugToolkit, "ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		return
	}

	mapWidgetsToolkits(w, newt)
}

func doButton(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newButton(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(true, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(true, "Button() ct.broken", ct)
		return
	}
	if (ct.uiButton == nil) {
		log(true, "Button() uiButton == nil", ct)
		return
	}
	log(true, "Going to attempt:", c.Action)
	switch c.Action {
	case "Enable":
		ct.uiButton.Enable()
	case "Disable":
		ct.uiButton.Disable()
	case "Show":
		ct.uiButton.Show()
	case "Hide":
		ct.uiButton.Hide()
	case "Set":
		ct.uiButton.SetText(c.S)
	default:
		log(true, "Can't do", c.Action, "to a Button")
	}
}
