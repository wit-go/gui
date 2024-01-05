/*

Package gui implements a abstraction layer for Go visual elements.

Definitions:

	* Toolkit: the underlying GUI library (MacOS gui, Windows gui, gtk, qt, etc)
	* Node: A binary tree of all the underlying widgets

Principles:

	* Make code using this package simple to use
	* Hide complexity internally here
	* Isolate the GUI toolkit
	* Widget names should try to match [Wikipedia Graphical widget]
	* When in doubt, search upward in the binary tree
	* It's ok to guess. Try to do something sensible.

Debian Build

This worked on debian sid (mate-desktop) on 2023/12/03
I didn't record the dependances needed (gtk-dev)

	export GO111MODULE="off"
	go get go.wit.com/gui

When I am working on toolkit plugins, then I work
directly from ~/go/src/go.wit.com/gui/

Hello World Example

	// This creates a simple hello world window
	package main

	import 	(
		"log"
		"go.wit.com/gui"
	)

	var myGui *gui.Node // This is the beginning of the binary tree of widgets

	// go will sit here until the window exits
	func main() {
		myGui = gui.New().Default()

		helloworld()
	}

	// This initializes the first window, a group and a button
	func helloworld() {
		window := myGui.NewWindow("hello world")

		group := window.NewGroup("foo bar")
		group.NewButton("hello", func() {
			log.Println("world")
		})
	}

Hopefully this code example will remain syntactically
consistant.

External Toolkits

	* andlabs - https://github.com/andlabs/ui
	* gocui - https://github.com/awesome-gocui/gocui

The next step is to allow this to work against go-gtk and go-qt.

TODO: Add Fyne, WASM, native macos & windows, android and
hopefully also things like libSDL, faiface/pixel, slint

Bugs

"The author's idea of friendly may differ to that of many other people."

-- quote from the minimalistic window manager 'evilwm'

References

Useful links and other
external things which might be useful

* [Wikipedia Graphical widget](https://en.wikipedia.org/wiki/Graphical_widget)
* [GO Style Guide](https://google.github.io/styleguide/go/index) Code this way
* [MS Windows Application Library Kit](https://github.com/lxn/walk)
* [Federated git pull](https://github.com/forgefed/forgefed) Hopefully this will work for me with gitea
* [Github mirror](https://github.com/wit-go/gui) This repo on mirror. Hopefully I won't have to use this.
* [WIT GO projects](https://go.wit.com/) Attempt to model go.uber.org

*/
package gui
