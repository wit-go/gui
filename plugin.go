package gui

// This is based off of the excellent example and documentation here:
// https://github.com/vladimirvivien/go-plugin-example
// There truly are great people in this world.
// It's a pleasure to be here with all of you

import (
	"os"
	"embed"
	"plugin"

	"go.wit.com/gui/gui/toolkit"
)

var err error
type Symbol any

type aplug struct {
	name string
	filename string
	plug *plugin.Plugin

	// this tells the toolkit plugin how to send events
	// back here
	//
	// This is how we are passed information like a user clicking a button
	// or a user changing a dropdown menu or a checkbox
	//
	// From this channel, the information is then passed into the main program
	// Custom() function
	//
	Callback func(chan toolkit.Action)

	// This is how actions are sent to the toolkit. 
	// For example:
	// If a program is using GTK, when a program tries to make a new button
	// "Open GIMP", then it would pass an action via this channel into the toolkit
	// plugin and the toolkit plugin would add a button to the parent widget
	//
	// each toolkit has it's own goroutine and each one is sent this
	// add button request
	//
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
		~/go/src/go.wit.com/gui/toolkit/
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
		filename = "/tmp/" + name + ".so"
		log(logError, "write out file here", name, filename, len(pfile))
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
		f.Write(pfile)
		f.Close()
		p := initToolkit(name, filename)
		if (p != nil) {
			return p
		}
	} else {
		log(logError, filename, "was not embedded in the binary. Error:", err)
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
		filename = homeDir + "/go/src/go.wit.com/gui/toolkits/" + name + ".so"
		p = initToolkit(name, filename)
		if (p != nil) {
			return p
		}
	}

	homeDir, err = os.UserHomeDir()
	if err != nil {
		log(logError, "searchPaths() error. exiting here?")
	} else {
		filename = homeDir + "/go/src/go.wit.com/toolkits/" + name + ".so"
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
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			log(true, "missing plugin", name, "as filename", filename)
			return nil
		}
	}
	log(true, "Found plugin", name, "as filename", filename)

	plug, err := plugin.Open(filename)
	if err != nil {
		log(debugError, "plugin FAILED =", filename, err)
		return nil
	}
	log(debugPlugin, "initToolkit() loading plugin =", filename)

	var newPlug *aplug
	newPlug = new(aplug)
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

	log(debugPlugin, "initToolkit() END", newPlug.name, filename)
	return newPlug
}

// 2023/05/09 pretty clean
// 2023/04/06 Queue() is also being used and channels are being used. memcopy() only
func newAction(n *Node, atype toolkit.ActionType) *toolkit.Action {
	var a toolkit.Action
	a.ActionType = atype
	if (n == nil) {
		return &a
	}
	a.Name = n.Name
	a.Text = n.Text
	a.WidgetId = n.id

	a.B = n.B
	a.I = n.I
	a.S = n.S

	a.X = n.X
	a.Y = n.Y

	a.AtW = n.AtW
	a.AtH = n.AtH

	if (n.parent != nil) {
		a.ParentId = n.parent.id
	}
	a.WidgetType = n.WidgetType
	return &a
}

// sends the action/event to each toolkit via a golang plugin channel
func sendAction(a *toolkit.Action) {
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
		sleep(.02) // this delay makes it so SetText() works on initial widget creation
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
	// sleep(.5) // temp hack until chan communication is setup

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
			// sleep(.5) // is this needed? TODO: properly close channel
			return true
		}
	}
	return false
}
