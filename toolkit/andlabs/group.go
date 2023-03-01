package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func newGroup(parentW *toolkit.Widget, w *toolkit.Widget) {
	log(debugToolkit, "gui.andlabs.NewGroup()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewGroup() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
	}
	newt := t.NewGroup(w.Name)
	mapWidgetsToolkits(w, newt)
}

// make new Group here
func (t andlabsT) NewGroup(title string) *andlabsT {
	var newt andlabsT

	log(debugToolkit, "NewGroup() create", title)

	g := ui.NewGroup(title)
	g.SetMargined(margin)

	if (t.uiBox != nil) {
		// TODO: temporary hack to make the output textbox 'fullscreen'
		if (title == "output") {
			t.uiBox.Append(g, true)
		} else {
			t.uiBox.Append(g, stretchy)
		}
	} else if (t.uiWindow != nil) {
		t.uiWindow.SetChild(g)
	} else {
		log(debugError, "NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugError, "probably could just make a box here?")
		exit("internal wit/gui error")
	}

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(padded)
	g.SetChild(hbox)

	newt.uiGroup = g
	newt.uiBox = hbox
	newt.uiWindow = t.uiWindow
	newt.Name = title

	return &newt
}
