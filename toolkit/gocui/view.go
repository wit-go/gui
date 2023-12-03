package main

import (
	"fmt"
	"errors"
	"strconv"
	"bufio"
	"strings"

	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func (n *node) textResize() {
	w := n.tk
	var width, height int

	for i, s := range splitLines(n.Text) {
		log(logNow, "textResize() len =", len(s), i, s)
		if (width < len(s)) {
			width = len(s)
		}
		height = i
	}
	w.gocuiSize.w1 = w.gocuiSize.w0 + width + me.FramePadW
	w.gocuiSize.h1 = w.gocuiSize.h0 + height + me.FramePadH
	n.showWidgetPlacement(logNow, "textResize()")
}

// display's the text of the widget in gocui
func (n *node) showView() {
	var err error
	w := n.tk

	if (w.cuiName == "") {
		log(logError, "showView() w.cuiName was not set for widget", w)
		w.cuiName = strconv.Itoa(n.WidgetId)
	}

	if (w.v == nil) {
		n.updateView()
	}	
	x0, y0, x1, y1, err := me.baseGui.ViewPosition(w.cuiName)
	log(logInfo, "showView() w.v already defined for widget", n.Name, err)
	if (x0 != w.gocuiSize.w0) || (y0 != w.gocuiSize.h0) {
		log(logError, "showView() w.v.w0 != x0", n.Name, w.gocuiSize.w0, x0)
		log(logError, "showView() w.v.h0 != y0", n.Name, w.gocuiSize.h0, y0)
		n.updateView()
		return
	}
	if (x1 != w.gocuiSize.w1) || (y1 != w.gocuiSize.h1) {
		log(logError, "showView() w.v.w1 != x1", n.Name, w.gocuiSize.w1, x1)
		log(logError, "showView() w.v.h1 != y1", n.Name, w.gocuiSize.h1, y1)
		n.updateView()
		return
	}

	if (w.v.Visible == false) {
		log(logInfo, "showView() w.v.Visible set to true ", n.Name)
		w.v.Visible = true
	}
}

func (n *node) updateView() {
	var err error
	w := n.tk
	if (me.baseGui == nil) {
		log(logError, "showView() ERROR: me.baseGui == nil", w)
		return
	}
	me.baseGui.DeleteView(w.cuiName)
	w.v = nil

	a := w.gocuiSize.w0
	b := w.gocuiSize.h0
	c := w.gocuiSize.w1
	d := w.gocuiSize.h1

	w.v, err = me.baseGui.SetView(w.cuiName, a, b, c, d, 0)
	if err == nil {
		n.showWidgetPlacement(logError, "drawView()")
		log(logError, "drawView() internal plugin error err = nil")
		return
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		n.showWidgetPlacement(logError, "drawView()")
		log(logError, "drawView() internal plugin error error.IS()", err)
		return
	}

	me.baseGui.SetKeybinding(w.v.Name(), gocui.MouseLeft, gocui.ModNone, click)

	w.v.Wrap = true
	w.v.Frame = w.frame
	w.v.Clear()
	fmt.Fprint(w.v, n.Text)
	n.showWidgetPlacement(logNow, "Window: " + n.Text)

	n.setDefaultHighlight()
	n.setDefaultWidgetColor()
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
		n.deleteView()
	}
	for _, child := range n.children {
		child.hideWidgets()
	}
}

func (n *node) hideFake() {
	w := n.tk
	if (w.isFake) {
		n.deleteView()
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
		if n.IsCurrent() {
			n.showWidgetPlacement(logInfo, "current:")
			n.showView()
		} else {
			n.showWidgetPlacement(logInfo, "not:")
			// w.drawView()
		}
	}
	for _, child := range n.children {
		child.showWidgets()
	}
}
