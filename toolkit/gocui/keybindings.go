// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/awesome-gocui/gocui"
)

func defaultKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	for _, n := range []string{"but1", "but2", "help", "but3"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}
	if err := g.SetKeybinding("", gocui.MouseRelease, gocui.ModNone, mouseUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, globalDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModMouseCtrl, ctrlDown); err != nil {
		return err
	}
	/*
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, globalDown); err != nil {
		return err
	}
	*/
	if err := g.SetKeybinding("msg", gocui.MouseLeft, gocui.ModNone, msgDown); err != nil {
		return err
	}
	addDebugKeys(g)
	return nil
}

var showDebug bool = true

func addDebugKeys(g *gocui.Gui) {
	// dump all widget info to the log
	g.SetKeybinding("", 'd', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			log(logNow, "gocui.SetKeyBinding() dumpTree() START")
			// me.rootNode.dumpTree(true)
			fakeStartWidth = me.DevelOffsetW
			fakeStartHeight = me.TabH + me.FramePadH
			if (showDebug) {
				me.rootNode.showFake()
				showDebug = false
			} else {
				me.rootNode.hideFake()
				showDebug = true
			}
			return nil
	})

	// hide all widgets
	g.SetKeybinding("", 'h', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.rootNode.hideWidgets()
			return nil
	})

	// show all widgets
	g.SetKeybinding("", 's', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.rootNode.showWidgets()
			return nil
	})

	// exit
	g.SetKeybinding("", 'q', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.baseGui.Close()
			exit("forced exit() from within gocui")
			return nil
	})
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.baseGui.Close()
			exit("forced exit() from within gocui")
			return nil
	})

	// panic
	g.SetKeybinding("", 'p', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.baseGui.Close()
			panic("forced panic in gocui")
			return nil
	})
}
