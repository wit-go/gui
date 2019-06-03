package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// add(nil, newbox, "")				// use this when the Window is created. Always called 'MAINBOX'
// add(gw.BoxMap["MAINBOX"], newbox, name)	// use this to add a box off the main box
// add(gw.BoxMap["BUTTONBOX"], newbox, name)	// use this to add something to the box called 'BUTTONBOX'
// add(box, newbox, name)			// add 'newbox' to 'box' and call it 'name'
func add(box *GuiBox, newbox *GuiBox) {
	log.Println("gui.add() START box =", box)
	log.Println("gui.add() START newbox =", newbox)
	if (box == nil) {
		log.Println("\tgui.add() add to Window as MAINBOX")
		if (newbox.Window.UiTab != nil) {
			// create a new tab here
			// add the box to it as MAINBOX
			log.Println("\tgui.add() add to Window as a UiTab")
			// TODO: allow passing where to append
			// newbox.Window.UiTab.InsertAt(newbox.Name, 0, newbox.UiBox)
			newbox.Window.UiTab.Append(newbox.Name, newbox.UiBox)
			newbox.Window.UiTab.SetMargined(0, true)

			// TODO: figure out how to make a new Tab/Window/Box here
			// window := InitGuiWindow(Data.Config, newbox.Name, gw.MakeWindow, gw.UiWindow, gw.UiTab)
			// window.UiTab.Delete(0)
			// window.MakeWindow(window)
			// newbox.Window = window

			newbox.Window.BoxMap["MAINBOX"] = newbox
			log.Println("gui.add() END")
			return
		} else {
			log.Println("\tgui.add() ERROR DONT KNOW HOW TO ADD TO A RAW WINDOW YET")
			// add this to the window
		}
		log.Println("\tgui.add() ERROR DON'T KNOW HOW TO add to Window as MAINBOX DONE")
		log.Println("gui.add() END")
		return
	}
	log.Println("\tgui.add() adding", newbox.Name, "to", box.Name)
	// copy the box settings over
	newbox.Window = box.Window
	box.UiBox.Append(newbox.UiBox, false)

	// add the newbox to the Window.BoxMap[]
	box.Window.BoxMap[newbox.Name] = newbox
	log.Println("gui.add() END")
}

func HardHorizontalBox(gw *GuiWindow) *GuiBox {
	log.Println("HardHorizontalBreak START")
	var newbox *GuiBox
	newbox = new(GuiBox)
	newbox.Window = gw

	box := gw.BoxMap["MAINBOX"]
	if (box != nil) {
		// There is already a box. Add a Seperator
		tmp := ui.NewHorizontalSeparator()
		box.UiBox.Append(tmp, true)
		add(gw.BoxMap["MAINBOX"], newbox)
	}

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(true)
	newbox.UiBox = hbox
	newbox.Name = "Hbox1"
	add(gw.BoxMap["MAINBOX"], newbox)
	log.Println("HardHorizontalBreak END")
	return newbox
}

func HardVerticalBreak(box *GuiBox) {
	log.Println("HardVerticalBreak START")
	gw := box.Window
	mainbox := gw.BoxMap["MAIN"]
	if (mainbox == nil) {
		log.Println("HardHorizontalBreak ERROR MAIN box == nil")
		return
	}

	tmp := ui.NewVerticalSeparator()
	mainbox.UiBox.Append(tmp, false)

	hbox := ui.NewVerticalBox()
	hbox.SetPadded(true)
	box.UiBox = hbox
	mainbox.UiBox.Append(hbox, false)
	log.Println("HardVerticalBreak END")
}

func HorizontalBreak(box *GuiBox) {
	tmp := ui.NewHorizontalSeparator()
	box.UiBox.Append(tmp, false)
}

func VerticalBreak(box *GuiBox) {
	tmp := ui.NewVerticalSeparator()
	box.UiBox.Append(tmp, false)
}

func AddGenericBox(gw *GuiWindow, name string) *GuiBox {
	log.Println("AddGenericBox() START name =", name)
	// create a new vertical box off of the mainbox
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	var newbox *GuiBox
	newbox = new(GuiBox)
	newbox.UiBox = vbox
	newbox.Window = gw
	newbox.Name = name
	add(gw.BoxMap["MAINBOX"], newbox)

	// create a new horizonal box off of the vertical box
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	var newhbox *GuiBox
	newhbox = new(GuiBox)
	newhbox.UiBox = hbox
	newhbox.Window = gw
	newhbox.Name = "Hbox1"
	add(newbox, newhbox)

	return newbox
}
