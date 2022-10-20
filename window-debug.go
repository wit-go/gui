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

func (n *Node) DebugTab(title string) *Node {
	var newN, gog, g1, g2, g3, dd *Node

	// time.Sleep(1 * time.Second)
	newN = n.NewTab(title)
	newN.Dump()

	gog = newN.NewGroup("GOLANG")
	gog.NewLabel("go language")
	gog.AddButton("GO Language Debug", func (*Node) {
		GolangDebugWindow()
	})

	gog.NewLabel("wit/gui package")
	gog.AddButton("WIT/GUI Package Debug", func (*Node) {
		Config.Width = 640
		Config.Height = 480
		Queue(DebugWindow)
	})
	gog.AddButton("Demo wit/gui", func (*Node) {
		DemoWindow()
	})
	gog.AddButton("Demo toolkit andlabs/ui", func (*Node) {
		DemoToolkitWindow()
	})

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
		dd.AddDropdown(child.Name)
	}
	dd.SetDropdown(0)

	g2 = newN.NewGroup("Debug Window")
	g2.AddButton("SetMargined(tab)", func (*Node) {
		log.Println("\tSTART")
		name := dd.GetText()
		log.Println("\tENDed with", name)
		// gw.UiTab.SetMargined(*gw.TabNumber, true)
	})
	g2.AddButton("Hide(tab)", func (*Node) {
		// gw.UiTab.Hide()
	})
	g2.AddButton("Show(tab)", func (*Node) {
		// gw.UiTab.Show()
	})
	g2.AddButton("Delete(tab)", func (*Node) {
		// gw.UiTab.Delete(*gw.TabNumber)
	})
	g2.AddButton("change Title", func (*Node) {
		// mainWindow.SetText("hello world")
	})

	/////////////////////////////////////////////////////
	g3 = newN.NewGroup("Node Debug")

	g3.AddButton("Node.Dump()", func (n *Node) {
		n.Dump()
	})
	g3.AddButton("Node.ListChildren(false)", func (n *Node) {
		n.ListChildren(false)
	})
	g3.AddButton("Node.ListChildren(true)", func (n *Node) {
		n.ListChildren(true)
	})
	g3.AddButton("AddDebugTab()", func (n *Node) {
		if (bugWin != nil) {
			bugWin.DebugTab("added this DebugTab")
		}
	})

	return newN
}
