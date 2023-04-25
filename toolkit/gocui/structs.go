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
	defaultWidth int
	defaultHeight int
	// nextW int	// where the next window or tab flag should go

	// the amount to put between winodw tab widgets
	TabPadW int `default:"4" dense:"2"`
	// PadH int `default:"3" dense:"2"`

	// the raw beginning of each window (or tab)
	rawW int `default:"7"`
	JWC int `default:"7"`
	rawH int `default:"3"`

	bookshelf bool // do you want things arranged in the box like a bookshelf or a stack?
	canvas bool // if set to true, the windows are a raw canvas
	menubar bool // for windows
	stretchy bool // expand things like buttons to the maximum size
	padded bool // add space between things like buttons
	margin bool // add space around the frames of windows

	horizontalPadding int
	groupPadding int `default:"6" dense:"2"` // this is supposed to be how far to indent to the left
	buttonPadding int `default:"4" dense:"3"` // if 3, buttons slightly overlap
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

	// the actual size
	width int
	height int

	// this is the gocui way
	w0, h0, w1, h1 int      // left top right bottom
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
