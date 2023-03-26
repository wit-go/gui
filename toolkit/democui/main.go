// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
)

func OnExit(f func(string)) {
	Custom = f
}

func Init() {
	log("Init() of democui")
}

func Exit() {
	g.Close()
}

func mouseClick(name string) {
	// output screws up the console. Need to fix this by redirecting all console output to a file from log.Println()
	// log.Println("g.Close()")
	// g.Close()

	log("Found andlabs Running custom function for the mouse click")
	Custom(name)
	// panic("got andlabs")
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

	MouseMain()
}

/*
func StartConsoleMouse() {
	defer g.Close()
	log("start Main()")

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		exit(err)
	}
	log("exit Main()")
}
*/
