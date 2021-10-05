package gui

import "log"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

var names = make([]string, 100)

func makeWindowDebug() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

/////////////////////////////////////////////////////
	vbox := addGroup(hbox, "Numbers")
	pbar := ui.NewProgressBar()
	vbox.Append(pbar, false)

/////////////////////////////////////////////////////
	vbox = addGroup(hbox, "WindowMap")
	cbox := ui.NewCombobox()

	for name, _ := range Data.WindowMap {
		log.Println("range Data.WindowMap() name =", name)
		addName(cbox, name)
	}

	vbox.Append(cbox, false)

	cbox.OnSelected(func(*ui.Combobox) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("names[x] =", names[x])
		dumpBox(names[x])
	})

/////////////////////////////////////////////////////
	vbox = addGroup(hbox, "Buttons")

	b1 := addButton(vbox, "dumpBox(name)")
	b1.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("names[x] =", names[x])
		dumpBox(names[x])
	})

	b2 := addButton(vbox, "SetMargined()")
	b2.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("findBox; names[x] =", names[x])
		findBox(names[x])
		gw := findBox(names[x])
		if (gw == nil) {
			return
		}
		if (gw.UiTab == nil) {
			return
		}
		if (gw.TabNumber == nil) {
			return
		}
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(gw)
		log.Println("gui.DumpBoxes()\tWindow.UiTab     =", gw.UiTab)
		log.Println("gui.DumpBoxes()\tWindow.TabNumber =", *gw.TabNumber)
		gw.UiTab.SetMargined(*gw.TabNumber, true)
	})

	return hbox
}

var x int = 0

func addName(c *ui.Combobox, s string) {
	c.Append(s)
	names[x] = s
	x = x + 1
}

func addGroup(b *ui.Box, name string) *ui.Box {
	group := ui.NewGroup(name)
	group.SetMargined(true)
	b.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	return vbox
}

func findBox(s string) *GuiWindow {
	for name, window := range Data.WindowMap {
		if (name == s) {
			return window
		}
	}
	return nil
}

func dumpBox(s string) {
	for name, window := range Data.WindowMap {
		if (name != s) {
			continue
		}
		log.Println("gui.DumpBoxes() MAP: ", name)
		if (window.TabNumber == nil) {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
		} else {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
		}
		log.Println("gui.DumpBoxes()\tWindow.name =", window.Name)
		// log.Println("gui.DumpBoxes()\tWindow.UiWindow type =", reflect.TypeOf(window.UiWindow))
		log.Println("gui.DumpBoxes()\tWindow.UiWindow =", window.UiWindow)
		log.Println("gui.DumpBoxes()\tWindow.UiTab    =", window.UiTab)
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			if (name == "MAINBOX") {
				if (Config.Debug) {
					scs := spew.ConfigState{MaxDepth: 1}
					scs.Dump(abox.UiBox)
				}
			}
		}
		if (window.UiTab != nil) {
			pages := window.UiTab.NumPages()
			log.Println("gui.DumpBoxes()\tWindow.UiTab.NumPages() =", pages)
			tabSetMargined(window.UiTab)
			if (Config.Debug) {
				scs := spew.ConfigState{MaxDepth: 2}
				scs.Dump(window.UiTab)
			}
		}
	}
}

func addButton(box *ui.Box, name string) *ui.Button {
	button := ui.NewButton(name)

	button.OnClicked(func(*ui.Button) {
		log.Println("Should do something here")
	})

	box.Append(button, false)
	return button
}
