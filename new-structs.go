package gui

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// https://ieftimov.com/post/golang-datastructures-trees/

type Node struct {
	id     string
	Name   string
	Width  int
	Height int

	children []*Node
	box	*GuiBox

	control  *ui.Control
	window  *ui.Window
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
