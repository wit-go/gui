package gui

import (
	"log"
//	"fmt"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

)

// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

func (n *Node) AddTabRaw(title string, uiC ui.Control) *Node {
	log.Println("gui.Node.AddTabRaw()")

	tab := n.uiTab
	if (tab == nil) {
		log.Println("gui.Node.AddTabRaw() FAIL tab == nil")
		return n
	}

	if (uiC == nil) {
		// hbox := ui.NewHorizontalBox()
		// hbox.SetPadded(true)
		// uiC = hbox
		log.Println("gui.Node.AddTabRaw() FAIL *ui.Control == nil")
		return n
	}
	tab.Append(title, uiC)

	/*
	newNode := parent.makeNode(title, 555, 600 + Config.counter)
	newNode.uiTab = tab
	newNode.uiBox = uiC
	// panic("gui.AddTab() after makeNode()")
	tabSetMargined(newNode.uiTab)
	*/
	return n
}
