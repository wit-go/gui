// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"git.wit.org/wit/gui/toolkit"

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
	helpLabel *gocui.View
	err error
	ch chan(func ())
)

func Init() {
	baseGui, err = gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}

	baseGui.Highlight = true
	baseGui.SelFgColor = gocui.ColorRed
	baseGui.SelFrameColor = gocui.ColorRed

	baseGui.SetManagerFunc(layout)

	if err := initKeybindings(baseGui); err != nil {
		log.Panicln(err)
	}

	viewWidget = make(map[*gocui.View]*toolkit.Widget)
	stringWidget = make(map[string]*toolkit.Widget)

	ch = make(chan func())
}

func Queue(f func()) {
	log.Println("QUEUEEEEE")
	f()
}

func Main(f func()) {
	if (baseGui == nil) {
		panic("WTF Main()")
	}
	defer baseGui.Close()
	// log.Println("ADDDDDDDD BUTTTTTTTTTON")
	// addButton("test 3")
	f()
	if err := baseGui.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
	baseGui.Close()
	os.Exit(0)
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
