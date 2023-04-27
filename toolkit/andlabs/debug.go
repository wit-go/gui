package main

import "git.wit.org/wit/gui/toolkit"

var defaultBehavior bool = true

var bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
var canvas bool // if set to true, the windows are a raw canvas
var menubar bool // for windows
var stretchy bool // expand things like buttons to the maximum size
var padded bool // add space between things like buttons
var margin bool // add space around the frames of windows

var debugToolkit bool = true
var debugChange bool = true
var debugPlugin bool = true
var debugAction bool = true
var debugFlags bool = true
var debugGrid bool = true
var debugNow bool = true
var debugError bool = true

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

func ShowDebug () {
	log(true, "debugToolkit =", debugToolkit)
	log(true, "debugChange  =", debugChange)
	log(true, "debugAction  =", debugPlugin)
	log(true, "debugFlags    =", debugFlags)
	log(true, "debugNow      =", debugNow)
	log(true, "debugError   =", debugError)
}

func (t *andlabsT) Dump(b bool) {
	if ! b {
		return
	}
	log(b, "Name  = ", t.Width, t.Height)
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
}

/*
func GetDebugToolkit () bool {
	return debugToolkit
}
*/

func flag(a *toolkit.Action) {
	// should set the checkbox to this value
	switch a.S {
	case "Toolkit":
		debugToolkit = a.B
	case "Change":
		debugChange = a.B
	case "Plugin":
		debugPlugin = a.B
	case "Flags":
		debugFlags = a.B
	case "Error":
		debugError = a.B
	case "Now":
		debugNow = a.B
	case "Show":
		ShowDebug()
	default:
		log(debugError, "Can't set unknown flag", a.S)
	}
}
