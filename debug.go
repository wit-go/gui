package gui

// Lots of debugging things:
// A function dump out the binary tree

import (
	"strconv"
	"git.wit.org/wit/gui/toolkit"
)

// various debugging flags
var debugNow bool = true	// useful for active development
var debugGui bool = false
var debugError bool = true
var debugDump bool = false
var debugNode bool = false
var debugTabs bool = false
var debugFlags bool = false
var debugChange bool = false	// shows user events like mouse and keyboard
var debugPlugin	bool = false
var debugAction	bool = false

// for printing out the binary tree
var listChildrenParent *Node
var listChildrenDepth int = 0
var defaultPadding = "  "

func SetDebug (s bool) {
	debugGui     = s
	debugTabs    = s

	logNow = s
	logInfo = s
	logWarn = s
	logError = s
	logVerbose = s

	SetFlag("Node", s)
	SetFlag("Tabs", s)
	SetFlag("Dump", s)
	SetFlag("Flags", s)
	SetFlag("Plugin", s)
	SetFlag("Change", s)
	SetFlag("Error", s)

	// This flag is only for the internal toolkit debugging
	SetFlag("Toolkit", s)
}

func SetFlag (s string, b bool) {
	switch s {
	case "Toolkit":
		// This flag is only for internal toolkit debugging
	case "Tabs":
		debugTabs = b
	case "Node":
		debugNode = b
	case "Dump":
		debugDump = b
	case "Error":
		debugError = b
	case "Change":
		debugChange = b
	case "Flags":
		debugFlags  = b
	case "Plugin":
		debugPlugin  = b
	case "Show":
		// ShowDebugValues() // print them here?
	default:
		log(debugGui, "Can't set unknown flag", s)
	}

	var a toolkit.Action
	a.ActionType = toolkit.Set
	a.WidgetType = toolkit.Flag
	a.S = s
	a.B = b
	// a.Widget = &newNode.widget
	// a.Where = &n.widget
	// action(&a)
	newaction(&a, nil, nil)
}

func ShowDebugValues() {
	// The order here should match the order in the GUI
	// TODO: get the order from the node binary tree
	log(true, "Debug        =", debugGui)
	log(true, "DebugError   =", debugError)
	log(true, "DebugChange  =", debugChange)
	log(true, "DebugDump    =", debugDump)
	log(true, "DebugTabs    =", debugTabs)
	log(true, "DebugPlugin  =", debugPlugin)
	log(true, "DebugNode    =", debugNode)

	SetFlag("Show", true)
}

func (n *Node) Dump() {
	b := true
	// log("Dump() dump =", b)
	Indent(b, "NODE DUMP START")
	Indent(b, "id           = ", n.id)
	Indent(b, "Name         = ", n.Name)
	Indent(b, "(X,Y)        = ", n.X, n.Y)
	Indent(b, "Next (X,Y)   = ", n.NextX, n.NextY)

	if (n.parent == nil) {
		Indent(b, "parent       = nil")
	} else {
		Indent(b, "parent.id    =", n.parent.id)
	}
	if (n.children != nil) {
		Indent(b, "children     = ", n.children)
	}
	if (n.Custom != nil) {
		Indent(b, "Custom       = ", n.Custom)
	}
	Indent(b, "NODE DUMP END")

	var a toolkit.Action
	a.ActionType = toolkit.Dump
	a.WidgetId = n.id
	newaction(&a, activeWidget, nil)
}

func Indent(b bool, a ...interface{}) {
	logindent(b, listChildrenDepth, defaultPadding, a...)
}

func (n *Node) dumpWidget(b bool) string {
	var info, d string

	if (n == nil) {
		log(debugError, "dumpWidget() node == nil")
		return ""
	}
	info = n.WidgetType.String()

	d = strconv.Itoa(n.id) + " " + info

	var tabs string
	for i := 0; i < listChildrenDepth; i++ {
		tabs = tabs + defaultPadding
	}
	d = tabs + d
	logindent(b, listChildrenDepth, defaultPadding, n.id, info)
	return d
}

// func (n *Node) ListChildren(dump bool, dropdown *Node, mapNodes map[string]*Node) {
func (n *Node) ListChildren(dump bool) {
	if (n == nil) {
		return
	}

	n.dumpWidget(dump)
	if len(n.children) == 0 {
		if (n.parent == nil) {
			return
		}
		log(debugNode, "\t\t\tparent =",n.parent.id)
		if (listChildrenParent != nil) {
			log(debugNode, "\t\t\tlistChildrenParent =",listChildrenParent.id)
			if (listChildrenParent.id != n.parent.id) {
				log("parent =",n.parent.id, n.parent.Name)
				log("listChildrenParent =",listChildrenParent.id, listChildrenParent.Name)
				log(listChildrenParent.id, "!=", n.parent.id)
				exit("parent.child does not match child.parent")
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
		if (child.children == nil) {
			log(debugNode, "\t\t", child.id, "has no children")
		} else {
			log(debugNode, "\t\t\tHas children:", child.children)
		}
		listChildrenParent = n
		listChildrenDepth += 1
		// child.ListChildren(dump, dropdown, mapNodes)
		child.ListChildren(dump)
		listChildrenDepth -= 1
	}
	return
}
