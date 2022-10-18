package gui

import (
	"log"
	"fmt"
//	"reflect"

	// "github.com/davecgh/go-spew/spew"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

)

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

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

// The Node is simply the name and the size of whatever GUI element exists
type Node struct {
	id     string

	Name   string
	Width  int
	Height int

	parent	*Node
	children []*Node

	window	*GuiWindow
	box	*GuiBox
	custom    func(*Node)
	OnChanged func(*Node)

	Toolkit	*toolkit.Toolkit

	uiControl *ui.Control
	uiButton  *ui.Button
	uiGroup   *ui.Group
	uiSlider  *ui.Slider
	uiSpinbox *ui.Spinbox
	uiWindow  *ui.Window
	uiTab  *ui.Tab
	uiBox  *ui.Box
	uiText *ui.EditableCombobox
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Window() *Node {
	return n.parent
}

func (n *Node) Dump() {
	IndentPrintln("id         = ", n.id)
	IndentPrintln("Name       = ", n.Name)
	IndentPrintln("Width      = ", n.Width)
	IndentPrintln("Height     = ", n.Height)

	if (n.parent == nil) {
		IndentPrintln("parent     = nil")
	} else {
		IndentPrintln("parent     =", n.parent.id)
	}
	if (n.children != nil) {
		IndentPrintln("children   = ", n.children)
	}

	if (n.window != nil) {
		IndentPrintln("window     = ", n.window)
	}
	if (n.box != nil) {
		IndentPrintln("box        = ", n.box)
	}

	if (n.uiWindow != nil) {
		IndentPrintln("uiWindow   = ", n.uiWindow)
	}
	if (n.uiTab != nil) {
		IndentPrintln("uiTab      = ", n.uiTab)
	}
	if (n.uiBox != nil) {
		IndentPrintln("uiBox      = ", n.uiBox)
	}
	if (n.Toolkit != nil) {
		IndentPrintln("Toolkit    = ", n.Toolkit)
		n.Toolkit.Dump()
	}
	if (n.uiControl != nil) {
		IndentPrintln("uiControl  = ", n.uiControl)
	}
	if (n.uiButton != nil) {
		IndentPrintln("uiButton   = ", n.uiButton)
	}
	if (n.custom != nil) {
		IndentPrintln("custom     = ", n.custom)
	}
	if (n.OnChanged != nil) {
		IndentPrintln("OnChanged  = ", n.OnChanged)
	}
	if (n.id == "") {
		// Node structs should never have a nil id.
		// I probably shouldn't panic here, but this is just to check the sanity of
		// the gui package to make sure it's not exiting
		panic("gui.Node.Dump() id == nil TODO: make a unigue id here in the golang gui library")
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

var listChildrenParent *Node
var listChildrenDepth int = 0
var defaultPadding = "  "

func IndentPrintln(a ...interface{}) {
	indentPrintln(listChildrenDepth, defaultPadding, a)
}

func indentPrintln(depth int, format string, a ...interface{}) {
	var tabs string
	for i := 0; i < depth; i++ {
		tabs = tabs + format
	}

	// newFormat := tabs + strconv.Itoa(depth) + " " + format
	newFormat := tabs + format
	log.Println(newFormat, a)
}

func (n *Node) ListChildren(dump bool) {
	indentPrintln(listChildrenDepth, defaultPadding, n.id, n.Width, n.Height, n.Name)

	if (dump == true) {
		n.Dump()
	}
	if len(n.children) == 0 {
		if (n.parent == nil) {
		} else {
			if (Config.DebugNode) {
				log.Println("\t\t\tparent =",n.parent.id)
			}
			if (listChildrenParent != nil) {
				if (Config.DebugNode) {
					log.Println("\t\t\tlistChildrenParent =",listChildrenParent.id)
				}
				if (listChildrenParent.id != n.parent.id) {
					log.Println("parent.child does not match child.parent")
					panic("parent.child does not match child.parent")
				}
			}
		}
		if (Config.DebugNode) {
			log.Println("\t\t", n.id, "has no children")
		}
		return
	}
	for _, child := range n.children {
		// log.Println("\t\t", child.id, child.Width, child.Height, child.Name)
		if (child.parent != nil) {
			if (Config.DebugNode) {
				log.Println("\t\t\tparent =",child.parent.id)
			}
		} else {
			log.Println("\t\t\tno parent")
			panic("no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		if (Config.DebugNode) {
			if (child.children == nil) {
				log.Println("\t\t", child.id, "has no children")
			} else {
				log.Println("\t\t\tHas children:", child.children)
			}
		}
		listChildrenParent = n
		listChildrenDepth += 1
		child.ListChildren(dump)
		listChildrenDepth -= 1
	}
	return
}

// The parent Node needs to be the raw Window
// The 'stuff' Node needs to be the contents of the tab
//
// This function should make a new node with the parent and
// the 'stuff' Node as a child
func (n *Node) AddTabNode(title string, b *GuiBox) *Node {
	var newNode *Node
	parent := n

	newNode = parent.makeNode(title, 444, 400 + Config.counter)
	newNode.uiTab = parent.uiTab
	newNode.box = b

	if (Config.DebugNode) {
		fmt.Println("")
		log.Println("parent:")
		parent.Dump()

		fmt.Println("")
		log.Println("newNode:")
		newNode.Dump()
	}

	if (newNode.uiTab == nil) {
		log.Println("wit/gui/ AddTabNode() Something went wrong tab == nil")
		// TODO: try to find the tab or window and make them if need be
		return newNode
	}
	newNode.uiTab.Append(title, b.UiBox)

	return newNode
}

func (n *Node) AddTab(title string, uiC *ui.Box) *Node {
	parent := n
	log.Println("gui.Node.AddTab() START name =", title)
	if parent.uiWindow == nil {
		parent.Dump()
		log.Println("gui.Node.AddTab() ERROR ui.Window == nil")
		return nil
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

	newNode := parent.makeNode(title, 555, 600 + Config.counter)
	newNode.uiTab = tab
	newNode.uiBox = uiC
	// panic("gui.AddTab() after makeNode()")
	tabSetMargined(newNode.uiTab)
	return newNode
}
