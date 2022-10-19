package gui

import (
	"log"
	"os"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

// This function should make a new node with the parent and
// the 'tab' as a child

func (n *Node) NewTab(title string) *Node {
	log.Println("gui.Node.AddTab() START name =", title)

	return n.AddTabNew(title)
}

func (n *Node) AddTabNew(title string) *Node {
	log.Println("gui.Node.AddTab() START name =", title)

	if (n.Toolkit == nil) {
		log.Println("FUCK TOOLKIT nil uiWindow =", n.uiWindow)
		log.Println("FUCK TOOLKIT nil uiTab =", n.uiTab)
		log.Println("FUCK TOOLKIT nil Toolkit =", n.Toolkit)
		// return n.AddTab(title) // need to make a toolkit here
		n.Dump()
		os.Exit(0)
	}
	log.Println("Make new node")
	newN := n.New(title)
	log.Println("Add tab to window")
	t := n.Toolkit.AddTab(title)
	newN.Toolkit = t

	n.Append(newN)
	return newN
}

func (n *Node) AddTab(title string, uiC *ui.Box) *Node {
	return n.AddTabNew(title)
}
/*
func (n *Node) AddTabBAD(title string, uiC *ui.Box) *Node {
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

	newNode := n.New(title)
	newNode.uiTab = tab
	newNode.uiBox = uiC
	// tabSetMargined(newNode.uiTab)
	return newNode
}
*/
