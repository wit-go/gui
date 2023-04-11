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
func Init() {
	log(logInfo, "Init() of awesome-gocui")
	me.defaultWidth = 10
	me.defaultHeight = 2 // this means by default one line of text in a button
	me.defaultBehavior = true

	me.horizontalPadding = 20
	me.horizontalPadding = 20
	me.groupPadding = 4
	me.buttonPadding = 3

	// the raw beginning of each window (or tab)
	me.rawW = 7
	me.rawH = 3

	me.padW = 3
	me.padH = 3

	me.pluginChan = make(chan toolkit.Action)

	log(logNow, "Init() start pluginChan")
	go catchActionChannel()
	sleep(.1)
	go main()
	// probably not needed, but in here for now under development
	sleep(.1)
}

// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	me.callback = guiCallback
}

func PluginChannel() chan toolkit.Action {
	return me.pluginChan
}

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
}

func main() {
	log(logInfo, "main() start Init()")

	outf, err := os.OpenFile("/tmp/witgui.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		exit("error opening file: %v", err)
	}
	defer outf.Close()

	setOutput(outf)
	log("This is a test log entry")

	MouseMain()
	me.baseGui.Close()
}
