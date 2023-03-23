package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

/*
	This adds a tab

	andlabs/ui is goofy in the sense that you have to determine
	if the ui.Window already has a tab in it. If it does, then
	you need to add this tab and not run SetChild() on the window
	or instead it replaces the existing tab with the new one

	I work around this by always sending a Toolkit that is a tab
	once there is one. If you send a Window here, it will replace
	any existing tabs rather than adding a new one
*/
func (t *andlabsT) newTab(name string) *andlabsT {
	// var w *ui.Window
	var newt *andlabsT

	log(debugToolkit, "gui.toolkit.AddTab()")

	if (t.uiWindow == nil) {
		log(debugToolkit, "gui.Toolkit.UiWindow == nil. I can't add a toolbar without window")
		return nil
	}

	if (t.uiTab == nil) {
		// this means you have to make a new tab
		log(debugToolkit, "gui.toolkit.NewTab() GOOD. This should be the first tab:", name)
		newt = rawTab(t.uiWindow, name)
		t.uiTab = newt.uiTab
	} else {
		// this means you have to append a tab
		log(debugToolkit, "gui.toolkit.NewTab() GOOD. This should be an additional tab:", name)
		newt = t.appendTab(name)
	}

	newt.Name = name

	log(debugToolkit, "t:")
	t.Dump(debugToolkit)
	log(debugToolkit, "newt:")
	newt.Dump(debugToolkit)

	return newt
}

// This sets _all_ the tabs to Margin = true
//
// TODO: do proper tab tracking (will be complicated). low priority
func tabSetMargined(tab *ui.Tab, b bool) {
	c := tab.NumPages()
	for i := 0; i < c; i++ {
		log(debugToolkit, "SetMargined", i, b)
		tab.SetMargined(i, b)
	}
}

func rawTab(w *ui.Window, name string) *andlabsT {
	var newt andlabsT
	log(debugToolkit, "gui.toolkit.NewTab() ADD", name)

	if (w == nil) {
		log(debugToolkit, "gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		log(debugToolkit, "gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		log(debugToolkit, "gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		sleep(1)
		return nil
	}
	log(debugToolkit, "gui.toolkit.AddTab() START name =", name)
	tab := ui.NewTab()
	w.SetMargined(margin)

	hbox := ui.NewHorizontalBox() // this makes everything go along the horizon
	hbox.SetPadded(padded)
	tab.Append(name, hbox)
	tabSetMargined(tab, margin) // TODO: run this in the right place(?)
	w.SetChild(tab)

	newt.uiWindow = w
	newt.uiTab = tab
	newt.uiControl = tab
	newt.uiBox = hbox
	return &newt
}

func (t *andlabsT) appendTab(name string) *andlabsT {
	var newT andlabsT
	log(debugToolkit, "gui.toolkit.NewTab() ADD", name)

	if (t.uiTab == nil) {
		log(debugToolkit, "gui.Toolkit.UiWindow == nil. I can't add a widget without a place to put it")
		panic("should never have happened. wit/gui/toolkit has ui.Tab == nil")
	}
	log(debugToolkit, "gui.toolkit.AddTab() START name =", name)

	var hbox *ui.Box
	if (defaultBehavior) {
		hbox = ui.NewHorizontalBox()
	} else {
		if (bookshelf) {
			hbox = ui.NewHorizontalBox()
		} else {
			hbox = ui.NewVerticalBox()
		}
	}
	hbox.SetPadded(padded)
	t.uiTab.Append(name, hbox)

	newT.uiWindow = t.uiWindow
	newT.uiTab = t.uiTab
	newT.uiBox = hbox
	return &newT
}

func newTab(a *toolkit.Action) {
	parentW := a.Where
	w := a.Widget
	var newt *andlabsT
	log(debugToolkit, "gui.andlabs.NewTab()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewTab() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.newTab(w.Name)
	mapWidgetsToolkits(a, newt)
}

func doTab(a *toolkit.Action) {
	newTab(a)
}
