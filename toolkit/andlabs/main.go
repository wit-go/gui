package main

import (
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	// the _ means we only need this for the init()
	_ "github.com/andlabs/ui/winmanifest"
)

func Main(f func()) {
	log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
	ui.Main( func() {
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		// time.Sleep(1 * time.Second)
		// NewWindow2("helloworld2", 200, 100)
		f()
	})
}

// Other goroutines must use this to access the GUI
//
// You can not acess / process the GUI thread directly from
// other goroutines. This is due to the nature of how
// Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
//
// For example: Queue(NewWindow())
//
func Queue(f func()) {
	log(debugToolkit, "Sending function to ui.QueueMain()")
	log(true, "THIS DOES BREAK. TODO: wrap this")
	ui.QueueMain(f)
	// f()
}

func Init() {
	log(debugToolkit, "Init()")

	mapWidgets = make(map[*andlabsT]*toolkit.Widget)
	mapToolkits = make(map[*toolkit.Widget]*andlabsT)
}

func Quit() {
	log(debugToolkit, "Quit() TODO: close the toolkit cleanly")
}
