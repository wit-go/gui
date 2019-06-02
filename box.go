package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import pb "git.wit.com/wit/witProtobuf"
// import "github.com/davecgh/go-spew/spew"

// THIS IS CLEAN

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
			newbox.Window.UiTab.InsertAt(newbox.Name, 0, newbox.UiBox)
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
	log.Println("gui.add() adding", newbox.Name, "to", box.Name)
	// copy the box settings over
	newbox.Window = box.Window
	box.UiBox.Append(newbox.UiBox, false)

	// add the newbox to the Window.BoxMap[]
	box.Window.BoxMap[newbox.Name] = newbox
	log.Println("gui.add() END")
}

func InitGuiBox(gw *GuiWindow, box *GuiBox, uiBox *ui.Box, name string) *GuiBox {
	log.Println("InitGuiBox() START")
	var newGuiBox GuiBox
	newGuiBox.UiBox = uiBox
	newGuiBox.Window = gw
	uiBox.SetPadded(true)

	if (box != nil) {
		log.Println("InitGuiBox() APPEND NEW BOX TO OLD BOX")
		box.UiBox.Append(uiBox, false)
	} else {
		log.Println("InitGuiBox() APPEND NEW BOX TO TAB")
		gw.UiTab.Append(name, uiBox)
	}
	gw.BoxMap[name] = &newGuiBox
	log.Println("InitGuiBox() END")
	return &newGuiBox
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

func CreateGenericBox(gw *GuiWindow, b *GuiButton, name string) *GuiBox{
	log.Println("CreateAddVmBox() START name =", name)

	var box *GuiBox
	box = new(GuiBox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	box.UiBox = vbox
	box.Window = gw
	gw.BoxMap["ADD VM" + name] = box

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	gw.UiTab.Append(name, vbox)
	gw.UiTab.SetMargined(0, true)
	return box
}

func CreateBox(gw *GuiWindow, name string) *GuiBox {
	log.Println("CreateVmBox() START")
	log.Println("CreateVmBox() vm.Name =", name)
	log.Println("CreateVmBox() gw =", gw)

	var box *GuiBox
	box = new(GuiBox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	log.Println("CreateVmBox() vbox =", vbox)
	log.Println("CreateVmBox() box.UiBox =", box.UiBox)
	box.UiBox = vbox
	log.Println("CreateVmBox() box.Window =", box.Window)
	box.Window = gw
	log.Println("CreateVmBox() gw.BoxMap =", gw.BoxMap)
	gw.BoxMap[name] = box

	hboxAccount := ui.NewHorizontalBox()
	hboxAccount.SetPadded(true)
	vbox.Append(hboxAccount, false)

	box.UiBox = hboxAccount

	gw.UiTab.Append(name, vbox)
	gw.UiTab.SetMargined(0, true)
	return box
}
