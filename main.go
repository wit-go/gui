package gui

import (
	"log"
	"os"
)

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

const Xaxis = 0 // stack things horizontally
const Yaxis = 1 // stack things vertically

func init() {
	log.Println("gui.init() has been run")

	Config.counter = 0
	Config.prefix = "wit"

	// Config.Options.Debug = true
	// Config.Options.DebugNode = true
	// Config.Options.DebugTabs = true

	title := "guiBinaryTree"
	w     := 640
	h     := 480

	// Populates the top of the binary tree
	Config.master = addNode(title, w, h)
	if (Config.Options.Debug) {
		Config.master.Dump()
	}

	// load the gocli plugin
	PlugGocli = LoadPlugin("../../toolkit/gocli.so")
}

func Main(f func()) {
	if (Config.Options.Debug) {
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	}
	toolkit.Main(f)
}

// Other goroutines must use this to access the GUI
//
// You can not acess / process the GUI thread directly from
// other goroutines. This is due to the nature of how
// Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
// For example: gui.Queue(NewWindow())
func Queue(f func()) {
	log.Println("Sending function to gui.Main() (using gtk via andlabs/ui)")
	toolkit.Queue(f)
}

// The window is destroyed but the application does not quit
func StandardClose(n *Node) {
	if (Config.Options.Debug) {
		log.Println("wit/gui Standard Window Close. name =", n.Name)
	}
}


// The window is destroyed but the application does not quit
func StandardExit(n *Node) {
	if (Config.Options.Debug) {
		log.Println("wit/gui Standard Window Exit. running os.Exit()")
	}
	os.Exit(0)
}
