package gui

// This is based off of the excellent example and documentation here:
// https://github.com/vladimirvivien/go-plugin-example
// There truly are great people in this world.
// It's a pleasure to be here with all of you

import (
	"os"
	"embed"
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

func newAction(n *Node, atype toolkit.ActionType) *toolkit.Action {
	var a toolkit.Action
	a.ActionType = atype
	if (n == nil) {
		return &a
	}
	a.Name = n.Name
	a.Text = n.Text
	a.WidgetId = n.id
	a.WidgetType = n.WidgetType

	return &a
}

func sendAction(a *toolkit.Action, n *Node, where *Node) {
	newaction(a, n, where)
}

// 2023/04/06 Queue() is also being used and channels are being used. memcopy() only
func newaction(a *toolkit.Action, n *Node, where *Node) {
	// remove this
	if (n != nil) {
		a.WidgetId = n.id
		a.WidgetType = n.WidgetType
		a.ActionType = a.ActionType
	}
	// end remove
	if (where != nil) {
		a.ParentId = where.id
		if (where.WidgetType == toolkit.Grid) {
			placeGrid(a, n, where)
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
		// added during debugging. might be a good idea in general for a tactile experience
		sleep(.02)
	}
}

func (n *Node) InitEmbed(resFS embed.FS) *Node {
	me.resFS = resFS
	return n
}

func (n *Node) LoadToolkitEmbed(name string, b []byte) *Node {
	for _, aplug := range allPlugins {
		log(logInfo, "LoadToolkitEmbed() already loaded toolkit plugin =", aplug.name)
		if (aplug.name == name) {
			log(logError, "LoadToolkitEmbed() SKIPPING", name, "as you can't load it twice")
			return n
		}
	}

	f, err := os.CreateTemp("", "sample." + name + ".so")
	if (err != nil) {
		return n
	}
	defer os.Remove(f.Name())
	f.Write(b)

	p := initToolkit(name, f.Name())
	if (p == nil) {
		log(logError, "LoadToolkitEmbed() embedded go file failed", name)
	}
	return n
}

func (n *Node) ListToolkits() {
	for _, aplug := range allPlugins {
		log(logNow, "ListToolkits() already loaded toolkit plugin =", aplug.name)
	}
}

func (n *Node) LoadToolkit(name string) *Node {
	log(logInfo, "LoadToolkit() START for name =", name)
	plug := initPlugin(name)
	if (plug == nil) {
		return n
	}

	log(logInfo, "LoadToolkit() sending InitToolkit action to the plugin channel")
	var a toolkit.Action
	a.ActionType = toolkit.InitToolkit
	plug.pluginChan <- a
	sleep(.5) // temp hack until chan communication is setup

	// TODO: find a new way to do this that is locking, safe and accurate
	me.rootNode.redraw(plug)
	log(logInfo, "LoadToolkit() END for name =", name)
	return n
}

func (n *Node) CloseToolkit(name string) bool {
	log(logInfo, "CloseToolkit() for name =", name)
	for _, plug := range allPlugins {
		log(debugGui, "CloseToolkit() found", plug.name)
		if (plug.name == name) {
			log(debugNow, "CloseToolkit() sending close", name)
			var a toolkit.Action
			a.ActionType = toolkit.CloseToolkit
			plug.pluginChan <- a
			sleep(.5)
			return true
		}
	}
	return false
}
