// LICENSE: same as the go language itself
// Copyright 2023 WIT.COM

// all structures and variables are local (aka lowercase)
// since the plugin should be isolated to access only
// by functions() to insure everything here is run
// inside a dedicated goroutine

package main

import (
	"fmt"
	"sync"
	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// const delta = 1

// It's probably a terrible idea to call this 'me'
var me config

type config struct {
	highest	int // highest widgetId
	baseGui *gocui.Gui // the main gocui handle
	widgets map[int]*cuiWidget
	callback func(int)
	helpLabel *gocui.View

	defaultBehavior bool
	defaultWidth int
	defaultHeight int
	nextW int	// where the next window or tab flag should go

	bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
	canvas bool // if set to true, the windows are a raw canvas
	menubar bool // for windows
	stretchy bool // expand things like buttons to the maximum size
	padded bool // add space between things like buttons
	margin bool // add space around the frames of windows
}

/*
// This is a map between the widgets in wit/gui and the internal structures of gocui
var viewWidget map[*gocui.View]*toolkit.Widget
var stringWidget map[string]*toolkit.Widget
*/

var (
//	g *gocui.Gui
//	Custom func(string)

	initialMouseX, initialMouseY, xOffset, yOffset int
	globalMouseDown, msgMouseDown, movingMsg bool

//	err error
)

// the gocui way
// the logical size of the widget
// corner starts at in the upper left corner
type rectType struct {
	// this is the gocui way
	w0, h0, w1, h1 int      // left top right bottom
}

/*
type realSizeT struct {
	width, height int
}
*/


type cuiWidget struct {
	id int	// widget ID
	parentId int
	widgetType   toolkit.WidgetType

	name   string // a descriptive name of the widget
	text   string // the current text being displayed

	visable bool // widget types like 'box' are 'false'
	realWidth int // the real width
	realHeight int // the real height
	realSize rectType  // the display size of this widget
	logicalSize rectType  // the logical size. Includes all the child widgets

	nextX	int
	nextY	int

	// horizontal=true  means layout widgets like books on a bookshelf
	// horizontal=false means layout widgets like books in a stack
	horizontal bool `default:false`

	tainted bool
	v *gocui.View
	baseGui *gocui.Gui // use gogui.Manager ? as 'workspaces?'

	// writeMutex protects locks the write process
	writeMutex sync.Mutex

	// deprecate
	// logicalWidth int  `default:8`
	// logicalHeight int  `default:2`
	// rect   rectType
	// current rectType // the logical size. Includes all the child widgets
	// width  int
	// height int
}

// from the gocui devs:
// Write appends a byte slice into the view's internal buffer. Because
// View implements the io.Writer interface, it can be passed as parameter
// of functions like fmt.Fprintf, fmt.Fprintln, io.Copy, etc. Clear must
// be called to clear the view's buffer.

func (w *cuiWidget) Write(p []byte) (n int, err error) {
	w.tainted = true
	w.writeMutex.Lock()
	defer w.writeMutex.Unlock()
	// v.makeWriteable(v.wx, v.wy)
	// v.writeRunes(bytes.Runes(p))
	fmt.Fprintln(w.v, p)
	log(logNow, "widget.Write()")

	return len(p), nil
}
