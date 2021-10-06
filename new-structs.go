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

	uiType   *ui.Control
	children []*Node
}

func (n Node) SetName(name string) {
	// n.uiType.SetName(name)
	log.Println("n.uiType =", n.uiType)
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
