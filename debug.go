package gui

import "log"
import "time"
import "fmt"
// import "reflect"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"
import "github.com/davecgh/go-spew/spew"
// import pb "git.wit.com/wit/witProtobuf"

//
// this watches the GUI primarily to process protobuf's
// this is pointless or wrong but I use it for debugging
//
func WatchGUI() {
	count := 0

	for {
		if (count > 20) {
			log.Println("Sleep() in watchGUI()")
			if (Config.Debug) {
				DumpBoxes()
			}
			count = 0
		}
		count += 1
		time.Sleep(200 * time.Millisecond)
	}
}

func DumpBoxes() {
	for name, window := range Data.WindowMap {
		log.Println("gui.DumpBoxes() MAP: ", name)
		if (window.TabNumber == nil) {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
		} else {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
		}
		log.Println("gui.DumpBoxes()\tWindow.name =", window.Name)
		// log.Println("gui.DumpBoxes()\tWindow.UiWindow type =", reflect.TypeOf(window.UiWindow))
		log.Println("gui.DumpBoxes()\tWindow.UiWindow =", window.UiWindow)
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			if (name == "MAINBOX") {
				if (Config.Debug) {
					scs := spew.ConfigState{MaxDepth: 1}
					scs.Dump(abox.UiBox)
				}
			}
		}
		if (window.UiTab != nil) {
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
			if (Config.Debug) {
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
		b.Heading  = fmt.Sprintf("heading%d", key)
		parts = append(parts, b)
	}

	log.Println("Sleep for 2 seconds, then try to add new tabs")
	time.Sleep(1 * 1000 * 1000 * 1000)
}
