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
var pluginChan chan toolkit.Action

func catchActionChannel() {
	log(logNow, "makeCallback() START")
	for {
		log(logNow, "makeCallback() for loop")
	    	select {
		case a := <-pluginChan:
			log(logNow, "makeCallback() SELECT widget id =", a.WidgetId, a.Name)
			// go Action(a)
			if (a.WidgetType == toolkit.Window) {
				log(logNow, "makeCallback() WINDOW START")
				go ui.Main( func() {
					log(logNow, "ui.Main() WINDOW START")
					rawAction(&a)
					log(logNow, "ui.Main() WINDOW END")
				})
				sleep(.5)
				log(logNow, "makeCallback() WINDOW END")
			} else {
				log(logNow, "makeCallback() STUFF")
				rawAction(&a)
				/*
				Queue( func() {
					rawAction(&a)
				})
				*/
				log(logNow, "makeCallback() STUFF END")
			}
			// sleep(.1)
		}
	}
}

func Main(f func()) {
	log(debugNow, "gui.Main() START (using gtk via andlabs/ui)")
	f() // support the old way. deprecate this
}

// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	callback = guiCallback
}

func PluginChannel() chan toolkit.Action {
	return pluginChan
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
	log(logNow, "Sending function to ui.QueueMain()")
	log(logNow, "using gui.Queue() in this plugin DOES BREAK. TODO: solve this with channels")
	ui.QueueMain(f)
}

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func Init() {
	log(logNow, "Init() START")
	log(debugToolkit, "Init()")
	// Can you pass values to a plugin init() ? Otherwise, there is no way to safely print
	// log(debugToolkit, "gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)

	// mapWidgets = make(map[*andlabsT]*toolkit.Widget)
	// mapToolkits = make(map[*toolkit.Widget]*andlabsT)

	andlabs = make(map[int]*andlabsT)
	pluginChan = make(chan toolkit.Action)

	log(logNow, "Init() ui.Main() start")
	go catchActionChannel()
	/*
	ui.Main( func() {
		log(logNow, "gui.Main() IN (using gtk via andlabs/ui)")
		var a toolkit.Action
		a.Name = "jcarr"
		a.Width = 640
		a.Height = 480
		a.WidgetId = 0
		newWindow(&a)
		// time.Sleep(1 * time.Second)
		// NewWindow2("helloworld2", 200, 100)
		log(logNow, "gui.Main() EXIT (using gtk via andlabs/ui)")
	})
	*/
	log(logNow, "Init() END")
}

// TODO: properly exit the plugin since Quit() doesn't do it
func Quit() {
	log(debugToolkit, "Quit() TODO: close the toolkit cleanly")
}
