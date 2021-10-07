package gui

import (
	"log"
	"os"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// https://ieftimov.com/post/golang-datastructures-trees/

type Node struct {
	id     string
	Name   string
	Width  int
	Height int

	parent	*Node
	children []*Node

	box	*GuiBox

	control  *ui.Control
	window  *ui.Window
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Window() *Node {
	return n.parent
}

func (n *Node) Dump() {
	log.Println("gui.Node.Dump() id       = ", n.id)
	log.Println("gui.Node.Dump() Name     = ", n.Name)
	log.Println("gui.Node.Dump() Width    = ", n.Width)
	log.Println("gui.Node.Dump() Height   = ", n.Height)
	log.Println("gui.Node.Dump() parent   = ", n.parent)
	log.Println("gui.Node.Dump() children = ", n.children)
	log.Println("gui.Node.Dump() box      = ", n.box)
	log.Println("gui.Node.Dump() control  = ", n.control)
	log.Println("gui.Node.Dump() window   = ", n.window)
}


func (n *Node) SetName(name string) {
	// n.uiType.SetName(name)
	if (n.window != nil) {
		log.Println("node is a window. setting title =", name)
		n.window.SetTitle(name)
		return
	}
	log.Println("*ui.Control =", n.control)
	return
}

func (n *Node) FindWindowBox() *GuiBox {
	if (n.box == nil) {
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
	}
	return n.box
}

func (n *Node) Append(child Node) {
	//	if (n.UiBox == nil) {
	//		return
	//	}
	// n.uiType.Append(child, x)
}
func (n *Node) List() {
	findByIdDFS(n, "test")
}

func findByIdDFS(node *Node, id string) *Node {
	log.Println("findByIdDFS()", id, node)
	if node.id == id {
		log.Println("Found node id =", id, node)
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			findByIdDFS(child, id)
		}
	}
	return nil
}

func (n *Node) InitTab(title string, custom func() ui.Control) *Node {
	boxs := n.box
	if boxs == nil {
		log.Println("gui.InitTab() 1 Fuck node = ", n)
		n.Dump()
		os.Exit(-1)
	}
	if boxs.Window == nil {
		log.Println("gui.InitTab() 2 Fuck node = ", n)
		n.Dump()
		os.Exit(-1)
		return nil
	}
	if boxs.Window.UiWindow == nil {
		log.Println("gui.InitTab() 3 Fuck node = ", n)
		n.Dump()
		os.Exit(-1)
		return nil
	}

	window := boxs.Window.UiWindow
	tab := ui.NewTab()
	window.SetChild(tab)
	window.SetMargined(true)

	tab.Append(title, custom())
	tab.SetMargined(0, true)
	// tab.SetMargined(1, true)

	boxs.Window.UiTab = tab
	if boxs.node == nil {
		log.Println("gui.InitTab() 4 Fuck node = ", n)
		n.Dump()
		os.Exit(-1)
	}
	return n
}
