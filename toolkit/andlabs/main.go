package main

import (
	"embed"
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	// the _ means we only need this for the init()
	_ "github.com/andlabs/ui/winmanifest"
)

//go:embed resources
var res embed.FS

func Main(f func()) {
	log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
	ui.Main( func() {
		log(debugToolkit, "Starting gui.Main() (using gtk via andlabs/ui)")
		// time.Sleep(1 * time.Second)
		// NewWindow2("helloworld2", 200, 100)
		f()
	})
}

// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	callback = guiCallback
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
	log(debugPlugin, "using gui.Queue() in this plugin DOES BREAK. TODO: solve this with channels")
	ui.QueueMain(f)
}

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func Init() {
	log(debugToolkit, "Init()")
	// Can you pass values to a plugin init() ? Otherwise, there is no way to safely print
	// log(debugToolkit, "gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)

	// mapWidgets = make(map[*andlabsT]*toolkit.Widget)
	// mapToolkits = make(map[*toolkit.Widget]*andlabsT)

	andlabs = make(map[int]*andlabsT)
}

// TODO: properly exit the plugin since Quit() doesn't do it
func Quit() {
	log(debugToolkit, "Quit() TODO: close the toolkit cleanly")
}
