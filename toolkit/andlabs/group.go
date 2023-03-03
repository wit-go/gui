package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func newGroup(parentW *toolkit.Widget, w *toolkit.Widget) {
	// log(debugToolkit, "gui.andlabs.NewGroup()", w.Name)
	log(true, "NewGroup()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "NewGroup() toolkit struct == nil. name=", parentW.Name, w.Name)
		listMap(debugToolkit)
	}
	newt := t.rawGroup(w.Name)
	mapWidgetsToolkits(w, newt)
}

// make new Group here
func (t *andlabsT) rawGroup(title string) *andlabsT {
	var newt andlabsT
	newt.Name = title

	log(debugToolkit, "NewGroup() create", newt.Name)

	g := ui.NewGroup(newt.Name)
	g.SetMargined(margin)
	newt.uiGroup = g

	t.doAppend(&newt, nil)
	/*
	if (t.uiBox != nil) {
		// TODO: temporary hack to make the output textbox 'fullscreen'
		if (newt.Name == "output") {
			t.uiBox.Append(g, true)
		} else {
			t.uiBox.Append(g, stretchy)
		}
	} else if (t.uiWindow != nil) {
		log(true, "This is a raw window without a box. probably make a box here and add the group to that")
		t.uiBox = ui.NewHorizontalBox()
		t.uiWindow.SetChild(t.uiBox)
		log(true, "tried to make a box")
		if (newt.Name == "output") {
			t.uiBox.Append(g, true)
		} else {
			t.uiBox.Append(g, stretchy)
		}
	} else {
		log(debugError, "NewGroup() node.UiBox == nil. I can't add a range UI element without a place to put it")
		log(debugError, "probably could just make a box here?")
		exit("internal wit/gui error")
	}
	*/

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(padded)
	g.SetChild(hbox)

	newt.uiBox = hbox
	newt.uiWindow = t.uiWindow
	newt.uiTab = t.uiTab

	return &newt
}
