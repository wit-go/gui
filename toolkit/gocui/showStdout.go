// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"
)

var outputW int = 80
var outputH int = 24

func moveMsg(g *gocui.Gui) {
	mx, my := g.MousePosition()
	if !movingMsg && (mx != initialMouseX || my != initialMouseY) {
		movingMsg = true
	}
	g.SetView("msg", mx-xOffset, my-yOffset, mx-xOffset+outputW, my-yOffset+outputH, 0)
}

func showMsg(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	// setOutput(me.rootNode)

	makeOutputWidget(g, l)
	return nil
}

func makeOutputWidget(g *gocui.Gui, stringFromMouseClick string) {
	maxX, maxY := g.Size()
	v, err := g.SetView("msg", maxX-32, maxY/2, maxX/2+outputW, maxY/2+outputH, 0)
	// help, err := g.SetView("help", maxX-32, 0, maxX-1, 13, 0)
	if errors.Is(err, gocui.ErrUnknownView) {
		log("this is supposed to happen?", err)
	}

	if (err != nil) {
		log("create output window failed", err)
		return
	}

	v.Clear()
	v.SelBgColor = gocui.ColorCyan
	v.SelFgColor = gocui.ColorBlack
	fmt.Fprintln(v, "figure out how to capture STDOUT to here\n" + stringFromMouseClick)
	g.SetViewOnBottom("msg")
	return
}
