package gui

import "log"
import "time"
import "regexp"
import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

// THIS IS NOT CLEAN (almost?)

func ShowTab(gw *GuiWindow, tabname string, title string) *GuiWindow {
	log.Println("ShowTab() gw =", gw)
	if (gw.UiTab == nil) {
		log.Println("ShowTab() gw.UiTab = nil THIS IS BAD")
		os.Exit(-1)
	}
	window := InitGuiWindow(Data.Config, tabname, gw.MakeWindow, gw.UiWindow, gw.UiTab)
	window.UiTab.Delete(0)

	abox := window.MakeWindow(window)
	window.BoxMap[tabname] = abox
	window.UiTab.InsertAt(title, 0, abox.UiBox)
	window.UiTab.SetMargined(0, true)
	return window
}

func GuiInit() {
	ui.OnShouldQuit(func() bool {
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})
}

func AddMainTab(gw *GuiWindow) *GuiBox {
	log.Println("ShowMainTab() gw =", gw)
	log.Println("ShowMainTab() gw.UiTab =", gw.UiTab)

	window := InitGuiWindow(Data.Config, "MAIN", nil, gw.UiWindow, gw.UiTab)

	box := InitGuiBox(window, nil, ui.NewHorizontalBox(), "MAIN")

	if (Data.Debug) {
		log.Println("makeCloudInfoBox() add debugging buttons")
		addDebuggingButtons(box)
		box.UiBox.Append(ui.NewVerticalSeparator(), false)
	}

	// box := gw.MakeWindow(gw)
	// abox := makeCloudInfoBox(gw, box)
	return box
}

func ShowMainTabShowBox(gw *GuiWindow, box *GuiBox) {
	log.Println("gui.ShowMainTabShowBox() box =", box)
	// gw.UiTab.Delete(0)
	gw.BoxMap["MAIN3"] = box
	// gw.UiTab.InsertAt("Main", 0, box.UiBox)
	gw.UiTab.SetMargined(0, true)
}

func InitGuiBox(gw *GuiWindow, box *GuiBox, uiBox *ui.Box, name string) *GuiBox {
	log.Println("InitGuiBox() START")
	var newGuiBox GuiBox
	newGuiBox.UiBox = uiBox
	newGuiBox.Window = gw
	uiBox.SetPadded(true)

	if (box != nil) {
		log.Println("InitGuiBox() APPEND NEW BOX TO OLD BOX")
		box.UiBox.Append(uiBox, false)
	} else {
		log.Println("InitGuiBox() APPEND NEW BOX TO TAB")
		gw.UiTab.Append(name, uiBox)
	}
	gw.BoxMap[name] = &newGuiBox
	log.Println("InitGuiBox() END")
	return &newGuiBox
}

func InitGuiWindow(c *pb.Config, action string, maketab func(*GuiWindow) *GuiBox, uiW *ui.Window, uiT *ui.Tab) *GuiWindow {
	log.Println("InitGuiWindow() START")
	var newGuiWindow GuiWindow
	newGuiWindow.Width	= int(c.Width)
	newGuiWindow.Height	= int(c.Height)
	newGuiWindow.Action	= action
	newGuiWindow.MakeWindow	= maketab
	newGuiWindow.UiWindow	= uiW
	newGuiWindow.UiTab	= uiT
	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	newGuiWindow.EntryMap["test"] = nil
	Data.Windows = append(Data.Windows, &newGuiWindow)

	log.Println("InitGuiWindow() END *GuiWindow =", &newGuiWindow)
	return &newGuiWindow
}


func StartNewWindow(c *pb.Config, bg bool, action string, maketab func(*GuiWindow) *GuiBox) {
	log.Println("InitNewWindow() Create a new window")
	window := InitGuiWindow(c, action, maketab, nil, nil)
	/*
	newGuiWindow.Width	= int(c.Width)
	newGuiWindow.Height	= int(c.Height)
	newGuiWindow.Action	= action
	newGuiWindow.MakeWindow	= maketab
	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	newGuiWindow.EntryMap["test"] = nil
	Data.Windows = append(Data.Windows, &newGuiWindow)
	*/

	if (bg) {
		log.Println("ShowWindow() IN NEW GOROUTINE")
		go ui.Main(func() {
			InitTabWindow(window)
		})
		time.Sleep(2000 * time.Millisecond)
	} else {
		log.Println("ShowWindow() WAITING for ui.Main()")
		ui.Main(func() {
			InitTabWindow(window)
		})
	}
}

func getSplashText(a string) *ui.AttributedString {
	var aText  *ui.AttributedString
	aText = ui.NewAttributedString(a)
	return aText
}

func InitTabWindow(gw *GuiWindow) {
	log.Println("InitTabWindow() THIS WINDOW IS NOT YET SHOWN")

	gw.UiWindow = ui.NewWindow("", int(gw.Width), int(gw.Height), true)
	gw.UiWindow.SetBorderless(false)

        // create a 'fake' button entry for the mouse clicks
	var newBM GuiButton
	newBM.Action	= "QUIT"
	newBM.GW	= gw
	Data.AllButtons = append(Data.AllButtons, &newBM)

	gw.UiWindow.OnClosing(func(*ui.Window) bool {
		log.Println("InitTabWindow() OnClosing() THIS WINDOW IS CLOSING gw=", gw)
                ui.Quit()
		return true
	})

	gw.UiTab = ui.NewTab()
	gw.UiWindow.SetChild(gw.UiTab)
	gw.UiWindow.SetMargined(true)

	log.Println("InitTabWindow() gw =", gw)

	gw.MakeWindow(gw)

//	abox := gw.MakeWindow(gw)
//	gw.UiTab.Append("WIT Splash", abox.UiBox)
//	gw.UiTab.SetMargined(0, true)

	gw.UiWindow.Show()
}

func AddBoxToTab(name string, tab *ui.Tab, box *ui.Box) {
	tab.Append(name, box)
	tab.SetMargined(0, true)
}

/*
// string handling examples that might be helpful for normalizeInt()
isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

for _, username := range []string{"userone", "user2", "user-three"} {
    if !isAlpha(username) {
        fmt.Printf("%q is not valid\n", username)
    }
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

func alphaOnly(s string) bool {
   for _, char := range s {
      if !strings.Contains(alpha, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}
*/

func normalizeInt(s string) string {
	// reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Println("normalizeInt() regexp.Compile() ERROR =", err)
		return s
	}
	clean := reg.ReplaceAllString(s, "")
	log.Println("normalizeInt() s =", clean)
	return clean
}

func defaultEntryChange(e *ui.Entry) {
	for key, em := range Data.AllEntries {
		if (Data.Debug) {
			log.Println("\tdefaultEntryChange() Data.AllEntries =", key, em)
		}
		if Data.AllEntries[key].UiEntry == e {
			log.Println("defaultEntryChange() FOUND", 
				"action =", Data.AllEntries[key].Action,
				"Last =", Data.AllEntries[key].Last,
				"e.Text() =", e.Text())
			Data.AllEntries[key].Last = e.Text()
			if Data.AllEntries[key].Normalize != nil {
				fixed := Data.AllEntries[key].Normalize(e.Text())
				e.SetText(fixed)
			}
			return
		}
	}
	log.Println("defaultEntryChange() ERROR. MISSING ENTRY MAP. e.Text() =", e.Text())
}

func defaultMakeEntry(startValue string, edit bool, action string) *GuiEntry {
	e := ui.NewEntry()
	e.SetText(startValue)
	if (edit == false) {
		e.SetReadOnly(true)
	}
	e.OnChanged(defaultEntryChange)

	// add the entry field to the global map
	var newEntry GuiEntry
	newEntry.UiEntry  = e
	newEntry.Edit     = edit
	newEntry.Action   = action
	if (action == "Memory") {
		newEntry.Normalize = normalizeInt
	}
	Data.AllEntries = append(Data.AllEntries, &newEntry)

	return &newEntry
}
