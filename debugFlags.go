package gui

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func (n *Node) DebugFlags(makeWindow bool) {
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
		w = n.NewTab("Flags")
	}

	g = w.NewGroup("Show")

	g.NewButton("Dump Flags", func () {
		ShowDebugValues()
	})

	g.NewButton("All On", func () {
		SetDebug(true)
	})

	g.NewButton("All Off", func () {
		SetDebug(false)
	})

	g = w.NewGroup("List")
	g = g.NewGrid("flags grid", 2, 2)
	// generally useful debugging
	cb1 := g.NewCheckbox("debug Gui")
	g.NewLabel("like verbose=1")
	cb1.Custom = func() {
		debugGui = cb1.widget.B
		log(debugGui, "Custom() n.widget =", cb1.Name, cb1.widget.B)
	}

	// errors. by default these always output somewhere
	cbE := g.NewCheckbox("debug Error")
	g.NewLabel("(bad things. default=true)")
	cbE.Custom = func() {
		SetFlag("Error",  cbE.widget.B)
	}

	// debugging that will show you things like mouse clicks, user inputing text, etc
	// also set toolkit.DebugChange
	cb2 := g.NewCheckbox("debug Change")
	g.NewLabel("keyboard and mouse events")
	cb2.Custom = func() {
		SetFlag("Change", cb2.widget.B)
	}

	// supposed to tell if you are going to dump full variable output
	cb3 := g.NewCheckbox("debug Dump")
	g.NewLabel("show lots of output")
	cb3.Custom = func() {
		SetFlag("Dump",  cbE.widget.B)
	}

	cb4 := g.NewCheckbox("debug Tabs")
	g.NewLabel("tabs and windows")
	cb4.Custom = func() {
		SetFlag("Tabs",  cb4.widget.B)
	}

	cb6 := g.NewCheckbox("debug Node")
	g.NewLabel("the binary tree)")
	cb6.Custom = func() {
		SetFlag("Node",  cb6.widget.B)
	}

	// should show you when things go into or come back from the plugin
	cb5 := g.NewCheckbox("debug Plugin")
	g.NewLabel("plugin interaction)")
	cb5.Custom = func() {
		SetFlag("Plugin",  cb5.widget.B)
	}

	// turns on debugging inside the plugin toolkit
	cb7 := g.NewCheckbox("debug Toolkit")
	g.NewLabel("the plugin internals)")
	cb7.Custom = func() {
		// SetDebugToolkit(cb7.widget.B)
		SetFlag("Toolkit", cb7.widget.B)
	}
}
