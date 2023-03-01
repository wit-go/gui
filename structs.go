package gui

import (
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
	master	*Node

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

	widget	toolkit.Widget

	// deprecate these and use toolkit.Widget
	Name   string
	Width  int
	Height int

	// this function is run when there are mouse or keyboard events
	Custom func()

	parent	*Node
	children []*Node

	// is keeping
	// deprecate these things if they don't really need to exist
	// checked   bool
	// text      string
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Window() *Node {
	return n.parent
}

func (n *Node) Append(child *Node) {
	n.children = append(n.children, child)
	if (debugGui) {
		log(debugNode, "child node:")
		child.Dump()
		log(debugNode, "parent node:")
		n.Dump()
	}
}
