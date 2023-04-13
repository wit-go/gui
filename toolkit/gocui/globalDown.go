// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func globalDown(g *gocui.Gui, v *gocui.View) error {
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
