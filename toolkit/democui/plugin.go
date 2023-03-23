package main

import (
	// if you include more than just this import
	// then your plugin might be doing something un-ideal (just a guess from 2023/02/27)
	"git.wit.org/wit/gui/toolkit"

	"github.com/awesome-gocui/gocui"
)

// This is a map between the widgets in wit/gui and the internal structures of gocui
var viewWidget map[*gocui.View]*toolkit.Widget
var stringWidget map[string]*toolkit.Widget

func Quit() {
	g.Close()
}

// This lists out the know mappings
func listMap() {
	for v, w := range viewWidget {
		log("view =", v.Name, "widget name =", w.Name)
	}
	for s, w := range stringWidget {
		log("string =", s, "widget =", w)
	}
}


//
// This should be called ?
// Pass() ?
// This handles all interaction between the wit/gui package (what golang knows about)
// and this plugin that talks to the OS and does scary and crazy things to make
// a GUI on whatever OS or whatever GUI toolkit you might have (GTK, QT, WASM, libcurses)
//
// Once you are here, you should be in a protected goroutine created by the golang wit/gui package
//
// TODO: make sure you can't escape this goroutine
//
func Send(p *toolkit.Widget, c *toolkit.Widget) {
	if (p == nil) {
		log(debugPlugin, "Send() parent = nil")
	} else {
		log(debugPlugin, "Send() parent =", p.Name, ",", p.Type)
	}
	log(debugPlugin, "Send() child  =", c.Name, ",", c.Type)

	/*
	if (c.Action == "SetMargin") {
		log(debugError, "need to implement SetMargin here")
		setMargin(c, c.B)
		return
	}
	*/

	switch c.Type {
		/*
	case toolkit.Window:
		// doWindow(c)
	case toolkit.Tab:
		// doTab(p, c)
	case toolkit.Group:
		newGroup(p, c)
	case toolkit.Button:
		newButton(p, c)
	case toolkit.Checkbox:
		// doCheckbox(p, c)
	case toolkit.Label:
		// doLabel(p, c)
	case toolkit.Textbox:
		// doTextbox(p, c)
	case toolkit.Slider:
		// doSlider(p, c)
	case toolkit.Spinner:
		// doSpinner(p, c)
	case toolkit.Dropdown:
		// doDropdown(p, c)
	case toolkit.Combobox:
		// doCombobox(p, c)
	case toolkit.Grid:
		// doGrid(p, c)
		*/
	/*
	case toolkit.Flag:
		// log(debugFlags, "plugin Send() flag parent =", p.Name, p.Type)
		// log(debugFlags, "plugin Send() flag child  =", c.Name, c.Type)
		// log(debugFlags, "plugin Send() flag child.Action  =", c.Action)
		// log(debugFlags, "plugin Send() flag child.S  =", c.S)
		// log(debugFlags, "plugin Send() flag child.B  =", c.B)
		// log(debugFlags, "plugin Send() what to flag?")
		// should set the checkbox to this value
		switch c.S {
		case "Toolkit":
			debugToolkit = c.B
		case "Change":
			debugChange = c.B
		case "Plugin":
			debugPlugin = c.B
		case "Flags":
			debugFlags = c.B
		case "Error":
			debugError = c.B
		case "Show":
			ShowDebug()
		default:
			log(debugError, "Can't set unknown flag", c.S)
		}
	*/
	default:
		log(debugError, "plugin Send() unknown parent =", p.Name, p.Type)
		log(debugError, "plugin Send() unknown child  =", c.Name, c.Type)
		log(debugError, "plugin Send() Don't know how to do", c.Type, "yet")
	}
}
