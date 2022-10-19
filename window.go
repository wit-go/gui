package gui

import (
	"log"
	"strconv"

//	"github.com/andlabs/ui"
//	_ "github.com/andlabs/ui/winmanifest"
)

import toolkit "git.wit.org/wit/gui/toolkit/andlabs"

/*
func MessageWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBox(gw.UiWindow, msg1, msg2)
}

func ErrorWindow(gw *GuiWindow, msg1 string, msg2 string) {
	ui.MsgBoxError(gw.UiWindow, msg1, msg2)
}
*/

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
	log.Println("gui.DeleteWindow() \tSHOULD DELETE TAB", tab, "HERE")
	log.Println("gui.DumpBoxes() \tUiTab =", window.UiTab)
	tabnum			:= window.UiTab.NumPages()
	log.Println("gui.DumpBoxes() \tUiTab.NumPages() =", tabnum)
	if (tabnum > 0) {
		window.UiTab.Delete(tab)
	}
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

// func mapWindowOld(parent *Node, window *ui.Window, title string, x int, y int) *Node {
func mapWindow(title string, w int, h int) *Node {
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
	newGuiWindow.Width = w
	newGuiWindow.Height = h
	newGuiWindow.Name = title
	// newGuiWindow.UiWindow = window

	newGuiWindow.BoxMap = make(map[string]*GuiBox)
	newGuiWindow.EntryMap = make(map[string]*GuiEntry)

	Data.WindowMap[newGuiWindow.Name] = &newGuiWindow

	var box GuiBox
	box.Window = &newGuiWindow
	box.Name = title

	node := addNode(title, w, h)
	node.box = &box
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
	var n *Node
	var t *toolkit.Toolkit

	title := Config.Title
	w     := Config.Width
	h     := Config.Height
	f     := Config.Exit

	n = mapWindow(title, w, h)
	n.custom = f
	box := n.box
	log.Println("gui.NewWindow() title = box.Name =", box.Name)

	t = toolkit.NewWindow(title, w, h)
	t.Custom = func () {
		log.Println("GOT TO MY CUSTOM EXIT!!!! for window name:", box.Name)
		f(n)
	}
	n.Toolkit = t
	n.uiWindow = t.UiWindowBad // this is temporary

	window := n.uiWindow

	/*
	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy() on node.Name =", n.Name)
		if (f != nil) {
			f(n)
		}
		return true
	})
	*/

	box.Window.UiWindow = window
	if(n.uiWindow == nil) {
		panic("node.uiWindow == nil. This should never happen")
	}
	return n
}
