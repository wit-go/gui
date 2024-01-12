package gui

import (
	"os"

	"go.wit.com/log"
	"go.wit.com/gui/widget"
)

// TODO: make a fake 'plugin' channel of communication to andlabs for mswindows
// Windows doesn't support plugins. How can I keep andlabs and only compile it on windows?
// https://forum.heroiclabs.com/t/setting-up-goland-to-compile-plugins-on-windows/594/5
// import toolkit "go.wit.com/gui/toolkit/andlabs"

const Xaxis = 0 // stack things horizontally
const Yaxis = 1 // stack things vertically

func init() {
	log.Log(NOW, "init() has been run")

	me.counter = 0
	me.prefix = "wit"

	// Populates the top of the binary tree
	me.rootNode = addNode("guiBinaryTree")
	me.rootNode.WidgetType = widget.Root
	me.rootNode.hidden = false // always send the rootNode to the toolkits

	// used to pass debugging flags to the toolkit plugins
	me.flag = me.rootNode.newNode("flag", 0)
	me.flag.WidgetType = widget.Flag

	me.flag = me.rootNode.newNode("stdout", 0)
	me.flag.WidgetType = widget.Stdout

	me.guiChan = make(chan widget.Action, 1)
	go watchCallback()
}

func watchCallback() {
	log.Info("watchCallback() START")
	for {
		log.Info("watchCallback() restarted select for toolkit user events")
	    	select {
		case a := <-me.guiChan:
			if (a.ActionType == widget.UserQuit) {
				log.Info("doUserEvent() User sent Quit()")
				me.rootNode.doCustom()
				log.Exit("wit/gui toolkit.UserQuit")
				break
			}
			if (a.ActionType == widget.EnableDebug) {
				log.Warn("doUserEvent() Enable Debugging Window")
				log.Warn("doUserEvent() TODO: not implemented")
				// DebugWindow()
				break
			}

			n := me.rootNode.FindId(a.WidgetId)
			if (n == nil) {
				log.Warn("watchCallback() UNKNOWN widget id =", a.WidgetId, a.ProgName)
			} else {
				log.Info("watchCallback() FOUND widget id =", n.id, n.progname)
				n.doUserEvent(a)
			}
			// this maybe a good idea?
			// TODO: Throttle user events somehow
			// sleep(.01) // hack that throttles user events
		}
	}
}

func (n *Node) doCustom() {
	log.Info("doUserEvent() widget =", n.id, n.progname, n.WidgetType)
	if (n.Custom == nil) {
		log.Warn("Custom() = nil. SKIPPING")
		return
	}
	go n.Custom()
}

func (n *Node) doUserEvent(a widget.Action) {
	log.Info("doUserEvent() node =", n.id, n.progname)
	if a.Value == nil {
		log.Warn("doUserEvent() a.A == nil", n.id, n.progname)
		return
	}
	n.value = a.Value
	n.doCustom()
	return
	/*
	switch n.WidgetType {
	case widget.Checkbox:
		log.Info("doUserEvent() node =", n.id, n.progname, "set to:", n.value)
		n.doCustom()
	case widget.Button:
		log.Info("doUserEvent() node =", n.id, n.progname, "button clicked")
		n.doCustom()
	case widget.Combobox:
		n.S = a.S
		log.Info("doUserEvent() node =", n.id, n.progname, "set to:", n.S)
		n.doCustom()
	case widget.Dropdown:
		n.S = a.S
		log.Info("doUserEvent() node =", n.id, n.progname, "set to:", n.S)
		n.doCustom()
	case widget.Textbox:
		n.S = a.S
		log.Info("doUserEvent() node =", n.id, n.progname, "set to:", n.S)
		n.doCustom()
	case widget.Spinner:
		n.I = a.I
		log.Info("doUserEvent() node =", n.id, n.progname, "set to:", n.I)
		n.doCustom()
	case widget.Slider:
		n.I = a.I
		log.Info("doUserEvent() node =", n.id, n.progname, "set to:", n.I)
		n.doCustom()
	case widget.Window:
		log.Info("doUserEvent() node =", n.id, n.progname, "window closed")
		n.doCustom()
	default:
		log.Info("doUserEvent() type =", n.WidgetType)
	}
	*/
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
	if (argGui.GuiPlugin != "") {
		log.Warn("New.Default() try toolkit =", argGui.GuiPlugin)
		return n.LoadToolkit(argGui.GuiPlugin)
	}
	// if DISPLAY isn't set, return since gtk can't load
	// TODO: figure out how to check what to do in macos and mswindows
	if (os.Getenv("DISPLAY") == "") {
		if (n.LoadToolkit("gocui") == nil) {
			log.Warn("New() failed to load gocui")
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
	log.Log(GUI, "wit/gui Standard Window Close. name =", n.progname)
	log.Log(GUI, "wit/gui Standard Window Close. n.Custom exit =", n.Custom)
}

// The window is destroyed and the application exits
// TODO: properly exit the plugin since Quit() doesn't do it
func StandardExit() {
	log.Log(NOW, "wit/gui Standard Window Exit. running os.Exit()")
	log.Log(NOW, "StandardExit() attempt to exit each toolkit plugin")
	for i, plug := range allPlugins {
		log.Log(NOW, "NewButton()", i, plug)
	}
	log.Exit(0)
}
