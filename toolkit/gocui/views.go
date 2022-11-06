// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
)

func newView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	name := fmt.Sprintf("v%v", idxView)
	v, err := g.SetView(name, maxX/2-5, maxY/2-5, maxX/2+5, maxY/2+5, 0)
	if err == nil {
		return err
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		return err
	}

	v.Wrap = true
	fmt.Fprintln(v, strings.Repeat(name+" ", 30))

	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	return nil
}

func delView(g *gocui.Gui) error {
	if len(views) <= 1 {
		return nil
	}

	if err := g.DeleteView(views[curView]); err != nil {
		return err
	}
	views = append(views[:curView], views[curView+1:]...)

	return nextView(g, false)
}

func nextView(g *gocui.Gui, disableCurrent bool) error {
	next := curView + 1
	if next > len(views)-1 {
		next = 0
	}

	if _, err := g.SetCurrentView(views[next]); err != nil {
		return err
	}

	curView = next
	return nil
}

func moveView(g *gocui.Gui, v *gocui.View, dx, dy int) error {
	name := v.Name()
	x0, y0, x1, y1, err := g.ViewPosition(name)
	if err != nil {
		return err
	}
	log.Println(x0, y0, x1, y1)
	if _, err := g.SetView(name, x0+dx, y0+dy, x1+dx, y1+dy, 0); err != nil {
		return err
	}
	x0, y0, x1, y1, err = g.ViewPosition(name)
	log.Println(x0, y0, x1, y1)
	return nil
}
