package gui

import (
	"log"
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func MessageWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBox(gw.UiWindow, msg1, msg2)
}

func ErrorWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBoxError(gw.UiWindow, msg1, msg2)
}

func DeleteWindow(name string) {
	log.Println("gui.DeleteWindow() START name =", name)
	window := Data.WindowMap[name]
	if window == nil {
		log.Println("gui.DeleteWindow() NO WINDOW WITH name =", name)
		return
	}

	log.Println("gui.DumpBoxes() MAP: ", name)
	log.Println("gui.DumpBoxes()\tWindow.name =", window.Name)
	if window.TabNumber == nil {
		log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
	}
	tab := *window.TabNumber
	log.Println("gui.DumpBoxes() \tWindows.TabNumber =", tab)
	log.Println("gui.DumpBoxes() \tSHOULD DELETE TAB", tab, "HERE")
	window.UiTab.Delete(tab)
	delete(Data.WindowMap, name)

	// renumber tabs here
	for name, window := range Data.WindowMap {
		log.Println("gui.DumpBoxes() MAP: ", name)
		if window.TabNumber == nil {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
		} else {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
			if tab < *window.TabNumber {
				log.Println("gui.DumpBoxes() \tSubtracting 1 from TabNumber")
				*window.TabNumber -= 1
				log.Println("gui.DumpBoxes() \tWindows.TabNumber is now =", *window.TabNumber)
			}
		}
	}
}

func CreateWindow(title string, tabname string, x int, y int, custom func() ui.Control) *Node {
	n := CreateBlankWindow(title, x, y)
	if (n.box == nil) {
		log.Println("SERIOUS ERROR n.box == nil in CreateWindow()")
		log.Println("SERIOUS ERROR n.box == nil in CreateWindow()")
		log.Println("SERIOUS ERROR n.box == nil in CreateWindow()")
		log.Println("SERIOUS ERROR n.box == nil in CreateWindow()")
	}
	n.AddTab(title, custom())
	// TODO: run custom() here // Oct 9
	return n
}

func (n *Node) Add(e Element) *Node {
	newNode := n.addNode("testingAdd")
	if(e == Tab) {
		log.Println("gui.Add() SHOULD ADD element =", e.String())
	}
	return newNode
}

//
// Create a new node
// if parent == nil, that means it is a new window and needs to be put
// in the window map (aka Data.NodeMap)
//
func (parent *Node) addNode(title string) *Node {
	var node Node
	node.Name = title
	node.Width = parent.Width
	node.Height = parent.Height
	node.parent = parent

	id := Config.prefix + strconv.Itoa(Config.counter)
	Config.counter += 1
	node.id = id

	parent.Append(&node)
	return &node
}

func makeNode(parent *Node, title string, x int, y int) *Node {
	var node Node
	node.Name = title
	node.Width = x
	node.Height = y

	id := Config.prefix + strconv.Itoa(Config.counter)
	Config.counter += 1
	node.id = id

	// panic("gui.makeNode() START")
	if (parent == nil) {
		if (Data.NodeMap[title] != nil) {
			log.Println("Duplicate uiNewWindow() name =", title)
			// TODO: just change the 'title' to something unique
			panic(fmt.Sprintf("Duplicate uiNewWindow() name = %s\n", title))
			return nil
		}
		// panic("gui.makeNode() before NodeMap()")
		Data.NodeMap[title] = &node
		Data.NodeArray = append(Data.NodeArray, &node)
		Data.NodeSlice = append(Data.NodeSlice, &node)
		// panic("gui.makeNode() after NodeMap()")
		return &node
	} else {
		panic("gui.makeNode() before Append()")
		parent.Append(&node)
		panic("gui.makeNode() after Append()")
	}
	node.parent = parent
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

func CreateBlankWindow(title string, x int, y int) *Node {
	node := mapWindow(nil, nil, title, x, y)
	box := node.box
	log.Println("gui.CreateBlankWindow() title = box.Name =", box.Name)

	node.uiNewWindow(box.Name, x, y)
	window := node.uiWindow

	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy()", box.Name)
		window.Destroy()
		return true
	})

	box.Window.UiWindow = window
	return node
}

/*
func (n *Node) initBlankWindow() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	return hbox
}
*/

func makeBlankNode(title string) *Node {
	log.Println("gui.makeBlankNode() title =", title)
	if Data.NodeMap[title] != nil {
		log.Println("gui.makeBlankNode() already exists title =", title)
		title = title + Config.prefix + strconv.Itoa(Config.counter)
		Config.counter += 1
	}
	if Data.NodeMap[title] != nil {
		panic("gui.makeBlankNode() already exists")
		return nil
	}

	node := makeNode(nil, title, x, y)
	return node
}

func mapWindow(parent *Node, window *ui.Window, title string, x int, y int) *Node {
	log.Println("gui.WindowMap START title =", title)
	if Data.WindowMap[title] != nil {
		log.Println("Data.WindowMap[title] already exists title =", title)
		title = title + Config.prefix + strconv.Itoa(Config.counter)
		Config.counter += 1
	}
	if Data.WindowMap[title] != nil {
		log.Println("Data.WindowMap[title] already exists title =", title)
		panic("Data.WindowMap[newGuiWindow.Name] already exists")
		return nil
	}

	var newGuiWindow GuiWindow
	newGuiWindow.Width = x
	newGuiWindow.Height = y
	newGuiWindow.Name = title
	newGuiWindow.UiWindow = window

	newGuiWindow.BoxMap = make(map[string]*GuiBox)
	newGuiWindow.EntryMap = make(map[string]*GuiEntry)

	Data.WindowMap[newGuiWindow.Name] = &newGuiWindow

	var box GuiBox
	box.Window = &newGuiWindow
	box.Name = title

	node := makeNode(parent, title, x, y)
	node.box = &box
	node.uiWindow = window
	box.node = node

	newGuiWindow.BoxMap["jcarrInitTest"] = &box

	return node
}

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

	var node *Node
	node = mapWindow(nil, nil, title, w, h)
	box := node.box
	log.Println("gui.NewWindow() title = box.Name =", box.Name)

	node.uiNewWindow(box.Name, w, h)
	window := node.uiWindow

	f := Config.Exit
	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy() on node.Name =", node.Name)
		if (f != nil) {
			f(node)
		}
		return true
	})

	box.Window.UiWindow = window
	return node
}
