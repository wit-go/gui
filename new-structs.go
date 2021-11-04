package gui

import (
	"log"
	"fmt"
	"image/color"
//	"reflect"

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
	OnChanged	func ()
	Color	color.RGBA

	parent	*Node
	children []*Node

	area	*GuiArea
	custom func(*Node)
	values	interface {}

	uiControl *ui.Control
	uiButton  *ui.Button
	uiFontButton  *ui.FontButton
	uiColorButton  *ui.ColorButton
	uiWindow  *ui.Window
	uiAttrstr *ui.AttributedString
	uiTab     *ui.Tab
	uiBox     *ui.Box
	uiArea    *ui.Area
	uiText    *ui.EditableCombobox
	uiMultilineEntry    *ui.MultilineEntry
	uiGroup	*ui.Group
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
	log.Println("gui.Node.Dump() OnChanged  = ", n.OnChanged)

	if (n.parent == nil) {
		log.Println("gui.Node.Dump() parent     = nil")
	} else {
		log.Println("gui.Node.Dump() parent     = ", n.parent.id)
	}
	log.Println("gui.Node.Dump() children   = ", n.children)

	// log.Println("gui.Node.Dump() window     = ", n.window)
	// log.Println("gui.Node.Dump() box        = ", n.box)

	log.Println("gui.Node.Dump() uiWindow   = ", n.uiWindow)
	log.Println("gui.Node.Dump() uiTab      = ", n.uiTab)
	log.Println("gui.Node.Dump() uiBox      = ", n.uiBox)
	log.Println("gui.Node.Dump() uiControl  = ", n.uiControl)
	log.Println("gui.Node.Dump() uiButton   = ", n.uiButton)
	log.Println("gui.Node.Dump() uiText     = ", n.uiText)
	if (n.id == "") {
		log.Println("THIS SHOULD NOT HAPPEN: gui.Node.Dump() id == nil")
	}
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

// var indent string = "\t"
var indent string = "    "

func indentPrintln(depth int, format string, a ...interface{}) {
	var space string
	for i := 0; i < depth; i++ {
		space = space + indent
	}

	// newFormat := tabs + strconv.Itoa(depth) + " " + format
	newFormat := space + format
	log.Println(newFormat, a)
}

func (n *Node) ListChildren(dump bool) {
	indentPrintln(listChildrenDepth, "", n.id, n.Width, n.Height, n.Name)

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
func (n *Node) AddTabNode(title string) *Node {
	var newNode *Node
	parent := n

	newNode = parent.makeNode(title, 444, 400 + Config.counter)
	newNode.uiTab = parent.uiTab
	// newNode.box = b

	if (Config.DebugNode) {
		fmt.Println("")
		log.Println("parent:")
		parent.Dump()

		fmt.Println("")
		log.Println("newNode:")
		newNode.Dump()
	}

	if (newNode.uiTab != nil) {
		log.Println("ERROR: wit/gui/ AddTabNode() Something went wrong tab == nil")
		// TODO: try to find the tab or window and make them if need be
		// newNode.uiTab.Append(title, b.UiBox)
	}

	return newNode
}

func (n *Node) AddHorizontalBreak() *Node {
	log.Println("AddHorizontalBreak  added to node =", n.Name)
	if (n.uiBox != nil) {
		tmp := ui.NewHorizontalSeparator()
		n.uiBox.Append(tmp, Config.Stretchy)
	} else {
		n.Dump()
		return nil
	}
	return n
}

func (n *Node) AddVerticalBreak() *Node {
	log.Println("AddVerticalBreak  added to node =", n.Name)
	if (n.uiBox != nil) {
		tmp := ui.NewVerticalSeparator()
		n.uiBox.Append(tmp, Config.Stretchy)
	} else {
		n.Dump()
		return nil
	}
	return n
}

func (n *Node) AddHorizontalBox(title string) *Node {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	if (n.uiBox != nil) {
		log.Println("add new hbox to uiBox =", n.uiBox)
		n.uiBox.Append(hbox, Config.Stretchy)
		newNode := n.makeNode(title, 333, 333 + Config.counter)
		newNode.parent = n
		newNode.uiBox = hbox
		// newNode.uiControl = hbox
		return newNode
	}
	if (n.uiTab != nil) {
		log.Println("add new hbox to uiTab =", n.uiTab)
		n.uiTab.Append(title, hbox)
		newNode := n.makeNode(title, 333, 333 + Config.counter)
		newNode.parent = n
		newNode.uiBox = hbox
		// newNode.uiControl = hbox
		return newNode
	}
	return n
}
func (n *Node) AddTab(title string, uiC *ui.Box) *Node {
	parent := n
	log.Println("gui.Node.AddTab() START name =", title)
	if parent.uiWindow == nil {
		parent.Dump()
		log.Println("gui.Node.AddTab() ERROR ui.Window == nil")
		return nil
	}
	if parent.uiTab == nil {
		inittab := ui.NewTab() // no, not that 'inittab'
		parent.uiWindow.SetChild(inittab)
		parent.uiWindow.SetMargined(true)
		parent.uiTab = inittab
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
	tabSetMargined(newNode.uiTab)
	return newNode
}
