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
	Config.Exit = StandardClose
	bugWin = NewWindow()
	bugWin.DebugTab("Debug Tab")
}

var checkd, checkdn, checkdt, checkdtk *Node

//////////////////////// debug flags //////////////////////////////////
func debugFlags(n *Node) {
	var df, checkdn, checkdd, changeCheckbox *Node
	df = n.NewGroup("Debug Flags")
	df.NewLabel("flags to control debugging output")

	cb := df.NewCheckbox("debugGui")
	cb.custom = func() {
		log(true, "custom ran correctly for =", n.Name)
		debugGui = true
	}

	checkdd = df.NewCheckbox("debugDump")
	checkdd.custom = func() {
		log(true, "debugDump() custom ran correctly for =", n.Name)
		debugDump = true
	}

	checkdn = df.NewCheckbox("debugNode")
	checkdn.custom = func() {
		log(true, "debugNode() custom ran correctly for =", n.Name)
		debugNode = true
	}


	cb = df.NewCheckbox("debugChange")
	cb.custom = func() {
		log(true, "checkbox: custom() ran correctly for =", cb.Name)
		log(true, "START debugChange =", debugChange)
		if (debugChange) {
			debugChange = false
			SetDebugChange(false)
			log(true, "debugToolkitChange turned off node.Name =", cb.Name)
		} else {
			debugChange = true
			SetDebugChange(true)
			log(true, "debugToolkitChange turned on Name =", cb.Name)
		}
		log(true, "END   debugChange =", debugChange)
	}

	cb = df.NewCheckbox("debugTabs")
	cb.custom = func() {
		log(true, "debugTabs() custom ran correctly for =", n.Name)
		debugTabs = true
	}

	cb = df.NewCheckbox("debugPlugin")
	cb.custom = func() {
		log(true, "debugPlugin() custom ran correctly for =", n.Name)
		debugPlugin = true
	}

	changeCheckbox = df.NewCheckbox("debugToolkit")
	changeCheckbox.custom = func() {
		SetDebugToolkit(true)
	}

	df.NewButton("Debug Toolkit", func() {
		if (debugToolkit) {
			SetDebugToolkit(false)
			log(true, "debugToolkit turned off node.Name =", n.Name)
		} else {
			SetDebugToolkit(true)
			log(true, "debugToolkit turned on Name =", n.Name)
		}
	})

	df.NewButton("Dump Debug Flags", func () {
		ShowDebugValues()
	})
}

func (n *Node) DebugTab(title string) *Node {
	var newN, gog, g1, g2, g3, dd *Node

	// time.Sleep(1 * time.Second)
	newN = n.NewTab(title)
	newN.Dump()

//////////////////////// main debug things //////////////////////////////////
	gog = newN.NewGroup("GOLANG")
	gog.NewLabel("go language")
	gog.NewButton("GO Language Debug", func () {
		GolangDebugWindow()
	})

	gog.NewLabel("wit/gui package")
	gog.NewButton("Demo toolkit andlabs/ui", func () {
		// DemoToolkitWindow()
	})

	debugFlags(newN)

//////////////////////// window debugging things //////////////////////////////////
	g1 = newN.NewGroup("Current Windows")
	dd = g1.NewDropdown("Window Dropdown")
	log(debugGui, "dd =", dd)

	// initialize the windows map if it hasn't been
	if (mapWindows == nil) {
		mapWindows = make(map[string]*Node)
	}

	var dump = false
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
		mapWindows[child.Name] = child
	}

	dd.SetDropdown(0)

	g2 = newN.NewGroup("Debug Window")
	g2.NewButton("SetMargined(tab)", func () {
		log(debugChange, "START SetMargined(tab)", g2.Name)
		// name := dd.GetText()
		name := dd.Widget.S
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
		log(debugChange, "END dd.Widget.S =", dd.Widget.S)
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
