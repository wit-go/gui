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
	var w *ui.Window
	var newt *Toolkit

	log.Println("gui.toolkit.AddTab() sleep 3")

	w = t.uiWindow
	if (w == nil) {
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		return nil
	}

	if (t.uiTab == nil) {
		// this means you have to make a new tab
		log.Println("gui.toolkit.NewTab() GOOD. This should be the first tab:", name)
		newt = newTab(w, name)
		t.uiTab = newt.uiTab
	} else {
		log.Println("gui.toolkit.NewTab() GOOD. This should be an additional tab:", name)
		newt = t.appendTab(name)
		// this means you have to append a tab
	}
	log.Println("t:")
	t.Dump()
	log.Println("newt:")
	newt.Dump()

	return newt
}

func (t *Toolkit) SetTabBox(box *ui.Box) {
	var tab *ui.Tab

	log.Println("wit/gui/toolkit SetTabBox()")
	t.Dump()
	if (t.uiTab == nil) {
		log.Println("wit/gui/toolkit SetTabBox() got uiTab == nil")
		panic("fucknuts")
		return
	}
	if (t.uiBox == nil) {
		log.Println("wit/gui/toolkit SetTabBox() got uiBox == nil. Appending a new tab here")
		tab = t.uiTab
		tab.Append(t.Name, box)
		tabSetMargined(tab)
		return
	} else {
		log.Println("wit/gui/toolkit SetTabBox() got uiBox != nil. Appending the box to the existing box strechy = true")
		t.uiBox.Append(box, true) // strechy == true
		t.uiBox2 = box
		// t.uiBox.Append(box, false) // strechy == false
		return
	}

}

// This sets _all_ the tabs to Margin = true
//
// TODO: do proper tab tracking (will be complicated). low priority
func tabSetMargined(tab *ui.Tab) {
	c := tab.NumPages()
	for i := 0; i < c; i++ {
		tab.SetMargined(i, true)
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
	// time.Sleep(2 * time.Second)
	tab := ui.NewTab()
	w.SetMargined(true)

	hbox := ui.NewHorizontalBox() // this makes everything go along the horizon
	// hbox := ui.NewVerticalBox()
	hbox.SetPadded(true)
	tab.Append(name, hbox)
	w.SetChild(tab)

	t.uiWindow = w
	t.uiTab = tab
	t.uiBox = hbox
	// tabSetMargined(newNode.uiTab)
	return &t
}

func (t *Toolkit) appendTab(name string) *Toolkit {
	log.Println("gui.toolkit.NewTab() ADD", name)
	var newT Toolkit

	if (t.uiWindow == nil) {
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		log.Println("gui.toolkit.NewTab() node.UiWindow == nil. I can't add a tab without a window")
		time.Sleep(1 * time.Second)
		return nil
	}
	log.Println("gui.toolkit.AddTab() START name =", name)

	hbox := ui.NewHorizontalBox() // this makes everything go along the horizon
	// hbox := ui.NewVerticalBox()
	hbox.SetPadded(true)
	t.uiTab.Append(name, hbox)
	// w.SetChild(tab)

	newT.uiWindow = t.uiWindow
	newT.uiTab = t.uiTab
	newT.uiBox = hbox
	// tabSetMargined(newNode.uiTab)
	return &newT
}
