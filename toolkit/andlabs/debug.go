package main

// import "git.wit.org/wit/gui/toolkit"

import "github.com/davecgh/go-spew/spew"

var defaultBehavior bool = true

var bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
var canvas bool // if set to true, the windows are a raw canvas
var menubar bool // for windows
var stretchy bool // expand things like buttons to the maximum size
var padded bool // add space between things like buttons
var margin bool // add space around the frames of windows

var debugToolkit bool
var debugChange bool
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

func ShowDebug () {
	log(true, "debugToolkit =", debugToolkit)
	log(true, "debugChange =", debugChange)
}

func GetDebugToolkit () bool {
	return debugToolkit
}

func (t *andlabsT) String() string {
	return t.GetText()
}

func forceDump(t *andlabsT) {
	tmp := debugToolkit
	debugToolkit = true
	t.Dump()
	debugToolkit = tmp
}

func (t *andlabsT) GetText() string {
	log(debugToolkit, "gui.Toolkit.GetText() Enter debugToolkit=", debugToolkit)
	if (t.uiEntry != nil) {
		log(debugToolkit, "gui.Toolkit.uiEntry.Text() =", t.uiEntry.Text())
		return t.uiEntry.Text()
	}
	if (t.uiMultilineEntry != nil) {
		log(debugToolkit, "gui.Toolkit.uiMultilineEntry.Text() =", t.uiMultilineEntry.Text())
		text := t.uiMultilineEntry.Text()
		log(debugToolkit, "gui.Toolkit.uiMultilineEntry.Text() =", text)
		t.text = text
		return text
	}
	if (t.uiCombobox != nil) {
		log(debugToolkit, "gui.Toolkit.uiCombobox() =", t.text)
		return t.text
	}
	return ""
}

func (t *andlabsT) SetText(s string) bool {
	log(debugToolkit, "gui.Toolkit.Text() SetText() Enter")
	if (t.uiEntry != nil) {
		log(debugToolkit, "gui.Toolkit.Value() =", t.uiEntry.Text)
		t.uiEntry.SetText(s)
		return true
	}
	if (t.uiMultilineEntry != nil) {
		log(debugToolkit, "gui.Toolkit.Value() =", t.uiMultilineEntry.Text)
		t.uiMultilineEntry.SetText(s)
		return true
	}
	return false
}

func sanity(t *andlabsT) bool {
	if (debugToolkit) {
		log(debugToolkit, "gui.Toolkit.Value() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t.uiEntry == nil) {
		log(debugToolkit, "gui.Toolkit.Value() =", t.uiEntry.Text)
		return false
	}
	return true
}

func (t *andlabsT) SetValue(i int) bool {
	log(debugToolkit, "gui.Toolkit.SetValue() START")
	if (sanity(t)) {
		return false
	}
	t.Dump()
	// panic("got to toolkit.SetValue")
	return true
}

func (t *andlabsT) Value() int {
	if (debugToolkit) {
		log(debugToolkit, "gui.Toolkit.Value() Enter")
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(t)
	}
	if (t == nil) {
		log(debugToolkit, "gui.Toolkit.Value() can not get value t == nil")
		return 0
	}
	if (t.uiSlider != nil) {
		log(debugToolkit, "gui.Toolkit.Value() =", t.uiSlider.Value)
		return t.uiSlider.Value()
	}
	if (t.uiSpinbox != nil) {
		log(debugToolkit, "gui.Toolkit.Value() =", t.uiSpinbox.Value)
		return t.uiSpinbox.Value()
	}
	log(debugToolkit, "gui.Toolkit.Value() Could not find a ui element to get a value from")
	return 0
}

func (t *andlabsT) Dump() {
	if ! debugToolkit {
		return
	}
	log(debugToolkit, "gui.Toolkit.Dump() Name  = ", t.Name, t.Width, t.Height)
	if (t.uiBox != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiBox      =", t.uiBox)
	}
	if (t.uiButton != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiButton    =", t.uiButton)
	}
	if (t.uiCombobox != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiCombobox  =", t.uiCombobox)
	}
	if (t.uiWindow != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiWindow    =", t.uiWindow)
	}
	if (t.uiTab != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiTab       =", t.uiTab)
	}
	if (t.uiGroup != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiGroup     =", t.uiGroup)
	}
	if (t.uiEntry != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiEntry     =", t.uiEntry)
	}
	if (t.uiMultilineEntry != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiMultilineEntry =", t.uiMultilineEntry)
	}
	if (t.uiSlider != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiSlider    =", t.uiSlider)
	}
	if (t.uiCheckbox != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() uiCheckbox  =", t.uiCheckbox)
	}
	if (t.OnExit != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() OnExit      =", t.OnExit)
	}
	if (t.Custom != nil) {
		log(debugToolkit, "gui.Toolkit.Dump() Custom      =", t.Custom)
	}
	log(debugToolkit, "gui.Toolkit.Dump() c         =", t.c)
	log(debugToolkit, "gui.Toolkit.Dump() val       =", t.val)
	log(debugToolkit, "gui.Toolkit.Dump() text      =", t.text)
}
