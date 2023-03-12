package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func newButton(parentW *toolkit.Widget, w *toolkit.Widget) {
	var t, newt *andlabsT
	var b *ui.Button
	log(debugToolkit, "newButton()", w.Name)

	t = mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "newButton() toolkit struct == nil. name=", parentW.Name, w.Name)
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

	log(debugToolkit, "newButton() about to append to Box parent t:", w.Name)
	log(debugToolkit, "newButton() about to append to Box new t:", w.Name)
	if (debugToolkit) {
		ShowDebug ()
	}

	if (t.uiBox != nil) {
		t.uiBox.Append(b, stretchy)
	} else if (t.uiWindow != nil) {
		t.uiWindow.SetChild(b)
	} else {
		log(debugError, "ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		log(debugError, "ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		return
	}

	mapWidgetsToolkits(w, newt)
}

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

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
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Button() ct.broken", ct)
		return
	}
	if (ct.uiButton == nil) {
		log(debugError, "Button() uiButton == nil", ct)
		return
	}
	log(debugToolkit, "Going to attempt:", c.Action)
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
		log(debugError, "Can't do", c.Action, "to a Button")
	}
}
