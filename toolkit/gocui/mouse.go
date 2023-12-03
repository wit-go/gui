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
	if vx, vy, _, _, err := g.ViewPosition("msg"); err == nil {
		xOffset = initialMouseX - vx
		yOffset = initialMouseY - vy
		msgMouseDown = true
	}
	return nil
}

func mouseUp(g *gocui.Gui, v *gocui.View) error {
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
	msg := fmt.Sprintf("Mouse really down at: %d,%d", mx, my) + "foo\n" + "bar\n"
	x := mx - len(msg)/2
	if x < 0 {
		x = 0
	} else if x+len(msg)+1 > maxX-1 {
		x = maxX - 1 - len(msg) - 1
	}
	if v, err := g.SetView("globalDown", x, my-1, x+len(msg)+1, my+1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.WriteString(msg)
	}
	return nil
}
