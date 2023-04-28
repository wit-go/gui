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
	name string
	filename string
	plug *plugin.Plugin
	// sym *plugin.Symbol
	InitOk bool

	// startup whatever might need to be setup in the plugin
//	Init func()

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
}

var allPlugins []*aplug

// loads and initializes a toolkit (andlabs/ui, gocui, etc)
// attempts to locate the .so file
func initPlugin(name string) *aplug {
	log(logInfo, "initPlugin() START")

	for _, aplug := range allPlugins {
		log(debugGui, "initPlugin() already loaded toolkit plugin =", aplug.name)
		if (aplug.name == name) {
			log(debugError, "initPlugin() SKIPPING", name, "as you can't load it twice")
			return nil
		}
	}

	return searchPaths(name)
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

/*
	TODO: clean this up. use command args?
	TODO: use LD_LIBRARY_PATH ?
	This searches in the following order for the plugin .so files:
		./toolkit/
		~/go/src/go.wit.org/gui/toolkit/
		/usr/lib/go-gui/
*/
func searchPaths(name string) *aplug {
	var filename string
	var pfile []byte
	var err error

	// first try to load the embedded plugin file
	filename = "plugins/" + name + ".so"
	pfile, err = me.resFS.ReadFile(filename)
	if (err == nil) {
		log(logError, "write out file here", name, filename, len(pfile))
		exit()
	} else {
		log(logError, filename, "was not embedded. Error:", err)
	}

	// attempt to write out the file from the internal resource
	filename = "toolkit/" + name + ".so"
	p := initToolkit(name, filename)
	if (p != nil) {
		return p
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log(logError, "searchPaths() error. exiting here?")
	} else {
		filename = homeDir + "/go/src/git.wit.org/wit/gui/toolkit/" + name + ".so"
		p = initToolkit(name, filename)
		if (p != nil) {
			return p
		}
	}

	filename = "/usr/lib/go-gui/" + name + ".so"
	p = initToolkit(name, filename)
	if (p != nil) {
		return p
	}

	filename = "/usr/local/lib/" + name + ".so"
	p = initToolkit(name, filename)
	if (p != nil) {
		return p
	}
	return nil
}

// load module
// 1. open the shared object file to load the symbols
func initToolkit(name string, filename string) *aplug {
	plug, err := plugin.Open(filename)
	if err != nil {
		log(debugGui, "plugin FAILED =", filename, err)
		return nil
	}
	log(debugGui, "initToolkit() loading plugin =", filename)

	var newPlug *aplug
	newPlug = new(aplug)
	newPlug.InitOk = false
	newPlug.name = name
	newPlug.filename = filename
	newPlug.plug = plug

	// this tells the toolkit plugin how to send user events back to us
	// for things like: the user clicked on the 'Check IPv6'
	newPlug.Callback = sendCallback(newPlug, "Callback")

	// this let's us know where to send requests to the toolkit
	// for things like: add a new button called 'Check IPv6'
	newPlug.PluginChannel = getPluginChannel(newPlug, "PluginChannel")

	// add it to the list of plugins
	allPlugins = append(allPlugins, newPlug)


	// set the communication to the plugins
	newPlug.pluginChan = newPlug.PluginChannel()
	if (newPlug.pluginChan == nil) {
		log(debugError, "initToolkit() ERROR PluginChannel() returned nil for plugin:", newPlug.name, filename)
		return nil
	}
	newPlug.Callback(me.guiChan)
	newPlug.InitOk = true

	log(debugPlugin, "initToolkit() END", newPlug.name, filename)
	return newPlug
}

// 2023/04/06 Queue() is also being used and channels are being used. memcopy() only
func newaction(a *toolkit.Action, n *Node, where *Node) {
	if (n != nil) {
		a.WidgetId = n.id
		a.WidgetType = n.WidgetType
		a.ActionType = a.ActionType
	}

	// TODO: redo this grid logic
	if (where != nil) {
		log(logInfo, "Action() START on where X,Y, Next X,Y =", where.Name, where.X, where.Y, where.NextX, where.NextY)
		a.ParentId = where.id
		switch where.WidgetType {
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
		switch where.WidgetType {
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
