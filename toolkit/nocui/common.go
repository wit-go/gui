package main

/*
	These code should be common to all gui plugins

	There are some helper functions that are probably going to be
	the same everywhere. Mostly due to handling the binary tree structure
	and the channel communication

	For now, it's just a symlink to the 'master' version in
	./toolkit/nocui/common.go
*/

import (
	"go.wit.com/gui/toolkit"
)

// this is the channel we send user events like
// mouse clicks or keyboard events back to the program
var callback chan toolkit.Action

// this is the channel we get requests to make widgets
var pluginChan chan toolkit.Action

type node struct {
	parent	*node
	children []*node

	WidgetId	int	// widget ID
	WidgetType	toolkit.WidgetType
	ParentId	int	// parent ID

	Name   string
	Text   string

	// This is how the values are passed back and forth
	// values from things like checkboxes & dropdown's
	B	bool
	I	int
	S	string

	A	any // switch to this or deprecate this? pros/cons?

	// This is used for things like a slider(0,100)
	X      int
	Y      int

	// This is for the grid size & widget position
	W      int
	H      int
	AtW    int
	AtH    int

	vals []string // dropdown menu items

	// horizontal=true  means layout widgets like books on a bookshelf
	// horizontal=false means layout widgets like books in a stack
	horizontal bool `default:false`

	hasTabs bool // does the window have tabs?
	currentTab bool // the visible tab

	// the internal plugin toolkit structure
	// in the gtk plugin, it has gtk things like margin & border settings
	// in the text console one, it has text console things like colors for menus & buttons
	tk *guiWidget
}

// searches the binary tree for a WidgetId
func (n *node) findWidgetId(id int) *node {
	if (n == nil) {
		return nil
	}

	if n.WidgetId == id {
		return n
	}

	for _, child := range n.children {
		newN := child.findWidgetId(id)
		if (newN != nil) {
			return newN
		}
	}
	return nil
}

func (n *node) doUserEvent() {
	if (callback == nil) {
		log(logError, "doUserEvent() callback == nil", n.WidgetId)
		return
	}
	var a toolkit.Action
	a.WidgetId = n.WidgetId
	a.Name = n.Name
	a.Text = n.Text
	a.S = n.S
	a.I = n.I
	a.B = n.B
	a.ActionType = toolkit.User
	log(logInfo, "doUserEvent() START: send a user event to the callback channel")
	callback <- a
	log(logInfo, "doUserEvent() END:   sent a user event to the callback channel")
	return
}

func addNode(a *toolkit.Action) *node {
	n := new(node)
	n.WidgetType = a.WidgetType
	n.WidgetId = a.WidgetId
	n.ParentId = a.ParentId

	// copy the data from the action message
	n.Name = a.Name
	n.Text = a.Text
	n.I = a.I
	n.S = a.S
	n.B = a.B

	n.X = a.X
	n.Y = a.Y

	n.W = a.W
	n.H = a.H
	n.AtW = a.AtW
	n.AtH = a.AtH

	// store the internal toolkit information
	n.tk = initWidget(n)
	// n.tk = new(guiWidget)

	if (a.WidgetType == toolkit.Root) {
		log(logInfo, "addNode() Root")
		return n
	}

	if (me.rootNode.findWidgetId(a.WidgetId) != nil) {
		log(logError, "addNode() WidgetId already exists", a.WidgetId)
		return me.rootNode.findWidgetId(a.WidgetId)
	}

	// add this new widget on the binary tree
	n.parent = me.rootNode.findWidgetId(a.ParentId)
	if n.parent != nil {
		n.parent.children = append(n.parent.children, n)
		//w := n.tk
		//w.parent = n.parent.tk
		//w.parent.children = append(w.parent.children, w)
	}
	return n
}

// Other goroutines must use this to access the GUI
//
// You can not acess / process the GUI thread directly from
// other goroutines. This is due to the nature of how
// Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
//
// this sets the channel to send user events back from the plugin
func Callback(guiCallback chan toolkit.Action) {
	callback = guiCallback
}

func PluginChannel() chan toolkit.Action {
	return pluginChan
}
