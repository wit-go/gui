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
	baseGui.SetManagerFunc(helplayout)
}

func helplayout(g *gocui.Gui) error {
	var err error
	maxX, _ := g.Size()

	helpLabel, err = g.SetView("help", maxX-32, 0, maxX-1, 12, 0)
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
		fmt.Fprintln(helpLabel, "STDOUT: /tmp/witgui.log")
		fmt.Fprintln(helpLabel, "Ctrl-C or Q: Exit")
	}
	return nil
}
