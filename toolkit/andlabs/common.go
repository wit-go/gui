package main

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func init() {
	// Can you pass values to a plugin init() ? Otherwise, there is no way to safely print
	// log(debugToolkit, "gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)
}

func (t andlabsT) commonChange(widget string) {
	s := t.String()
	log(debugChange, "commonChange() START widget   =", widget)
	log(debugChange, "commonChange() t.String =", s)
	if (t.OnChanged != nil) {
		// log(debugChange, "commonChange() toolkit.OnChanged() START")
		// t.OnChanged(&t)
		exit("OnChanged is not implemented. TODO: FIX THIS")
		return
	}
	if (t.Custom != nil) {
		log(debugChange, "commonChange() START toolkit.Custom()")
		t.Custom()
		log(debugChange, "commonChange() END toolkit.Custom()")
		return
	}
	if (widget == "Checkbox") {
		log(debugChange, "commonChange() END Need to read the Checkbox value")
		return
	}
	if (widget == "Dropdown") {
		t.getDropdown()
		if (t.tw == nil) {
			log(debugChange, "commonChange() END tw.Custom == nil")
		}
		if (t.tw.Custom == nil) {
			log(debugChange, "commonChange() END Dropdown (no custom())")
		}
		t.tw.Custom()
		log(debugChange, "commonChange() END Dropdown")
		return
	}
	log(debugChange, "commonChange() t.String =", s)
	log(debugChange, "commonChange() ENDED without finding any callback")
}

func (t *andlabsT) getDropdown() {
	log(debugChange, "commonChange() Need to read the dropdown menu")
	if (t.uiCombobox == nil) {
		log(debugChange, "commonChange() END BAD NEWS. t.uiCombobox == nil")
		return
	}
	i := t.uiCombobox.Selected()
	log(debugChange, "commonChange() t.uiCombobox = ", i)
	if (t.tw == nil) {
		log(debugChange, "commonChange() END tw = nil")
		return
	}
	t.tw.S = t.String()
	log(debugChange, "commonChange() END tw = ", t.tw)
	return
}

// does some sanity checks on the internal structs of the binary tree
// TODO: probably this should not panic unless it's running in devel mode (?)
func (t *andlabsT) broken() bool {
	if (t.uiBox == nil) {
		if (t.uiWindow != nil) {
			log(debugToolkit, "gui.Toolkit.UiBox == nil. This is an empty window. Try to add a box")
			t.NewBox()
			return false
		}
		log(debugToolkit, "gui.Toolkit.UiBox == nil. I can't add a widget without a place to put it")
		// log(debugToolkit, "probably could just make a box here?")
		// corruption or something horrible?
		panic("wit/gui toolkit/andlabs func broken() invalid goroutine access into this toolkit?")
		panic("wit/gui toolkit/andlabs func broken() this probably should not cause the app to panic here (?)")
		return true
	}
	if (t.uiWindow == nil) {
		log(debugToolkit, "gui.Toolkit.UiWindow == nil. I can't add a widget without a place to put it (IGNORING FOR NOW)")
		forceDump(t)
		return false
	}
	return false
}
