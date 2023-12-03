// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"github.com/awesome-gocui/gocui"
)

// This initializes the gocui package
// it runs SetManagerFunc which passes every input
// event (keyboard, mouse, etc) to the function "gocuiEvent()"
func gocuiMain() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	me.baseGui = g

	g.Cursor = true
	g.Mouse = true

	// this sets the function that is run on every event. For example:
	// When you click the mouse, move the mouse, or press a key on the keyboard
	// This is equivalent to xev or similar to cat /dev/input on linux
	g.SetManagerFunc(gocuiEvent)

	if err := defaultKeybindings(g); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		panic(err)
	}
}

// Thanks to the gocui developers -- your package kicks ass
// This function is called on every event. It is a callback function from the gocui package
// which has an excellent implementation. While gocui handles things like text highlighting
// and the layout of the text areas -- also things like handling SIGWINCH and lots of really
// complicated console handling, it sends events here in a clean way.
// This is equivalent to the linux command xev (apt install x11-utils)
func gocuiEvent(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	mx, my := g.MousePosition()
	log(logVerbose, "handleEvent() START", maxX, maxY, mx, my, msgMouseDown)
	if _, err := g.View("msg"); msgMouseDown && err == nil {
		moveMsg(g)
	}
	if widgetView, _ := g.View("msg"); widgetView == nil {
		log(logNow, "handleEvent() create output widget now", maxX, maxY, mx, my)
		makeOutputWidget(g, "this is a create before a mouse click")
		if (me.logStdout != nil) {
			// setOutput(me.logStdout)
		}
	} else {
		log(logInfo, "output widget already exists", maxX, maxY, mx, my)
	}
	mouseMove(g)
	log(logVerbose, "handleEvent() END  ", maxX, maxY, mx, my, msgMouseDown)
	return nil
}

func dragOutputWindow() {
}

// turns off the frame on the global window
func setFrame(b bool) {
	// TODO: figure out what this might be useful for
	// what is this do? I made it just 2 lines for now. Is this useful for something?
	v := SetView("global", 15, 5, 80, 8, 10)
	if (v == nil) {
		log(logError, "setFrame() global failed")
	}
	v.Frame = b
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func SetView(name string, x0, y0, x1, y1 int, overlaps byte) *gocui.View {
	if (me.baseGui == nil) {
		log(logError, "SetView() ERROR: me.baseGui == nil")
		return nil
	}

	v, err := me.baseGui.SetView(name, x0, y0, x1, y1, overlaps)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			log(logError, "SetView() global failed on name =", name)
		}
		return nil
	}
	return v
}
