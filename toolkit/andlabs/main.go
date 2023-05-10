package main

import (
	"sync"
	"git.wit.org/wit/gui/toolkit"

	"github.com/andlabs/ui"
	// the _ means we only need this for the init()
	_ "github.com/andlabs/ui/winmanifest"
)

// this is the channel we get requests to make widgets
var pluginChan chan toolkit.Action

// the starting point of the binary tree
var rootNode *node

var uiMainUndef bool = true
var uiMain sync.Once
var muAction sync.Mutex

func catchActionChannel() {
	log(logInfo, "catchActionChannel() START")
	for {
		log(logInfo, "catchActionChannel() for loop")
	    	select {
		case a := <-pluginChan:
			log(logInfo, "catchActionChannel() SELECT widget id =", a.WidgetId, a.Name)
			log(logInfo, "catchActionChannel() STUFF", a.WidgetId, a.ActionType, a.WidgetType)
			muAction.Lock()
			// TODO ui.QueueMain(f)
			// TODO ui.QueueMain( func() {rawAction(a)} )
			rawAction(a)
			muAction.Unlock()
			log(logInfo, "catchActionChannel() STUFF END", a.WidgetId, a.ActionType, a.WidgetType)
		}
	}
}

// Other goroutines must use this to access the GUI
//
// You can not acess / process the GUI thread directly from
// other goroutines. This is due to the nature of how
// Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
//
// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	callback = guiCallback
}

func PluginChannel() chan toolkit.Action {
	return pluginChan
}

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func init() {
	log(logNow, "Init() START")
	log(debugToolkit, "Init()")
	// Can you pass values to a plugin init() ? Otherwise, there is no way to safely print
	// log(debugToolkit, "gui/toolkit init() Setting defaultBehavior = true")
	setDefaultBehavior(true)

	// andlabs = make(map[int]*andlabsT)
	pluginChan = make(chan toolkit.Action, 1)

	log(logNow, "Init() start channel reciever")
	go ui.Main(func() {
		demoUI()
	})
	go catchActionChannel()
	log(logNow, "Init() END")
}
