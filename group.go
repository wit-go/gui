package gui

import "log"

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func (n *Node) NewGroup(name string) *Node {
	var t *toolkit.Toolkit
	var gNode, bNode *Node
	log.Println("toolkit.NewGroup() START", name)

	// make a *Node with a *toolkit.Group
	gNode = n.AddNode(name + " part1")
	t = toolkit.NewGroup(n.uiBox, name)
	gNode.Toolkit = t
	gNode.Dump()

	// make a *Node with a *toolkit.Box
	bNode = n.AddNode(name + " part2")
	bNode.Toolkit = t.NewBox()
	bNode.uiBox = bNode.Toolkit.GetBox()
	bNode.Dump()
	return bNode
}

func (n *Node) AddGroup(title string) *Node {
	return n.NewGroup(title)
}

func (n *Node) AddGroupOld(title string) *Node {
	if (n == nil) {
		return nil
	}
	hbox := n.uiBox
	if (hbox == nil) {
		return n
	}
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	newNode := n.AddNode(title)
	newNode.uiBox = vbox
	return newNode
}

/*
func (n *Node) NewGroup(title string) *Node {
	group := ui.NewGroup(title)
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)
*/
