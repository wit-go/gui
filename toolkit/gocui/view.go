package main

import (
	"fmt"
	"errors"
	"strconv"
	"bufio"
	"strings"

	"github.com/awesome-gocui/gocui"
//	"git.wit.org/wit/gui/toolkit"
)

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func (w *cuiWidget) textResize() {
	var width, height int

	for i, s := range splitLines(w.text) {
		log(logNow, "textResize() len =", len(s), i, s)
		if (width < len(s)) {
			width = len(s)
		}
		height = i
	}
	w.realWidth = width + 3
	w.realHeight = me.defaultHeight + height
	w.gocuiSize.w1 = w.gocuiSize.w0 + w.realWidth
	w.gocuiSize.h1 = w.gocuiSize.h0 + w.realHeight
	w.showWidgetPlacement(logNow, "textResize()")
}

func (w *cuiWidget) drawView() {
	var err error
	if (w.cuiName == "") {
		log(logError, "drawView() w.cuiName was not set for widget", w)
		w.cuiName = strconv.Itoa(w.id)
	}

	if (w.v != nil) {
		log(logInfo, "drawView() w.v already defined for widget", w)
		v, _ := me.baseGui.View(w.cuiName)
		if (v == nil) {
			log(logError, "drawView() ERROR view does not really exist", w)
			w.v = nil
		} else {
			return
		}
	}

	v, _ := me.baseGui.View(w.cuiName)
	if (v != nil) {
		log(logError, "drawView() already defined for name", w.cuiName)
		w.v = v
		return
	}

	a := w.gocuiSize.w0
	b := w.gocuiSize.h0
	c := w.gocuiSize.w1
	d := w.gocuiSize.h1

	w.v, err = me.baseGui.SetView(w.cuiName, a, b, c, d, 0)
	if err == nil {
		w.showWidgetPlacement(logError, "drawView()")
		log(logError, "drawView() internal plugin error err = nil")
		return
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		w.showWidgetPlacement(logError, "drawView()")
		log(logError, "drawView() internal plugin error error.IS()", err)
		return
	}

	me.baseGui.SetKeybinding(w.v.Name(), gocui.MouseLeft, gocui.ModNone, click)

	w.v.Wrap = true
	fmt.Fprintln(w.v, " " + w.text)

	w.setDefaultWidgetColor()
}
