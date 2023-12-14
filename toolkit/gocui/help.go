// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
)

var helpText []string = []string{"KEYBINDINGS",
	"",
	"?: toggle help",
	"d: toggle debugging",
	"r: redraw widgets",
	"s/h: show/hide all widgets",
	"L: list all widgets",
	"q: quit()",
	"p: panic()",
	"o: show Stdout",
	"l: log to /tmp/witgui.log",
	"Ctrl-D: Toggle Debugging",
	"Ctrl-V: Toggle Verbose Debugging",
	"Ctrl-C: Exit",
	"",
}

func hidehelplayout() {
	me.baseGui.DeleteView("help")
	// n.deleteView()
	// child.hideFake()
}

func helplayout() error {
	g := me.baseGui
	var err error
	maxX, _ := g.Size()

	var newW int = 8
	for _, s := range(helpText) {
		if newW < len(s) {
			newW = len(s)
		}
	}

	help, err := g.SetView("help", maxX-(newW + me.FramePadW), 0, maxX-1, len(helpText) + me.FramePadH, 0)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		help.SelBgColor = gocui.ColorGreen
		help.SelFgColor = gocui.ColorBlack
		// fmt.Fprintln(help, "Enter: Click Button")
		// fmt.Fprintln(help, "Tab/Space: Switch Buttons")
		// fmt.Fprintln(help, "Backspace: Delete Button")
		// fmt.Fprintln(help, "Arrow keys: Move Button")

		fmt.Fprintln(help, strings.Join(helpText, "\n"))

		if _, err := g.SetCurrentView("help"); err != nil {
			return err
		}
	}
	me.helpLabel = help
	return nil
}
