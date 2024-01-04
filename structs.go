package gui

import (
	"sync"
	"embed"
	"go.wit.com/gui/toolkits"
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

type guiConfig struct {
	initOnce sync.Once

	// This is the master node. The Binary Tree starts here
	rootNode *Node

	// if the user prefers new windows or 'windows within windows' tabs
	makeTabs bool

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

	// for NewLabel("hello"), Text = 'hello'
	Text string  // what is visable to the user

	// for NewLabel("hello"), if Name = 'HELLO'
	// this can programatically identify the widget
	// The name must be unique
	Name string  // a name useful for debugging

	// used for Windows in toolkits measured in pixels
	width  int
	height int

	// used for anything that needs a range (for example: a slider)
	X	int
	Y	int

	// the grid widget max width and height
	// the max height can be implemented in the toolkit plugin
	// to restrict the number of rows to display
	W	int
	H	int

	// where the next widget should be put in this grid
	NextW	int
	NextH	int

	// if this widget is in a grid, this is the position of a widget
	AtW	int
	AtH	int

	// the current widget value.
	I	int
	S	string
	B	bool

	// this function is run when there are mouse or keyboard events
	Custom func()

	parent	*Node
	children []*Node
}
