package gui

import (
	"log"
//	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func findUiWindow() *ui.Window {
	for _, node := range Data.NodeMap {
		if (node.uiWindow != nil) {
			return node.uiWindow
		}
	}
	return nil
}

func MessageWindow(msg1 string, msg2 string) (*Node) {
	uiW := findUiWindow()
	ui.MsgBox(uiW, msg1, msg2)
	// TODO: make new node
	return nil
}

func ErrorWindow(msg1 string, msg2 string) (*Node) {
	uiW := findUiWindow()
	ui.MsgBoxError(uiW, msg1, msg2)
	return nil
}

func initNode(title string, x int, y int) *Node {
	var node Node
	node.Name = title
	node.Width = x
	node.Height = y

	id := Config.prefix + strconv.Itoa(Config.counter)
	Config.counter += 1
	node.id = id

	if (Data.NodeMap[title] != nil) {
		log.Println("Duplicate window name =", title)
		// TODO: just change the 'title' to something unique
		// panic(fmt.Sprintf("Duplicate window name = %s\n", title))
		return Data.NodeMap[title]
	}
	Data.NodeMap[title] = &node
	Data.NodeArray = append(Data.NodeArray, &node)
	Data.NodeSlice = append(Data.NodeSlice, &node)
	return &node
	//	parent.Append(&node)
	//node.parent = parent
	return &node
}

func (parent *Node) makeNode(title string, x int, y int) *Node {
	var node Node
	node.Name = title
	node.Width = x
	node.Height = y

	id := Config.prefix + strconv.Itoa(Config.counter)
	Config.counter += 1
	node.id = id

	parent.Append(&node)
	node.parent = parent
	return &node
}

func (n *Node) AddNode(title string) *Node {
	var node Node
	node.Name = title
	node.Width = n.Width
	node.Height = n.Height

	id := Config.prefix + strconv.Itoa(Config.counter)
	Config.counter += 1
	node.id = id

	n.Append(&node)
	node.parent = n
	return &node
}

func (n *Node) uiNewWindow(title string, x int, y int) {
	w := ui.NewWindow(title, x, y, false)
	w.SetBorderless(false)
	f := Config.Exit
	w.OnClosing(func(*ui.Window) bool {
		if (Config.Debug) {
			log.Println("ui.Window().OnClosing()")
		}
		if (f != nil) {
			f(n)
		}
		return true
	})
	w.SetMargined(true)
	w.Show()
	n.uiWindow = w
	// w.node = &node
	return
}

/*
func mapWindow(parent *Node, window *ui.Window, title string, x int, y int) *Node {
	log.Println("gui.WindowMap START title =", title)

	node := makeNode(parent, title, x, y)
	node.uiWindow = window

	return node
}
*/

// This routine creates a blank window with a Title and size (W x H)
//
// This routine can not have any arguements due to the nature of how
// it can be passed via the 'andlabs/ui' queue which, because it is
// cross platform, must pass UI changes into the OS threads (that is
// my guess).
func NewWindow() *Node {
	title := Config.Title
	w     := Config.Width
	h     := Config.Height

	if (Data.NodeMap[title] != nil) {
		log.Println("Duplicate window name =", title)
		Data.NodeMap[title].Dump()
		Data.NodeMap[title].ListChildren(false)
		uiW := Data.NodeMap[title].uiWindow
		if (uiW != nil) {
			uiW.Show()
		}
		panic("check here to see if window is really alive")
		return Data.NodeMap[title]
	}

	var n *Node
	n = initNode(title, w, h)
	n.uiNewWindow(title, w, h)
	window := n.uiWindow

	f := Config.Exit
	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy() on node.Name =", n.Name)
		if (f != nil) {
			f(n)
		}
		return true
	})

	n.uiWindow = window
	if(n.uiWindow == nil) {
		panic("node.uiWindow == nil. This should never happen")
	}
	return n
}
