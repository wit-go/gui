package gui

import (
	"log"
	"reflect"
)

// import toolkit "git.wit.org/wit/gui/toolkit/andlabs"
import newtoolkit	"git.wit.org/wit/gui/toolkit"

//
// All GUI Data Structures and functions that are external
// within the toolkit/ abstraction layer
//
// More than one Window is not supported in a cross platform
// sense & may never be. On many toolkits you have to have 'tabs'
// Native Windows and MacOS toolkits work with tabs
//
// If that is the case, this code should abstract the concept of
// windows and make everything 'tabs'
//

var Config GuiConfig

func GetDebug () bool {
	return Config.Debug.Debug
}

func SetDebug (s bool) {
	Config.Debug.Debug = s
	// also set these
	Config.Debug.Dump = s
	Config.Debug.Node = s
	// toolkit.DebugToolkit = s
}

func GetDebugToolkit () bool {
	return Config.Debug.Toolkit
}

func SetDebugToolkit (s bool) {
	Config.Debug.Toolkit = s
}

func ShowDebugValues() {
	log.Println("\t wit/gui Debug =", Config.Debug.Debug)
	log.Println("\t wit/gui DebugDump =", Config.Debug.Dump)
	log.Println("\t wit/gui DebugNode =", Config.Debug.Node)
	log.Println("\t wit/gui DebugTabs =", Config.Debug.Tabs)
	log.Println("\t wit/gui DebugPlugin =", Config.Debug.Plugin)
	log.Println("\t wit/gui DebugChange =", Config.Debug.Change)
	log.Println("\t wit/gui DebugToolkit =", Config.Debug.Toolkit)
}

// This struct can be used with go-arg
type GuiDebug struct {
	// These are global debugging settings
	// TODO: move to a standard logging system
	Debug   bool
	Dump    bool
	Node    bool
	Tabs    bool
	Plugin  bool
	Change  bool `help:"debug mouse clicks and keyboard input"`
	Toolkit bool `help:"debug toolkit"`
}

type GuiConfig struct {
	// This is the master node. The Binary Tree starts here
	master	*Node

	// These are shortcuts to pass default values to make a new window
	Title      string
	Width      int
	Height     int
	Exit       func(*Node)

	Debug GuiDebug

	// hacks
	depth      int
	counter    int  // used to make unique ID's
	prefix     string
}

// The Node is simply the name and the size of whatever GUI element exists
type Node struct {
	id     int

	// deprecate these and use toolkit.Widget
	Name   string
	Width  int
	Height int

	Widget	newtoolkit.Widget

	// this function is run when there are mouse or keyboard events
	OnChanged func(*Node)

	parent	*Node
	// TODO: make children a double linked list since some toolkits require order (?)
	children []*Node

	// hmm. how do you handle this when the toolkits are plugins?
	// toolkit	*toolkit.Toolkit

	// things that may not really be needed (?)
	custom    func()
	checked   bool
	text      string
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Window() *Node {
	return n.parent
}

func (n *Node) Dump() {
	if ! Config.Debug.Dump {
		return
	}
	IndentPrintln("NODE DUMP START")
	IndentPrintln("id         = ", n.id)
	IndentPrintln("Name       = ", n.Name)
	IndentPrintln("Width      = ", n.Width)
	IndentPrintln("Height     = ", n.Height)

	if (n.parent == nil) {
		IndentPrintln("parent     = nil")
	} else {
		IndentPrintln("parent.id  =", n.parent.id)
	}
	if (n.children != nil) {
		IndentPrintln("children   = ", n.children)
	}
	if (n.custom != nil) {
		IndentPrintln("custom     = ", n.custom)
	}
	IndentPrintln("checked    = ", n.checked)
	if (n.OnChanged != nil) {
		IndentPrintln("OnChanged  = ", n.OnChanged)
	}
	IndentPrintln("text       = ", reflect.ValueOf(n.text).Kind(), n.text)
//	if (n.toolkit != nil) {
//		IndentPrintln("toolkit    = ", reflect.ValueOf(n.toolkit).Kind())
//		n.toolkit.Dump()
//	}
//	if (n.id == nil) {
//		// Node structs should never have a nil id.
//		// I probably shouldn't panic here, but this is just to check the sanity of
//		// the gui package to make sure it's not exiting
//		panic("gui.Node.Dump() id == nil TODO: make a unigue id here in the golang gui library")
//	}
	IndentPrintln("NODE DUMP END")
}

/*
func (n *Node) SetName(name string) {
	n.toolkit.SetWindowTitle(name)
	return
}
*/

func (n *Node) Append(child *Node) {
	n.children = append(n.children, child)
	if (Config.Debug.Debug) {
		log.Println("child node:")
		child.Dump()
		log.Println("parent node:")
		n.Dump()
	}
	// time.Sleep(3 * time.Second)
}

/*
func (n *Node) List() {
	findByIdDFS(n, "test")
}
*/

var listChildrenParent *Node
var listChildrenDepth int = 0
var defaultPadding = "  "

func IndentPrintln(a ...interface{}) {
	indentPrintln(listChildrenDepth, defaultPadding, a)
}

func indentPrintln(depth int, format string, a ...interface{}) {
	var tabs string
	for i := 0; i < depth; i++ {
		tabs = tabs + format
	}

	// newFormat := tabs + strconv.Itoa(depth) + " " + format
	newFormat := tabs + format
	log.Println(newFormat, a)
}

func (n *Node) ListChildren(dump bool) {
	indentPrintln(listChildrenDepth, defaultPadding, n.id, n.Width, n.Height, n.Name)

	if (dump == true) {
		n.Dump()
	}
	if len(n.children) == 0 {
		if (n.parent == nil) {
		} else {
			if (Config.Debug.Node) {
				log.Println("\t\t\tparent =",n.parent.id)
			}
			if (listChildrenParent != nil) {
				if (Config.Debug.Node) {
					log.Println("\t\t\tlistChildrenParent =",listChildrenParent.id)
				}
				if (listChildrenParent.id != n.parent.id) {
					log.Println("parent.child does not match child.parent")
					panic("parent.child does not match child.parent")
				}
			}
		}
		if (Config.Debug.Node) {
			log.Println("\t\t", n.id, "has no children")
		}
		return
	}
	for _, child := range n.children {
		// log.Println("\t\t", child.id, child.Width, child.Height, child.Name)
		if (child.parent != nil) {
			if (Config.Debug.Node) {
				log.Println("\t\t\tparent =",child.parent.id)
			}
		} else {
			log.Println("\t\t\tno parent")
			panic("no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		if (Config.Debug.Node) {
			if (child.children == nil) {
				log.Println("\t\t", child.id, "has no children")
			} else {
				log.Println("\t\t\tHas children:", child.children)
			}
		}
		listChildrenParent = n
		listChildrenDepth += 1
		child.ListChildren(dump)
		listChildrenDepth -= 1
	}
	return
}
