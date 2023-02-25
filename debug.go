package gui

// Lots of debugging things:
// A function dump out the binary tree

import (
	// "fmt"
	"reflect"
)

// various debugging flags
var debugGui bool = false
var debugDump bool = false
var debugNode bool = false
var debugTabs bool = false
var debugChange bool = false	// shows user events like mouse and keyboard
var debugPlugin	bool = false
var debugToolkit bool = false

func GetDebug () bool {
	return debugGui
}

func SetDebug (s bool) {
	debugGui = s
	// debugDump = s
	// debugNode = s
}

/*
func GetDebugToolkit () bool {
	return debugToolkit
}
*/

// This passes the debugToolkit flag to the toolkit plugin
func SetDebugToolkit (s bool) {
	debugToolkit = s
	for _, aplug := range allPlugins {
		log(debugPlugin, "gui.SetDebugToolkit() aplug =", aplug.name)
		if (aplug.SetDebugToolkit == nil) {
			log(debugPlugin, "\tgui.SetDebugToolkit() = nil", aplug.name)
			continue
		}
		aplug.SetDebugToolkit(s)
		return
	}
	log(debugPlugin, "\tgui.SetDebugToolkit() = nil in all plugins")
}

// This passes the debugChange flag to the toolkit plugin
func SetDebugChange (s bool) {
	// debugToolkit = s
	for _, aplug := range allPlugins {
		log(debugPlugin, "gui.SetDebugChange() aplug =", aplug.name)
		if (aplug.SetDebugChange == nil) {
			log(debugPlugin, "\tgui.SetDebugChange() = nil", aplug.name)
			continue
		}
		aplug.SetDebugChange(s)
		return
	}
	log(debugPlugin, "\tgui.SetDebugChange() = nil in all plugins")
}

func ShowDebugValues() {
	log(true, "Debug =", debugGui)
	log(true, "DebugDump =", debugDump)
	log(true, "DebugNode =", debugNode)
	log(true, "DebugTabs =", debugTabs)
	log(true, "DebugPlugin =", debugPlugin)
	log(true, "DebugChange =", debugChange)
	log(true, "DebugToolkit =", debugToolkit)

	// dump out the debugging flags for the plugins
	for _, aplug := range allPlugins {
		log(debugPlugin, "gui.ShowDebug() aplug =", aplug.name)
		if (aplug.ShowDebug == nil) {
			log(debugPlugin, "\tgui.ShowDebug() = nil", aplug.name)
			continue
		}
		aplug.ShowDebug()
		return
	}
}

func (n *Node) Dump() {
	if ! debugDump {
		return
	}
	Indent("NODE DUMP START")
	Indent("id         = ", n.id)
	Indent("Name       = ", n.Name)
	Indent("Width      = ", n.Width)
	Indent("Height     = ", n.Height)

	if (n.parent == nil) {
		Indent("parent     = nil")
	} else {
		Indent("parent.id  =", n.parent.id)
	}
	if (n.children != nil) {
		Indent("children   = ", n.children)
	}
	if (n.custom != nil) {
		Indent("custom     = ", n.custom)
	}
	Indent("checked    = ", n.checked)
	if (n.OnChanged != nil) {
		Indent("OnChanged  = ", n.OnChanged)
	}
	Indent("text       = ", reflect.ValueOf(n.text).Kind(), n.text)
	Indent("NODE DUMP END")
}

var listChildrenParent *Node
var listChildrenDepth int = 0
var defaultPadding = "  "

func Indent(a ...interface{}) {
	logindent(listChildrenDepth, defaultPadding, a...)
}

func (n *Node) dumpWidget() {
	var info string

	if (n.Widget.Type == "") {
		n.Widget.Type = "UNDEF"
	}
	info = n.Widget.Type

	info += ", " + n.Widget.Name
	if (n.Name != n.Widget.Name) {
		info += " NAME MISMATCH"
	}

	logindent(listChildrenDepth, defaultPadding, n.id, info)
}

func (n *Node) ListChildren(dump bool) {
	n.dumpWidget()

	if (dump == true) {
		n.Dump()
	}
	if len(n.children) == 0 {
		if (n.parent == nil) {
		} else {
			log(debugNode, "\t\t\tparent =",n.parent.id)
			if (listChildrenParent != nil) {
				log(debugNode, "\t\t\tlistChildrenParent =",listChildrenParent.id)
				if (listChildrenParent.id != n.parent.id) {
					// log("parent.child does not match child.parent")
					exit("parent.child does not match child.parent")
				}
			}
		}
		log(debugNode, "\t\t", n.id, "has no children")
		return
	}
	for _, child := range n.children {
		if (child.parent != nil) {
			log(debugNode, "\t\t\tparent =",child.parent.id)
		} else {
			log(debugGui, "\t\t\tno parent")
			// memory corruption? non-threadsafe access?
			// can all binary tree changes to Node.parent & Node.child be forced into a singular goroutine?
			panic("something is wrong with the wit golang gui logic and the binary tree is broken. child has no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		if (child.children == nil) {
			log(debugNode, "\t\t", child.id, "has no children")
		} else {
			log(debugNode, "\t\t\tHas children:", child.children)
		}
		listChildrenParent = n
		listChildrenDepth += 1
		child.ListChildren(dump)
		listChildrenDepth -= 1
	}
	return
}
