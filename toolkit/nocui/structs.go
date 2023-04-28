package main

import "git.wit.org/wit/gui/toolkit"

var callback chan toolkit.Action

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

	// the internal plugin toolkit structure
	tk *nocuiT
}

// stores the raw toolkit internals
type nocuiT struct {
	Width  int
	Height int

	c int
	val map[int]string
}
