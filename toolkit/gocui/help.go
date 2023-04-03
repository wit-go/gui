// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func addHelp() {
	me.baseGui.SetManagerFunc(helplayout)
}

func helplayout(g *gocui.Gui) error {
	var err error
	maxX, _ := g.Size()

	help, err := g.SetView("help", maxX-32, 0, maxX-1, 17, 0)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		help.SelBgColor = gocui.ColorGreen
		help.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(help, "KEYBINDINGS")
		fmt.Fprintln(help, "Enter: Click Button")
		fmt.Fprintln(help, "Tab/Space: Switch Buttons")
		fmt.Fprintln(help, "")
		fmt.Fprintln(help, "h: Help")
		fmt.Fprintln(help, "Backspace: Delete Button")
		fmt.Fprintln(help, "Arrow keys: Move Button")
		fmt.Fprintln(help, "t: Move Button to the top")
		fmt.Fprintln(help, "b: Move Button to the button")
		fmt.Fprintln(help, "h: hide buttons")
		fmt.Fprintln(help, "s: show buttons")
		fmt.Fprintln(help, "p: panic()")
		fmt.Fprintln(help, "STDOUT: /tmp/witgui.log")
		fmt.Fprintln(help, "Ctrl-C or Q: Exit")
		if _, err := g.SetCurrentView("help"); err != nil {
			return err
		}
	}
	me.helpLabel = help
	return nil
}
