package gui

import (
	"strconv"
	"git.wit.org/wit/gui/toolkit"
)


// global var for checking to see if this
// window/tab for debugging a widget exists
// check the binary tree instead (?) for a window called "Widgets" (bad idea)
var bugWidget *Node

// the widget all these actions are run against
var activeWidget *Node

// for testing move, this is the node things are put on
var activeJunk *Node

// the label where the user can see which widget is active
var activeLabel *Node
var activeLabelType *Node
var activeLabelNewName *Node
var activeLabelNewType *Node
var activeLabelNewX *Node
var activeLabelNewY *Node
var activeLabelNewB *Node

// tmp junk
var debugGrid *Node
var debugGridLabel *Node
var debugWidgetBut1, debugWidgetBut2 *Node

func setActiveWidget(w *Node) {
	if (w == nil) {
		log(debugError, "setActiveWidget() was sent nil !!!")
		return
	}
	activeWidget = w
	log(true, "The Widget is set to", w.id, w.Name)
	if (activeLabel == nil) {
		// the debug window doesn't exist yet so you can't display the change
		// TODO: make a fake binary tree for this(?)
		return
	}
	title := "ID =" + strconv.Itoa(w.id) + " " + w.Name
	activeLabel.SetText(title)
	activeLabelType.SetText("widget.Type = " + w.WidgetType.String())
	return
}

func DebugWidgetWindow(w *Node) {
	if (bugWidget != nil) {
		// this window was already created. Just change the widget we are working against
		setActiveWidget(w)
		return
	}

	// Either:
	// make a new window
	// make a new tab in the existing window
	if (makeTabs) {
		bugWidget = me.rootNode.NewWindow("Widgets")
		bugWidget.Custom = bugWidget.StandardClose
	} else {
		bugWidget = bugWin.NewTab("Widgets")
	}

	g := bugWidget.NewGroup("widget:")

	g2 := g.NewGroup("widget:")
	activeLabel = g2.NewLabel("undef")
	g2 = g.NewGroup("type:")
	activeLabelType = g2.NewLabel("undef")
	g2 = g.NewGroup("New name:")
	activeLabelNewName = g2.NewCombobox("newthing")
	activeLabelNewName.AddText("penguin")
	activeLabelNewName.AddText("snow")
	activeLabelNewName.AddText("GO")
	activeLabelNewName.AddText("debian")
	activeLabelNewName.AddText("RiscV")

	g2 = g.NewGroup("At X:")
	activeLabelNewX = g2.NewSpinner("tmp spinner", -1, 100)

	g2 = g.NewGroup("At Y:")
	activeLabelNewY = g2.NewSpinner("tmp spinner", -1, 100)

	g2 = g.NewGroup("bool B:")
	activeLabelNewB = g2.NewCheckbox("tmp bool")


	// common things that should work against each widget
	g = bugWidget.NewGroup("common things")
	g.NewButton("Enable()", func () {
		activeWidget.Enable()
	})
	g.NewButton("Disable()", func () {
		activeWidget.Disable()
	})
	g.NewButton("Show()", func () {
		activeWidget.Show()
	})
	g.NewButton("Hide()", func () {
		activeWidget.Hide()
	})
	g.NewButton("Dump()", func () {
		activeWidget.Dump()
	})

	g = bugWidget.NewGroup("add things")
	g.debugAddWidgetButton()
	g.NewLabel("experiments:")
	g.debugAddWidgetButtons()

	g = bugWidget.NewGroup("change things")
	g.NewButton("AddText()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.AddText
		a.S = activeLabelNewName.S
		newaction(&a, activeWidget, nil)
	})
	g.NewButton("SetText()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.SetText
		a.S = activeLabelNewName.S
		newaction(&a, activeWidget, nil)
	})
	g.NewButton("Margin()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.Margin
		newaction(&a, activeWidget, nil)
	})
	g.NewButton("Unmargin()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.Unmargin
		newaction(&a, activeWidget, nil)
	})
	g.NewButton("Pad()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.Pad
		newaction(&a, activeWidget, nil)
	})
	g.NewButton("Unpad()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.Unpad
		newaction(&a, activeWidget, nil)
	})
	g.NewButton("Move(junk)", func () {
		var a toolkit.Action
		a.ActionType = toolkit.Move
		newaction(&a, activeWidget, activeJunk)
	})
	g.NewButton("Delete()", func () {
		var a toolkit.Action
		a.ActionType = toolkit.Delete
		newaction(&a, activeWidget, activeJunk)
	})

	g = bugWidget.NewGroup("not working?")
	activeJunk = bugWidget.NewGroup("junk:")
	activeJunk.NewLabel("test junk")

	if (activeWidget == nil) {
		setActiveWidget(me.rootNode)
	}
}

func (n *Node) debugAddWidgetButtons() {
	n.NewButton("Dropdown", func () {
		a := activeWidget.NewDropdown("tmp dropdown")
		a.AddText("this is better than tcl/tk")
		a.AddText("make something for tim for qflow")
		a.AddText("and for riscv")
		a.Custom = func () {
			log("custom dropdown() a =", a.Name, a.S, "id=", a.id)
		}
	})
	n.NewButton("Combobox", func () {
		a := activeWidget.NewCombobox("tmp combobox")
		a.AddText("mirrors.wit.com")
		a.AddText("go.wit.com")
		a.Custom = func () {
			log("custom combobox() a =", a.Name, a.S, "id=", a.id)
		}
	})
	n.NewButton("Grid", func () {
		// Grid numbering by (X,Y)
		// -----------------------------
		// -- (1,1) -- (2,1) -- (3,1) --
		// -- (1,2) -- (2,1) -- (3,1) --
		// -----------------------------

		// SetDebug(true)
		debugGrid = activeWidget.NewGrid("tmp grid", 2, 3)
		debugGridLabel = debugGrid.NewLabel("mirrors.wit.com")
		/*
		debugGrid.SetNext(0,1)
		debugGrid.NewLabel("foo (0,1)")
		debugGrid.SetNext(1,1)
		debugGrid.NewLabel("foo (1,1)")
		debugGrid.SetNext(2,1)
		debugGrid.NewLabel("foo (2,1)")
		*/
		// SetDebug(false)
		DebugWidgetWindow(debugGrid)
	})
	n.NewButton("Image", func () {
		activeWidget.NewImage("image")
	})
	n.NewButton("Box(horizontal)", func () {
		a := activeWidget.NewBox("hBox", true)
		a.NewLabel("hBox")
		a.NewLabel("hBox 2")
	})
	n.NewButton("Box(vertical)", func () {
		a := activeWidget.NewBox("vBox", false)
		a.NewLabel("vBox")
		a.NewLabel("vBox 2")
	})
}

func (n *Node) debugAddWidgetButton() {
	activeLabelNewType = n.NewDropdown("tmp dropdown")
	activeLabelNewType.AddText("Window")
	activeLabelNewType.AddText("Tab")
	activeLabelNewType.AddText("Frame")
	activeLabelNewType.AddText("Grid")
	activeLabelNewType.AddText("Group")
	activeLabelNewType.AddText("Box")
	activeLabelNewType.AddText("Button")
	activeLabelNewType.AddText("Checkbox")
	activeLabelNewType.AddText("Dropdown")
	activeLabelNewType.AddText("Combobox")
	activeLabelNewType.AddText("Label")
	activeLabelNewType.AddText("Textbox")
	activeLabelNewType.AddText("Slider")
	activeLabelNewType.AddText("Spinner")
	activeLabelNewType.AddText("Image")
	activeLabelNewType.AddText("Area")
	activeLabelNewType.AddText("Form")
	activeLabelNewType.AddText("Font")
	activeLabelNewType.AddText("Color")
	activeLabelNewType.AddText("Dialog")

	n.NewButton("Add", func () {
		name :=  activeLabelNewName.S
		newX :=  activeLabelNewX.I
		newY :=  activeLabelNewY.I
		newB :=  activeLabelNewB.B

		if (newY == -1) {
			name = name + " (" + strconv.Itoa(activeWidget.NextX) + "," + strconv.Itoa(activeWidget.NextY) + ")"
		} else {
			activeWidget.SetNext(newX, newY)
			name = name + " (" + strconv.Itoa(newX) + "," + strconv.Itoa(newY) + ")"
		}
		log("New Name =", name)
		log("New Type =", activeLabelNewType.S)
		log("New X    =", newX)
		log("New Y    =", newY)
		log("activeWidget.NextX    =", activeWidget.NextX)
		log("activeWidget.NextY    =", activeWidget.NextY)
		log(debugNow, "Add() size (X,Y)", activeWidget.X, activeWidget.Y, "put next thing at (X,Y) =", activeWidget.NextX, activeWidget.NextY)
		activeWidget.Dump()

		// activeWidget.X = newX
		// activeWidget.Y = newY

		switch activeLabelNewType.S {
		case "Grid":
			activeWidget.NewGrid(name, newX, newY)
		case "Group":
			activeWidget.NewGroup(name)
		case "Box":
			activeWidget.NewBox(name, newB)
		case "Button":
			var n *Node
			n = activeWidget.NewButton(name, func () {
				log("got to button", name, n.id)
			})
		case "Checkbox":
			a := activeWidget.NewCheckbox(name)
			a.Custom = func () {
				log("custom checkox func a=", a.B, "id=", a.id)
			}
		case "Dropdown":
			a := activeWidget.NewDropdown(name)
			a.AddText(name + " yay")
			a.AddText(name + " haha")
			a.Custom = func () {
				log("WTF a=", a.B, "id=", a.id)
			}
		case "Combobox":
			a := activeWidget.NewCombobox(name)
			a.AddText(name + " foo")
			a.AddText(name + " bar")
		case "Label":
			activeWidget.NewLabel(name)
		case "Textbox":
			activeWidget.NewTextbox(name)
		case "Slider":
			activeWidget.NewSlider(name, newX, newY)
		case "Spinner":
			activeWidget.NewSpinner(name, newX, newY)
		default:
			log(debugError, "make what type?")
		}
	})
}
