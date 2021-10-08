package gui

import (
	"fmt"
	"os"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// import "reflect"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"

//
// Dump out debugging information every 4 seconds
//
func WatchGUI() {
	count := 0

	for {
		if count > 20 {
			log.Println("Sleep() in watchGUI()")
			if Config.Debug {
				DumpBoxes()
			}
			count = 0
		}
		count += 1
		time.Sleep(200 * time.Millisecond)
	}
}

func DumpWindows() {
	for name, _ := range Data.WindowMap {
		log.Println("gui.DumpWindows() window =", name)
	}
}

func DumpMap() {
	for name, window := range Data.WindowMap {
		log.Println("gui.DumpBoxes() MAP: ", name)
		log.Println("gui.DumpBoxes()     BOXES:", name)
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
		}
	}
}

func DumpBoxes() {
	for name, window := range Data.WindowMap {
		log.Println("gui.DumpBoxes() MAP: ", name)
		if window.TabNumber == nil {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
		} else {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
		}
		log.Println("gui.DumpBoxes()\tWindow.name =", window.Name)
		// log.Println("gui.DumpBoxes()\tWindow.UiWindow type =", reflect.TypeOf(window.UiWindow))
		log.Println("gui.DumpBoxes()\tWindow.UiWindow =", window.UiWindow)
		log.Println("gui.DumpBoxes()\tWindow.UiTab =", window.UiTab)
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			if name == "MAINBOX" {
				if Config.Debug {
					scs := spew.ConfigState{MaxDepth: 1}
					scs.Dump(abox.UiBox)
				}
			}
		}
		if window.UiTab != nil {
			// log.Println("gui.DumpBoxes()\tWindow.UiTab type =", reflect.TypeOf(window.UiTab))
			// log.Println("gui.DumpBoxes()\tWindow.UiTab =", window.UiTab)
			pages := window.UiTab.NumPages()
			log.Println("gui.DumpBoxes()\tWindow.UiTab.NumPages() =", pages)
			// for i := 0; i < pages; i++ {
			// 	log.Println("gui.DumpBoxes()\t\tWindow.UiTab.Margined(", i, ") =", window.UiTab.Margined(i))
			// }
			// tmp := spew.NewDefaultConfig()
			// tmp.MaxDepth = 2
			// tmp.Dump(window.UiTab)
			if Config.Debug {
				scs := spew.ConfigState{MaxDepth: 2}
				scs.Dump(window.UiTab)
			}
		}
	}
	/*
		for i, window := range Data.Windows {
			if (window.TabNumber == nil) {
				log.Println("gui.DumpBoxes() Data.Windows", i, "Name =", window.Name, "TabNumber = nil")
			} else {
				log.Println("gui.DumpBoxes() Data.Windows", i, "Name =", window.Name, "TabNumber =", *window.TabNumber)
			}
		}
	*/
}

func addTableTab() {
	var parts []TableColumnData

	for key, foo := range []string{"BG", "TEXTCOLOR", "BUTTON", "TEXTCOLOR", "TEXTCOLOR", "TEXT", "BUTTON", "TEXT", "BUTTON"} {
		log.Println(key, foo)

		var b TableColumnData
		b.CellType = foo
		b.Heading = fmt.Sprintf("heading%d", key)
		parts = append(parts, b)
	}

	log.Println("Sleep for 1 second, then try to add new tabs")
	time.Sleep(1 * time.Second)
}

func DebugDataNodeMap() {
	if Data.NodeMap == nil {
		log.Println("DebugDataNodeMap() NodeMap == nil")
		return
	}
	log.Println("DebugDataNodeMap():")
	for name, node := range Data.NodeMap {
		log.Println("\tNode name =", node.Width, node.Height, name)
		if (node.children == nil) {
			log.Println("\t\tNo children")
		} else {
			log.Println("\t\tHas children:", node.children)
		}
		// node.SetName("yahoo")
		// log.Println("\tData.NodeMap node =", node)
	}
}

func DebugDataNodeChildren() {
	if Data.NodeMap == nil {
		log.Println("DebugDataNodeChildren() NodeMap == nil")
		return
	}
	log.Println("DebugDataNodeChildren():")
	for name, node := range Data.NodeMap {
		log.Println("\tNode name =", node.Width, node.Height, name)
		if (node.children == nil) {
			log.Println("\t\tNo children")
			break
		}
		log.Println("\t\tHas children:", node.children)
	}
}

func FindNode(name string) *Node {
	if Data.NodeMap == nil {
		log.Println("gui.FindNode() gui.Data.NodeMap == nil")
		return nil
	}
	log.Println("gui.FindNode() searching Data.NodeMap:")
	for id, node := range Data.NodeMap {
		log.Println("\tData.NodeMap name =", node.Width, node.Height, id)
		node.Dump()
		if (name == node.Name) {
			return node
		}
		newNode := findByName(node, name)
		if (newNode != nil) {
			return newNode
		}
		log.Println("gui.FindNode() could not find node name =", name)
		os.Exit(-1)
	}
	log.Println("gui.FindNode() could not find node name =", name)
	return nil
}

func DebugNodeChildren() {
	if Data.NodeMap == nil {
		log.Println("Data.NodeMap == nil")
		return
	}
	log.Println("Dumping Data.NodeMap:")
	for name, node := range Data.NodeMap {
		log.Println("\tData.NodeMap name =", node.id, node.Width, node.Height, name)
		// node.Dump()
		node.ListChildren()
		// node.SetName("yahoo")
		// log.Println("\tData.NodeMap node =", node)
	}
}
