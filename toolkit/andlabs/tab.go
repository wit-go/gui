package toolkit

import (
	"log"
	"time"

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
func (t *Toolkit) AddTab(name string) *Toolkit {
	// var w *ui.Window
	var newt *Toolkit

	log.Println("gui.toolkit.AddTab() sleep 3")

	if (t.uiWindow == nil) {
		log.Println("gui.Toolkit.UiWindow == nil. I can't add a toolbar without window")
		return nil
	}

	if (t.uiTab == nil) {
		// this means you have to make a new tab
		log.Println("gui.toolkit.NewTab() GOOD. This should be the first tab:", name)
		newt = newTab(t.uiWindow, name)
		t.uiTab = newt.uiTab
	} else {
		// this means you have to append a tab
		log.Println("gui.toolkit.NewTab() GOOD. This should be an additional tab:", name)
		newt = t.appendTab(name)
	}

	newt.Name = name

	if (DebugToolkit) {
		log.Println("t:")
		t.Dump()
		log.Println("newt:")
		newt.Dump()
	}

	return newt
}

// This sets _all_ the tabs to Margin = true
//
// TODO: do proper tab tracking (will be complicated). low priority
func tabSetMargined(tab *ui.Tab) {
	c := tab.NumPages()
	for i := 0; i < c; i++ {
		log.Println("SetMargined", i, margin)
		tab.SetMargined(i, margin)
	}
}

func newTab(w *ui.Window, name string) *Toolkit {
	log.Println("gui.toolkit.NewTab() ADD", name)
	var t Toolkit

	if (w == nil) {
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		time.Sleep(1 * time.Second)
		return nil
	}
	log.Println("gui.toolkit.AddTab() START name =", name)
	tab := ui.NewTab()
	w.SetMargined(margin)

	hbox := ui.NewHorizontalBox() // this makes everything go along the horizon
	hbox.SetPadded(padded)
	tab.Append(name, hbox)
	tabSetMargined(tab) // TODO: run this in the right place(?)
	w.SetChild(tab)

	t.uiWindow = w
	t.uiTab = tab
	t.uiBox = hbox
	return &t
}

func (t *Toolkit) appendTab(name string) *Toolkit {
	log.Println("gui.toolkit.NewTab() ADD", name)
	var newT Toolkit

	if (t.uiTab == nil) {
		log.Println("gui.Toolkit.UiWindow == nil. I can't add a widget without a place to put it")
		panic("should never have happened. wit/gui/toolkit has ui.Tab == nil")
	}
	log.Println("gui.toolkit.AddTab() START name =", name)

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
