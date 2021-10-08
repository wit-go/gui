package gui

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func Main(f func()) {
	log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	log.Println("Starting gui.Main() (using gtk via andlabs/ui)")
	ui.Main(f)
}

// Other goroutines must use this
//
// You can not acess / process the GUI thread directly from
// other goroutines. This is due to the nature of how
// Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
// For example: gui.Queue(addNewTabForColorSelection())
func Queue(f func()) {
	log.Println("Sending function to gui.Main() (using gtk via andlabs/ui)")
	ui.QueueMain(f)
}

func ExampleWindow() {
	log.Println("START gui.ExampleWindow()")

	title := "Test Window"
	node := InitWindow(nil, nil, title, 0)
	box := node.box
	window := box.Window
	log.Println("box =", box)
	log.Println("window =", window)
	box.AddDebugTab("jcarr Debug")

	window.UiWindow.Show()
}

func DebugWindow() {
	log.Println("START gui.ExampleWindow()")

	title := "Debug Window"
	node := InitWindow(nil, nil, title, 0)
	box := node.box
	window := box.Window
	log.Println("box =", box)
	log.Println("window =", window)
	box.AddDebugTab("jcarr Debug")

	window.UiWindow.Show()
}
