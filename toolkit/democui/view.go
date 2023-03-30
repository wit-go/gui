package main

import (
	"fmt"
	"errors"
	"strconv"

	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

var adjusted bool = false

// expands the logical size of the parents
func (w *cuiWidget) setParentLogical(p *cuiWidget) {
	if (w.visable) {
		// expand the parent logicalsize to include the widget realSize
		if (p.logicalSize.w0 > w.realSize.w0) {
			p.logicalSize.w0 = w.realSize.w0
			adjusted = true
		}
		if (p.logicalSize.h0 > w.realSize.h0) {
			p.logicalSize.h0 = w.realSize.h0
			adjusted = true
		}
		if (p.logicalSize.w1 < w.realSize.w1) {
			p.logicalSize.w1 = w.realSize.w1
			adjusted = true
		}
		if (p.logicalSize.h1 < w.realSize.h1) {
			p.logicalSize.h1 = w.realSize.h1
			adjusted = true
		}
	} else {
		// expand the parent logicalsize to include the widget logicalsize
		if (p.logicalSize.w0 > w.logicalSize.w0) {
			p.logicalSize.w0 = w.logicalSize.w0
			adjusted = true
		}
		if (p.logicalSize.h0 > w.logicalSize.h0) {
			p.logicalSize.h0 = w.logicalSize.h0
			adjusted = true
		}
		if (p.logicalSize.w1 < w.logicalSize.w1) {
			p.logicalSize.w1 = w.logicalSize.w1
			adjusted = true
		}
		if (p.logicalSize.h1 < w.logicalSize.h1) {
			p.logicalSize.h1 = w.logicalSize.h1
			adjusted = true
		}
	}
	if (w.visable) {
		// adjust the widget realSize to the top left corner of the logicalsize
		if (w.logicalSize.w0 > w.realSize.w0) {
			w.realSize.w0 = w.logicalSize.w0
			w.realSize.w1 = w.realSize.w0 + w.realWidth
			adjusted = true
		}
		if (w.logicalSize.h0 > w.realSize.h0) {
			w.realSize.h0 = w.logicalSize.h0
			w.realSize.h1 = w.realSize.h0 + w.realHeight
			adjusted = true
		}
	}
	w.showWidgetPlacement(logNow, "setParentLogical() widget")
	p.showWidgetPlacement(logNow, "setParentLogical() parent")
	if (w.id == 0) || (p.id == 0) {
		// stop resizing when you hit the root widget
		return
	}
	// pass the logical resizing up
	pP := me.widgets[p.parentId]
	if (pP != nil) {
		p.setParentLogical(pP)
	}
}

var fakeStartWidth int = 80
var fakeStartHeight int = 0
func (w *cuiWidget) setFake() {
	if (w.visable) {
		return
	}
	// setup fake labels for non-visable things off screen
	w.realWidth = me.defaultWidth
	w.realHeight = me.defaultHeight
	w.realSize.w0 = fakeStartWidth
	w.realSize.h0 = fakeStartHeight
	w.realSize.w1 =  w.realSize.w0 + me.defaultWidth
	w.realSize.h1 =  w.realSize.h0 + me.defaultHeight
	fakeStartHeight += 2
	w.showWidgetPlacement(logNow, "setFake()")
}

func drawView(w *cuiWidget) *gocui.View {
	var newName string = ""
	newName = strconv.Itoa(w.id)
	if (me.baseGui == nil) {
		log(logError, "drawView() me.baseGui == nil")
		return nil
	}

	a := w.realSize.w0
	b := w.realSize.h0
	c := w.realSize.w1
	d := w.realSize.h1
	v, err := me.baseGui.SetView(newName, a, b, c, d, 0)
	if err == nil {
		log(logError, "drawView() internal plugin error err = nil")
		return nil
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		log(logError, "drawView() internal plugin error error.IS()", err)
		return nil
	}
	w.v = v

	return v
}

func boxedPlace(w *cuiWidget) {
	t := len(w.name)
	if (w.id == 0) {
		w.realWidth = 0
		w.realHeight = 0
		return
	}
	p := me.widgets[w.parentId]
	if (p == nil) {
		log(logError, "ERRRRRRRRRRRORRRRRRRRRRRRR: parentId widget == nil")
		return
	}

	switch p.widgetType {
	case toolkit.Box:
		w.realWidth = t + 3
		w.realHeight = me.defaultHeight
		w.realSize.w0 = p.nextX
		w.realSize.h0 = p.nextY
		w.realSize.w1 = p.nextX + w.realWidth
		w.realSize.h1 = p.nextY + w.realHeight

		w.logicalSize.w0 = p.nextX
		w.logicalSize.h0 = p.nextY
		w.logicalSize.w1 = p.nextX + w.realWidth
		w.logicalSize.h1 = p.nextY + w.realHeight

		w.nextX = p.nextX
		w.nextY = p.nextY
		if (w.horizontal) {
			log(logNow, "PARENT BOX IS HORIZONTAL")
			p.nextX += w.realWidth
		} else {
			log(logNow, "PARENT BOX IS VERTICAL")
			p.nextY += w.realHeight
		}
	case toolkit.Group:
		w.realWidth = t + 3
		w.realHeight = me.defaultHeight

		w.realSize.w0 = p.nextX
		w.realSize.h0 = p.nextY
		w.realSize.w1 = p.nextX + w.realWidth
		w.realSize.h1 = p.nextY + w.realHeight

		w.logicalSize.w0 = p.nextX
		w.logicalSize.h0 = p.nextY
		w.logicalSize.w1 = p.nextX + w.realWidth
		w.logicalSize.h1 = p.nextY + w.realHeight

		w.nextX = w.logicalSize.w0 + 3 // default group padding
		w.nextY = w.logicalSize.h1

		// increment parent
		p.nextY += w.realHeight
	default:
		w.realWidth = t + 3
		w.realHeight = me.defaultHeight
		w.realSize.w0 = p.nextX
		w.realSize.h0 = p.nextY
		w.realSize.w1 = w.realSize.w0 + w.realWidth
		w.realSize.h1 = w.realSize.h0 + w.realHeight

		// increment parent
		p.nextY += w.realHeight
	}
	p.showWidgetPlacement(logNow, "bP parent")
	w.showWidgetPlacement(logNow, "bP widget")
}

func findPlace(w *cuiWidget, a *toolkit.Action) {
	t := len(w.name)
	w.visable = true
	switch w.widgetType {
	case toolkit.Root:
		w.visable = false
		w.setFake()
		w.showWidgetPlacement(logNow, "Root:")
	case toolkit.Flag:
		w.visable = false
		w.setFake()
		w.showWidgetPlacement(logNow, "Flag:")
	case toolkit.Window:
		w.realWidth = t + 3
		w.realHeight = me.defaultHeight

		w.realSize.w0 = me.nextW
		w.realSize.h0 = 0 
		w.realSize.w1 = w.realSize.w0 + w.realWidth
		w.realSize.h1 = w.realHeight

		w.logicalSize.w0 = me.nextW
		w.logicalSize.h0 = 0
		w.logicalSize.w1 = w.logicalSize.w0 + w.realWidth
		w.logicalSize.h1 = w.realHeight

		w.nextX = w.logicalSize.w0 + t // default group padding
		w.nextY = w.logicalSize.h1

		me.nextW += w.realWidth
		w.showWidgetPlacement(logNow, "window:")
	case toolkit.Tab:
		w.realWidth = t + 3
		w.realHeight = me.defaultHeight

		w.realSize.w0 = me.nextW
		w.realSize.h0 = 0
		w.realSize.w1 = w.realSize.w0 + w.realWidth
		w.realSize.h1 = w.realHeight

		w.logicalSize.w0 = me.nextW
		w.logicalSize.h0 = 0
		w.logicalSize.w1 = w.logicalSize.w0 + w.realWidth
		w.logicalSize.h1 = w.realHeight

		w.nextX = w.logicalSize.w0 + t // default group padding
		w.nextY = w.logicalSize.h1
		me.nextW += w.realWidth
		w.showWidgetPlacement(logNow, "tab:")
	case toolkit.Grid:
		p := me.widgets[w.parentId]
		w.horizontal = a.B
		w.visable = false
		w.setFake()

		if (p == nil) {
			log(logError, "ERRRRRRRRRRRORRRRRRRRRRRRR: parentId widget == nil")
			return
		}
		w.logicalSize.w0 = p.nextX
		w.logicalSize.h0 = p.nextY
		w.logicalSize.w1 = p.nextX
		w.logicalSize.h1 = p.nextY

		w.nextX = p.nextX
		w.nextY = p.nextY
		w.showWidgetPlacement(logNow, "grid:")
	case toolkit.Box:
		p := me.widgets[w.parentId]
		w.horizontal = a.B
		w.visable = false
		w.setFake()

		if (p == nil) {
			log(logError, "ERRRRRRRRRRRORRRRRRRRRRRRR: parentId widget == nil")
			return
		}
		w.logicalSize.w0 = p.nextX
		w.logicalSize.h0 = p.nextY
		w.logicalSize.w1 = p.nextX
		w.logicalSize.h1 = p.nextY

		w.nextX = p.nextX
		w.nextY = p.nextY
		w.showWidgetPlacement(logNow, "box:")
	case toolkit.Group:
		p := me.widgets[w.parentId]
		w.horizontal = a.B
		w.visable = false
		w.setFake()

		if (p == nil) {
			log(logError, "ERRRRRRRRRRRORRRRRRRRRRRRR: parentId widget == nil")
			return
		}
		w.logicalSize.w0 = p.nextX
		w.logicalSize.h0 = p.nextY
		w.logicalSize.w1 = p.nextX
		w.logicalSize.h1 = p.nextY

		w.nextX = p.nextX
		w.nextY = p.nextY
		w.showWidgetPlacement(logNow, "group:")
	default:
		boxedPlace(w)
	}
}

func place(w *cuiWidget, a *toolkit.Action) {
	log(logInfo, "place() START")
	findPlace(w, a)
	v := drawView(w)
	if (v == nil) {
		log(logError, "place() drawView(w) returned nil")
		return
	}
	me.baseGui.SetKeybinding(v.Name(), gocui.MouseLeft, gocui.ModNone, click)

	v.Wrap = true
	fmt.Fprintln(v, " " + w.name)

	w.SetDefaultWidgetColor()

	log(logInfo, "place() END")
	return
}
