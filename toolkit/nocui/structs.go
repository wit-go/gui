package main

// import "go.wit.com/gui/toolkit"

// stores the raw toolkit internals
type guiWidget struct {
	Width  int
	Height int

	c int
	val map[int]string
}

// It's probably a terrible idea to call this 'me'
var me config

type config struct {
	rootNode *node // the base of the binary tree. it should have id == 0
}
