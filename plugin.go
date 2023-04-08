package gui

// This is based off of the excellent example and documentation here:
// https://github.com/vladimirvivien/go-plugin-example
// There truly are great people in this world.
// It's a pleasure to be here with all of you

import (
	"os"
	"plugin"

	"git.wit.org/wit/gui/toolkit"
)

var err error
type Symbol any

type aplug struct {
	// Ok bool
	name string
	filename string
	plug *plugin.Plugin
	sym *plugin.Symbol
	LoadOk bool
	InitOk bool
	MainOk bool

	Init func()

	// This passes the go channel to the plugin
	// the plugin then passes actions back to
	// here where they are processed. If you wit/gui this is how
	// you are passed information like a user clicking a button
	// or a user changing a dropdown menu or a checkbox
	//
	// from this channel, the information is then passed into your
	// Custom() function
	//
	// the custom functions are run from inside of your own goroutine
	// where you initialize the gui
	Callback func(chan toolkit.Action)

	// This is how actions are sent to the plugin
	//
	// for example, when you you create a new button, it sends
	// a structure to the goroutine that is handling the gui
	// each toolkit has it's own goroutine and each one is sent this
	// add button request
	pluginChan chan toolkit.Action

	PluginChannel func() chan toolkit.Action

	// deprecate all this
	// TODO: make Main() main() and never allow the user to call it
	// run plugin.Main() when the plugin is loaded
	Main func(func ())	// this never returns. Each plugin must have it's own goroutine
	Quit func()

	// simplifies passing to the plugin
	// Send func(*toolkit.Widget, *toolkit.Widget)
	// should replace Send()
	// Action func(*toolkit.Action)
}

var allPlugins []*aplug

// loads and initializes a toolkit (andlabs/ui, gocui, etc)
func LoadToolkit(name string) *aplug {
	var newPlug *aplug
	newPlug = new(aplug)

	log(logInfo, "LoadToolkit() START")
	newPlug.LoadOk = false

	for _, aplug := range allPlugins {
		log(debugGui, "LoadToolkit() already loaded toolkit plugin =", aplug.name)
		if (aplug.name == name) {
			log(debugError, "LoadToolkit() SKIPPING", name, "as you can't load it twice")
			return aplug
		}
	}

	// locate the shared library file
	filename := name + ".so"
	loadPlugin(newPlug, filename)
	if (newPlug.plug == nil) {
		log(true, "attempt to find plugin", filename, "failed")
		return nil
	}
	// newPlug.Ok = true
	newPlug.name = name

	// deprecate Init(?)
	newPlug.Init = loadFuncE(newPlug, "Init")

	// should make a goroutine that never exits
	newPlug.Main  = loadFuncF(newPlug, "Main")

	// should send things to the goroutine above
	// newPlug.Queue = loadFuncF(&newPlug, "Queue")

	// unload the plugin and restore state
	newPlug.Quit = loadFuncE(newPlug, "Quit")

	// Sends instructions like "Add", "Delete", "Disable", etc
	// Sends a widget (button, checkbox, etc) and it's parent widget
	// newPlug.Action = loadFuncA(newPlug, "Action")

	// this tells the toolkit plugin how to send user events back to us
	// for things like: the user clicked on the 'Check IPv6'
	newPlug.Callback = sendCallback(newPlug, "Callback")

	// this let's us know where to send requests to the toolkit
	// for things like: add a new button called 'Check IPv6'
	newPlug.PluginChannel = getPluginChannel(newPlug, "PluginChannel")

	allPlugins = append(allPlugins, newPlug)

	log(debugPlugin, "LoadToolkit() END", newPlug.name, filename)
	newPlug.Init()

	// set the communication to the plugins
	newPlug.pluginChan = newPlug.PluginChannel()
	newPlug.Callback(Config.guiChan)

	newPlug.LoadOk = true
	return newPlug
}

// TODO: All these functions need to be done a smarter way
// but I haven't worked out the golang syntax to make it smarter
func loadFuncE(p *aplug, funcName string) func() {
	var newfunc func()
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log(debugGui, "DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func())
	if !ok {
		log(debugGui, "function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

//	newPlug.PluginChannel = getPluginChannel(newPlug, "PluginChannel")
func getPluginChannel(p *aplug, funcName string) func() chan toolkit.Action {
	var newfunc func() chan toolkit.Action
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log(debugGui, "DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func() chan toolkit.Action)
	if !ok {
		log(debugGui, "function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

func sendCallback(p *aplug, funcName string) func(chan toolkit.Action) {
	var newfunc func(chan toolkit.Action)
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log(debugGui, "DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(chan toolkit.Action))
	if !ok {
		log(debugGui, "function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

func loadFunc2(p *aplug, funcName string) func(*toolkit.Widget, *toolkit.Widget) {
	var newfunc func(*toolkit.Widget, *toolkit.Widget)
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log(debugGui, "DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(*toolkit.Widget, *toolkit.Widget))
	if !ok {
		log(debugGui, "function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

// does this fix loadFuncE problems?
// TODO: still need to move to channels here
func loadFuncA(p *aplug, funcName string) func(*toolkit.Action) {
	var newfunc func(*toolkit.Action)
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log(debugGui, "DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(*toolkit.Action))
	if !ok {
		log(debugGui, "function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

// This is probably dangerous and should never be done
// executing arbitrary functions will cause them to run inside the goroutine that
// the GUI toolkit itself is running in. TODO: move to channels here
func loadFuncF(p *aplug, funcName string) func(func ()) {
	var newfunc func(func ())
	var ok bool
	var test plugin.Symbol

	test, err = p.plug.Lookup(funcName)
	if err != nil {
		log(debugGui, "DID NOT FIND: name =", test, "err =", err)
		return nil
	}

	newfunc, ok = test.(func(func ()))
	if !ok {
		log(debugGui, "function name =", funcName, "names didn't map correctly. Fix the plugin name =", p.name)
		return nil
	}
	return newfunc
}

/*
	This searches in the following order for the plugin .so files:
		./toolkit/
		~/go/src/go.wit.org/gui/toolkit/
		/usr/lib/go-gui/
*/
func loadPlugin(p *aplug, name string) {
	var filename string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log(logError, "loadPlugin() error. exiting here?")
		return
	}

	// attempt to write out the file from the internal resource
	filename = "toolkit/" + name
	p.plug = loadfile(filename)
	if (p.plug != nil) {
		p.filename = filename
		return
	}

	filename = homeDir + "/go/src/git.wit.org/wit/gui/toolkit/" + name
	p.plug = loadfile(filename)
	if (p.plug != nil) {
		p.filename = filename
		return
	}

	filename = "/usr/lib/go-gui/" + name
	p.plug = loadfile(filename)
	if (p.plug != nil) {
		p.filename = filename
		return
	}
	return
}

// load module
// 1. open the shared object file to load the symbols
func loadfile(filename string) *plugin.Plugin {
	plug, err := plugin.Open(filename)
	if err != nil {
		log(debugGui, "plugin FAILED =", filename, err)
		return nil
	}
	log(debugGui, "plugin WORKED =", filename)
	log(true, "loading plugin", filename, "worked")
	return plug
}

// 2023/04/06 Queue() is also being used and channels are being used. memcopy() only
func newaction(a *toolkit.Action, n *Node, where *Node) {
	if (n != nil) {
		a.WidgetId = n.id
		a.WidgetType = n.widget.Type
		a.ActionType = a.ActionType
	}

	// TODO: redo this grid logic
	if (where != nil) {
		log(logInfo, "Action() START on where X,Y, Next X,Y =", where.Name, where.X, where.Y, where.NextX, where.NextY)
		a.ParentId = where.id
		switch where.widget.Type {
		case toolkit.Grid:
			// where.Dump(true)
			log(logInfo, "Action() START on Grid (X,Y)", where.X, where.Y, "put next thing at (X,Y) =", where.NextX, where.NextY)
			//
			// fix values here if they are invalid. Index starts at 1
			if (where.NextX < 1) {
				where.NextX = 1
			}
			if (where.NextY < 1) {
				where.NextY = 1
			}
			//
			a.X = where.NextX
			a.Y = where.NextY
			log(logInfo, "Action() END   on Grid (X,Y)", where.X, where.Y, "put next thing at (X,Y) =", where.NextX, where.NextY)
		default:
		}
	}

	for _, aplug := range allPlugins {
		log(debugPlugin, "Action() aplug =", aplug.name, "Action type=", a.ActionType)
		if (aplug.pluginChan == nil) {
			log(logInfo, "Action() retrieving the aplug.PluginChannel()", aplug.name)
			aplug.pluginChan = aplug.PluginChannel()
			log(logInfo, "Action() retrieved", aplug.pluginChan)
		}
		log(logInfo, "Action() SEND to pluginChan", aplug.name)
		aplug.pluginChan <- *a
		sleep(.02)
	}
	// increment where to put the next widget in a grid or table
	if (where != nil) {
		switch where.widget.Type {
		case toolkit.Grid:
			log(logInfo, "Action() START size (X,Y)", where.X, where.Y, "put next thing at (X,Y) =", where.NextX, where.NextY)
			where.NextY += 1
			if (where.NextY > where.Y) {
				where.NextX += 1
				where.NextY = 1
			}
			log(logInfo, "Action() END size (X,Y)", where.X, where.Y, "put next thing at (X,Y) =", where.NextX, where.NextY)
		default:
		}
	}
}
