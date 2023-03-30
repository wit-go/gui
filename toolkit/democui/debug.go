package main

import (
	"fmt"
	"strconv"

	"git.wit.org/wit/gui/toolkit"
	"github.com/awesome-gocui/gocui"
)

// var debugError bool = true

// This is important. This sets the defaults for the gui. Without this, there isn't correct padding, etc
func setDefaultBehavior(s bool) {
	me.defaultBehavior = s
	if (me.defaultBehavior) {
		log(logInfo, "Setting this toolkit to use the default behavior.")
		log(logInfo, "This is the 'guessing' part as defined by the wit/gui 'Principles'. Refer to the docs.")
		me.stretchy = false
		me.padded = true
		me.menubar = true
		me.margin = true
		me.canvas = false
		me.bookshelf = true // 99% of the time, things make a vertical stack of objects
	} else {
		log(logInfo, "This toolkit is set to ignore the default behavior.")
	}
}

func actionDump(b bool, a *toolkit.Action) {
	if (a == nil) {
		log(b, "action = nil")
		return
	}

	log(b, "a.Name             =", a.Name)
	log(b, "a.Text             =", a.Text)
	log(b, "a.WidgetId         =", a.WidgetId)
	log(b, "a.ParentId         =", a.ParentId)
	log(b, "a.B                =", a.B)
	log(b, "a.S                =", a.S)
	widgetDump(b, a.Widget)
}

func widgetDump(b bool, w *toolkit.Widget) {
	if (w == nil) {
		log(b, "widget = nil")
		return
	}

	/*
	log(b, "widget.Name        =", w.Name)
	log(b, "widget.Type        =", w.Type)
	log(b, "widget.Custom      =", w.Custom)
	log(b, "widget.B           =", w.B)
	log(b, "widget.I           =", w.I)
	log(b, "widget.Width       =", w.Width)
	log(b, "widget.Height      =", w.Height)
	log(b, "widget.X           =", w.X)
	log(b, "widget.Y           =", w.Y)
	*/
}

func dumpWidgets(g *gocui.Gui, v *gocui.View) {
	for _, view := range g.Views() {
		i, _ := strconv.Atoi(view.Name())
		if (me.widgets[i] != nil) {
			continue
		}
		log(logNow, "dump() not a widget. view.Name =", view.Name())
	}

	for i := 0; i <= me.highest; i++ {
		w := me.widgets[i]
		if (w == nil) {
			continue
		}
		w.showWidgetPlacement(logNow, "")

		if (w.v == nil) {
			log(logError, "dump()        ERROR w.v == nil")
		} else {
			if (strconv.Itoa(i) != w.v.Name()) {
				log(logError, "dump()        ERROR unequal str.Itoa(i) =", strconv.Itoa(i))
				log(logError, "dump()        ERROR unequal w.v.Name()  =", w.v.Name())
			}
		}
	}
}

func (w *cuiWidget) showWidgetPlacement(b bool, s string) {
	log(b, "dump()", s,
		fmt.Sprintf("(wId,pId)=(%3d,%3d)", w.id, w.parentId),
		fmt.Sprintf("real()=(%3d,%3d,%3d,%3d)", w.realSize.w0, w.realSize.h0, w.realSize.w1, w.realSize.h1),
		"next()=(", w.nextX, ",", w.nextY, ")",
		"logical()=(", w.logicalSize.w0, ",", w.logicalSize.h0, ",", w.logicalSize.w1, ",", w.logicalSize.h1, ")",
		w.widgetType, ",", w.name, "text=", w.text)

	if (w.realWidth != (w.realSize.w1 - w.realSize.w0)) {
		log(b, "dump()", s,
			"badsize()=(", w.realWidth, ",", w.realHeight, ")",
			"badreal()=(", w.realSize.w0, ",", w.realSize.h0, ",", w.realSize.w1, ",", w.realSize.h1, ")",
			w.widgetType, ",", w.name)
	}
	if (w.realHeight != (w.realSize.h1 - w.realSize.h0)) {
		log(b, "dump()", s,
			"badsize()=(", w.realWidth, ",", w.realHeight, ")",
			"badreal()=(", w.realSize.w0, ",", w.realSize.h0, ",", w.realSize.w1, ",", w.realSize.h1, ")",
			w.widgetType, ",", w.name)
	}
}
