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
// the label where the user can see which widget is active
var activeLabel *Node
var activeLabelType *Node

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
	title := "ID =" + strconv.Itoa(w.id) + " " + w.widget.Name
	activeLabel.SetText(title)
	activeLabelType.SetText("widget.Type = " + w.widget.Type.String())

	// temporary stuff
	if (w.widget.Type == toolkit.Window) {
		debugWidgetBut1.widget.Action = "Enable"
		send(debugWidgetBut1.parent, debugWidgetBut1)
		debugWidgetBut2.widget.Action = "Enable"
		send(debugWidgetBut2.parent, debugWidgetBut2)
	} else {
		debugWidgetBut1.widget.Action = "Disable"
		send(debugWidgetBut1.parent, debugWidgetBut1)
		debugWidgetBut2.widget.Action = "Disable"
		send(debugWidgetBut2.parent, debugWidgetBut2)
	}
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
		Config.Title = "Widgets"
		Config.Width = 300
		Config.Height = 400
		bugWidget = NewWindow()
		bugWidget.Custom = bugWidget.StandardClose
	} else {
		bugWidget = bugWin.NewTab("Widgets")
	}

	g := bugWidget.NewGroup("widget:")

	activeLabel = g.NewLabel("undef")
	activeLabelType = g.NewLabel("undef")

	// common things that should work against each widget
	g = bugWidget.NewGroup("common things")
	g.NewButton("Disable()", func () {
		activeWidget.widget.Action = "Disable"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Enable()", func () {
		activeWidget.widget.Action = "Enable"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Show()", func () {
		activeWidget.widget.Action = "Show"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Hide()", func () {
		activeWidget.widget.Action = "Hide"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Delete()", func () {
		Delete(activeWidget)
	})
	g.NewButton("Dump()", func () {
		g := debugGui
		d := debugDump
		debugGui = true
		debugDump = true
		activeWidget.Dump()
		debugGui = g
		debugDump = d
	})

	newG := bugWidget.NewGroup("add things")
	newG.debugAddWidgetButtons()

	g = bugWidget.NewGroup("change things")
	g.NewButton("SetMargin(true)", func () {
		activeWidget.widget.Action = "SetMargin"
		activeWidget.widget.B = true
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("SetMargin(false)", func () {
		activeWidget.widget.Action = "SetMargin"
		activeWidget.widget.B = false
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Value()", func () {
		log("activeWidget.B =", activeWidget.widget.B)
		log("activeWidget.I =", activeWidget.widget.I)
		log("activeWidget.S =", activeWidget.widget.S)
	})
	g.NewButton("Set(true)", func () {
		activeWidget.widget.Action = "Set"
		activeWidget.widget.B = true
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Set(false)", func () {
		activeWidget.widget.Action = "Set"
		activeWidget.widget.B = false
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Set(20)", func () {
		activeWidget.widget.Action = "Set"
		activeWidget.widget.B = true
		activeWidget.widget.I = 20
		activeWidget.widget.S = "20"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("SetText('foo')", func () {
		activeWidget.widget.Action = "Set"
		activeWidget.widget.S = "foo"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Delete()", func () {
		activeWidget.widget.Action = "Delete"
		send(activeWidget.parent, activeWidget)
	})
	debugWidgetBut1 = g.NewButton("SetRaw(true)", func () {
		activeWidget.widget.Action = "SetRaw"
		activeWidget.widget.B = true
		send(activeWidget.parent, activeWidget)
	})
	debugWidgetBut2 = g.NewButton("SetRaw(false)", func () {
		activeWidget.widget.Action = "SetRaw"
		activeWidget.widget.B = false
		send(activeWidget.parent, activeWidget)
	})

	g = bugWidget.NewGroup("not working?")
	g.NewButton("Add('foo')", func () {
		activeWidget.widget.Action = "Add"
		activeWidget.widget.S = "foo"
		send(activeWidget.parent, activeWidget)
	})
	g.NewButton("Add button to (1,1)", func () {
		activeWidget.widget.Action = "AddGrid"
		activeWidget.widget.B = false
		send(activeWidget, debugGridLabel)
		// debugGrid = gShoactiveWidget.NewGrid("tmp grid", 2, 3)
	})

	if (activeWidget == nil) {
		setActiveWidget(Config.master)
	}
}

func (n *Node) debugAddWidgetButtons() {
	n.NewButton("Button", func () {
		a := activeWidget.NewButton("myButton", nil)
		a.Custom = func () {
			log("this code is more better", a.widget.B, "id=", a.id)
		}
	})
	n.NewButton("Checkbox", func () {
		a := activeWidget.NewCheckbox("myCheckbox")
		a.Custom = func () {
			log("custom checkox func a=", a.widget.B, "id=", a.id)
		}
	})
	n.NewButton("Label", func () {
		activeWidget.NewLabel("mylabel")
	})
	n.NewButton("Textbox", func () {
		a := activeWidget.NewTextbox("mytext")
		a.Custom = func () {
			log("custom TextBox() a =", a.widget.S, "id=", a.id)
		}
	})
	n.NewButton("Slider", func () {
		a := activeWidget.NewSlider("tmp slider", 10, 55)
		a.Custom = func () {
			log("custom slider() a =", a.widget.I, "id=", a.id)
		}
	})
	n.NewButton("Spinner", func () {
		a := activeWidget.NewSpinner("tmp spinner", 6, 32)
		a.Custom = func () {
			log("custom spinner() a =", a.widget.I, "id=", a.id)
		}
	})
	n.NewButton("Dropdown", func () {
		a := activeWidget.NewDropdown("tmp dropdown")
		a.AddDropdownName("this is better than tcl/tk")
		a.AddDropdownName("make something for tim")
		a.AddDropdownName("for qflow")
		a.Add("and for riscv")
		a.Custom = func () {
			log("custom dropdown() a =", a.widget.Name, a.widget.S, "id=", a.id)
		}
	})
	n.NewButton("Combobox", func () {
		a := activeWidget.NewCombobox("tmp combobox")
		a.Add("mirrors.wit.com")
		a.Add("go.wit.org")
		a.Custom = func () {
			log("custom combobox() a =", a.widget.Name, a.widget.S, "id=", a.id)
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
		// SetDebug(false)
		DebugWidgetWindow(debugGrid)
	})
	n.NewButton("Image", func () {
		activeWidget.NewImage("image")
	})
	n.NewButton("Tab", func () {
		activeWidget.NewTab("myTab")
	})
	n.NewButton("Group", func () {
		a := activeWidget.NewGroup("myGroup")
		a.Custom = func () {
			log("this code is more better", a.widget.B, "id=", a.id)
		}
	})
	n.NewButton("Box(horizontal)", func () {
		a := activeWidget.NewBox("hBox", true)
		a.Custom = func () {
			log("this code is more better", a.widget.B, "id=", a.id)
		}
	})
	n.NewButton("Box(vertical)", func () {
		a := activeWidget.NewBox("vBox", true)
		a.Custom = func () {
			log("this code is more better", a.widget.B, "id=", a.id)
		}
	})
}
