package gui

import (
	"git.wit.org/wit/gui/toolkit"
	"sync"
)

//
// All GUI Data Structures and functions that are external
// within the toolkit/ abstraction layer
//
// More than one Window does not exist in every GUI situtaion and
// can never be. On many toolkits you have to have 'tabs', like
// Native Windows and MacOS toolkits
//
// If that is the case, this code abstracts the concept of
// windows and makes each window a 'tabs' in a single window.
//
// Reminder from Goals: This is for simple GUI's.
// For example, a "Mouse Control Panel" not the GIMP or blender.
//

var Config GuiConfig

// This struct can be used with the go-arg package
type GuiArgs struct {
	Toolkit []string `arg:"--toolkit" help:"The order to attempt loading plugins [gocui,andlabs,gtk,qt]"`
	GuiDebug bool `arg:"--gui-debug" help:"debug the GUI"`
	GuiVerbose bool `arg:"--gui-verbose" help:"enable all GUI flags"`
}

// var verbose GuiArgs.GuiDebug

type GuiConfig struct {
	// This is the master node. The Binary Tree starts here
	rootNode *Node

	// A node off of rootNode for passing debugging flags
	flag	*Node

	// These are shortcuts to pass default values to make a new window
	Title      string
	Width      int
	Height     int
	Exit       func(*Node)

	// hacks
	depth      int
	counter    int  // used to make unique ID's
	prefix     string

	ActionCh1 chan int
	ActionCh2 chan int
}

// The Node is a binary tree. This is how all GUI elements are stored
// simply the name and the size of whatever GUI element exists
type Node struct {
	id     int
	initOnce sync.Once

	widget	toolkit.Widget
	WidgetType	toolkit.WidgetType

//	Title  string  // what is visable to the user
//	Desc   string  // a name useful for programming

	Text string  // what is visable to the user
	Name string  // a name useful for programming

	// used for Windows
	Width  int
	Height int

	// used for anything that needs a range
	X	int
	Y	int

	// used for grids and tables
	NextX	int
	NextY	int

	// used for values
	I	int
	S	string
	B	bool

	// this function is run when there are mouse or keyboard events
	Custom func()

	parent	*Node
	children []*Node
}
