# gui

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

* andlabs - [https://github.com/andlabs/ui](https://github.com/andlabs/ui)
* gocui - [https://github.com/awesome-gocui/gocui](https://github.com/awesome-gocui/gocui)

The next step is to allow this to work against go-gtk and go-qt.

TODO: Add Fyne, WASM, native macos & windows, android and
hopefully also things like libSDL, faiface/pixel, slint

## Bugs

"The author's idea of friendly may differ to that of many other people."

-- quote from the minimalistic window manager 'evilwm'

## References

Useful links and other
external things
which might be useful

* [Wikipedia Graphical widget](https://en.wikipedia.org/wiki/Graphical_widget)
* [Github mirror](https://github.com/witorg/gui)
* [Federated git pull](https://github.com/forgefed/forgefed)

## Functions

### func [GetDebug](/structs.go#L25)

`func GetDebug() bool`

### func [GetDebugToolkit](/structs.go#L37)

`func GetDebugToolkit() bool`

### func [IndentPrintln](/structs.go#L188)

`func IndentPrintln(a ...interface{})`

### func [Init](/main.go#L41)

`func Init()`

### func [LoadToolkit](/plugin.go#L37)

`func LoadToolkit(name string)`

loads and initializes a toolkit (andlabs/ui, gocui, etc)

### func [Main](/main.go#L56)

`func Main(f func())`

### func [Queue](/main.go#L77)

`func Queue(f func())`

Other goroutines must use this to access the GUI

You can not acess / process the GUI thread directly from
other goroutines. This is due to the nature of how
Linux, MacOS and Windows work (they all work differently. suprise. surprise.)
For example: gui.Queue(NewWindow())

### func [SetDebug](/structs.go#L29)

`func SetDebug(s bool)`

### func [SetDebugToolkit](/structs.go#L41)

`func SetDebugToolkit(s bool)`

### func [ShowDebugValues](/structs.go#L45)

`func ShowDebugValues()`

### func [StandardClose](/main.go#L83)

`func StandardClose(n *Node)`

The window is destroyed but the application does not quit

### func [StandardExit](/main.go#L90)

`func StandardExit(n *Node)`

The window is destroyed but the application does not quit

### func [Watchdog](/watchdog.go#L16)

`func Watchdog()`

This program sits here.
If you exit here, the whole thing will os.Exit()

This goroutine can be used like a watchdog timer

## Types

### type [GuiConfig](/structs.go#L68)

`type GuiConfig struct { ... }`

#### Variables

```golang
var Config GuiConfig
```

### type [GuiOptions](/structs.go#L56)

`type GuiOptions struct { ... }`

This struct can be used with go-arg

### type [Node](/structs.go#L87)

`type Node struct { ... }`

The Node is simply the name and the size of whatever GUI element exists

#### func [NewWindow](/window.go#L15)

`func NewWindow() *Node`

This routine creates a blank window with a Title and size (W x H)

This routine can not have any arguements due to the nature of how
it can be passed via the 'andlabs/ui' queue which, because it is
cross platform, must pass UI changes into the OS threads (that is
my guess).

### type [Symbol](/plugin.go#L17)

`type Symbol any`

## Sub Packages

* [need-to-redo](./need-to-redo)

* [toolkit](./toolkit)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
