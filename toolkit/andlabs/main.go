package main

import (
	"log"
// 	"time"

	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	// the _ means we only need this for the init()
	_ "github.com/andlabs/ui/winmanifest"
)

func Main(f func()) {
	if (DebugToolkit) {
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	}
	ui.Main( func() {
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
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
	if (DebugToolkit) {
		log.Println("Sending function to ui.QueueMain() (using gtk via andlabs/ui)")
	}
	//ui.QueueMain(f)
	f()
}

func Init() {
	log.Println("should Init() here")

	mapWidgets = make(map[*andlabsT]*toolkit.Widget)
	mapToolkits = make(map[*toolkit.Widget]*andlabsT)
}

func Quit() {
	log.Println("should Quit() here")
	// myExit(nil)
}
