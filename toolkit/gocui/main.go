// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"git.wit.org/wit/gui/toolkit"
)

// sets defaults and establishes communication
// to this toolkit from the wit/gui golang package
func init() {
	log(logInfo, "Init() of awesome-gocui")

	// init the config struct default values
	Set(&me, "default")

	me.pluginChan = make(chan toolkit.Action)

	log(logNow, "Init() start pluginChan")
	go catchActionChannel()
	sleep(.1) // probably not needed, but in here for now under development
	go main()
	sleep(.1) // probably not needed, but in here for now under development
}

// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	me.callback = guiCallback
}

func PluginChannel() chan toolkit.Action {
	return me.pluginChan
}

/*
recieves requests from the program to do things like:
* add new widgets
* change the text of a label
* etc..
*/
func catchActionChannel() {
	log(logInfo, "catchActionChannel() START")
	for {
		log(logInfo, "catchActionChannel() infinite for() loop restarted select on channel")
	    	select {
		case a := <-me.pluginChan:
			if (me.baseGui == nil) {
				// something went wrong initializing the gocui
				log(logError,"ERROR: console did not initialize")
				continue
			}
			log(logNow, "catchActionChannel()", a.WidgetId, a.ActionType, a.WidgetType, a.Name)
			action(&a)
		}
	}
}

func Exit() {
	// TODO: what should actually happen here?
	me.baseGui.Close()
	sendBackQuit()
}

func sendBackQuit() {
	// send 'Quit' back to the program (?)
	var a toolkit.Action
	a.ActionType = toolkit.UserQuit
	me.callback <- a
}

var outf *os.File

func main() {
	var err error
	log(logInfo, "main() start Init()")

	outf, err = os.OpenFile("/tmp/witgui.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		exit("error opening file: %v", err)
	}
	defer outf.Close()

	// setOutput(outf)
	// log("This is a test log entry")

	MouseMain()

	log(true, "MouseMain() closed")
	me.baseGui.Close()

	sendBackQuit()
}
