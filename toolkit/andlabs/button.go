package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"git.wit.org/wit/gui/toolkit"
)

func NewButton(parentW *toolkit.Widget, w *toolkit.Widget) {
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

	b.OnClicked(func(*ui.Button) {
		log(debugChange, "TODO: SHOULD LEAVE Button click HERE VIA channels. button name =", w.Name)
		log(debugChange, "FOUND WIDGET =", w)
		if (w.Custom == nil) {
			log(debugChange, "WIDGET DOES NOT have Custom()")
			log(debugChange, "TODO: NOTHING TO DO button name =", w.Name)
			return
		}
		// t.Dump()
		// newt.Dump()
		log(debugChange, "Running w.Custom()")
		w.Custom()
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
