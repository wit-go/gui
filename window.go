package gui

import (
	"log"
)

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

// This routine creates a blank window with a Title and size (W x H)
//
// This routine can not have any arguements due to the nature of how
// it can be passed via the 'andlabs/ui' queue which, because it is
// cross platform, must pass UI changes into the OS threads (that is
// my guess).
func NewWindow() *Node {
	var n *Node
	var t *toolkit.Toolkit

	title := Config.Title
	w     := Config.Width
	h     := Config.Height
	f     := Config.Exit

	// Windows are created off of the master node of the Binary Tree
	n = Config.master.New(title)
	n.custom = f

	t = toolkit.NewWindow(title, w, h)
	t.Custom = func () {
		log.Println("Got to wit/gui Window Close for window:", title)
		if (n.custom == nil) {
			log.Println("Got to wit/gui Window Close custom() == nil")
		}
		log.Println("Got to wit/gui Window Close START custom()")
		n.custom(n)
		log.Println("Got to wit/gui Window Close END custom()")
	}
	n.toolkit = t

	return n
}
