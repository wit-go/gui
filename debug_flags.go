package gui

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func (n *Node) debugFlags(makeWindow bool) {
	var w, g *Node

	// Either:
	// make a new window
	// make a new tab in the existing window
	if (makeWindow) {
		Config.Title = "Debug Flags"
		Config.Width = 300
		Config.Height = 400
		w = NewWindow()
		w.Custom = w.StandardClose
	} else {
		w = n.NewTab("Debug Flags")
	}
	w.Dump()

	g = w.NewGroup("Debug Flags")

	g.NewButton("Turn on all Debug Flags", func () {
		SetDebug(true)
	})

	g.NewButton("Turn off all Debug Flags", func () {
		SetDebug(false)
	})

	// generally useful debugging
	cb1 := g.NewCheckbox("debugGui")
	cb1.Custom = func() {
		debugGui = cb1.widget.B
		log(debugGui, "Custom() n.widget =", cb1.widget.Name, cb1.widget.B)
	}

	// debugging that will show you things like mouse clicks, user inputing text, etc
	// also set toolkit.DebugChange
	cb2 := g.NewCheckbox("debugChange")
	cb2.Custom = func() {
		debugChange = cb2.widget.B
		SetDebugChange(cb2.widget.B)
		log(debugGui, "Custom() n.widget =", cb2.widget.Name, cb2.widget.B)
	}

	// supposed to tell if you are going to dump full variable output
	cb3 := g.NewCheckbox("debugDump")
	cb3.Custom = func() {
		debugDump = cb3.widget.B
		log(debugGui, "Custom() n.widget =", cb3.widget.Name, cb3.widget.B)
	}

	cb4 := g.NewCheckbox("debugTabs")
	cb4.Custom = func() {
		debugTabs = cb4.widget.B
		log(debugGui, "Custom() n.widget =", cb4.widget.Name, cb4.widget.B)
	}

	// should show you when things go into or come back from the plugin
	cb5 := g.NewCheckbox("debugPlugin")
	cb5.Custom = func() {
		debugPlugin = cb5.widget.B
		log(debugGui, "Custom() n.widget =", cb5.widget.Name, cb5.widget.B)
	}

	cb6 := g.NewCheckbox("debugNode")
	cb6.Custom = func() {
		debugNode = cb6.widget.B
		log(debugGui, "Custom() n.widget =", cb6.widget.Name, cb6.widget.B)
	}

	// turns on debugging inside the plugin toolkit
	cb7 := g.NewCheckbox("debugToolkit")
	cb7.Custom = func() {
		SetDebugToolkit(cb7.widget.B)
		log(debugGui, "Custom() n.widget =", cb7.widget.Name, cb7.widget.B)
	}

	g.NewButton("Dump Debug Flags", func () {
		ShowDebugValues()
	})
}
