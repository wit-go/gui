package gui

import (
	"log"
	"fmt"
	// "time"

	// "github.com/davecgh/go-spew/spew"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

type Element int

// https://ieftimov.com/post/golang-datastructures-trees/
const (
	Unknown Element = iota
	Window
	Tab
	Box
	Label
	Combo
)

func (s Element) String() string {
	switch s {
	case Window:
		return "window"
	case Tab:
		return "tab"
	case Box:
		return "box"
	case Label:
		return "label"
	case Combo:
		return "combo"
	}
	return "unknown"
}

type Node struct {
	id     string
	Name   string
	Width  int
	Height int

	parent	*Node
	children []*Node

	window	*GuiWindow
	box	*GuiBox

	uiControl  *ui.Control
	uiWindow  *ui.Window
	uiTab  *ui.Tab
	uiBox  *ui.Box
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Window() *Node {
	return n.parent
}

func (n *Node) Dump() {
	log.Println("gui.Node.Dump() id         = ", n.id)
	log.Println("gui.Node.Dump() Name       = ", n.Name)
	log.Println("gui.Node.Dump() Width      = ", n.Width)
	log.Println("gui.Node.Dump() Height     = ", n.Height)

	if (n.parent == nil) {
		log.Println("gui.Node.Dump() parent     = nil")
	} else {
		log.Println("gui.Node.Dump() parent     = ", n.parent.id)
	}
	log.Println("gui.Node.Dump() children   = ", n.children)

	log.Println("gui.Node.Dump() window     = ", n.window)
	log.Println("gui.Node.Dump() box        = ", n.box)

	log.Println("gui.Node.Dump() uiWindow   = ", n.uiWindow)
	log.Println("gui.Node.Dump() uiTab      = ", n.uiTab)
	log.Println("gui.Node.Dump() uiBox      = ", n.uiBox)
	log.Println("gui.Node.Dump() uiControl  = ", n.uiControl)
	if (n.id == "") {
		panic("gui.Node.Dump() id == nil")
	}
}


func (n *Node) SetBox(box *GuiBox) {
	n.box = box
}

func (n *Node) SetName(name string) {
	// n.uiType.SetName(name)
	if (n.uiWindow != nil) {
		log.Println("node is a window. setting title =", name)
		n.uiWindow.SetTitle(name)
		return
	}
	log.Println("*ui.Control =", n.uiControl)
	return
}

func (n *Node) Append(child *Node) {
	//	if (n.UiBox == nil) {
	//		return
	//	}
	n.children = append(n.children, child)
	if (Config.Debug) {
		log.Println("child node:")
		child.Dump()
		log.Println("parent node:")
		n.Dump()
	}
	// time.Sleep(3 * time.Second)
}

func (n *Node) List() {
	findByIdDFS(n, "test")
}

func (n *Node) ListChildren(dump bool) {
	log.Println("\tListChildren() node =", n.id, n.Name, n.Width, n.Height)

	if (dump == true) {
		n.Dump()
	}
	if len(n.children) == 0 {
		if (n.parent != nil) {
			log.Println("\t\t\tparent =",n.parent.id)
		}
		log.Println("\t\t", n.id, "has no children")
		return
	}
	for _, child := range n.children {
		log.Println("\t\tListChildren() child =",child.id,  child.Name, child.Width, child.Height)
		if (child.parent != nil) {
			log.Println("\t\t\tparent =",child.parent.id)
		} else {
			log.Println("\t\t\tno parent")
			panic("no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		if (child.children == nil) {
			log.Println("\t\t", child.id, "has no children")
		} else {
			log.Println("\t\t\tHas children:", child.children)
		}
		child.ListChildren(dump)
	}
	return
}

// The parent Node needs to be the raw Window
// The 'stuff' Node needs to be the contents of the tab
//
// This function should make a new node with the parent and
// the 'stuff' Node as a child
func (parent *Node) AddTabNode(title string, b *GuiBox) *Node {
	// Ybox := gui.NewBox(box, gui.Yaxis, "Working Stuff")
	// var baseControl ui.Control
	// baseControl = Ybox.UiBox
	// return baseControl

	var newNode *Node
	// var newControl ui.Control

	/*
	if (parent.box == nil) {
		// TODO: fix this to use a blank box
		// uiC := parent.initBlankWindow()
		hbox := ui.NewHorizontalBox()
		hbox.SetPadded(true)
		newNode.uiBox = hbox
		panic("node.AddTabNode() can not add a tab if the box == nil")
	}
	if (parent.uiTab == nil) {
		panic("node.AddTabNode() can not add a tab if parent.uiTab == nil")
	}
	*/

	newNode = parent.makeNode(title, 444, 400 + Config.counter)
	newNode.uiTab = parent.uiTab
	newNode.box = b

	/*
	newControl = b.UiBox
	newNode.uiTab.Append(title, newControl)
	*/
	newNode.uiTab.Append(title, b.UiBox)

	fmt.Println("")
	log.Println("parent:")
	parent.Dump()

	fmt.Println("")
	log.Println("newNode:")
	newNode.Dump()

	// panic("node.AddTabNode()")

	return newNode
}

func (parent *Node) AddTab(title string, uiC ui.Control) *Node {
	log.Println("gui.Node.AddTab() START name =", title)
	if parent.uiWindow == nil {
		parent.Dump()
		panic("gui.AddTab() ERROR ui.Window == nil")
	}
	if parent.box == nil {
		parent.Dump()
		panic("gui.AddTab() ERROR box == nil")
	}
	if parent.uiTab == nil {
		inittab := ui.NewTab() // no, not that 'inittab'
		parent.uiWindow.SetChild(inittab)
		parent.uiWindow.SetMargined(true)
		parent.uiTab = inittab

		// parent.Dump()
		// panic("gui.AddTab() ERROR uiTab == nil")
	}

	tab := parent.uiTab
	parent.uiWindow.SetMargined(true)

	if (uiC == nil) {
		hbox := ui.NewHorizontalBox()
		hbox.SetPadded(true)
		uiC = hbox
	}
	tab.Append(title, uiC)
	tab.SetMargined(0, true)

	// panic("gui.AddTab() before makeNode()")
	newNode := parent.makeNode(title, 555, 600 + Config.counter)
	newNode.uiTab = tab
	// panic("gui.AddTab() after makeNode()")
	return newNode
}
