package main

import (
	"git.wit.org/wit/gui/toolkit"
)

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func init() {
	// Can you pass values to a plugin init() ? Otherwise, there is no way to safely print
	// log(debugToolkit, "gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)
}

func (t andlabsT) commonChange(tw *toolkit.Widget) {
	log(debugChange, "commonChange() START widget   =", t.Name, t.Type)
	if (tw == nil) {
		log(true, "commonChange() What the fuck. there is no widget t.tw == nil")
		return
	}
	if (tw.Custom == nil) {
		log(debugChange, "commonChange() END    Widget.Custom() = nil", t.tw.Name, t.tw.Type)
		return
	}
	tw.Custom()
	log(debugChange, "commonChange() END   Widget.Custom()", t.tw.Name, t.tw.Type)
}

// does some sanity checks on the internal structs of the binary tree
// TODO: probably this should not panic unless it's running in devel mode (?)
// TODO: redo this now that WidgetType is used and send() is used to package plugins
func (t *andlabsT) broken() bool {
	if (t.parent != nil) {
		return false
	}
	if (t.uiBox == nil) {
		if (t.uiWindow != nil) {
			log(debugToolkit, "UiBox == nil. This is an empty window. Try to add a box")
			t.NewBox()
			return false
		}
		log(true, "UiBox == nil. I can't add a widget without a place to put it")
		// log(debugToolkit, "probably could just make a box here?")
		// corruption or something horrible?
		t.Dump(true)
		panic("wit/gui toolkit/andlabs func broken() invalid goroutine access into this toolkit?")
		panic("wit/gui toolkit/andlabs func broken() this probably should not cause the app to panic here (?)")
		return true
	}
	if (t.uiWindow == nil) {
		log(debugToolkit, "UiWindow == nil. I can't add a widget without a place to put it (IGNORING FOR NOW)")
		t.Dump(debugToolkit)
		return false
	}
	return false
}
func broken(w *toolkit.Widget) bool {
	if (w == nil) {
		log(true, "widget == nil. I can't do anything widget")
		return true
	}
	return false
}
