package gui

import (
	"sync"
	"embed"
	"go.wit.com/gui/widget"
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
// windows and makes each window a 'tab' in a single window.
//
// Reminder from Goals: This is for simple GUI's.
// For example, a "Mouse Control Panel" not the GIMP or blender.
//

var me guiConfig

// Range(1, 10) includes the values 1 and 10
// almost all toolkits use integers so there doesn't
// seem to be a good idea to use 'type any' here as it
// just makes things more complicated for no good reason
type Range struct {
	Low int
	High int
}

type List []string

type guiConfig struct {
	initOnce sync.Once

	// This is the master node. The Binary Tree starts here
	rootNode *Node

	// A node off of rootNode for passing debugging flags
	flag	*Node

	counter    int  // used to make unique WidgetId's

	// sets the chan for the plugins to call back too
	guiChan chan widget.Action

	// option to pass in compiled plugins as embedded files
	resFS	embed.FS

	// used to beautify logging to Stdout
	depth      int
	prefix     string
}

// The Node is a binary tree. This is how all GUI elements are stored
// simply the name and the size of whatever GUI element exists
type Node struct {
	id	int // should be unique
	hidden	bool // Sierpinski Carpet mode. It's there, but you can't see it.
	pad	bool // the toolkit may use this. it's up to the toolkit
	margin	bool // the toolkit may use this. it's up to the toolkit
	expand	bool // the toolkit may use this. it's up to the toolkit

	WidgetType	widget.WidgetType

	// the current widget value.
	value	any

	// this can programatically identify the widget
	// The name must be unique
	progname string  // a name useful for debugging

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


	// this function is run when there are mouse or keyboard events
	Custom func()

	parent	*Node
	children []*Node
}
