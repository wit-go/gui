package main

import (
	"fmt"
	"errors"
	"strconv"
	"bufio"
	"strings"

	"github.com/awesome-gocui/gocui"
	"go.wit.com/gui/toolkit"
)

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func (n *node) textResize() bool {
	w := n.tk
	var width, height int = 0, 0
	var changed bool = false

	for i, s := range splitLines(n.Text) {
		log(logInfo, "textResize() len =", len(s), i, s)
		if (width < len(s)) {
			width = len(s)
		}
		height += 1
	}
	if (w.gocuiSize.w1 != w.gocuiSize.w0 + width + me.FramePadW) {
		w.gocuiSize.w1 = w.gocuiSize.w0 + width + me.FramePadW
		changed = true
	}
	if (w.gocuiSize.h1 != w.gocuiSize.h0 + height + me.FramePadH) {
		w.gocuiSize.h1 = w.gocuiSize.h0 + height + me.FramePadH
		changed = true
	}
	if (changed) {
		n.showWidgetPlacement(logNow, "textResize() changed")
	}
	return changed
}

func (n *node) hideView() {
	n.SetVisible(false)
}

// display's the text of the widget in gocui
// will create a new gocui view if there isn't one or if it has been moved
func (n *node) showView() {
	var err error
	w := n.tk

	if (w.cuiName == "") {
		log(logError, "showView() w.cuiName was not set for widget", w)
		w.cuiName = strconv.Itoa(n.WidgetId)
	}

	// if the gocui element doesn't exist, create it
	if (w.v == nil) {
		n.recreateView()
	}	
	x0, y0, x1, y1, err := me.baseGui.ViewPosition(w.cuiName)
	log(logInfo, "showView() w.v already defined for widget", n.Name, err)

	// n.smartGocuiSize()
	changed := n.textResize()

	if (changed) {
		log(logNow, "showView() textResize() changed. Should recreateView here wId =", w.cuiName)
	} else {
		log(logNow, "showView() Clear() and Fprint() here wId =", w.cuiName)
		w.v.Clear()
		fmt.Fprint(w.v, n.Text)
		n.SetVisible(false)
		n.SetVisible(true)
		return
	}

	// if the gocui element has changed where it is supposed to be on the screen
	// recreate it
	if (x0 != w.gocuiSize.w0) {
		n.recreateView()
		return
	}
	if (y0 != w.gocuiSize.h0) {
		log(logError, "showView() start hight mismatch id=", w.cuiName, "gocui h vs computed h =", w.gocuiSize.h0, y0)
		n.recreateView()
		return
	}
	if (x1 != w.gocuiSize.w1) {
		log(logError, "showView() too wide", w.cuiName, "w,w", w.gocuiSize.w1, x1)
		n.recreateView()
		return
	}
	if (y1 != w.gocuiSize.h1) {
		log(logError, "showView() too high", w.cuiName, "h,h", w.gocuiSize.h1, y1)
		n.recreateView()
		return
	}

	n.SetVisible(true)
}

// create or recreate the gocui widget visible
// deletes the old view if it exists and recreates it
func (n *node) recreateView() {
	var err error
	w := n.tk
	log(logError, "recreateView() START", n.WidgetType, n.Name)
	if (me.baseGui == nil) {
		log(logError, "recreateView() ERROR: me.baseGui == nil", w)
		return
	}

	// this deletes the button from gocui
	me.baseGui.DeleteView(w.cuiName)
	w.v = nil

	if (n.Name == "CLOUDFLARE_EMAIL") {
		n.showWidgetPlacement(logNow, "n.Name=" + n.Name + " n.Text=" + n.Text + " " + w.cuiName)
		n.dumpWidget("jwc")
		n.textResize()
		n.showWidgetPlacement(logNow, "n.Name=" + n.Name + " n.Text=" + n.Text + " " + w.cuiName)
	}

	a := w.gocuiSize.w0
	b := w.gocuiSize.h0
	c := w.gocuiSize.w1
	d := w.gocuiSize.h1

	w.v, err = me.baseGui.SetView(w.cuiName, a, b, c, d, 0)
	if err == nil {
		n.showWidgetPlacement(logError, "recreateView()")
		log(logError, "recreateView() internal plugin error err = nil")
		return
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		n.showWidgetPlacement(logError, "recreateView()")
		log(logError, "recreateView() internal plugin error error.IS()", err)
		return
	}

	// this sets up the keybinding for the name of the window
	// does this really need to be done? I think we probably already
	// know everything about where all the widgets are so we could bypass
	// the gocui package and just handle all the mouse events internally here (?)
	// for now, the w.v.Name is a string "1", "2", "3", etc from the widgetId

	// set the binding for this gocui view now that it has been created
	// gocui handles overlaps of views so it will run on the view that is clicked on
	me.baseGui.SetKeybinding(w.v.Name(), gocui.MouseLeft, gocui.ModNone, click)

	// this actually sends the text to display to gocui
	w.v.Wrap = true
	w.v.Frame = w.frame
	w.v.Clear()
	fmt.Fprint(w.v, n.Text)
	// n.showWidgetPlacement(logNow, "n.Name=" + n.Name + " n.Text=" + n.Text + " " + w.cuiName)
	// n.dumpWidget("jwc 2")

	// if you don't do this here, it will be black & white only
	if (w.color != nil) {
		w.v.FrameColor = w.color.frame
		w.v.FgColor = w.color.fg
		w.v.BgColor = w.color.bg
		w.v.SelFgColor = w.color.selFg
		w.v.SelBgColor = w.color.selBg
	}
	if (n.Name == "CLOUDFLARE_EMAIL") {
		n.showWidgetPlacement(logNow, "n.Name=" + n.Name + " n.Text=" + n.Text + " " + w.cuiName)
		n.dumpTree(true)
	}
	log(logError, "recreateView() END")
}

func (n *node) hideWidgets() {
	w := n.tk
	w.isCurrent = false
	switch n.WidgetType {
	case toolkit.Root:
	case toolkit.Flag:
	case toolkit.Window:
	case toolkit.Box:
	case toolkit.Grid:
	default:
		n.hideView()
	}
	for _, child := range n.children {
		child.hideWidgets()
	}
}

func (n *node) hideFake() {
	w := n.tk
	if (w.isFake) {
		n.hideView()
	}
	for _, child := range n.children {
		child.hideFake()
	}
}

func (n *node) showFake() {
	w := n.tk
	if (w.isFake) {
		n.setFake()
		n.showWidgetPlacement(logNow, "showFake:")
		n.showView()
	}
	for _, child := range n.children {
		child.showFake()
	}
}

func (n *node) showWidgets() {
	w := n.tk
	if (w.isFake) {
		// don't display by default
	} else {
		n.showWidgetPlacement(logInfo, "current:")
		n.showView()
	}
	for _, child := range n.children {
		child.showWidgets()
	}
}
