// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"github.com/awesome-gocui/gocui"
)

const delta = 1

var (
	g *gocui.Gui
	Custom func(string)

	initialMouseX, initialMouseY, xOffset, yOffset int
	globalMouseDown, msgMouseDown, movingMsg bool

	views   = []string{}
	curView = -1
	idxView = 0
	currentX = 5
	currentY = 2
	groupSize = 0
	baseGui *gocui.Gui
	helpLabel *gocui.View
	err error
	ch chan(func ())
	outf *os.File
)
