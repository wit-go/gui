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

	control  *ui.Control
	window  *ui.Window
}

func (n Node) SetName(name string) {
	// n.uiType.SetName(name)
	if (n.window != nil) {
		log.Println("node is a window. setting title =", name)
		n.window.SetTitle(name)
		return
	}
	log.Println("*ui.Control =", n.control)
	return
}

func (n Node) Append(child Node) {
	//	if (n.UiBox == nil) {
	//		return
	//	}
	// n.uiType.Append(child, x)
}

func findByIdDFS(node *Node, id string) *Node {
	if node.id == id {
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			findByIdDFS(child, id)
		}
	}
	return nil
}
