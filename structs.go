package gui

import (
	"sync"
	"embed"
	"git.wit.org/wit/gui/toolkit"
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

var me guiConfig

var GuiArg GuiArgs

// This struct can be used with the go-arg package
type GuiArgs struct {
	Gui string `arg:"--gui" help:"Use this gui toolkit [andlabs,gocui,nocui]"`
	GuiDebug bool `arg:"--gui-debug" help:"open the GUI debugger"`
	GuiVerbose bool `arg:"--gui-verbose" help:"enable all logging"`
}

type guiConfig struct {
	initOnce sync.Once

	// This is the master node. The Binary Tree starts here
	rootNode *Node

	// A node off of rootNode for passing debugging flags
	flag	*Node

	counter    int  // used to make unique WidgetId's

	// sets the chan for the plugins to call back too
	guiChan chan toolkit.Action

	// option to pass in compiled plugins as embedded files
	resFS	embed.FS

	// used to beautify logging to Stdout
	depth      int
	prefix     string
}

// The Node is a binary tree. This is how all GUI elements are stored
// simply the name and the size of whatever GUI element exists
type Node struct {
	id     int

	WidgetType	toolkit.WidgetType

	Text string  // what is visable to the user
	Name string  // a name useful for programming

	// used for Windows in toolkits measured in pixels
	width  int
	height int

	// used for anything that needs a range (for example: a slider)
	X	int
	Y	int

	// the grid max width and height
	// ignore max height when there is no space left?
	W	int
	H	int
	// where the next widget should be put in this grid
	NextW	int
	NextH	int

	// if this widget is in a grid, this is the position
	AtW	int
	AtH	int

	// used for values
	I	int
	S	string
	B	bool

	// this function is run when there are mouse or keyboard events
	Custom func()

	parent	*Node
	children []*Node
}
