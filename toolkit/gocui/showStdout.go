package main

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"
	"go.wit.com/gui/toolkit"
)

var outputW int = 180
var outputH int = 24

func moveMsg(g *gocui.Gui) {
	mx, my := g.MousePosition()
	if !movingMsg && (mx != initialMouseX || my != initialMouseY) {
		movingMsg = true
	}
	g.SetView("msg", mx-xOffset, my-yOffset, mx-xOffset+outputW, my-yOffset+outputH + me.FramePadH, 0)
	g.SetViewOnBottom("msg")
}

func showMsg(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	log(true, "showMsg() v.name =", v.Name())
	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	makeOutputWidget(g, l)
	return nil
}

func makeOutputWidget(g *gocui.Gui, stringFromMouseClick string) *gocui.View {
	maxX, maxY := g.Size()

	if (me.rootNode == nil) {
		// keep skipping this until the binary tree is initialized
		return nil
	}

	if (me.logStdout == nil) {
		a := new(toolkit.Action)
		a.Name = "stdout"
		a.Text = "stdout"
		a.WidgetType = toolkit.Stdout
		a.WidgetId = -3
		a.ParentId = 0
		n := addNode(a)
		me.logStdout = n
		me.logStdout.tk.gocuiSize.w0 = maxX - 32
		me.logStdout.tk.gocuiSize.h0 = maxY/2
		me.logStdout.tk.gocuiSize.w1 = me.logStdout.tk.gocuiSize.w0 + outputW
		me.logStdout.tk.gocuiSize.h1 = me.logStdout.tk.gocuiSize.h0 + outputH
	}
	v, err := g.View("msg")
	if (v == nil) {
		log("makeoutputwindow() this is supposed to happen. v == nil", err)
	} else {
		log("makeoutputwindow() msg != nil. WTF now? err =", err)
	}

	// help, err := g.SetView("help", maxX-32, 0, maxX-1, 13, 0)
	// v, err = g.SetView("msg", 3, 3, 30, 30, 0)

	v, err = g.SetView("msg", maxX-32, maxY/2, maxX/2+outputW, maxY/2+outputH, 0)
	if errors.Is(err, gocui.ErrUnknownView) {
		log("makeoutputwindow() this is supposed to happen?", err)
	}

	if (err != nil) {
		log("makeoutputwindow() create output window failed", err)
		return nil
	}

	if (v == nil) {
		log("makeoutputwindow() msg == nil. WTF now? err =", err)
		return nil
	} else {
		me.logStdout.tk.v = v
	}

	v.Clear()
	v.SelBgColor = gocui.ColorCyan
	v.SelFgColor = gocui.ColorBlack
	fmt.Fprintln(v, "figure out how to capture STDOUT to here\n" + stringFromMouseClick)
	g.SetViewOnBottom("msg")
	// g.SetViewOnBottom(v.Name())
	return v
}
