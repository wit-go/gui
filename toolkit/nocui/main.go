package main

import (
	"sync"
	"git.wit.org/wit/gui/toolkit"
)

var muAction sync.Mutex

func catchActionChannel() {
	log(logNow, "catchActionChannel() START")
	for {
		log(logNow, "catchActionChannel() for loop")
	    	select {
		case a := <-pluginChan:
			log(logNow, "catchActionChannel() SELECT widget id =", a.WidgetId, a.Name)
			log(logNow, "catchActionChannel() STUFF", a.WidgetId, a.ActionType, a.WidgetType)
			muAction.Lock()
			doAction(&a)
			muAction.Unlock()
			log(logNow, "catchActionChannel() STUFF END", a.WidgetId, a.ActionType, a.WidgetType)
		}
	}
}

/*
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
*/

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func init() {
	log(logNow, "Init() START")
	log(logInfo, "Init()")

	// andlabs = make(map[int]*andlabsT)
	pluginChan = make(chan toolkit.Action, 1)

	log(logNow, "Init() start channel reciever")
	go catchActionChannel()
	go simpleStdin()
	log(logNow, "Init() END")
}
