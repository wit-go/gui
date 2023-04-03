// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
)

func Init() {
	log(logInfo, "Init() of awesome-gocui")
	me.defaultWidth = 10
	me.defaultHeight = 2
	me.defaultBehavior = true

	me.horizontalPadding = 20
	me.groupPadding = 2
	me.buttonPadding = 2
}

func Exit() {
	// TODO: exit correctly
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

	MouseMain()
}
