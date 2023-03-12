package gui

// Lots of debugging things:
// A function dump out the binary tree

import (
	"strconv"
	"git.wit.org/wit/gui/toolkit"
)

// various debugging flags
var debugGui bool = false
var debugError bool = false
var debugDump bool = false
var debugNode bool = false
var debugTabs bool = false
var debugFlags bool = false
var debugChange bool = false	// shows user events like mouse and keyboard
var debugPlugin	bool = false

// for printing out the binary tree
var listChildrenParent *Node
var listChildrenDepth int = 0
var defaultPadding = "  "

func SetDebug (s bool) {
	debugGui     = s
	debugTabs    = s

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

	// send the flag to the toolkit
	n := Config.flag
	log(debugChange, "Set() toolkit flag", s, "to", b)
	n.widget.Action = "Set"
	n.widget.S = s
	n.widget.B = b
	send(nil, n)
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
	if ! debugDump {
		return
	}
	Indent("NODE DUMP START")
	Indent("id           = ", n.id)
	Indent("Name         = ", n.Name)
	Indent("Width        = ", n.Width)
	Indent("Height       = ", n.Height)
	Indent("Widget Name  = ", n.widget.Name)
	Indent("Widget Type  = ", n.widget.Type)
	Indent("Widget Id    = ", n.widget.GetId())

	if (n.parent == nil) {
		Indent("parent       = nil")
	} else {
		Indent("parent.id    =", n.parent.id)
	}
	if (n.children != nil) {
		Indent("children     = ", n.children)
	}
	if (n.Custom != nil) {
		Indent("Custom       = ", n.Custom)
	}
	Indent("NODE DUMP END")
}

func Indent(a ...interface{}) {
	logindent(listChildrenDepth, defaultPadding, a...)
}

func (n *Node) dumpWidget() string {
	var info, d string

	info = n.widget.Type.String()

	info += ", " + n.widget.Name
	if (n.Name != n.widget.Name) {
		info += " NAME MISMATCH"
	}
	if (n.widget.Type == toolkit.Checkbox) {
		info += " = " + strconv.FormatBool(n.widget.B)
	}

	d = strconv.Itoa(n.id) + " " + info

	var tabs string
	for i := 0; i < listChildrenDepth; i++ {
		tabs = tabs + defaultPadding
	}
	d = tabs + d
	logindent(listChildrenDepth, defaultPadding, n.id, info)
	return d
}

func (n *Node) ListChildren(dump bool, dropdown *Node, mapNodes map[string]*Node) {
	s := n.dumpWidget()
	if (dropdown != nil) {
		dropdown.AddDropdownName(s)
		if (mapNodes != nil) {
			mapNodes[s] = n
		}
	}

	if (dump == true) {
		n.Dump()
	}
	if len(n.children) == 0 {
		if (n.parent == nil) {
			return
		}
		log(debugNode, "\t\t\tparent =",n.parent.id)
		if (listChildrenParent != nil) {
			log(debugNode, "\t\t\tlistChildrenParent =",listChildrenParent.id)
			if (listChildrenParent.id != n.parent.id) {
				// log("parent.child does not match child.parent")
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
		child.ListChildren(dump, dropdown, mapNodes)
		listChildrenDepth -= 1
	}
	return
}
