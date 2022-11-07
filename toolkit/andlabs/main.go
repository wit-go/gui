package toolkit

import (
	"log"

	"github.com/andlabs/ui"
	// the _ means we only need this for the init()
	_ "github.com/andlabs/ui/winmanifest"
)

func Main(f func()) {
	if (DebugToolkit) {
		log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	}
	ui.Main(f)
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
		log.Println("Sending function to gui.Main() (using gtk via andlabs/ui)")
	}
	ui.QueueMain(f)
}
