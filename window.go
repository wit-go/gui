package gui

import (
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

	title := Config.Title
	// Windows are created off of the master node of the Binary Tree
	newNode = Config.master.New(title, "Window")

	newNode.Widget.Width     = Config.Width
	newNode.Widget.Height    = Config.Height

	if (Config.Exit != nil) {
		log("setting a custom exit")
		newNode.custom = func() {
			Config.Exit(newNode)
		}
	} else {
		log("not setting a custom exit")
	}

	if (newNode.custom == nil) {
		newNode.custom = func () {StandardExit(newNode)}
	}

	newNode.Widget.Custom = newNode.custom

	log(debugGui, "gui.Node.Window()", title)

	for _, aplug := range allPlugins {
		log(debugGui, "gui.Node.NewWindow() toolkit plugin =", aplug.name)
		if (aplug.NewWindow == nil) {
			log(debugGui, "gui.Node.NewWindow() is nil")
			continue
		}
		aplug.NewWindow(&newNode.Widget)
	}

	return newNode
}
