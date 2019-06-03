package gui

import "log"
import "time"
import "fmt"
// import "reflect"

// import "github.com/andlabs/ui"
// import _ "github.com/andlabs/ui/winmanifest"
// import "github.com/davecgh/go-spew/spew"
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
			for i, window := range Data.Windows {
				log.Println("watchGUI() Data.Windows", i, "Name =", window.Name)
				for name, abox := range window.BoxMap {
					log.Printf("\twatchGUI() BOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
					/*
					if (name == "DEBUG") {
						log.Println("\t\twatchGUI() BOX abox =", reflect.TypeOf(abox))
						win := abox.Window
						log.Println("\t\twatchGUI() BOX win =", reflect.TypeOf(win))
						area := win.Area
						log.Println("\t\twatchGUI() BOX area =", reflect.TypeOf(area), area.UiArea)
						// spew.Dump(area.UiArea)
						// area.UiArea.Show()
						// time.Sleep(2000 * time.Millisecond)
						// os.Exit(0)
					}
					*/
				}
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
		b.Heading  = fmt.Sprintf("heading%d", key)
		parts = append(parts, b)
	}

	log.Println("Sleep for 2 seconds, then try to add new tabs")
	time.Sleep(1 * 1000 * 1000 * 1000)
	// AddTableTab(Data.Window1.T, 1, "test seven", 7, parts, nil)
}

/*
func runTestHide(b *GuiButton) {
	log.Println("runTestHide START")
	Data.Window1.Box1.Hide()
	Data.Window1.Box2.Hide()
	// time.Sleep(2000 * time.Millisecond)
	Data.State = "HIDE"
	log.Println("runTestHide END")
}
*/
