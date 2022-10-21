package toolkit

import "log"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"

func init() {
	log.Println("gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)
}

func (t Toolkit) commonChange(widget string) {
	s := t.String()
	log.Println("gui.Toolkit.ui.OnChanged() =", s)
	if (DebugToolkit) {
		log.Println("gui.Toolkit.ui.OnChanged() =", s)
	}
	if (t.OnChanged != nil) {
		log.Println("gui.Toolkit.OnChanged() trying to run toolkit.OnChanged() entered val =", s)
		t.OnChanged(&t)
		return
	}
	if (t.Custom != nil) {
		log.Println("gui.Toolkit.OnChanged() Running toolkit.Custom()")
		t.Dump()
		t.Custom()
		return
	}
	log.Println("gui.Toolkit.OnChanged() ENDED without finding any callback")
}

func (t Toolkit) broken() bool {
	if (t.uiBox == nil) {
		log.Println("gui.Toolkit.UiBox == nil. I can't add a widget without a place to put it")
		// log.Println("probably could just make a box here?")
		// corruption or something horrible?
		panic("wit/gui toolkit/andlabs func broken() invalid goroutine access into this toolkit?")
		return true
	}
	if (t.uiWindow == nil) {
		log.Println("gui.Toolkit.UiWindow == nil. I can't add a widget without a place to put it (IGNORING FOR NOW)")
		return false
	}
	return false
}
