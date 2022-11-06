// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
)

const delta = 1

var (
	views   = []string{}
	curView = -1
	idxView = 0
	currentX = 5
	currentY = 2
	groupSize = 0
	baseGui *gocui.Gui
)

var helpLabel *gocui.View

func Init() {
	// setup log to write to a file
//	logInit()

	g, err := gocui.NewGui(gocui.OutputNormal, true)
	baseGui = g
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.SelFrameColor = gocui.ColorRed

	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := newView(g); err != nil {
		log.Panicln(err)
	}

	AddButton("hello")
	AddButton("world")
	AddButton("foo")

	AddGroup("blank")
	AddButton("bar")
	AddButton("bar none")
	AddButton("bar going")

	AddGroup("te")
	AddButton("world 2")
	AddButton("foo 2")

	if err := baseGui.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func ToolkitMain() {
	if err := baseGui.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	var err error
	maxX, _ := g.Size()
	helpLabel, err = g.SetView("help", maxX-32, 0, maxX-1, 11, 0)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		fmt.Fprintln(helpLabel, "KEYBINDINGS")
		fmt.Fprintln(helpLabel, "Enter: Click Button")
		fmt.Fprintln(helpLabel, "Tab/Space: Switch Buttons")
		fmt.Fprintln(helpLabel, "")
		fmt.Fprintln(helpLabel, "h: Help")
		fmt.Fprintln(helpLabel, "Backspace: Delete Button")
		fmt.Fprintln(helpLabel, "Arrow keys: Move Button")
		fmt.Fprintln(helpLabel, "t: Move Button to the top")
		fmt.Fprintln(helpLabel, "b: Move Button to the button")
		fmt.Fprintln(helpLabel, "Ctrl-C or Q: Exit")
	}
	return nil
}
