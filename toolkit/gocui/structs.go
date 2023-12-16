// LICENSE: same as the go language itself
// Copyright 2023 WIT.COM

// all structures and variables are local (aka lowercase)
// since the plugin should be isolated to access only
// by functions() to insure everything here is run
// inside a dedicated goroutine

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"strings"
	"github.com/awesome-gocui/gocui"
//	"git.wit.org/wit/gui/toolkit"
)

// It's probably a terrible idea to call this 'me'
var me config

var showDebug bool = true
var showHelp bool = true
var redoWidgets bool = true

// This is the window that is currently active
var currentWindow *node

type config struct {
	baseGui *gocui.Gui // the main gocui handle
	rootNode *node // the base of the binary tree. it should have id == 0

	ctrlDown *node // shown if you click the mouse when the ctrl key is pressed
	currentWindow *node // this is the current tab or window to show
	logStdout *node // where to show STDOUT
	helpLabel *gocui.View
	ddview *node // the gocui view to select dropdrown lists
	ddClicked bool // the dropdown menu view was clicked
	ddNode *node // the dropdown menu is for this widget

	/*
	// this is the channel we send user events like
	// mouse clicks or keyboard events back to the program
	callback chan toolkit.Action

	// this is the channel we get requests to make widgets
	pluginChan chan toolkit.Action
	*/

	// When the widget has a frame, like a button, it adds 2 lines runes on each side
	// so you need 3 char spacing in each direction to not have them overlap
	// the amount of padding when there is a frame
	FramePadW int `default:"1" dense:"0"`
	FramePadH int `default:"1" dense:"0"`

	PadW int `default:"1" dense:"0"`
	PadH int `default:"1" dense:"0"`

	// how far down to start Window or Tab headings
	WindowW int `default:"8" dense:"0"`
	WindowH int `default:"-1"`
	TabW int `default:"5" dense:"0"`
	TabH int `default:"1" dense:"0"`

	// additional amount of space to put between window & tab widgets
	WindowPadW int `default:"8" dense:"0"`
	TabPadW int `default:"4" dense:"0"`

	// additional amount of space to indent on a group
	GroupPadW int `default:"6" dense:"2"`

	// the raw beginning of each window (or tab)
	RawW int `default:"1"`
	RawH int `default:"5"`

	// offset for the hidden widgets
	FakeW int `default:"20"`

	padded bool // add space between things like buttons
	bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
	canvas bool // if set to true, the windows are a raw canvas
	menubar bool // for windows
	stretchy bool // expand things like buttons to the maximum size
	margin bool // add space around the frames of windows

	// writeMutex protects locks the write process
	writeMutex sync.Mutex

	// used for listWidgets() debugging
	depth int
}

// deprecate these
var (
	initialMouseX, initialMouseY, xOffset, yOffset int
	globalMouseDown, msgMouseDown, movingMsg bool
)

// this is the gocui way
// corner starts at in the upper left corner
type rectType struct {
	w0, h0, w1, h1 int      // left top right bottom
}

func (r *rectType) Width() int {
	return r.w1 - r.w0
}

func (r *rectType) Height() int {
	return r.h1 - r.h0
}

type guiWidget struct {
	// the gocui package variables
	v *gocui.View // this is nil if the widget is not displayed
	cuiName string // what gocui uses to reference the widget

	// the logical size of the widget
	// For example, 40x12 would be the center of a normal terminal
	// size rectType

	// the actual gocui display view of this widget
	// sometimes this isn't visible like with a Box or Grid
	gocuiSize rectType

	isCurrent bool // is this the currently displayed Window or Tab?
	isFake bool // widget types like 'box' are 'false'

	// used to track the size of grids
	widths map[int]int // how tall each row in the grid is
	heights map[int]int // how wide each column in the grid is

	tainted bool
	frame bool

	// for a window, this is currently selected tab
	selectedTab *node

	// what color to use
	color *colorT
}

// from the gocui devs:
// Write appends a byte slice into the view's internal buffer. Because
// View implements the io.Writer interface, it can be passed as parameter
// of functions like fmt.Fprintf, fmt.Fprintln, io.Copy, etc. Clear must
// be called to clear the view's buffer.

func (w *guiWidget) Write(p []byte) (n int, err error) {
	w.tainted = true
	me.writeMutex.Lock()
	defer me.writeMutex.Unlock()
	if (me.logStdout.tk.v == nil) {
		// optionally write the output to /tmp
		s := fmt.Sprint(string(p))
		s = strings.TrimSuffix(s, "\n")
		fmt.Fprintln(outf, s)
		v, _ := me.baseGui.View("msg")
		if (v != nil) {
			// fmt.Fprintln(outf, "found msg")
			me.logStdout.tk.v = v
		}
	} else {
		// display the output in the gocui window
		me.logStdout.tk.v.Clear()

		s := fmt.Sprint(string(p))
		s = strings.TrimSuffix(s, "\n")
		tmp := strings.Split(s, "\n")
		outputS = append(outputS, tmp...)
		if (len(outputS) > outputH) {
			l := len(outputS) - outputH
			outputS = outputS[l:]
		}
		fmt.Fprintln(me.logStdout.tk.v, strings.Join(outputS, "\n"))
	}

	return len(p), nil
}

func Set(ptr interface{}, tag string) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		log(logError, "Set() Not a pointer", ptr, "with tag =", tag)
		return fmt.Errorf("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		defaultVal := t.Field(i).Tag.Get(tag)
		name := t.Field(i).Name
		// log("Set() try name =", name, "defaultVal =", defaultVal)
		setField(v.Field(i), defaultVal, name)
	}
	return nil
}

func setField(field reflect.Value, defaultVal string, name string) error {

	if !field.CanSet() {
		// log("setField() Can't set value", field, defaultVal)
		return fmt.Errorf("Can't set value\n")
	} else {
		log("setField() Can set value", name, defaultVal)
	}

	switch field.Kind() {
	case reflect.Int:
		val, _ := strconv.Atoi(defaultVal)
		field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	case reflect.Bool:
		if defaultVal == "true" {
			field.Set(reflect.ValueOf(true))
		} else {
			field.Set(reflect.ValueOf(false))
		}
	}

	return nil
}
