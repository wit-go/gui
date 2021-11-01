package gui

import (
	"log"
	// "fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
//	"github.com/davecgh/go-spew/spew"
)

var names = make([]string, 100)
var nodeNames = make([]string, 100)

func DebugWindow() {
	Config.Title = "DebugWindow()"
	node := NewWindow()
	node.DebugTab("WIT GUI Debug Tab")
}

// TODO: remove this crap
// What does this actually do?
// It populates the nodeNames in a map. No, not a map, an array.
// What is the difference again? (other than one being in order and a predefined length)
func addNodeName(c *ui.Combobox, s string, id string) {
	c.Append(s)
	nodeNames[y] = id
	y = y + 1
}

func makeWindowDebug() *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	/////////////////////////////////////////////////////
	nodeBox := addGroup(hbox, "Windows:")
	nodeCombo := ui.NewCombobox()

	for name, node := range Data.NodeMap {
		if (Config.Debug) {
			log.Println("range Data.NodeMap() name =", name)
		}
		tmp := node.id + " (" + name + ")"
		addNodeName(nodeCombo, tmp, node.id)
	}
	nodeCombo.SetSelected(0)

	nodeBox.Append(nodeCombo, false)

	nodeCombo.OnSelected(func(*ui.Combobox) {
		y := nodeCombo.Selected()
		log.Println("y =", y)
		log.Println("nodeNames[y] =", nodeNames[y])
		node := Data.findId(nodeNames[y])
		if (node != nil) {
			node.Dump()
		}
	})

	/////////////////////////////////////////////////////
	vbox := addGroup(hbox, "Node Debug")

	n1 := addButton(vbox, "Data.DumpNodeMap()")
	n1.OnClicked(func(*ui.Button) {
		Data.DumpNodeMap()
	})

	n1 = addButton(vbox, "Data.ListChildren(false)")
	n1.OnClicked(func(*ui.Button) {
		Data.ListChildren(false)
	})

	n1 = addButton(vbox, "Data.ListChildren(true)")
	n1.OnClicked(func(*ui.Button) {
		Data.ListChildren(true)
	})

	n1 = addButton(vbox, "Node.Dump()")
	n1.OnClicked(func(*ui.Button) {
		y := nodeCombo.Selected()
		log.Println("y =", y)
		log.Println("nodeNames[y] =", nodeNames[y])
		node := Data.findId(nodeNames[y])
		if (node != nil) {
			node.Dump()
		}
	})

	n1 = addButton(vbox, "Node.ListChildren(false)")
	n1.OnClicked(func(*ui.Button) {
		y := nodeCombo.Selected()
		log.Println("y =", y)
		log.Println("nodeNames[y] =", nodeNames[y])
		node := Data.findId(nodeNames[y])
		if (node != nil) {
			node.ListChildren(false)
		}
	})

	n1 = addButton(vbox, "Node.ListChildren(true)")
	n1.OnClicked(func(*ui.Button) {
		y := nodeCombo.Selected()
		log.Println("y =", y)
		log.Println("nodeNames[y] =", nodeNames[y])
		node := Data.findId(nodeNames[y])
		if (node != nil) {
			node.ListChildren(true)
		}
	})

	n1 = addButton(vbox, "Node.AddDebugTab")
	n1.OnClicked(func(*ui.Button) {
		y := nodeCombo.Selected()
		log.Println("y =", y)
		log.Println("nodeNames[y] =", nodeNames[y])
		node := Data.findId(nodeNames[y])
		if (node != nil) {
			node.DebugTab("added this DebugTab")
		}
	})

	n1 = addButton(vbox, "Node.DemoAndlabsUiTab")
	n1.OnClicked(func(*ui.Button) {
		y := nodeCombo.Selected()
		log.Println("y =", y)
		log.Println("nodeNames[y] =", nodeNames[y])
		node := Data.findId(nodeNames[y])
		if (node != nil) {
			node.DemoAndlabsUiTab("ran gui.AddDemoAndlabsUiTab() " + strconv.Itoa(Config.counter))
		}
	})

	return hbox
}

// TODO: remove this crap
var x int = 0
var y int = 0

// TODO: remove this crap
func addName(c *ui.Combobox, s string) {
	c.Append(s)
	names[x] = s
	x = x + 1
}

func addGroup(b *ui.Box, name string) *ui.Box {
	group := ui.NewGroup(name)
	group.SetMargined(true)
	b.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	return vbox
}

func addButton(box *ui.Box, name string) *ui.Button {
	button := ui.NewButton(name)

	button.OnClicked(func(*ui.Button) {
		log.Println("Should do something here")
	})

	box.Append(button, false)
	return button
}

func (n *Node) DebugTab(title string) {
	newNode := n.AddTab(title, makeWindowDebug())
	if (Config.DebugNode) {
		newNode.Dump()
	}
	tabSetMargined(newNode.uiTab)
}

// This sets _all_ the tabs to Margin = true
//
// TODO: do proper tab tracking (will be complicated). low priority
func tabSetMargined(tab *ui.Tab) {
	if (Config.DebugTabs) {
		log.Println("tabSetMargined() IGNORE THIS")
	}
	c := tab.NumPages()
	for i := 0; i < c; i++ {
		if (Config.DebugTabs) {
			log.Println("tabSetMargined() i =", i)
		}
		tab.SetMargined(i, true)
	}
}
