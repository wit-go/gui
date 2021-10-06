package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

type Node struct {
	Name		string
	tag		string
	Width		int
	Height		int

	uiType		*ui.Control
	Children	[]*Node
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
