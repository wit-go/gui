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

## Functions

### func [DebugWidgetWindow](/debugWidget.go#L52)

`func DebugWidgetWindow(w *Node)`

### func [DebugWindow](/debugWindow.go#L21)

`func DebugWindow()`

Creates a window helpful for debugging this package

### func [ExampleCatcher](/chan.go#L37)

`func ExampleCatcher(f func())`

### func [Indent](/debug.go#L120)

`func Indent(b bool, a ...interface{})`

### func [LoadPlugin](/main.go#L114)

`func LoadPlugin(name string) bool`

### func [LoadToolkit](/plugin.go#L68)

`func LoadToolkit(name string) *aplug`

loads and initializes a toolkit (andlabs/ui, gocui, etc)

### func [SetDebug](/debug.go#L28)

`func SetDebug(s bool)`

### func [SetFlag](/debug.go#L44)

`func SetFlag(s string, b bool)`

### func [ShowDebugValues](/debug.go#L79)

`func ShowDebugValues()`

### func [StandardExit](/main.go#L155)

`func StandardExit()`

The window is destroyed and the application exits
TODO: properly exit the plugin since Quit() doesn't do it

### func [Watchdog](/watchdog.go#L15)

`func Watchdog()`

This program sits here.
If you exit here, the whole thing will os.Exit()

This goroutine can be used like a watchdog timer

## Types

### type [GuiArgs](/structs.go#L26)

`type GuiArgs struct { ... }`

This struct can be used with the go-arg package

### type [GuiConfig](/structs.go#L34)

`type GuiConfig struct { ... }`

#### Variables

```golang
var Config GuiConfig
```

### type [Node](/structs.go#L61)

`type Node struct { ... }`

The Node is a binary tree. This is how all GUI elements are stored
simply the name and the size of whatever GUI element exists

#### func [New](/main.go#L132)

`func New() *Node`

There should only be one of these per application
This is due to restrictions by being cross platform
some toolkit's on some operating systems don't support more than one
Keep things simple. Do the default expected thing whenever possible

### type [Symbol](/plugin.go#L16)

`type Symbol any`

## Sub Packages

* [log](./log)

* [toolkit](./toolkit)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
