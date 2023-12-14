package gui

import (
	"os"
	"git.wit.org/wit/gui/toolkit"
)

// TODO: make a fake 'plugin' channel of communication to andlabs for mswindows
// Windows doesn't support plugins. How can I keep andlabs and only compile it on windows?
// https://forum.heroiclabs.com/t/setting-up-goland-to-compile-plugins-on-windows/594/5
// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

const Xaxis = 0 // stack things horizontally
const Yaxis = 1 // stack things vertically

func init() {
	log("init() has been run")

	me.counter = 0
	me.prefix = "wit"

	// Populates the top of the binary tree
	me.rootNode = addNode("guiBinaryTree")
	me.rootNode.WidgetType = toolkit.Root

	// used to pass debugging flags to the toolkit plugins
	me.flag = me.rootNode.newNode("flag", 0)
	me.flag.WidgetType = toolkit.Flag

	me.flag = me.rootNode.newNode("stdout", 0)
	me.flag.WidgetType = toolkit.Stdout

	me.guiChan = make(chan toolkit.Action, 1)
	go watchCallback()
}

func watchCallback() {
	log(logInfo, "watchCallback() START")
	for {
		log(logNow, "watchCallback() restarted select for toolkit user events")
	    	select {
		case a := <-me.guiChan:
			if (a.ActionType == toolkit.UserQuit) {
				log(logNow, "doUserEvent() User sent Quit()")
				me.rootNode.doCustom()
				exit("wit/gui toolkit.UserQuit")
				break
			}
			if (a.ActionType == toolkit.EnableDebug) {
				log(logNow, "doUserEvent() Enable Debugging Window")
				DebugWindow()
				break
			}

			n := me.rootNode.FindId(a.WidgetId)
			if (n == nil) {
				log(logError, "watchCallback() UNKNOWN widget id =", a.WidgetId, a.Name)
			} else {
				log(logNow, "watchCallback() FOUND widget id =", n.id, n.Name)
				n.doUserEvent(a)
			}
			// this maybe a good idea?
			// TODO: Throttle user events somehow
			// sleep(.01) // hack that throttles user events
		}
	}
}

func (n *Node) doCustom() {
	log(logNow, "doUserEvent() widget =", n.id, n.Name, n.WidgetType, n.B)
	if (n.Custom == nil) {
		log(debugError, "Custom() = nil. SKIPPING")
		return
	}
	go n.Custom()
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

// There should only be one of these per application
// This is due to restrictions by being cross platform
// some toolkit's on some operating systems don't support more than one
// Keep things simple. Do the default expected thing whenever possible
func New() *Node {
	return me.rootNode
}

// try to load andlabs, if that doesn't work, fall back to the console
func (n *Node) Default() *Node {
	if (GuiArg.Gui != "") {
		log(logError, "New.Default() try toolkit =", GuiArg.Gui)
		return n.LoadToolkit(GuiArg.Gui)
	}
	// if DISPLAY isn't set, return since gtk can't load
	// TODO: figure out how to check what to do in macos and mswindows
	if (os.Getenv("DISPLAY") == "") {
		if (n.LoadToolkit("gocui") == nil) {
			log(logError, "New() failed to load gocui")
		}
		return n
	}
	if (n.LoadToolkit("andlabs") != nil) {
		return n
	}
	n.LoadToolkit("gocui")
	return n
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
	for i, plug := range allPlugins {
		log("NewButton()", i, plug)
	}
	exit(0)
}
