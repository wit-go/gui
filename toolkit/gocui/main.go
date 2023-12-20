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

	pluginChan = make(chan toolkit.Action)

	log(logNow, "Init() start pluginChan")
	go catchActionChannel()
	sleep(.1) // probably not needed, but in here for now under development
	go main()
	sleep(.1) // probably not needed, but in here for now under development
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
		case a := <-pluginChan:
			if (me.baseGui == nil) {
				// something went wrong initializing the gocui
				log(logError,"ERROR: console did not initialize")
				continue
			}
			log(logInfo, "catchActionChannel()", a.WidgetId, a.ActionType, a.WidgetType, a.Name)
			action(&a)
		}
	}
}

func Exit() {
	// TODO: what should actually happen here?
	log(true, "Exit() here. doing standardExit()")
	standardExit()
}

func standardExit() {
	log(true, "standardExit() doing baseGui.Close()")
	me.baseGui.Close()
	log(true, "standardExit() doing outf.Close()")
	outf.Close()
	// log(true, "standardExit() setOutput(os.Stdout)")
	// setOutput(os.Stdout)
	log(true, "standardExit() send back Quit()")
	go sendBackQuit() // don't stall here in case the
	// induces a delay in case the callback channel is broken
	sleep(1)
	log(true, "standardExit() exit()")
	exit()
}
func sendBackQuit() {
	// send 'Quit' back to the program (?)
	var a toolkit.Action
	a.ActionType = toolkit.UserQuit
	callback <- a
}

var outf *os.File

func main() {
	var err error
	log(logInfo, "main() start Init()")

	outf, err = os.OpenFile("/tmp/witgui.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		exit("error opening file: %v", err)
	}
	os.Stdout = outf
	defer outf.Close()

	// setOutput(outf)
	// log("This is a test log entry")

	ferr, _ := os.OpenFile("/tmp/witgui.err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	os.Stderr = ferr
	gocuiMain()

	log(true, "MouseMain() closed")
	standardExit()
}
