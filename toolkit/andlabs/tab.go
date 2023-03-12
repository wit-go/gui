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
func tabSetMargined(tab *ui.Tab) {
	c := tab.NumPages()
	for i := 0; i < c; i++ {
		log(debugToolkit, "SetMargined", i, margin)
		tab.SetMargined(i, margin)
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
	tabSetMargined(tab) // TODO: run this in the right place(?)
	w.SetChild(tab)

	newt.uiWindow = w
	newt.uiTab = tab
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

func newTab(parentW *toolkit.Widget, w *toolkit.Widget) {
	var newt *andlabsT
	log(debugToolkit, "gui.andlabs.NewTab()", w.Name)

	t := mapToolkits[parentW]
	if (t == nil) {
		log(debugToolkit, "go.andlabs.NewTab() toolkit struct == nil. name=", parentW.Name, w.Name)
		return
	}
	newt = t.newTab(w.Name)
	mapWidgetsToolkits(w, newt)
}

func doTab(p *toolkit.Widget, c *toolkit.Widget) {
	if broken(c) {
		return
	}
	if (c.Action == "New") {
		newTab(p, c)
		return
	}
	ct := mapToolkits[c]
	if (ct == nil) {
		log(debugError, "Trying to do something on a widget that doesn't work or doesn't exist or something", c)
		return
	}
	if ct.broken() {
		log(debugError, "Tab() ct.broken", ct)
		return
	}
	if (ct.uiTab == nil) {
	
		log(debugError, "Tab() uiTab == nil", ct)
		return
	}
	log(debugChange, "Going to attempt:", c.Action)
	switch c.Action {
	case "Enable":
		ct.uiTab.Enable()
	case "Disable":
		ct.uiTab.Disable()
	case "Show":
		ct.uiTab.Show()
	case "Hide":
		ct.uiTab.Hide()
	case "Get":
		c.I = ct.uiTab.NumPages()
	case "Add":
		log(true, "how do I add a tab here in doTab()?")
		dump(p, c, true)
	case "SetMargin":
		i := ct.uiTab.NumPages()
		log(true, "tab.NumPages() =", i)
		for i > 0 {
			i -= 1
			log(true, "uiTab.SetMargined(true) for i =", i)
			ct.uiTab.SetMargined(i, c.B)
		}
	default:
		log(debugError, "Can't do", c.Action, "to a Tab")
	}
}
