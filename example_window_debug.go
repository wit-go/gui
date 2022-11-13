package gui

import (
	"log"
)

var names = make([]string, 100)
var nodeNames = make([]string, 100)

var bugWin *Node
/*
	Creates a window helpful for debugging this package
*/
func DebugWindow() {
	Config.Title = "git.wit.org/wit/gui debug fixme"
	Config.Width = 300
	Config.Height = 200
	Config.Exit = StandardClose
	bugWin = NewWindow()
	bugWin.DebugTab("WIT GUI Debug Tab")
}

// this function is used by the examples to add a tab
// dynamically to the bugWin node
// TODO: make this smarter once this uses toolkit/
func DebugTab() {
	if (bugWin == nil) {
		log.Println("Not sure what window to add this to? Use node.DebugTab() instead")
		return;
	}
	bugWin.DebugTab("does this work?")
}

var checkd, checkdn, checkdt, checkdtk *Node

//////////////////////// debug flags //////////////////////////////////
func debugFlags(n *Node) {
	var df, checkd, checkdn, checkdd, changeCheckbox *Node
	df = n.NewGroup("Debug Flags")
	df.NewLabel("flags to control debugging output")

	checkd = df.NewCheckbox("Debug")
	checkd.OnChanged = func(*Node) {
		// checkd.checked = checkd.toolkit.Checked()
		Config.Debug.Debug = true
		if (Config.Debug.Debug) {
			log.Println("Debug turned on")
		} else {
			log.Println("Debug turned off")
		}
	}

	checkdn = df.NewCheckbox("Debug Node")
	checkdn.OnChanged = func(*Node) {
		Config.Debug.Node = true
	}

	checkdd = df.NewCheckbox("Debug node.Dump()")
	checkdd.OnChanged = func(*Node) {
		Config.Debug.Dump = true
	}

	changeCheckbox = df.NewCheckbox("Debug Change")
	changeCheckbox.OnChanged = func(*Node) {
		Config.Debug.Change = true
	}

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
		// GolangDebugWindow()
	})

	gog.NewLabel("wit/gui package")
	gog.NewButton("WIT/GUI Package Debug", func () {
		Config.Width = 640
		Config.Height = 480
		// Queue(DebugWindow)
	})
	gog.NewButton("Demo wit/gui", func () {
		// DemoWindow()
	})
	gog.NewButton("Demo toolkit andlabs/ui", func () {
		// DemoToolkitWindow()
	})

	debugFlags(newN)

//////////////////////// window debugging things //////////////////////////////////
	g1 = newN.NewGroup("Current Windows")
	dd = g1.NewDropdown("Window Dropdown")
	log.Println("dd =", dd)

	var dump = false
	for _, child := range Config.master.children {
		log.Println("\t\t", child.id, child.Width, child.Height, child.Name)
		if (child.parent != nil) {
			log.Println("\t\t\tparent =",child.parent.id)
		} else {
			log.Println("\t\t\tno parent")
			panic("no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		dd.AddDropdownName(child.Name)
	}
	dd.SetDropdown(0)

	g2 = newN.NewGroup("Debug Window")
	g2.NewButton("SetMargined(tab)", func () {
		log.Println("\tSTART")
		name := dd.GetText()
		log.Println("\tENDed with", name)
		// gw.UiTab.SetMargined(*gw.TabNumber, true)
	})
	g2.NewButton("Hide(tab)", func () {
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
		bugWin.Dump()
	})
	g3.NewButton("Node.ListChildren(false)", func () {
		bugWin.ListChildren(false)
	})
	g3.NewButton("Node.ListChildren(true)", func () {
		bugWin.ListChildren(true)
	})
	g3.NewButton("AddDebugTab()", func () {
		if (bugWin != nil) {
			bugWin.DebugTab("added this DebugTab")
		}
	})

	return newN
}
