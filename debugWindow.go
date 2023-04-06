package gui

import (
//	"git.wit.org/wit/gui/toolkit"
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

//////////////////////// main debug things //////////////////////////////////
	gog = newN.NewGroup("Debugging Windows:")

	// generally useful debugging
	cb := gog.NewCheckbox("Seperate windows")
	cb.Custom = func() {
		makeTabs = cb.widget.B
		log(debugGui, "Custom() n.widget =", cb.Name, cb.widget.B)
	}
	makeTabs = false
	cb.Set(false)

	gog.NewButton("Debug Flags", func () {
		bugWin.DebugFlags(makeTabs)
	})
	gog.NewButton("Debug Widgets", func () {
		DebugWidgetWindow(newN)
	})
	gog.NewButton("GO Language Internals", func () {
		bugWin.DebugGolangWindow(makeTabs)
	})
	gog.NewButton("GO Channels debug", func () {
		bugWin.DebugGoChannels(makeTabs)
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

	g2 := newN.NewGroup("more things")

	g2.NewButton("Node.ListChildren(true)", func () {
		if (activeWidget == nil) {
			activeWidget = Config.rootNode
		}
		activeWidget.ListChildren(true)
	})

	g2.NewButton("test conc", func () {
		makeConc()
	})

	g2.NewButton("List Plugins", func () {
		for _, aplug := range allPlugins {
			log("Loaded plugin:", aplug.name, aplug.filename)
		}
	})

	g2.NewButton("load plugin 'gocui'", func () {
		StartS("gocui")
	})

	g2.NewButton("Redraw(gocui)", func () {
		Redraw("gocui")
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
	if (activeWidget == nil) {
		// the debug window doesn't exist yet so you can't display the change
		// TODO: make a fake binary tree for this(?)
		return
	}

	// var last = ""
	for _, child := range Config.rootNode.children {
		log(debugGui, "\t\t", child.id, child.Width, child.Height, child.Name)
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

	// log("dumpWidget() ", b, listChildrenDepth, defaultPadding, n.id, info)

	var addDropdowns func (*Node)
	addDropdowns = func (n *Node) {
		s := n.dumpWidget(true)
		dd.AddDropdownName(s)
		mapWindows[s] = n

		for _, child := range n.children {
			listChildrenDepth += 1
			addDropdowns(child)
			listChildrenDepth -= 1
		}
	}

	// list everything in the binary tree
	addDropdowns(Config.rootNode)
}
