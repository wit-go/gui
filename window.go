package gui

import (
	"git.wit.org/wit/gui/toolkit"
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
	var custom func()

	// If the user didn't set a custom Exit() use the standard exit() function
	// This makes sure the GUI properly closes everything (GTK, QT, console ui, etc exit)
	if (Config.Exit != nil) {
		log(debugGui, "setting a custom exit")
		custom = func() {
			log(debugChange, "Running a custom exit()", Config.Exit)
			log(debugChange, "Running a custom exit() Config.Title =", Config.Title)
			log(debugChange, "Running a custom exit() Config.Width =", Config.Width)
			Config.Exit(newNode)
		}
	} else {
		log(debugGui, "setting the standard exit")
		custom = func () {
			log(debugChange, "Running StandardExit()")
			StandardExit()
		}
	}
	// Windows are created off of the master node of the Binary Tree
	newNode = Config.master.New(Config.Title, toolkit.Window, custom)

	newNode.widget.Width     = Config.Width
	newNode.widget.Height    = Config.Height

	log(debugGui, "Window()", Config.Title)

	send(nil, newNode)
	return newNode
}
