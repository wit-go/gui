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

func (n *Node) FindWindowBox() *GuiBox {
	if (n.box == nil) {
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		log.Println("SERIOUS ERROR n.box == nil in FindWindowBox()")
		os.Exit(-1)
	}
	return n.box
}

func (n *Node) Append(child *Node) {
	//	if (n.UiBox == nil) {
	//		return
	//	}
	n.children = append(n.children, child)
}
func (n *Node) List() {
	findByIdDFS(n, "test")
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

func (n *Node) InitTab(title string, custom func() ui.Control) *Node {
	if n.uiWindow == nil {
		log.Println("gui.InitTab() ERROR ui.Window == nil")
		n.Dump()
		os.Exit(-1)
	}
	if n.box != nil {
		log.Println("gui.InitTab() ERROR box already exists")
		n.Dump()
		os.Exit(-1)
	}

	tab := ui.NewTab()
	n.uiWindow.SetChild(tab)
	n.uiWindow.SetMargined(true)

	tab.Append(title, custom())
	tab.SetMargined(0, true)

	var newNode Node
	newNode.Name = title
	newNode.parent = n
	n.Append(&newNode)
	newNode.uiTab = tab
	/*
	if boxs.node == nil {
		log.Println("gui.InitTab() 4 Fuck node = ", n)
		n.Dump()
		os.Exit(-1)
	}
	*/
	return &newNode
}
