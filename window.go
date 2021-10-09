package gui

import (
	"log"
	"strconv"
	"time"

	"github.com/andlabs/ui"

	// import "regexp"

	_ "github.com/andlabs/ui/winmanifest"
)

func initUI(name string, callback func(*GuiBox) *GuiBox) {
	ui.Main(func() {
		log.Println("gui.initUI() inside ui.Main()")

		node := InitWindow(nil, nil, "StartNewWindow"+name, 0)
		box := node.box
		box = callback(box)
		window := box.Window
		log.Println("StartNewWindow() box =", box)

		window.UiWindow.Show()
	})
}

func StartNewWindow(bg bool, name string, axis int, callback func(*GuiBox) *GuiBox) {
	log.Println("StartNewWindow() ui.Main() Create a new window")

	if bg {
		go initUI(name, callback)
		time.Sleep(500 * time.Millisecond) // this might make it more stable on windows?
	} else {
		initUI(name, callback)
	}
}

func MessageWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBox(gw.UiWindow, msg1, msg2)
}

func ErrorWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBoxError(gw.UiWindow, msg1, msg2)
}

//
// This creates a new 'window' (which is just a tab in the window)
// This is this way because on Linux you can have more than one
// actual window but that does not appear to work on the MacOS or Windows
//
func InitWindow(parent *Node, gw *GuiWindow, name string, axis int) *Node {
	log.Println("gui.InitWindow() START")

	var box *GuiBox
	var node *Node

	if gw == nil {
		node = mapWindow(parent, nil, name, Config.Width, Config.Height)
		box = node.box
	} else {
		node = mapWindow(parent, gw.UiWindow, name, Config.Width, Config.Height)
		box = node.box
	}

	// box.Window = &newGuiWindow
	newGuiWindow := box.Window

	// This is the first window. One must create it here
	if gw == nil {
		log.Println("gui.initWindow() ADDING ui.NewWindow()")
		node = uiNewWindow(node, name, Config.Width, Config.Height)
		box.node = node
		if (node.box == nil) {
			node.box = box
		}
		w := node.uiWindow
		newGuiWindow.UiWindow = w

		// newGuiWindow.UiWindow.SetTitle("test")
		w.OnClosing(func(*ui.Window) bool {
			log.Println("gui.InitWindow() OnClosing() THIS WINDOW IS CLOSING newGuiWindow=", newGuiWindow)
			// newGuiWindow.UiWindow.Destroy()
			if Config.Exit == nil {
				ui.Quit()
			} else {
				// allow a custom exit function
				Config.Exit(newGuiWindow)
			}
			return true
		})

		newGuiWindow.UiTab = ui.NewTab()
		newGuiWindow.UiWindow.SetChild(newGuiWindow.UiTab)
		newGuiWindow.UiWindow.SetMargined(true)
		tmp := 0
		newGuiWindow.TabNumber = &tmp
		node.uiTab = newGuiWindow.UiTab
	} else {
		newGuiWindow.UiWindow = gw.UiWindow
		newGuiWindow.UiTab = gw.UiTab
		node.uiTab = newGuiWindow.UiTab
	}

	newGuiWindow.BoxMap = make(map[string]*GuiBox)
	newGuiWindow.EntryMap = make(map[string]*GuiEntry)
	// Data.Windows = append(Data.Windows, &newGuiWindow)

	if newGuiWindow.UiTab == nil {
		tabnum := 0
		newGuiWindow.TabNumber = &tabnum
	} else {
		tabnum := newGuiWindow.UiTab.NumPages()
		newGuiWindow.TabNumber = &tabnum
		if (node.uiTab == nil) {
			node.uiTab = newGuiWindow.UiTab
		}
	}

	Data.WindowMap[newGuiWindow.Name] = newGuiWindow

	if (box.node == nil) {
		fn := FindNode("full initTab")
		log.Println("InitWindow() fn =", fn)
		log.Println("InitWindow() mapping node <=> box")
		box.node = fn
		if (fn.box == nil) {
			log.Println("InitWindow() mapping node <=> box")
			fn.box = box
		}
	}
	if (box.node == nil) {
		Data.ListChildren(true)
		log.Println("InitWindow() box has a FUCKING nil node")
		fn := FindNode("full initTab")
		log.Println("InitWindow() fn =", fn)
		panic(-1)
	}

	if (newGuiWindow.node == nil) {
		Data.ListChildren(true)
		log.Println("InitWindow() newGuiWindow has a FUCKING nil node")
		// panic(-1)
	}

	log.Println("InitWindow() END *GuiWindow =", newGuiWindow)
	if (box.node == nil) {
		box.Dump()
		panic(-1)
	}
	box.Dump()
	box.node.Dump()
	if (box.node != node) {
		log.Println("InitWindow() box.node != node. Hmmm....")
		log.Println("InitWindow() box.node != node. Hmmm....")
		log.Println("InitWindow() box.node != node. Hmmm....")
		panic(-1)
	}
	if (node.box != box) {
		log.Println("InitWindow() node.box != box. Hmmm....")
		log.Println("InitWindow() node.box != box. Hmmm....")
		log.Println("InitWindow() node.box != box. Hmmm....")
		panic(-1)
	}
	if (node.uiTab == nil) {
		// DebugNodeChildren()
		// panic("node.uiTab = nil")
	}
	return node
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
	n.InitTab(title)
	// TODO: run custom() here // Oct 9
	return n
}

//
// Create a new node
// if parent == nil, that means it is a new window and needs to be put
// in the window map (aka Data.NodeMap)
//
func makeNode(parent *Node, title string, x int, y int) *Node {
	var node Node
	node.Name = title
	node.Width = x
	node.Height = y

	id := "jwc" + strconv.Itoa(Config.counter)
	Config.counter += 1
	node.id = id

	if (parent == nil) {
		if (Data.NodeMap[title] != nil) {
			log.Println("Duplicate uiNewWindow() name =", title)
			// TODO: just change the 'title' to something unique
			return nil
		}
		Data.NodeMap[title] = &node
		return &node
	} else {
		parent.Append(&node)
	}
	node.parent = parent
	return &node
}

func uiNewWindow(node *Node, title string, x int, y int) *Node {
	if (node == nil) {
		node = makeNode(nil, title, x, y)
	}

	w := ui.NewWindow(title, x, y, false)
	w.SetBorderless(false)
	w.OnClosing(func(*ui.Window) bool {
		log.Println("ui.Window().OnClosing() IS EMPTY FOR window name =", title)
		return true
	})
	w.SetMargined(true)
	w.Show()
	node.uiWindow = w
	// w.node = &node
	return node
}

func CreateBlankWindow(title string, x int, y int) *Node {
	n := mapWindow(nil, nil, title, x, y)
	box := n.box
	log.Println("gui.CreateBlankWindow() title = box.Name =", box.Name)

	n = uiNewWindow(n, box.Name, x, y)
	window := n.uiWindow

	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy()", box.Name)
		window.Destroy()
		return true
	})

	box.Window.UiWindow = window
	return n
}

func initBlankWindow() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	return hbox
}

var master = 0

func mapWindow(parent *Node, window *ui.Window, title string, x int, y int) *Node {
	log.Println("gui.WindowMap START title =", title)
	if Data.WindowMap[title] != nil {
		log.Println("Data.WindowMap[title] already exists title =", title)
		master = master + 1
		title = title + " jcarr " + strconv.Itoa(master)
	}
	if Data.WindowMap[title] != nil {
		log.Println("Data.WindowMap[title] already exists title =", title)
		panic("Data.WindowMap[newGuiWindow.Name] already exists")
		return nil
	}

	log.Println("gui.WindowMap START title =", title)
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

	// func makeNode(parent *Node, title string, x int, y int) *Node {
	node := makeNode(parent, title, x, y)
	node.box = &box
	node.uiWindow = window
	box.node = node

	newGuiWindow.BoxMap["jcarrInitTest"] = &box

	return node
}

func NewWindow(title string, x int, y int) *GuiBox {
	n := mapWindow(nil, nil, title, x, y)
	box := n.box
	log.Println("gui.NewWindow() title = box.Name =", box.Name)

	n = uiNewWindow(n, box.Name, x, y)
	window := n.uiWindow

	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy()", box.Name)
		window.Destroy()
		return true
	})

	box.Window.UiWindow = window
	return box
}
