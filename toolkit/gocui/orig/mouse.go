// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toolkit

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/awesome-gocui/gocui"
)

var g *gocui.Gui
var err error
var Custom func(string)

func OnExit(f func(string)) {
	Custom = f
}

func Exit() {
	g.Close()
}

func mouseClick(name string) {
	// output screws up the console. Need to fix this by redirecting all console output to a file from log.Println()
	// log.Println("g.Close()")
	// g.Close()

	log.Println("Found andlabs Running custom function for the mouse click")
	Custom(name)
	// panic("got andlabs")
}

func Init() {
	log.Println("start Init()")

	f, err := os.OpenFile("/tmp/guilogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	g, err = gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}
	log.Println("exit Init()")
}

func StartConsoleMouse() {
	defer g.Close()
	log.Println("start Main()")

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
	log.Println("exit Main()")
}

func layout(g *gocui.Gui) error {
	if v, err := g.SetView("but1", 2, 2, 22, 17, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "andlabs")
		fmt.Fprintln(v, "addDemoTab")
		fmt.Fprintln(v, "DemoToolkitWindow")
		fmt.Fprintln(v, "DebugWindow")
		fmt.Fprintln(v, "do nothing")
		fmt.Fprintln(v, "exit")
		if _, err := g.SetCurrentView("but1"); err != nil {
			return err
		}
	}
	if v, err := g.SetView("but2", 24, 2, 44, 4, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Button 2 - line 1")
	}
	if v, err := g.SetView("but3", 24, 2, 44, 4, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Button 2 - line 1")
	}
	if v, err := g.SetView("but4", 24, 2, 44, 4, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Button 2 - line 1")
	}
	if v, err := g.SetView("but5", 24, 2, 44, 4, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Button 2 - line 1")
	}
	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	for _, n := range []string{"but1", "but2"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}
	if err := g.SetKeybinding("msg", gocui.MouseLeft, gocui.ModNone, delMsg); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseRight, gocui.ModNone, delMsg); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseMiddle, gocui.ModNone, delMsg); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
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

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		mouseClick(l)
		fmt.Fprintln(v, l)
	}
	return nil
}

func delMsg(g *gocui.Gui, v *gocui.View) error {
	// Error check removed, because delete could be called multiple times with the above keybindings
	g.DeleteView("msg")
	return nil
}
