# gui

Package gui implements a abstraction layer for Go visual elements.

Definitions:

```go
* Toolkit: the underlying GUI library (MacOS gui, Windows gui, gtk, qt, etc)
* Node: A binary tree of all the underlying widgets
```

Principles:

```go
* Make code using this package simple to use
* Hide complexity internally here
* Isolate the GUI toolkit
* Widget names should try to match [Wikipedia Graphical widget]
* When in doubt, search upward in the binary tree
* It's ok to guess. Try to do something sensible.
```

Quick Start

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

Toolkits

```go
* andlabs - [https://github.com/andlabs/ui](https://github.com/andlabs/ui)
* gocui - [https://github.com/awesome-gocui/gocui](https://github.com/awesome-gocui/gocui)
```

The next step is to allow this to work against go-gtk and go-qt.

TODO: Add Fyne, WASM, native macos & windows, android and
hopefully also things like libSDL, faiface/pixel, slint

## Bugs

"The author's idea of friendly may differ to that of many other people."

-- quote from the minimalistic window manager 'evilwm'

## References

Useful links and other
external things which might be useful

[Wikipedia Graphical widget]: [https://en.wikipedia.org/wiki/Graphical_widget](https://en.wikipedia.org/wiki/Graphical_widget)
[Github mirror]: [https://github.com/witorg/gui](https://github.com/witorg/gui)
[Federated git pull]: [https://github.com/forgefed/forgefed](https://github.com/forgefed/forgefed)
[GO Style Guide]: [https://google.github.io/styleguide/go/index](https://google.github.io/styleguide/go/index)

```go
* [Wikipedia Graphical widget]
* [Github mirror]
* [Federated git pull]
* [GO Style Guide]
```

version v1.3

I like things to be easy.

this means all the log settings are in one place. it should allow
things to be over-ridden externally to the library
but still allow command line --args to pass debugging settings

## I also have a generic sleep() and exit() in here because it's simple

Usage:

log("something", foo, bar)
var DEBUG bool = true
log(DEBUG, "something else", someOtherVariable)  # if DEBUG == false, return doing nothing
log(SPEW, "something else", someOtherVariable)   # this get's sent to spew.Dump(). Very useful for debugging!

## Variables

```golang
var INFO bool
```

```golang
var LOGOFF bool = false // turn this off, all logging stops

```

```golang
var SPEW spewt
```

```golang
var WARN bool
```

## Functions

### func [DebugWindow](/debug_window.go#L9)

`func DebugWindow()`

Creates a window helpful for debugging this package

### func [GetDebug](/debug.go#L20)

`func GetDebug() bool`

### func [GolangDebugWindow](/example_window_golang_debug.go#L10)

`func GolangDebugWindow()`

### func [Indent](/debug.go#L120)

`func Indent(a ...interface{})`

### func [InitPlugins](/main.go#L35)

`func InitPlugins(names []string)`

### func [LoadToolkit](/plugin.go#L53)

`func LoadToolkit(name string) bool`

loads and initializes a toolkit (andlabs/ui, gocui, etc)

### func [Main](/main.go#L76)

`func Main(f func())`

This should not pass a function

### func [Queue](/main.go#L106)

`func Queue(f func())`

Other goroutines must use this to access the GUI

You can not acess / process the GUI thread directly from
other goroutines. This is due to the nature of how
Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
For example: gui.Queue(NewWindow())

### func [SetDebug](/debug.go#L24)

`func SetDebug(s bool)`

### func [SetDebugChange](/debug.go#L52)

`func SetDebugChange(s bool)`

This passes the debugChange flag to the toolkit plugin

### func [SetDebugToolkit](/debug.go#L37)

`func SetDebugToolkit(s bool)`

This passes the debugToolkit flag to the toolkit plugin

### func [ShowDebugValues](/debug.go#L66)

`func ShowDebugValues()`

### func [StandardClose](/main.go#L119)

`func StandardClose(n *Node)`

The window is destroyed but the application does not quit

### func [StandardExit](/main.go#L125)

`func StandardExit(n *Node)`

The window is destroyed but the application does not quit

### func [Watchdog](/watchdog.go#L15)

`func Watchdog()`

This program sits here.
If you exit here, the whole thing will os.Exit()

This goroutine can be used like a watchdog timer

## Types

### type [GuiArgs](/structs.go#L25)

`type GuiArgs struct { ... }`

This struct can be used with the go-arg package

### type [GuiConfig](/structs.go#L30)

`type GuiConfig struct { ... }`

#### Variables

```golang
var Config GuiConfig
```

### type [Node](/structs.go#L48)

`type Node struct { ... }`

The Node is a binary tree. This is how all GUI elements are stored
simply the name and the size of whatever GUI element exists

#### func [NewWindow](/window.go#L14)

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

### type [Symbol](/plugin.go#L16)

`type Symbol any`

## Sub Packages

* [toolkit](./toolkit)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
