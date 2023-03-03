package main

import "git.wit.org/wit/gui/toolkit"

// import "github.com/davecgh/go-spew/spew"

var defaultBehavior bool = true

var bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
var canvas bool // if set to true, the windows are a raw canvas
var menubar bool // for windows
var stretchy bool // expand things like buttons to the maximum size
var padded bool // add space between things like buttons
var margin bool // add space around the frames of windows

var debugToolkit bool
var debugChange bool
var debugPlugin bool
var debugFlag bool
var debugError bool = true
// var DebugToolkit bool

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func setDefaultBehavior(s bool) {
	defaultBehavior = s
	if (defaultBehavior) {
		log(debugToolkit, "Setting this toolkit to use the default behavior.")
		log(debugToolkit, "This is the 'guessing' part as defined by the wit/gui 'Principles'. Refer to the docs.")
		stretchy = false
		padded = true
		menubar = true
		margin = true
		canvas = false
		bookshelf = true // 99% of the time, things make a vertical stack of objects
	} else {
		log(debugToolkit, "This toolkit is set to ignore the default behavior.")
	}
}

/*
func SetDebugToolkit (s bool) {
	debugToolkit = s
	log(true, "debugToolkit =", debugToolkit)
	log(true, "debugChange =", debugChange)
}

func SetDebugChange (s bool) {
	debugChange = s
	log(true, "debugToolkit =", debugToolkit)
	log(true, "debugChange =", debugChange)
}
*/

func ShowDebug () {
	log(true, "debugToolkit =", debugToolkit)
	log(true, "debugError   =", debugError)
	log(true, "debugChange  =", debugChange)
}

func (t *andlabsT) Dump(b bool) {
	if ! b {
		return
	}
	log(b, "Name  = ", t.Name, t.Width, t.Height)
	if (t.uiBox != nil) {
		log(b, "uiBox      =", t.uiBox)
	}
	if (t.uiButton != nil) {
		log(b, "uiButton    =", t.uiButton)
	}
	if (t.uiCombobox != nil) {
		log(b, "uiCombobox  =", t.uiCombobox)
	}
	if (t.uiWindow != nil) {
		log(b, "uiWindow    =", t.uiWindow)
	}
	if (t.uiTab != nil) {
		log(b, "uiTab       =", t.uiTab)
	}
	if (t.uiGroup != nil) {
		log(b, "uiGroup     =", t.uiGroup)
	}
	if (t.uiEntry != nil) {
		log(b, "uiEntry     =", t.uiEntry)
	}
	if (t.uiMultilineEntry != nil) {
		log(b, "uiMultilineEntry =", t.uiMultilineEntry)
	}
	if (t.uiSlider != nil) {
		log(b, "uiSlider    =", t.uiSlider)
	}
	if (t.uiCheckbox != nil) {
		log(b, "uiCheckbox  =", t.uiCheckbox)
	}
	widgetDump(b, t.tw)
}

func widgetDump(b bool, w *toolkit.Widget) {
	if (w == nil) {
		log(b, "widget = nil")
		return
	}

	log(b, "widget.Name        =", w.Name)
	log(b, "widget.Action      =", w.Action)
	log(b, "widget.Type        =", w.Type)
	log(b, "widget.Custom      =", w.Custom)
	log(b, "widget.B           =", w.B)
	log(b, "widget.I           =", w.I)
	log(b, "widget.Width       =", w.Width)
	log(b, "widget.Height      =", w.Height)
	log(b, "widget.X           =", w.X)
	log(b, "widget.Y           =", w.Y)
}

/*
func GetDebugToolkit () bool {
	return debugToolkit
}
*/
