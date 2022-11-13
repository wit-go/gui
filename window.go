package gui

import (
	"log"
)

//import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

// This routine creates a blank window with a Title and size (W x H)
//
// This routine can not have any arguements due to the nature of how
// it can be passed via the 'andlabs/ui' queue which, because it is
// cross platform, must pass UI changes into the OS threads (that is
// my guess).
func NewWindow() *Node {
	var newNode *Node
//	var t *toolkit.Toolkit

	title := Config.Title
	// Windows are created off of the master node of the Binary Tree
	newNode = Config.master.New(title)

	newNode.Widget.Name      = title
	newNode.Widget.Width     = Config.Width
	newNode.Widget.Height    = Config.Height

	if (Config.Exit != nil) {
		newNode.custom = func() {
			Config.Exit(newNode)
		}
	}

	if (newNode.custom == nil) {
		newNode.custom = func () {StandardExit(newNode)}
	}

	newNode.Widget.Custom = newNode.custom

	log.Println("gui.Node.Window()", title)

	// t = toolkit.NewWindow(title, w, h)
	// n.toolkit = t

	for _, aplug := range allPlugins {
		log.Println("gui.Node.NewWindow() toolkit plugin =", aplug.name)
		if (aplug.NewWindow == nil) {
			log.Println("gui.Node.NewWindow() is nil")
			continue
		}
		aplug.NewWindow(&newNode.Widget)
	}

	// TODO: this is still confusing and probably wrong. This needs to communicate through a channel
	// newNode.toolkit = n.toolkit.NewButton(name)
	// newNode.toolkit.Custom = newNode.Widget.Custom

	return newNode
}
