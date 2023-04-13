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
	if _, err := g.View("msg"); msgMouseDown && err == nil {
		moveMsg(g)
	}
	if v, err := g.SetView("global", -1, -1, maxX, maxY, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Frame = false
	}
	/*
	if v, err := g.SetView("but1", 2, 2, 22, 7, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Button 1 - line 1")
		fmt.Fprintln(v, "Button 1 - line 2")
		fmt.Fprintln(v, "Button 1 - line 3")
		fmt.Fprintln(v, "Button 1 - line 4")
		if _, err := g.SetCurrentView("but1"); err != nil {
			return err
		}
	}
	*/
	helplayout(g)
	updateHighlightedView(g)
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
