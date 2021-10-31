/*
Package wit/gui implements a abstraction layer for Go visual elements in
a cross platform way. Right now, this abstraction is built on top of
the GUI toolkit 'andlabs/ui' which does the cross platform support.

A quick overview of the features, some general design guidelines
and principles for how this package should generally work:

	* GUI elements are stored in a tree of nodes
	* When in doubt, it's ok to guess. We will return something close.
	* It tries to make your code simple

Quick Start

This section demonstrates how to quickly get started with spew.  See the
sections below for further details on formatting and configuration options.

To dump a variable with full newlines, indentation, type, and pointer
information use Dump, Fdump, or Sdump:

	package main

	import 	(
		"git.wit.org/wit/gui"
	)

	func main() {
		gui.Main(initGUI)
	}

	// This initializes the first window
	func initGUI() {
		gui.Config.Title = "WIT GUI Window 1"
		gui.Config.Width = 640
		gui.Config.Height = 480
		node1 := gui.NewWindow()
		addDemoTab(node1, "A Simple Tab Demo")
	}

	func addDemoTab(n *gui.Node, title string) {
		newNode := n.AddTab(title, nil)

		groupNode1 := newNode.AddGroup("group 1")
		groupNode1.AddComboBox("demoCombo2", "more 1", "more 2", "more 3")
	}

Configuration Options

Configuration of the GUI is handled by fields in the ConfigType type.  For
convenience, all of the top-level functions use a global state available
via the gui.Config global.

The following configuration options are available:
	* Width
		When creating a new window, this is the width

	* Height
		When creating a new window, this is the height

	* Debug
		When 'true' log more output

GUI Usage

Errors

Since it is possible for custom Stringer/error interfaces to panic, spew
detects them and handles them internally by printing the panic information
inline with the output.  Since spew is intended to provide deep pretty printing
capabilities on structures, it intentionally does not return any errors.
*/
package gui
