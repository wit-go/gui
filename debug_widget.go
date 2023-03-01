package gui

import (
	"strconv"
)

func DebugWidgetWindow(w *Node) {
	var win, g *Node

	title := "ID =" + strconv.Itoa(w.id) + " " + w.widget.Name

	Config.Title = title
	Config.Width = 300
	Config.Height = 400
	win = NewWindow()
	win.Custom = w.StandardClose

	g = win.NewGroup("Actions")

	g.NewLabel(title)
	g.NewButton("Dump()", func () {
		w.Dump()
	})
	g.NewButton("Disable()", func () {
		w.widget.Action = "Disable"
		send(w.parent, w)
	})
	g.NewButton("Enable()", func () {
		w.widget.Action = "Enable"
		send(w.parent, w)
	})
	g.NewButton("Show()", func () {
		w.widget.Action = "Show"
		send(w.parent, w)
	})
	g.NewButton("Hide()", func () {
		w.widget.Action = "Hide"
		send(w.parent, w)
	})
	g.NewButton("Value()", func () {
		log("w.B =", w.widget.B)
		log("w.I =", w.widget.I)
		log("w.S =", w.widget.S)
	})
	g.NewButton("Set Value(20)", func () {
		w.widget.Action = "Set"
		w.widget.B = true
		w.widget.I = 20
		w.widget.S = "Set Value(20)"
		send(w.parent, w)
	})
	g.NewButton("Delete()", func () {
		Delete(w)
	})
}

func (n *Node) debugWidgets(makeWindow bool) {
	var w, gList, gShow *Node

	// Either:
	// make a new window
	// make a new tab in the existing window
	if (makeWindow) {
		Config.Title = "Widgets"
		Config.Width = 300
		Config.Height = 400
		w = NewWindow()
		w.Custom = w.StandardClose
	} else {
		w = n.NewTab("Widgets")
	}
	w.Dump()

	gList = w.NewGroup("Pick a widget to debug")
	gShow = w.NewGroup("Added Widgets go here")

	gList.NewButton("Button", func () {
		a := gShow.NewButton("myButton", func () {
			log("this code is more better")
		})
		DebugWidgetWindow(a)
	})
	gList.NewButton("Checkbox", func () {
		a := gShow.NewCheckbox("myCheckbox")
		a.Custom = func () {
			log("custom checkox func a =", a.widget.B, a.id)
		}
		DebugWidgetWindow(a)
	})
	gList.NewButton("Label", func () {
		a := gShow.NewLabel("mylabel")
		DebugWidgetWindow(a)
	})
	gList.NewButton("Textbox", func () {
		a := gShow.NewTextbox("mytext")
		a.Custom = func () {
			log("custom TextBox() a =", a.widget.S, a.id)
		}
		DebugWidgetWindow(a)
	})
	gList.NewButton("Slider", func () {
		a := gShow.NewSlider("tmp slider", 10, 55)
		a.Custom = func () {
			log("custom slider() a =", a.widget.S, a.id)
		}
		DebugWidgetWindow(a)
	})
	gList.NewButton("Spinner", func () {
		a := gShow.NewSpinner("tmp spinner", 6, 32)
		a.Custom = func () {
			log("custom spinner() a =", a.widget.S, a.id)
		}
		DebugWidgetWindow(a)
	})
}
