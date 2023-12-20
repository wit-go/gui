// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"github.com/awesome-gocui/gocui"
	"go.wit.com/gui/toolkit"
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
	// mouseDown() runs whenever you click on an unknown view (?)
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, mouseDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModMouseCtrl, ctrlDown); err != nil {
		return err
	}
//	if err := g.SetKeybinding(w.v.Name(), gocui.MouseLeft, gocui.ModNone, click); err != nil {
//		return err
//	}
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

func addDebugKeys(g *gocui.Gui) {
	// show debugging buttons
	g.SetKeybinding("", 'd', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			fakeStartWidth = me.FakeW
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

	// display the help menu
	g.SetKeybinding("", '?', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if (showHelp) {
				helplayout()
				showHelp = false
			} else {
				me.baseGui.DeleteView("help")
				showHelp = true
			}
			return nil
	})

	// redraw all the widgets
	g.SetKeybinding("", 'r', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if (redoWidgets) {
				redoWindows(0,0)
				redoWidgets = false
			} else {
				me.rootNode.hideWidgets()
				redoWidgets = true
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

	// list all widgets
	g.SetKeybinding("", 'L', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.rootNode.listWidgets()
			return nil
	})

	// list all widgets with positions
	g.SetKeybinding("", 'M', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			me.rootNode.dumpTree(true)
			return nil
	})

	// log to output window
	g.SetKeybinding("", 'o', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if me.logStdout.Visible() {
				me.logStdout.SetVisible(false)
				setOutput(os.Stdout)
			} else {
				me.logStdout.SetVisible(true)
				setOutput(me.logStdout.tk)
			}
			return nil
	})

	// exit
	g.SetKeybinding("", 'q', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			standardExit()
			return nil
	})
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			standardExit()
			return nil
	})
	g.SetKeybinding("", gocui.KeyCtrlD, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if (showDebug) {
				var a toolkit.Action
				a.B = true
				a.ActionType = toolkit.EnableDebug
				callback <- a
				logInfo = true
				logVerbose = true
			} else {
				logInfo = false
				logVerbose = false
			}
			return nil
	})
	g.SetKeybinding("", gocui.KeyCtrlV, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if (logVerbose) {
				logInfo = false
				logVerbose = false
			} else {
				logInfo = true
				logVerbose = true
			}
			return nil
	})

	// panic
	g.SetKeybinding("", 'p', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			standardExit()
			panic("forced panic in gocui")
			return nil
	})
}
