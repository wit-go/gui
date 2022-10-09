package gui

import (
	"fmt"
	"log"
	"time"

	// "github.com/davecgh/go-spew/spew"
)

// WatchGUI() opens a goroutine
//
// From that goroutine, it dumps out debugging information every 4 seconds
/*
	TODO: add configuration triggers on what to dump out
	TODO: allow this to be sent to /var/log, syslogd, systemd's journalctl, etc
*/
func WatchGUI() {
	count := 0

	for {
		if count > 20 {
			log.Println("Sleep() in watchGUI()")
			if Config.Debug {
				// DumpBoxes()
			}
			count = 0
		}
		count += 1
		time.Sleep(200 * time.Millisecond)
	}
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

func (dn *GuiData) ListChildren(dump bool) {
	if Data.NodeMap == nil {
		log.Println("gui.Data.ListChildren() Data.NodeMap == nil")
		return
	}
	log.Println("gui.Data.ListChildren() Data.NodeMap:")
	for name, node := range Data.NodeMap {
		listChildrenDepth = 0
		if (dump == true) {
			log.Println("tgui.Data.ListChildren() node =", node.id, node.Width, node.Height, name)
			node.Dump()
		}
		node.ListChildren(dump)
	}
}
