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

Quick Start

	// This creates a simple hello world window
	package main

	import 	(
		"log"
		"git.wit.org/wit/gui"
	)

	var window *gui.Node // This is the beginning of the binary tree of widgets

	// go will sit here until the window exits
	func main() {
		gui.Init()
		gui.Main(helloworld)
	}

	// This initializes the first window and 2 tabs
	func helloworld() {
		gui.Config.Title = "Hello World golang wit/gui Window"
		gui.Config.Width = 640
		gui.Config.Height = 480

		window := gui.NewWindow()
		addTab(window, "A Simple Tab Demo")
		addTab(window, "A Second Tab")
	}

	func addTab(w *gui.Node, title string) {
		tab := w.NewTab(title)

		group := tab.NewGroup("foo bar")
		group.NewButton("hello", func() {
			log.Println("world")
		})
	}


Debian Build

This worked on debian sid on 2022/10/20
I didn't record the dependances needed

	GO111MODULE="off" go get -v -t -u git.wit.org/wit/gui
	cd ~/go/src/git.wit.org/wit/gui/cmds/helloworld/
	GO111MODULE="off" go build -v -x
	./helloworld

Toolkits

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
external things
which might be useful

[Wikipedia Graphical widget]: https://en.wikipedia.org/wiki/Graphical_widget
[Github mirror]: https://github.com/witorg/gui
[Federated git pull]: https://github.com/forgefed/forgefed

	* [Wikipedia Graphical widget]
	* [Github mirror]
	* [Federated git pull]


*/
package gui
