# gui

Package gui implements a abstraction layer for Go visual elements in
a cross platform and library independent way. (hopefully this is will work)

A quick overview of the features, some general design guidelines
and principles for how this package should generally work:

Definitions:

```go
* Toolkit: the underlying library (MacOS gui, Windows gui, gtk, qt, etc)
* Node: A binary tree of all the underlying GUI toolkit elements
```

Principles:

```go
* Make code using this package simple to use
* When in doubt, search upward in the binary tree
* It's ok to guess. We will return something close.
* Hide complexity internally here
* Isolate the GUI toolkit
* Try to use [Wikipedia Graphical widget] names
```

## Quick Start

This section demonstrates how to quickly get started with spew.  See the
sections below for further details on formatting and configuration options.

```go
// This creates a simple hello world window
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
)

var window *gui.Node // This is the beginning of the binary tree of widgets

// go will sit here until the window exits
func main() {
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
```

## Debian Build

This worked on debian sid on 2022/10/20
I didn't record the dependances needed

```go
GO111MODULE="off" go get -v -t -u git.wit.org/wit/gui
cd ~/go/src/git.wit.org/wit/gui/cmds/helloworld/
GO111MODULE="off" go build -v -x
[./helloworld](./helloworld)
```

## Toolkits

* Andlabs - [https://github.com/andlabs/ui](https://github.com/andlabs/ui)
* gocui - [https://github.com/awesome-gocui/gocui](https://github.com/awesome-gocui/gocui)

The goal is to design something that will work with more than one.

Right now, this abstraction is built on top of the go package 'andlabs/ui'
which does the cross platform support.
The next step is to intent is to allow this to work directly against GTK and QT.

It should be able to add Fyne, WASM, native macos & windows, android and
hopefully also things like libSDL, faiface/pixel, slint

## Errors

Since it is possible for custom Stringer/error interfaces to panic, spew
detects them and handles them internally by printing the panic information
inline with the output.  Since spew is intended to provide deep pretty printing
capabilities on structures, it intentionally does not return any errors.

## Debugging

To dump variables with full newlines, indentation, type, and pointer
information this uses spew.Dump()

## Bugs

"The author's idea of friendly may differ to that of many other people."

-- manpage quote from the excellent minimalistic window manager 'evilwm'

## References

Useful links and other
external things
which might be useful

[Wikipedia Graphical widget]: [https://en.wikipedia.org/wiki/Graphical_widget](https://en.wikipedia.org/wiki/Graphical_widget)
[Github mirror]: [https://github.com/witorg/gui](https://github.com/witorg/gui)

```go
* [Wikipedia Graphical widget]
* [Github mirror]
```

## Variables

```golang
var PlugGocli *plugin.Plugin
```

```golang
var PlugGocliOk bool
```

```golang
var PlugHello *plugin.Plugin
```

## Functions

### func [DebugTab](/example_window_debug.go#L26)

`func DebugTab()`

this function is used by the examples to add a tab
dynamically to the bugWin node
TODO: make this smarter once this uses toolkit/

### func [DebugWindow](/example_window_debug.go#L14)

`func DebugWindow()`

Creates a window helpful for debugging this package

### func [DemoToolkitWindow](/example_window_demo_toolkit.go#L24)

`func DemoToolkitWindow()`

This creates a window that shows how the toolkit works
internally using it's raw unchanged code for the toolkit itself

This is a way to test and see if the toolkit is working at all
right now it shows the andlabs/ui/DemoNumbersPage()

### func [DemoWindow](/example_window_demo.go#L10)

`func DemoWindow()`

This creates a window that shows how this package works

### func [GetDebug](/structs.go#L24)

`func GetDebug() bool`

### func [GetDebugToolkit](/structs.go#L36)

`func GetDebugToolkit() bool`

### func [GocuiAddButton](/plugin.go#L108)

`func GocuiAddButton(name string)`

### func [GolangDebugWindow](/example_window_golang_debug.go#L12)

`func GolangDebugWindow()`

### func [IndentPrintln](/structs.go#L216)

`func IndentPrintln(a ...interface{})`

### func [LoadPlugin](/plugin.go#L36)

`func LoadPlugin(name string) *plugin.Plugin`

### func [LookupJcarrButton](/plugin.go#L98)

`func LookupJcarrButton()`

### func [Main](/main.go#L38)

`func Main(f func())`

### func [Queue](/main.go#L51)

`func Queue(f func())`

Other goroutines must use this to access the GUI

You can not acess / process the GUI thread directly from
other goroutines. This is due to the nature of how
Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
For example: gui.Queue(NewWindow())

### func [RunGreet](/plugin.go#L88)

`func RunGreet()`

### func [SetDebug](/structs.go#L28)

`func SetDebug(s bool)`

### func [SetDebugToolkit](/structs.go#L40)

`func SetDebugToolkit(s bool)`

### func [ShowDebugValues](/structs.go#L44)

`func ShowDebugValues()`

### func [StandardClose](/main.go#L57)

`func StandardClose(n *Node)`

The window is destroyed but the application does not quit

### func [StandardExit](/main.go#L65)

`func StandardExit(n *Node)`

The window is destroyed but the application does not quit

## Types

### type [Greeter](/plugin.go#L17)

`type Greeter interface { ... }`

TODO: could a protobuf work here?

### type [GuiConfig](/structs.go#L67)

`type GuiConfig struct { ... }`

#### Variables

```golang
var Config GuiConfig
```

### type [GuiOptions](/structs.go#L56)

`type GuiOptions struct { ... }`

This struct can be used with go-arg

### type [Node](/structs.go#L117)

`type Node struct { ... }`

The Node is simply the name and the size of whatever GUI element exists

#### func [NewStandardWindow](/example_window_demo_toolkit.go#L7)

`func NewStandardWindow(title string) *Node`

#### func [NewWindow](/window.go#L15)

`func NewWindow() *Node`

This routine creates a blank window with a Title and size (W x H)

This routine can not have any arguements due to the nature of how
it can be passed via the 'andlabs/ui' queue which, because it is
cross platform, must pass UI changes into the OS threads (that is
my guess).

This example demonstrates how to create a NewWindow()

Interacting with a GUI in a cross platform fashion adds some
unusual problems. To obvuscate those, andlabs/ui starts a
goroutine that interacts with the native gui toolkits
on the Linux, MacOS, Windows, etc.

Because of this oddity, to initialize a new window, the
function is not passed any arguements and instead passes
the information via the Config type.

```golang
package main

import (
	"git.wit.org/wit/gui"
)

func main() {
	// Define the name and size
	gui.Config.Title = "WIT GUI Window 1"
	gui.Config.Width = 640
	gui.Config.Height = 480

	// Create the Window
	gui.NewWindow()

}

```

 Output:

```
You get a window
```

### type [Widget](/widget.go#L12)

`type Widget struct { ... }`

what names should be used? This is not part of [[Graphical Widget]]
Event() seems like a good name.
Could a protobuf be used here? (Can functions be passed?)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
