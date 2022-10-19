package gui

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// WatchGUI() opens a goroutine
//
// From that goroutine, it dumps out debugging information every 4 seconds
/*
	TODO: add configuration triggers on what to dump out
	TODO: allow this to be sent to /var/log, syslogd, systemd's journalctl, etc
*/
/*
func watchGUI() {
	count := 0

	for {
		if count > 20 {
			log.Println("Sleep() in watchGUI()")
			if Config.Debug {
				dumpBoxes()
			}
			count = 0
		}
		count += 1
		time.Sleep(200 * time.Millisecond)
	}
}
*/

func dumpWindows() {
	for name, _ := range Data.WindowMap {
		log.Println("gui.DumpWindows() window =", name)
	}
}

func dumpMap() {
	for name, window := range Data.WindowMap {
		log.Println("gui.DumpBoxes() MAP: ", name)
		log.Println("gui.DumpBoxes()     BOXES:", name)
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
		}
	}
}

func dumpBoxes() {
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

func (dn *GuiData) DumpNodeMap() {
	log.Println("DebugDataNodeMap():")
	for name, node := range dn.NodeMap {
		log.Println("\tNode =", node.id, node.Width, node.Height, name)
		if (node.children == nil) {
			log.Println("\t\tNo children")
		} else {
			log.Println("\t\tHas children:", node.children)
		}
		// node.SetName("yahoo")
		// log.Println("\tData.NodeMap node =", node)
	}
}

/*
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
*/

func (dn *GuiData) ListChildren(dump bool) {
	if Data.NodeMap == nil {
		log.Println("gui.Data.ListChildren() Data.NodeMap == nil")
		return
	}
	log.Println("gui.Data.ListChildren() Data.NodeMap:")
	for name, node := range Data.NodeMap {
		log.Println("\tgui.Data.ListChildren() node =", node.id, node.Width, node.Height, name)
		if (dump == true) {
			node.Dump()
		}
		node.ListChildren(dump)
	}
}
