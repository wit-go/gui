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
	var newN, gog, g1, g2, g3, dd, junk, newThing *Node

	// time.Sleep(1 * time.Second)
	newN = n.NewTab(title)
	newN.Dump()

//////////////////////// main debug things //////////////////////////////////
	gog = newN.NewGroup("GOLANG")
	gog.NewLabel("go language")
	gog.NewButton("GO Language Debug", func () {
		newN.GolangDebugWindow(false)
	})
	gog.NewButton("Debug Flags", func () {
		newN.debugFlags(false)
	})
	gog.NewButton("Debug Widgets", func () {
		newN.debugWidgets(false)
	})

	gog.NewLabel("wit/gui package")
	gog.NewButton("Demo toolkit andlabs/ui", func () {
		// DemoToolkitWindow()
	})

	junk = gog.NewButton("junk", func () {
		log("click junk, get junk")
	})

	gog.NewLabel("tmp label")


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
	dd.NewButton("Delete(junk)", func () {
		Delete(junk)
	})
	dd.NewButton("myButton", func () {
		gog.NewButton("myButton", func () {
			log("this code is better")
		})
	})
	dd.NewButton("add Hope", func () {
		var i int = 1
		log("add hope?", i)
		gog.NewButton("hope", func () {
			i += 1
			log("write better code", i)
		})
	})
	dd.NewButton("add newThing", func () {
		var i, j int = 1, 1
		newThing = gog.NewThing("NewThing")
		newThing.Custom = func() {
			f := i + j
			log("newThing!!! n.widget =", newThing.widget.Name, newThing.widget.B, f)
			j = i
			i = f
		}
		log("newThing!!! n.widget")
	})

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
