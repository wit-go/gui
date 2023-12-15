// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"github.com/awesome-gocui/gocui"
)

// this function uses the mouse position to highlight & unhighlight things
// this is run every time the user moves the mouse over the terminal window
func mouseMove(g *gocui.Gui) {
	mx, my := g.MousePosition()
	for _, view := range g.Views() {
		view.Highlight = false
	}
	if v, err := g.ViewByPosition(mx, my); err == nil {
		v.Highlight = true
	}
}

func msgDown(g *gocui.Gui, v *gocui.View) error {
	initialMouseX, initialMouseY = g.MousePosition()
	log(true, "msgDown() X,Y", initialMouseX, initialMouseY)
	if vx, vy, _, _, err := g.ViewPosition("msg"); err == nil {
		xOffset = initialMouseX - vx
		yOffset = initialMouseY - vy
		msgMouseDown = true
	}
	return nil
}

func hideDDview() error {
	w, h := me.baseGui.MousePosition()
	log(true, "hide dropdown menu() view msgMouseDown (w,h) =", w, h)
	if (me.ddview == nil) {
		return gocui.ErrUnknownView
	}
	if (me.ddview.tk.v == nil) {
		return gocui.ErrUnknownView
	}
	me.ddview.SetVisible(false)
	return nil
}

func showDDview() error {
	w, h := me.baseGui.MousePosition()
	log(true, "show dropdown menu() view msgMouseDown (w,h) =", w, h)
	if (me.ddview == nil) {
		return gocui.ErrUnknownView
	}
	if (me.ddview.tk.v == nil) {
		return gocui.ErrUnknownView
	}
	me.ddview.SetVisible(true)
	return nil
}

func mouseUp(g *gocui.Gui, v *gocui.View) error {
	w, h := g.MousePosition()
	log(true, "mouseUp() view msgMouseDown (check here for dropdown menu click) (w,h) =", w, h)
	if (me.ddClicked) {
		me.ddClicked = false
		log(true, "mouseUp() ddview is the thing that was clicked", w, h)
		log(true, "mouseUp() find out what the string is here", w, h, me.ddview.tk.gocuiSize.h1)

		var newZone string = ""
		if (me.ddNode != nil) {
			value := h - me.ddview.tk.gocuiSize.h0 - 1
			log(true, "mouseUp() me.ddview.tk.gocuiSize.h1 =", me.ddview.tk.gocuiSize.h1)
			log(true, "mouseUp() me.ddNode.vals =", me.ddNode.vals)
			valsLen := len(me.ddNode.vals)
			log(true, "mouseUp() value =", value, "valsLen =", valsLen)
			log(true, "mouseUp() me.ddNode.vals =", me.ddNode.vals)
			if ((value >= 0) && (value < valsLen)) {
				newZone = me.ddNode.vals[value]
				log(true, "mouseUp() value =", value, "newZone =", newZone)
			}
		}
		hideDDview()
		if (newZone != "") {
			if (me.ddNode != nil) {
				me.ddNode.SetText(newZone)
				me.ddNode.S = newZone
				me.ddNode.doUserEvent()
			}
		}
		return nil
	}
	/*
	// if there is a drop down view active, treat it like a dialog box and close it
	if (hideDDview() == nil) {
		return nil
	}
	*/
	if msgMouseDown {
		msgMouseDown = false
		if movingMsg {
			movingMsg = false
			return nil
		} else {
			g.DeleteView("msg")
		}
	} else if globalMouseDown {
		globalMouseDown = false
		g.DeleteView("globalDown")
	}
	return nil
}

func mouseDown(g *gocui.Gui, v *gocui.View) error {
	mx, my := g.MousePosition()
	if vx0, vy0, vx1, vy1, err := g.ViewPosition("msg"); err == nil {
		if mx >= vx0 && mx <= vx1 && my >= vy0 && my <= vy1 {
			return msgDown(g, v)
		}
	}
	globalMouseDown = true
	maxX, _ := g.Size()
	test := findUnderMouse()
	msg := fmt.Sprintf("Mouse really down at: %d,%d", mx, my) + "foobar"
	if (test == me.ddview) {
		if (me.ddview.Visible()) {
			log(true, "hide DDview() Mouse really down at:", mx, my)
			hideDDview()
		} else {
			log(true, "show DDview() Mouse really down at:", mx, my)
			showDDview()
		}
		return nil
	}
	x := mx - len(msg)/2
	if x < 0 {
		x = 0
	} else if x+len(msg)+1 > maxX-1 {
		x = maxX - 1 - len(msg) - 1
	}
	log(true, "mouseDown() about to write out message to 'globalDown' view. msg =", msg)
	if v, err := g.SetView("globalDown", x, my-1, x+len(msg)+1, my+1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.WriteString(msg)
	}
	return nil
}
