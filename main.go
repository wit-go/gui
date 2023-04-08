package gui

import (
	"os"
	// "embed"
	"git.wit.org/wit/gui/toolkit"
)

// Windows doesn't support plugins. How can I keep andlabs and only compile it on windows?
// https://forum.heroiclabs.com/t/setting-up-goland-to-compile-plugins-on-windows/594/5
// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

const Xaxis = 0 // stack things horizontally
const Yaxis = 1 // stack things vertically

/*
	// TODO: 2023/03/03 rethink how to get a plugin or figure out how
	// golang packages can include a binary. Pull from /usr/go/go-gui/ ?
	// may this plugin work when all other plugins fail

	// if this is in the plugin, the packages can't work with go.mod builds
	# don't do this in the plugin // go:embed /usr/lib/go-gui/toolkit/gocui.so
	# don't do this in the plugin var res embed.FS
*/

func init() {
	log("init() has been run")

	Config.counter = 0
	Config.prefix = "wit"
	Config.Width = 640
	Config.Height = 480

	// Populates the top of the binary tree
	Config.rootNode = addNode("guiBinaryTree")
	Config.rootNode.WidgetType = toolkit.Root

	// used to pass debugging flags to the toolkit plugins
	Config.flag = Config.rootNode.New("flag", 0, nil)
	Config.flag.WidgetType = toolkit.Flag

	Config.guiChan = make(chan toolkit.Action)
	go watchCallback()
}

func doGuiChan() {
	for {
		select {
		case <-Config.ActionCh1:
			log(true, "CHANNEL ACTION 1  !!!!!")
			return
		case <-Config.ActionCh2:
			log(true, "CHANNEL ACTION 2  !!!!!")
			return
		default:
			log(true, "doGuiChan() nothing")
		}
		log(true, "doGuiChan() for()")
	}
}

// TODO: add logic to just load the 1st 'most common' gui toolkit
// and allow the 'go-arg' command line args to override the defaults
func InitPlugins(names []string) []string {
	log(debugGui, "Starting gui.Init()")

	for _, aplug := range allPlugins {
		log(debugGui, "LoadToolkit() already loaded toolkit plugin =", aplug.name)
		for _, name := range names {
			if (name == aplug.name) {
				return []string{name}
			}
		}
	}

	// try to load each plugin in the order passed to this function
	for _, name := range names {
		aPlug := LoadToolkit(name)
		if (aPlug != nil) {
			// exit because it worked!
			return []string{name}
		}
	}

	// the program didn't specify a plugin. Try to load one
	// TODO: detect the OS & user preferences to load the best one
	// TODO: commented out Init() on 02/26/2023 because I'm not sure how to run it correctly
	andlabsPlug := LoadToolkit("andlabs")
	if (andlabsPlug != nil) {
		return []string{}
	}

	gocuiPlug := LoadToolkit("andlabs")
	if (gocuiPlug != nil) {
		return []string{}
	}
	return []string{}
}

func Start() *Node {
	log(logInfo, "Start() Main(f)")
	f := func() {
	}
	go Main(f)
	sleep(1)
	return Config.rootNode
}

func watchCallback() {
	log(logNow, "makeCallback() START")
	for {
		log(logNow, "makeCallback() for loop")
	    	select {
		case a := <-Config.guiChan:
			log(logNow, "makeCallback() SELECT widget id =", a.WidgetId, a.Name)
			n := Config.rootNode.FindId(a.WidgetId)
			if (n == nil) {
				log(logError, "makeCallback() SELECT widget id =", a.WidgetId, a.Name)
			} else {
				go n.doUserEvent(a)
			}
			// this maybe a good idea?
			// TODO: Throttle user events somehow
			sleep(.1)
		}
	}
}

func (n *Node) doCustom() {
	log(logNow, "doUserEvent() widget =", n.id, n.Name, n.WidgetType, n.B)
	if (n.Custom == nil) {
		log(debugError, "Custom() = nil. SKIPPING")
		return
	}
	n.Custom()
}

func (n *Node) doUserEvent(a toolkit.Action) {
	log(logNow, "doUserEvent() node =", n.id, n.Name)
	switch n.WidgetType {
	case toolkit.Checkbox:
		n.B = a.B
		log(logNow, "doUserEvent() node =", n.id, n.Name, "set to:", n.B)
		n.doCustom()
	case toolkit.Button:
		log(logNow, "doUserEvent() node =", n.id, n.Name, "button clicked")
		n.doCustom()
	case toolkit.Combobox:
		n.S = a.S
		log(logNow, "doUserEvent() node =", n.id, n.Name, "set to:", n.S)
		n.doCustom()
	case toolkit.Dropdown:
		n.S = a.S
		log(logNow, "doUserEvent() node =", n.id, n.Name, "set to:", n.S)
		n.doCustom()
	case toolkit.Textbox:
		n.S = a.S
		log(logNow, "doUserEvent() node =", n.id, n.Name, "set to:", n.S)
		n.doCustom()
	case toolkit.Spinner:
		n.I = a.I
		log(logNow, "doUserEvent() node =", n.id, n.Name, "set to:", n.I)
		n.doCustom()
	case toolkit.Slider:
		n.I = a.I
		log(logNow, "doUserEvent() node =", n.id, n.Name, "set to:", n.I)
		n.doCustom()
	case toolkit.Window:
		log(logNow, "doUserEvent() node =", n.id, n.Name, "window closed")
		n.doCustom()
	default:
		log(logNow, "doUserEvent() type =", n.WidgetType)
	}
}

func (n *Node) LoadPlugin(name string) bool {
	StartS(name)
	return true
}

func StartS(name string) *Node {
	log(logInfo, "Start() Main(f) for name =", name)
	aplug := LoadToolkit(name)
	if (aplug == nil) {
		return Config.rootNode
	}
	// will this really work on mswindows & macos?
	f := func() {
	}
	go Main(f)
	sleep(1) // temp hack until chan communication is setup
	Config.rootNode.Redraw(aplug)
	return Config.rootNode
}

// This should not pass a function
func Main(f func()) {
	log(debugGui, "Starting gui.Main() (using gtk via andlabs/ui)")

	// TODO: this is linux only
	// TODO: detect if this was run from the command line (parent == bash?)
	// if DISPLAY is not set, don't even bother with loading andlabs
	if (os.Getenv("DISPLAY") == "") {
		InitPlugins([]string{"gocui"})
	} else {
		//InitPlugins([]string{"andlabs", "gocui"})
		InitPlugins([]string{"gocui", "andlabs"})
	}

	for _, aplug := range allPlugins {
		log(debugGui, "NewButton() toolkit plugin =", aplug.name)
		if (aplug.MainOk) {
			log(debugGui, "Main() Already Ran Main()", aplug.name)
			continue
		}
		if (aplug.Main == nil) {
			log(debugGui, "Main() Main == nil", aplug.name)
			continue
		}
		aplug.MainOk = true
		if (aplug.Callback == nil) {
			// TODO: don't load the module if this failed ?
			// if Callback() isn't set in the plugin, no information can be sent to it!
			log(debugError, "SERIOUS ERROR: plugin Callback() == nil. nothing will work for toolkit", aplug.name)
		} else {
			aplug.Callback(Config.guiChan)
		}

		if (aplug.PluginChannel == nil) {
			// TODO: don't load the module if this failed ?
			// if Callback() isn't set in the plugin, no information can be sent to it!
			log(debugError, "ERROR: plugin does not implement a send channel. toolkit =", aplug.name)
		} else {
			aplug.pluginChan = aplug.PluginChannel()
		}

		aplug.Main(f)
	}

}

// The window is destroyed but the application does not quit
func (n *Node) StandardClose() {
	log(debugGui, "wit/gui Standard Window Close. name =", n.Name)
	log(debugGui, "wit/gui Standard Window Close. n.Custom exit =", n.Custom)
}

// The window is destroyed and the application exits
// TODO: properly exit the plugin since Quit() doesn't do it
func StandardExit() {
	log("wit/gui Standard Window Exit. running os.Exit()")
	log("StandardExit() attempt to exit each toolkit plugin")
	for i, aplug := range allPlugins {
		log("NewButton()", i, aplug)
		if (aplug.Quit != nil) {
			aplug.Quit()
		}
	}
	exit(0)
}
