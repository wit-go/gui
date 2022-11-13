package main

import "log"

func init() {
	if (DebugToolkit) {
		log.Println("gui/toolkit init() Setting defaultBehavior = true")
	}
	setDefaultBehavior(true)
}

func (t andlabsT) commonChange(widget string) {
	s := t.String()
	if (DebugToolkit) {
		log.Println("gui.Toolkit.ui.OnChanged() =", s)
	}
	if (t.OnChanged != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.OnChanged() trying to run toolkit.OnChanged() entered val =", s)
		}
		t.OnChanged(&t)
		return
	}
	if (t.Custom != nil) {
		if (DebugToolkit) {
			log.Println("gui.Toolkit.OnChanged() Running toolkit.Custom()")
			t.Dump()
		}
		t.Custom()
		return
	}
	if (DebugToolkit) {
		log.Println("gui.Toolkit.OnChanged() ENDED without finding any callback")
	}
}

// does some sanity checks on the internal structs of the binary tree
// TODO: probably this should not panic unless it's running in devel mode (?)
func (t *andlabsT) broken() bool {
	if (t.uiBox == nil) {
		if (t.uiWindow != nil) {
			if (DebugToolkit) {
				log.Println("gui.Toolkit.UiBox == nil. This is an empty window. Try to add a box")
			}
			t.NewBox()
			return false
		}
		log.Println("gui.Toolkit.UiBox == nil. I can't add a widget without a place to put it")
		// log.Println("probably could just make a box here?")
		// corruption or something horrible?
		panic("wit/gui toolkit/andlabs func broken() invalid goroutine access into this toolkit?")
		panic("wit/gui toolkit/andlabs func broken() this probably should not cause the app to panic here (?)")
		return true
	}
	if (t.uiWindow == nil) {
		log.Println("gui.Toolkit.UiWindow == nil. I can't add a widget without a place to put it (IGNORING FOR NOW)")
		forceDump(t)
		return false
	}
	return false
}
