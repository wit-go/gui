package main

import (
//	"git.wit.org/wit/gui/toolkit"

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
func (p *node) newTab(n *node) {
	// var w *ui.Window
	var newt *andlabsT

	t := p.tk

	log(debugToolkit, "newTab() START", n.WidgetId, n.ParentId)

	if (t.uiTab == nil) {
		if (t.uiWindow == nil) {
			log(debugToolkit, "newTab() uiWindow == nil. I can't add a toolbar without window", n.WidgetId, n.ParentId)
			return
		}
		// this means you have to make a new tab
		log(debugToolkit, "newTab() GOOD. This should be the first tab:", n.WidgetId, n.ParentId)
		newt = rawTab(t.uiWindow, n.Text)
		t.uiTab = newt.uiTab
	} else {
		// this means you have to append a tab
		log(debugToolkit, "newTab() GOOD. This should be an additional tab:", n.WidgetId, n.ParentId)
		newt = t.appendTab(n.Text)
	}

	// add the structure to the array
	if (andlabs[n.WidgetId] == nil) {
		log(logInfo, "newTab() MAPPED", n.WidgetId, n.ParentId)
		andlabs[n.WidgetId] = newt
		newt.WidgetType = n.WidgetType
	} else {
		log(debugError, "newTab() DO WHAT?", n.WidgetId, n.ParentId)
		log(debugError, "THIS IS BAD")
	}

	newt.Name = n.Name

	log(debugToolkit, "t:")
	t.Dump(debugToolkit)
	log(debugToolkit, "newt:")
	newt.Dump(debugToolkit)
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
	log(debugToolkit, "rawTab() START", name)

	if (w == nil) {
		log(debugError, "UiWindow == nil. I can't add a tab without a window")
		log(debugError, "UiWindow == nil. I can't add a tab without a window")
		log(debugError, "UiWindow == nil. I can't add a tab without a window")
		sleep(1)
		return nil
	}

	tab := ui.NewTab()
	w.SetChild(tab)
	newt.uiTab = tab
	newt.uiControl = tab
	log(debugToolkit, "rawTab() END", name)
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

func newTab(n *node) {
	log(logInfo, "newTab() add to parent id:", n.ParentId)

	p := n.parent
	p.newTab(n)
}
