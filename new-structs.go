package gui

import (
	"log"
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

	box	*GuiBox

	uiControl  *ui.Control
	uiWindow  *ui.Window
	uiTab  *ui.Tab
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
	log.Println("gui.Node.Dump() parent     = ", n.parent)
	log.Println("gui.Node.Dump() children   = ", n.children)
	log.Println("gui.Node.Dump() box        = ", n.box)
	log.Println("gui.Node.Dump() uiControl  = ", n.uiControl)
	log.Println("gui.Node.Dump() uiWindow   = ", n.uiWindow)
	log.Println("gui.Node.Dump() uiTab      = ", n.uiTab)
	if (n.id == "") {
		panic("gui.Node.Dump() id == nil")
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

func (n *Node) FindTab() *ui.Tab {
	return n.uiTab
}

func (n *Node) FindControl() *ui.Control {
	return n.uiControl
}

func (n *Node) FindBox() *GuiBox {
	return n.box
}

func (n *Node) FindWindowBox() *GuiBox {
	if (n.box == nil) {
		panic("SERIOUS ERROR n.box == nil in FindWindowBox()")
	}
	return n.box
}

func (n *Node) Append(child *Node) {
	//	if (n.UiBox == nil) {
	//		return
	//	}
	n.children = append(n.children, child)
	log.Println("child node:")
	child.Dump()
	log.Println("parent node:")
	n.Dump()
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

func findByIdDFS(node *Node, id string) *Node {
	log.Println("findByIdDFS()", id, node)
	node.Dump()
	if node.id == id {
		log.Println("Found node id =", id, node)
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			newNode := findByIdDFS(child, id)
			if (newNode != nil) {
				return newNode
			}
		}
	}
	return nil
}

func findByName(node *Node, name string) *Node {
	log.Println("findByName()", name, node)
	node.Dump()
	if node.Name == name {
		log.Println("findByName() Found node name =", name, node)
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			newNode := findByName(child, name)
			if (newNode != nil) {
				return newNode
			}
		}
	}
	return nil
}

/*
func (parent *Node) InitTab(title string) *Node {
	if parent.uiWindow == nil {
		parent.Dump()
		panic("gui.InitTab() ERROR ui.Window == nil")
	}
	if parent.box == nil {
		parent.Dump()
		panic("gui.InitTab() ERROR box == nil")
	}

	tab := ui.NewTab()
	parent.uiWindow.SetChild(tab)
	parent.uiWindow.SetMargined(true)
	parent.uiTab = tab

	tab.Append(title, initBlankWindow())
	tab.SetMargined(0, true)

	newNode := makeNode(parent, title, 555, 600 + Config.counter)
	return newNode
}
*/

func (parent *Node) AddTab(title string) *Node {
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

		parent.Dump()
		// panic("gui.AddTab() ERROR uiTab == nil")
	}

	tab := parent.uiTab
	parent.uiWindow.SetMargined(true)

	tab.Append(title, initBlankWindow())
	tab.SetMargined(0, true)

	newNode := makeNode(parent, title, 555, 600 + Config.counter)
	newNode.uiTab = tab
	return newNode
}
