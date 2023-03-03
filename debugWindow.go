package gui

var bugWin *Node
var mapWindows map[string]*Node

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

var checkd, checkdn, checkdt, checkdtk, lb1, lb2 *Node
var myButton *Node

func (n *Node) DebugTab(title string) *Node {
	var newN, gog, g1, g2, g3, dd *Node

	// time.Sleep(1 * time.Second)
	newN = n.NewTab(title)
	newN.Dump()

//////////////////////// main debug things //////////////////////////////////
	gog = newN.NewGroup("Debugging")

	gog.NewButton("Debug Flags", func () {
		newN.debugFlags(false)
	})
	gog.NewButton("Debug Widgets", func () {
		newN.debugWidgets(false)
	})
	gog.NewButton("GO Language Internals", func () {
		newN.debugGolangWindow(false)
	})
	gog.NewButton("GO Channels debug", func () {
		newN.debugGoChannels(false)
	})

//////////////////////// window debugging things //////////////////////////////////
	g1 = newN.NewGroup("Current Windows")
	dd = g1.NewDropdown("Window Dropdown")
	dd.Custom = func() {
		name := dd.widget.S
		bugWin = mapWindows[name]
		log("The Window was set to", name)
	}
	log(debugGui, "dd =", dd)

	// initialize the windows map if it hasn't been
	if (mapWindows == nil) {
		mapWindows = make(map[string]*Node)
	}

	var dump = false
	var last = ""
	for _, child := range Config.master.children {
		log(debugGui, "\t\t", child.id, child.Width, child.Height, child.Name)
		if (child.parent != nil) {
			log(debugGui, "\t\t\tparent =",child.parent.id)
		} else {
			log(debugGui, "\t\t\tno parent")
			panic("no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		dd.AddDropdownName(child.Name)
		last = child.Name
		mapWindows[child.Name] = child
	}
	dd.SetDropdownName(last)

	g2 = newN.NewGroup("Debug Window")
	g2.NewButton("SetMargined(tab)", func () {
		log(debugChange, "START SetMargined(tab)", g2.Name)
		// name := dd.GetText()
		name := dd.widget.S
		log(true, "name =", name)
		log(debugChange, "name =", name)
		log(debugChange, "mapWindows[name] =", mapWindows[name])
		/*
		for s, n := range mapWindows {
			log(debugChange, "\tname =", name)
			log(debugChange, "\tmapWindows s =", s)
			log(debugChange, "\tmapWindows[s] =", n)
		}
		*/
		bugWin = mapWindows[name]
		log(debugChange, "END dd.widget.S =", dd.widget.S)
		// gw.UiTab.SetMargined(*gw.TabNumber, true)
	})
	g2.NewButton("Hide(tab)", func () {
		log(debugChange, "\tclick() START", g2.Name)
		// gw.UiTab.Hide()
	})
	g2.NewButton("Show(tab)", func () {
		// gw.UiTab.Show()
	})
	g2.NewButton("Delete(tab)", func () {
		// gw.UiTab.Delete(*gw.TabNumber)
	})
	g2.NewButton("change Title", func () {
		// mainWindow.SetText("hello world")
	})
	g2.NewButton("Quit", func () {
		exit()
	})

	/////////////////////////////////////////////////////
	g3 = newN.NewGroup("Node Debug")

	g3.NewButton("Node.Dump()", func () {
		debugGui = true
		debugDump = true
		bugWin.Dump()
	})
	g3.NewButton("Node.ListChildren(false)", func () {
		debugGui = true
		debugDump = true
		bugWin.ListChildren(false)
	})
	g3.NewButton("Node.ListChildren(true)", func () {
		debugGui = true
		debugDump = true
		bugWin.ListChildren(true)
	})

	return newN
}
