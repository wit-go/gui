package gui

import (
	"git.wit.org/wit/gui/toolkit"
)

// TODO: move all this shit into somewhere not global

// main debugging window
var bugWin *Node
// if there should be new windows or just tabs
var makeTabs bool = true

var mapWindows map[string]*Node
var checkd, checkdn, checkdt, checkdtk, lb1, lb2 *Node
var myButton *Node

/*
	Creates a window helpful for debugging this package
*/
func DebugWindow() {
	Config.Title = "go.wit.org/gui debug window"
	Config.Width = 300
	Config.Height = 200
	bugWin = NewWindow()
	bugWin.Custom = bugWin.StandardClose
	bugWin.DebugTab("Debug Tab")
}

func (n *Node) DebugTab(title string) *Node {
	var newN, gog, g1 *Node

	// time.Sleep(1 * time.Second)
	newN = n.NewTab(title)
	newN.Dump()

//////////////////////// main debug things //////////////////////////////////
	gog = newN.NewGroup("Debugging Windows:")

	// generally useful debugging
	cb := gog.NewCheckbox("Seperate windows")
	cb.Custom = func() {
		makeTabs = cb.widget.B
		log(debugGui, "Custom() n.widget =", cb.widget.Name, cb.widget.B)
	}
	makeTabs = false
	cb.Set(false)

	gog.NewButton("Debug Flags", func () {
		newN.DebugFlags(makeTabs)
	})
	gog.NewButton("Debug Widgets", func () {
		DebugWidgetWindow(newN)
	})
	gog.NewButton("GO Language Internals", func () {
		newN.DebugGolangWindow(makeTabs)
	})
	gog.NewButton("GO Channels debug", func () {
		newN.DebugGoChannels(makeTabs)
	})

	gog.NewLabel("Force Quit:")

	gog.NewButton("os.Exit()", func () {
		exit()
	})

//////////////////////// window debugging things //////////////////////////////////
	g1 = newN.NewGroup("list things")

	g1.NewButton("List Windows", func () {
		dropdownWindow(g1)
	})
	g1.NewButton("List Window Widgets", func () {
		dropdownWindowWidgets(g1)
	})

	g2 := newN.NewGroup("node things")

	g2.NewButton("Node.ListChildren(false)", func () {
		g := debugGui
		d := debugDump
		debugGui = true
		debugDump = true
		activeWidget.ListChildren(false, nil, nil)
		debugGui = g
		debugDump = d
	})
	g2.NewButton("Node.ListChildren(true)", func () {
		g := debugGui
		d := debugDump
		debugGui = true
		debugDump = true
		activeWidget.ListChildren(true, nil, nil)
		debugGui = g
		debugDump = d
	})

	return newN
}

func dropdownWindow(p *Node) {
	var mapWindows map[string]*Node
	mapWindows = make(map[string]*Node)

	dd := p.NewDropdown("Window Dropdown")
	dd.Custom = func() {
		name := dd.widget.S
		activeWidget = mapWindows[name]
		setActiveWidget(activeWidget)
		log("The Window was set to", name)
	}
	log(debugGui, "dd =", dd)

	// var last = ""
	for _, child := range Config.master.children {
		log(debugGui, "\t\t", child.id, child.Width, child.Height, child.Name)
		// skip the fake "Flag" node
		if (child.widget.Type == toolkit.Flag) {
			continue
		}
		dd.AddDropdownName(child.Name)
		// last = child.Name
		mapWindows[child.Name] = child
		if (activeWidget == nil) {
			activeWidget = child
		}
	}
	// dd.SetDropdownName(last)
}

func dropdownWindowWidgets(p *Node) {
	var mapWindows map[string]*Node
	mapWindows = make(map[string]*Node)

	dd := p.NewDropdown("Window Widgets Dropdown")
	dd.Custom = func() {
		name := dd.widget.S
		activeWidget = mapWindows[name]
		setActiveWidget(activeWidget)
	}
	log(debugGui, "dd =", dd)

	activeWidget.ListChildren(true, dd, mapWindows)
}
