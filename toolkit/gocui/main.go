// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"git.wit.org/wit/gui/toolkit"
)

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
}

// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	me.callback = guiCallback
}

func PluginChannel() chan toolkit.Action {
	return me.pluginChan
}

func catchActionChannel() {
	log(logNow, "makeCallback() START")
	for {
		log(logNow, "makeCallback() for loop")
	    	select {
		case a := <-me.pluginChan:
			log(logNow, "makeCallback() SELECT widget id =", a.WidgetId, a.Name)
			Action(&a)
			sleep(.1)
		}
	}
}

func Exit() {
	// TODO: send exit to the plugin
	me.baseGui.Close()
}

func Main(f func()) {
	log("start Init()")

	outf, err := os.OpenFile("/tmp/witgui.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		exit("error opening file: %v", err)
	}
	defer outf.Close()

	setOutput(outf)
	log("This is a test log entry")

	if (me.pluginChan == nil) {
		me.pluginChan = make(chan toolkit.Action)
	}
	go catchActionChannel()
	MouseMain()
	me.baseGui.Close()
}
