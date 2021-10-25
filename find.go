package gui

import (
	"log"
	"os"

//	"github.com/andlabs/ui"
//	_ "github.com/andlabs/ui/winmanifest"
//	"github.com/davecgh/go-spew/spew"
)


func FindWindow(s string) *GuiWindow {
	for name, window := range Data.WindowMap {
		if name == s {
			return window
		}
	}
	log.Printf("COULD NOT FIND WINDOW", s)
	return nil
}

func FindBox(s string) *GuiBox {
	for name, window := range Data.WindowMap {
		if name != s {
			continue
		}
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			return abox
		}
		log.Println("gui.FindBox() NEED TO INIT WINDOW name =", name)
	}
	log.Println("gui.FindBox() COULD NOT FIND BOX", s)
	return nil
}

func FindNode(name string) *Node {
	if Data.NodeMap == nil {
		log.Println("gui.FindNode() gui.Data.NodeMap == nil")
		return nil
	}
	log.Println("gui.FindNode() searching Data.NodeMap:")
	for id, node := range Data.NodeMap {
		log.Println("\tData.NodeMap name =", node.Width, node.Height, id)
		node.Dump()
		if (name == node.Name) {
			return node
		}
		newNode := findByName(node, name)
		if (newNode != nil) {
			return newNode
		}
		log.Println("gui.FindNode() could not find node name =", name)
		os.Exit(-1)
	}
	log.Println("gui.FindNode() could not find node name =", name)
	return nil
}

func (dn *GuiData) findId(id string) *Node {
	if Data.NodeMap == nil {
		log.Println("gui.Data.findId() map == nil")
		return nil
	}
	// log.Println("Dumping Data.NodeMap:")
	for name, node := range Data.NodeMap {
		// log.Println("\tData.NodeMap name =", node.id, node.Width, node.Height, name)
		if (id == node.id) {
			log.Println("\tgui.Data.findId() found node =", node.id, node.Width, node.Height, name)
			return node
		}
		// TODO: fix // Oct 9
		// node.findId(id)
	}
	return nil
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
