package main

import "log"
// import "os"


import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "git.wit.org/wit/gui/toolkit"
func NewButton(parentW *toolkit.Widget, w *toolkit.Widget) {
	var t, newt *andlabsT
	var b *ui.Button
	log.Println("gui.andlabs.NewButton()", w.Name)

	t = mapToolkits[parentW]
	if (t == nil) {
		log.Println("go.andlabs.NewButton() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}

	if t.broken() {
		return
	}
	newt = new(andlabsT)

	b = ui.NewButton(w.Name)
	newt.uiButton = b

	b.OnClicked(func(*ui.Button) {
		if (DebugToolkit) {
			log.Println("TODO: IN TOOLKIT GOROUTINE. SHOULD LEAVE HERE VIA channels. button name =", w.Name)
			log.Println("FOUND WIDGET!", w)
		}
		if (w.Custom != nil) {
			w.Custom()
			return
		}
		if (w.Event != nil) {
			w.Event(w)
			return
		}
		t.Dump()
		newt.Dump()
		if (DebugToolkit) {
			log.Println("TODO: LEFT TOOLKIT GOROUTINE WITH NOTHING TO DO button name =", w.Name)
		}
	})

	if (DebugToolkit) {
		log.Println("gui.Toolbox.NewButton() about to append to Box parent t:", w.Name)
		t.Dump()
		log.Println("gui.Toolbox.NewButton() about to append to Box new t:", w.Name)
		newt.Dump()
	}
	if (t.uiBox != nil) {
		t.uiBox.Append(b, stretchy)
	} else if (t.uiWindow != nil) {
		t.uiWindow.SetChild(b)
	} else {
		log.Println("ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		log.Println("ERROR: wit/gui andlabs couldn't place this button in a box or a window")
		return
	}

	mapWidgetsToolkits(w, newt)
}
