package gui

import "log"
import "os"
// import "reflect"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

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
			// newbox.Window.UiTab.SetMargined(0, true)

			// TODO: figure out how to make a new Tab/Window/Box here
			// window := InitGuiWindow(Data.Config, newbox.Name, gw.MakeWindow, gw.UiWindow, gw.UiTab)
			// window.UiTab.Delete(0)
			// window.MakeWindow(window)
			// newbox.Window = window

			newbox.Window.BoxMap["MAINBOX"] = newbox
			log.Println("gui.add() END")
			panic("gui.add() MAINBOX gui.add() END")
			return
		} else {
			log.Println("\tgui.add() ERROR DONT KNOW HOW TO ADD TO A RAW WINDOW YET")
			// add this to the window
		}
		log.Println("\tgui.add() ERROR DON'T KNOW HOW TO add to Window as MAINBOX DONE")
		log.Println("gui.add() END")
		panic("gui.add() gui.add() END")
		return
	}
	log.Println("\tgui.add() adding", newbox.Name, "to", box.Name)
	// copy the box settings over
	newbox.Window = box.Window
	if (box.node == nil) {
		box.Dump()
		panic("gui.add() ERROR box.node == nil")
	}

	if (newbox.UiBox == nil) {
		panic("gui.add() ERROR newbox.UiBox == nil")
	}

	if (box.UiBox == nil) {
		box.Dump()
		// panic("gui.add() ERROR box.UiBox == nil")
		return
		// TODO: fix this whole add() function // Oct 9
	}
	box.UiBox.Append(newbox.UiBox, false)
	box.Dump()
	panic("gui.add()")

	// add the newbox to the Window.BoxMap[]
	box.Window.BoxMap[newbox.Name] = newbox
	log.Println("gui.add() END")
}

func (n *Node) NewBox(axis int, name string) *Node {
	if (n.box == nil) {
		log.Println("box == nil. I can't add a box!")
		panic("gui.Node.NewBox() node.box == nil")
	}

	newBox		:= new(GuiBox)
	newBox.Window	= n.window
	newBox.Name	= name

	// make a new box & a new node
	newNode := n.makeNode(name, 111, 100 + Config.counter)
	Config.counter += 1

	var uiBox *ui.Box
	if (axis == Xaxis) {
		uiBox = ui.NewHorizontalBox()
	} else {
		uiBox = ui.NewVerticalBox()
	}
	uiBox.SetPadded(true)
	newBox.UiBox = uiBox
	newNode.uiBox = uiBox

	n.Append(newNode)
	// add(n.box, newBox)
	return newNode
}

func NewBox(box *GuiBox, axis int, name string) *GuiBox {
	log.Println("gui.NewBox() START")
	n := box.FindNode()
	if (n == nil) {
		log.Println("gui.NewBox() SERIOUS ERROR. CAN NOT FIND NODE")
		os.Exit(0)
	} else {
		log.Println("gui.NewBox() node =", n.Name)
	}
	var newbox *GuiBox
	newbox		= new(GuiBox)
	newbox.Window	= box.Window
	newbox.Name	= name

	var uiBox *ui.Box
	if (axis == Xaxis) {
		uiBox = ui.NewHorizontalBox()
	} else {
		uiBox = ui.NewVerticalBox()
	}
	uiBox.SetPadded(true)
	newbox.UiBox = uiBox
	add(box, newbox)
	// panic("gui.NewBox")
	return newbox
}

func HardBox(gw *GuiWindow, axis int, name string) *GuiBox {
	log.Println("HardBox() START axis =", axis)

	if (gw.node == nil) {
		gw.Dump()
		panic("gui.HardBox() gw.node == nil")
	}
	// add a Vertical Seperator if there is already a box
	// Is this right?
	box := gw.BoxMap["MAINBOX"]
	if (box != nil) {
		if (axis == Xaxis) {
			VerticalBreak(box)
		} else {
			HorizontalBreak(box)
		}
	}

	// make the new vbox
	var uiBox *ui.Box
	if (axis == Xaxis) {
		uiBox = ui.NewHorizontalBox()
	} else {
		uiBox = ui.NewVerticalBox()
	}
	uiBox.SetPadded(true)

	// Init a new GuiBox
	newbox		:= new(GuiBox)
	newbox.Window	= gw
	newbox.UiBox	= uiBox
	newbox.Name	= name

	add(gw.BoxMap["MAINBOX"], newbox)

	log.Println("HardBox END")
	return newbox
}

func HorizontalBreak(box *GuiBox) {
	log.Println("VerticalSeparator added to box =", box.Name)
	tmp := ui.NewHorizontalSeparator()
	if (box == nil) {
		return
	}
	if (box.UiBox == nil) {
		return
	}
	box.UiBox.Append(tmp, false)
}

func VerticalBreak(box *GuiBox) {
	log.Println("VerticalSeparator  added to box =", box.Name)
	tmp := ui.NewVerticalSeparator()
	box.UiBox.Append(tmp, false)
}
