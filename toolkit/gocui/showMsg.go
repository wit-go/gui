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

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2, maxY/2, maxX/2+outputW, maxY/2+outputH, 0); err == nil || errors.Is(err, gocui.ErrUnknownView) {
		v.Clear()
		v.SelBgColor = gocui.ColorCyan
		v.SelFgColor = gocui.ColorBlack
		l += "foo\n" + "bar\n"
		fmt.Fprintln(v, l)
	}
	// g.SetViewOnTop("msg")
	g.SetViewOnBottom("msg")
	return nil
}
