// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"

	"github.com/awesome-gocui/gocui"
)

func MouseMain() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	me.baseGui = g

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := defaultKeybindings(g); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		panic(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	mx, my := g.MousePosition()
	if _, err := g.View("msg"); msgMouseDown && err == nil {
		moveMsg(g)
	}
	// TODO: figure out what this might be useful for
	// what is this do? I made it just 2 lines for now. Is this useful for something?
	if v, err := g.SetView("global", 15, 5, maxX, 8, 10); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			log("global failed", maxX, maxY)
			return err
		}
		v.Frame = false
	}
	helplayout(g)
	if widgetView, _ := g.View("msg"); widgetView == nil {
		log(logInfo, "create output widget now", maxX, maxY, mx, my)
		makeOutputWidget(g, "this is a create before a mouse click")
		if (me.logStdout != nil) {
			setOutput(me.logStdout)
		}
	} else {
		log(logInfo, "output widget already exists", maxX, maxY, mx, my)
	}
	updateHighlightedView(g)
	log(logInfo, "layout() END", maxX, maxY, mx, my)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func updateHighlightedView(g *gocui.Gui) {
	mx, my := g.MousePosition()
	for _, view := range g.Views() {
		view.Highlight = false
	}
	if v, err := g.ViewByPosition(mx, my); err == nil {
		v.Highlight = true
	}
}

func msgDown(g *gocui.Gui, v *gocui.View) error {
	initialMouseX, initialMouseY = g.MousePosition()
	if vx, vy, _, _, err := g.ViewPosition("msg"); err == nil {
		xOffset = initialMouseX - vx
		yOffset = initialMouseY - vy
		msgMouseDown = true
	}
	return nil
}

func mouseUp(g *gocui.Gui, v *gocui.View) error {
	if msgMouseDown {
		msgMouseDown = false
		if movingMsg {
			movingMsg = false
			return nil
		} else {
			g.DeleteView("msg")
		}
	} else if globalMouseDown {
		globalMouseDown = false
		g.DeleteView("globalDown")
	}
	return nil
}
