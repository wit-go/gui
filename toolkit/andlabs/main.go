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

// this is the channel we get requests to make widgets
var pluginChan chan toolkit.Action

var uiMainUndef bool = true

func catchActionChannel() {
	log(logNow, "catchActionChannel() START")
	for {
		log(logNow, "catchActionChannel() for loop")
	    	select {
		case a := <-pluginChan:
			log(logNow, "catchActionChannel() SELECT widget id =", a.WidgetId, a.Name)
			// go Action(a)
			if (uiMainUndef) {
				log(logError,"catchActionChannel() main() was not run yet")
				log(logError,"catchActionChannel() main() was not run yet")
				log(logError,"catchActionChannel() main() was not run yet")
				log(logError,"catchActionChannel() ui.Main() START")
				log(logError,"catchActionChannel() ui.Main() START")
				log(logError,"catchActionChannel() ui.Main() START")
				log(logError,"catchActionChannel() ui.Main() START")
				sleep(1)
				go ui.Main(demoUI)
				// go ui.Main( func() {
				// 	rawAction(a)
				// })
				// probably not needed, but in here for now under development
				uiMainUndef = false
				sleep(1)
			} else {
				log(logNow, "catchActionChannel() STUFF", a.WidgetId, a.ActionType, a.WidgetType)
				rawAction(a)
				log(logNow, "catchActionChannel() STUFF END", a.WidgetId, a.ActionType, a.WidgetType)
			}
		}
	}
}

/*
func main(f func()) {
	log(debugNow, "Main() START (using gtk via andlabs/ui)")
	f() // support the old way. deprecate this
}
*/

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
func queue(f func()) {
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

	andlabs = make(map[int]*andlabsT)
	pluginChan = make(chan toolkit.Action)

	log(logNow, "Init() start channel reciever")
	go catchActionChannel()
	log(logNow, "Init() END")
}

// TODO: properly exit the plugin since Quit() doesn't do it
func Quit() {
	log(debugToolkit, "Quit() TODO: close the toolkit cleanly")
}
