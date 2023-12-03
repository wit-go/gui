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

var topX int = 2
var bottomX int = 20
var topY int = 2
var bottomY int = 7

func newJ(g *gocui.Gui) error {
	// maxX, maxY := g.Size()
	name := fmt.Sprintf("jcarr %v test ", idxView)
	v, err := g.SetView(name, topX, topY, bottomX, bottomY, 0)
	if err == nil {
		return err
	}
	if !errors.Is(err, gocui.ErrUnknownView) {
		return err
	}

	v.Wrap = true
	fmt.Fprintln(v, name)
	fmt.Fprintln(v, strings.Repeat("foo\n", 2))
	// fmt.Fprintln(v, strings.Repeat(name+" ", 30))
	log.Println("newJ added a new view", v.Name())

	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
	return nil
}
