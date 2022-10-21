package gui

import (
	"log"
)

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"


// the _ means we only need this for the init()

const Xaxis = 0 // box that is horizontal
const Yaxis = 1 // box that is vertical

func init() {
	log.Println("gui.init() has been run")

	Config.counter = 0
	Config.prefix = "wit"
	Config.DebugNode = false
	Config.DebugTabs = false

	title := "master"
	w     := 640
	h     := 480
	// f     := StandardClose

	Config.master = addNode(title, w, h)
	// Config.master.custom = f

	Config.master.Dump()
}

func Main(f func()) {
	log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
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
