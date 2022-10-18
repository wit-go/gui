/*
Package gui implements a abstraction layer for Go visual elements in
a cross platform way.

A quick overview of the features, some general design guidelines
and principles for how this package should generally work:

	* GUI elements are stored in a tree of nodes
	* When in doubt, it's ok to guess. We will return something close.
	* It tries to make your code simple

Quick Start

This section demonstrates how to quickly get started with spew.  See the
sections below for further details on formatting and configuration options.

	// This creates a simple hello world window
	package main

	import 	(
		"git.wit.org/wit/gui"
	)

	func main() {
		gui.Main(initGUI)
	}

	// This initializes the first window
	func initGUI() {
		gui.Config.Title = "Hello World golang wit/gui Window"
		gui.Config.Width = 640
		gui.Config.Height = 480
		node1 := gui.NewWindow()
		addDemoTab(node1, "A Simple Tab Demo")
		addDemoTab(node1, "A Second Tab")
	}

	func addDemoTab(n *gui.Node, title string) {
		newNode := n.AddTab(title, nil)

		groupNode1 := newNode.NewGroup("group 1")
		groupNode1.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
	}

Toolkit Usage

Right now, this abstraction is built on top of the go package 'andlabs/ui'
which does the cross platform support.
The next step is to intent is to allow this to work directly against GTK and QT.

It should be able to add Fyne, WASM, native macos & windows, android and
hopefully also things like libSDL, faiface/pixel, slint

Errors

Since it is possible for custom Stringer/error interfaces to panic, spew
detects them and handles them internally by printing the panic information
inline with the output.  Since spew is intended to provide deep pretty printing
capabilities on structures, it intentionally does not return any errors.

Debugging

To dump variables with full newlines, indentation, type, and pointer
information this uses spew.Dump()

*/
package gui
