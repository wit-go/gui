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
	"github.com/awesome-gocui/gocui"
	"git.wit.org/wit/gui/toolkit"
)

// It's probably a terrible idea to call this 'me'
var me config

type config struct {
	baseGui *gocui.Gui // the main gocui handle
	rootNode *cuiWidget // the base of the binary tree. it should have id == 0
	ctrlDown *cuiWidget // shown if you click the mouse when the ctrl key is pressed

	// this is the channel we send user events like
	// mouse clicks or keyboard events back to the program
	callback chan toolkit.Action

	// this is the channel we get requests to make widgets
	pluginChan chan toolkit.Action

	helpLabel *gocui.View

	DefaultBehavior bool `default:"true"`

	// Buttons, Group, Tabs, Windows, etc are by default assumed to be a single line
	// as a default, we make buttons 8 chars wide
	DefaultWidth int `default:"8"`
	DefaultHeight int `default:"1"`

	// When the widget has a frame, like a button, it adds 2 lines runes on each side
	// so you need 3 char spacing in each direction to not have them overlap
	// the amount of padding when there is a frame
	FramePadW int `default:"4" dense:"0"`
	FramePadH int `default:"1" dense:"0"`

	PadW int `default:"1" dense:"0"`
	PadH int `default:"1" dense:"0"`

	// additional amount of space to put between window & tab widgets
	WindowPadW int `default:"8" dense:"0"`
	TabPadW int `default:"4" dense:"0"`

	// how far down to start Window or Tab headings
	WindowW int `default:"20" dense:"0"`
	WindowH int `default:"0" dense:"0"`
	TabW int `default:"2" dense:"0"`
	TabH int `default:"2" dense:"0"`

	// additional amount of space to indent on a group
	GroupPadW int `default:"6" dense:"2"`

	// the raw beginning of each window (or tab)
	RawW int `default:"7"`
	RawH int `default:"3"`

	// offset for the hidden widgets
	DevelOffsetW int `default:"20"`

	bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
	canvas bool // if set to true, the windows are a raw canvas
	menubar bool // for windows
	stretchy bool // expand things like buttons to the maximum size
	padded bool // add space between things like buttons
	margin bool // add space around the frames of windows
}

var (
	initialMouseX, initialMouseY, xOffset, yOffset int
	globalMouseDown, msgMouseDown, movingMsg bool
)

// the gocui way
// the logical size of the widget
// corner starts at in the upper left corner
type rectType struct {
	// where the widget should calculate it's existance from
	// startW int
	// startH int

	// the is a shortcut to access
//	width int  // this is always w1 - w0
	height int // this is always h1 - h0

	// this is the gocui way
	w0, h0, w1, h1 int      // left top right bottom
}

func (r *rectType) Width() int {
	return r.w1 - r.w0
}

func (r *rectType) Height() int {
	return r.h1 - r.h0
}

type cuiWidget struct {
	id int	// widget ID
	// parentId int
	widgetType   toolkit.WidgetType

	name   string // a descriptive name of the widget
	text   string // the current text being displayed
	cuiName string // what gocui uses to reference the widget

	vals []string // dropdown menu options

	// visable bool // track if it's currently supposed to be shown
	isFake bool // widget types like 'box' are 'false'

	// where the widget's real corner is 
	startW int
	startH int

	// where the next child should be placed
	nextW	int
	nextH	int

	// the widget size to reserve or things will overlap
	realWidth int
	realHeight int

	gocuiSize rectType  // the display size of this widget
	// logicalSize rectType  // the logical size. Includes all the child widgets

	// used to track the size of grids
	widths map[int]int // how tall each row in the grid is
	heights map[int]int // how wide each column in the grid is

	// deprecate // where in the parent grid this widget should go
	parentW int
	parentH int

	// things from toolkit/action
	b bool
	i int
	s string
	x int
	y int
	width int
	height int

	// horizontal=true  means layout widgets like books on a bookshelf
	// horizontal=false means layout widgets like books in a stack
	horizontal bool `default:false`

	tainted bool
	v *gocui.View
	frame bool

	// writeMutex protects locks the write process
	writeMutex sync.Mutex

	parent	*cuiWidget
	children []*cuiWidget
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
	if (w.v == nil) {
		return
	}
	w.v.Clear()
	fmt.Fprintln(w.v, p)
	log(logNow, "widget.Write()", p)

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
		// name := t.Field(i).Name
		// log("Set() try name =", name, "defaultVal =", defaultVal)
		setField(v.Field(i), defaultVal)
	}
	return nil
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		// log("setField() Can't set value", field, defaultVal)
		return fmt.Errorf("Can't set value\n")
	} else {
		log("setField() Can set value", field, defaultVal)
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
