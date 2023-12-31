package gui

import 	(
	newlog "go.wit.com/log"
)

type LogSettings struct {
	ready	bool
	hidden	bool
	err	error
	name	string

	parent	*Node // should be the root of the 'gui' package binary tree
	window	*Node // our window for displaying the log package settings
	group	*Node //
	grid	*Node //
	checkbox *Node
	label *Node

}

func (ls *LogSettings) Set(b bool) {
	newlog.Set(ls.name, b)
	ls.checkbox.Set(b)
}

func (p *Node) NewLogFlag(name string) *LogSettings {
	ls := new(LogSettings)
	ls.parent = p
	ls.ready = false
	ls.name = name

	ls.checkbox = p.NewCheckbox(name)
	ls.label = p.NewLabel("Enable log." + name)
	ls.checkbox.Set(newlog.Get(name))
	ls.checkbox.Custom = func() {
		newlog.Set(name, ls.checkbox.B)
	}
	return ls
}

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func (n *Node) DebugFlags(makeWindow bool) {
	var w, g *Node

	logGadgets := make(map[string]*LogSettings)

	// Either:
	// make a new window
	// make a new tab in the existing window
	if (makeWindow) {
		w = me.rootNode.NewWindow("Debug Flags")
		w.Custom = w.StandardClose
		w = w.NewBox("hBox", true)
	} else {
		w = n.NewTab("Flags")
	}

	g = w.NewGroup("Show").Pad()

	g.NewButton("log.SetTmp()", func () {
		newlog.SetTmp()
	})

	g.NewButton("log.All(true)", func () {
		for _, lf := range logGadgets {
			lf.Set(true)
		}
	})

	g.NewButton("log.All(false)", func () {
		for _, lf := range logGadgets {
			lf.Set(false)
		}
	})

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

	logGadgets["INFO"] = g.NewLogFlag("INFO")
	logGadgets["WARN"] = g.NewLogFlag("WARN")
	logGadgets["SPEW"] = g.NewLogFlag("SPEW")
	logGadgets["ERROR"] = g.NewLogFlag("ERROR")

	// generally useful debugging
	cb1 := g.NewCheckbox("debug Gui")
	g.NewLabel("like verbose=1")
	cb1.Custom = func() {
		debugGui = cb1.B
		log(debugGui, "Custom() n.widget =", cb1.Name, cb1.B)
	}

	// turns on debugging inside the plugin toolkit
	cb7 := g.NewCheckbox("debug Toolkit")
	g.NewLabel("the plugin internals)")
	cb7.Custom = func() {
		// SetDebugToolkit(cb7.B)
		SetFlag("Toolkit", cb7.B)
	}
}
