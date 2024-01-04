package gui

// Lots of debugging things:
// A function dump out the binary tree

import (
	"errors"
	"strconv"

	"go.wit.com/log"
	"go.wit.com/gui/gui/toolkit"
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
		log.Log(GUI, "Can't set unknown flag", s)
	}

	a := new(toolkit.Action)
	a.ActionType = toolkit.Set
	a.WidgetType = toolkit.Flag
	a.S = s
	a.B = b
	sendAction(a)
}

func ShowDebugValues() {
	// The order here should match the order in the GUI
	// TODO: get the order from the node binary tree
	log.Log(true, "Debug        =", debugGui)
	log.Log(true, "DebugError   =", debugError)
	log.Log(true, "DebugChange  =", debugChange)
	log.Log(true, "DebugDump    =", debugDump)
	log.Log(true, "DebugTabs    =", debugTabs)
	log.Log(true, "DebugPlugin  =", debugPlugin)
	log.Log(true, "DebugNode    =", debugNode)

	SetFlag("Show", true)
}

func (n *Node) Dump() {
	b := true
	Indent(b, "NODE DUMP START")
	Indent(b, "id           = ", n.id)
	Indent(b, "Name         = ", n.Name)
	Indent(b, "(X,Y)        = ", n.X, n.Y)
	Indent(b, "Next (W,H)   = ", n.NextW, n.NextH)

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

	a := new(toolkit.Action)
	a.ActionType = toolkit.Dump
	a.WidgetId = n.id
	sendAction(a)
}

func Indent(b bool, a ...interface{}) {
	logindent(b, listChildrenDepth, defaultPadding, a...)
}

func (n *Node) dumpWidget(b bool) string {
	var info, d string

	if (n == nil) {
		log.Error(errors.New("dumpWidget() node == nil"))
		return ""
	}
	info = n.WidgetType.String()

	d = strconv.Itoa(n.id) + " " + info + " " + n.Name

	var tabs string
	for i := 0; i < listChildrenDepth; i++ {
		tabs = tabs + defaultPadding
	}
	logindent(b, listChildrenDepth, defaultPadding, d)
	return tabs + d
}

func (n *Node) Children() []*Node {
	return n.children
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
		log.Log(NODE, "\t\t\tparent =",n.parent.id)
		if (listChildrenParent != nil) {
			log.Log(NODE, "\t\t\tlistChildrenParent =",listChildrenParent.id)
			if (listChildrenParent.id != n.parent.id) {
				log.Log(true, "parent =",n.parent.id, n.parent.Name)
				log.Log(true, "listChildrenParent =",listChildrenParent.id, listChildrenParent.Name)
				log.Log(true, listChildrenParent.id, "!=", n.parent.id)
				log.Exit("parent.child does not match child.parent")
			}
		}
		log.Log(NODE, "\t\t", n.id, "has no children")
		return
	}
	for _, child := range n.children {
		if (child.parent != nil) {
			log.Log(NODE, "\t\t\tparent =",child.parent.id)
		} else {
			log.Log(GUI, "\t\t\tno parent")
			// memory corruption? non-threadsafe access?
			// can all binary tree changes to Node.parent & Node.child be forced into a singular goroutine?
			panic("something is wrong with the wit golang gui logic and the binary tree is broken. child has no parent")
		}
		if (child.children == nil) {
			log.Log(NODE, "\t\t", child.id, "has no children")
		} else {
			log.Log(NODE, "\t\t\tHas children:", child.children)
		}
		listChildrenParent = n
		listChildrenDepth += 1
		// child.ListChildren(dump, dropdown, mapNodes)
		child.ListChildren(dump)
		listChildrenDepth -= 1
	}
	return
}

// b bool, print if true
func logindent(b bool, depth int, format string, a ...any) {
	var tabs string
	for i := 0; i < depth; i++ {
		tabs = tabs + format
	}

	// newFormat := tabs + strconv.Itoa(depth) + " " + format
	newFormat := tabs + format

	// array prepend(). Why isn't this a standard function. It should be:
	// a.prepend(debugGui, newFormat)
	a = append([]any{b, newFormat}, a...)
	log.Log(b, a...)
}
