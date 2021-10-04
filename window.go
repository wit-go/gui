package gui

import "log"
import "time"
// import "regexp"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func StartNewWindow(bg bool, name string, axis int, callback func(*GuiBox) *GuiBox) {
	log.Println("StartNewWindow() Create a new window")

	if (bg) {
		log.Println("StartNewWindow() START NEW GOROUTINE for ui.Main()")
		go ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main() in NEW goroutine")

			// InitWindow must be called from within ui.Main()
			box := InitWindow(nil, name, axis)
			box = callback(box)
			window := box.Window
			log.Println("StartNewWindow() box =", box)

			runWindow(window.UiWindow)
		})
		time.Sleep(500 * time.Millisecond) // this might make it more stable on windows?
	} else {
		log.Println("StartNewWindow() WAITING for ui.Main()")
		ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")

			// InitWindow must be called from within ui.Main()
			box := InitWindow(nil, name, axis)
			box = callback(box)
			window := box.Window
			log.Println("StartNewWindow() box =", box)

			runWindow(window.UiWindow)
		})
	}
}

// This creates the raw andlabs/ui Window
func runWindow(uiWindow *ui.Window) {
	log.Println("runWindow() START ui.Window.Show()")
	uiWindow.Show()
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
func InitWindow(gw *GuiWindow, name string, axis int) *GuiBox {
	window := Data.WindowMap[name]
	if (window != nil) {
		box := window.BoxMap["MAINBOX"]
		log.Println("gui.InitWindow() tab already exists name =", name)
		ErrorWindow(box.Window, "Create Window Error", "Window " + name + " already exists")
		return nil
	}

	log.Println("InitGuiWindow() START")
	var newGuiWindow GuiWindow
	newGuiWindow.Height	= Config.Height
	newGuiWindow.Width	= Config.Width
	newGuiWindow.Axis	= axis
	newGuiWindow.Name       = name

	// This is the first window. One must create it here
	if (gw == nil) {
		log.Println("initWindow() ADDING ui.NewWindow()")
		newGuiWindow.UiWindow = ui.NewWindow(name, int(newGuiWindow.Width), int(newGuiWindow.Height), true)
		newGuiWindow.UiWindow.SetBorderless(false)

		newGuiWindow.UiWindow.OnClosing(func(*ui.Window) bool {
			log.Println("initTabWindow() OnClosing() THIS WINDOW IS CLOSING newGuiWindow=", newGuiWindow)
			// newGuiWindow.UiWindow.Destroy()
			if (Config.Exit == nil) {
				ui.Quit()
			} else {
				// allow a custom exit function
				Config.Exit(&newGuiWindow)
			}
			return true
		})

		newGuiWindow.UiTab = ui.NewTab()
		newGuiWindow.UiWindow.SetChild(newGuiWindow.UiTab)
		newGuiWindow.UiWindow.SetMargined(true)
		tmp := 0
		newGuiWindow.TabNumber = &tmp
	} else {
		newGuiWindow.UiWindow	= gw.UiWindow
		newGuiWindow.UiTab	= gw.UiTab
	}


	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	// Data.Windows = append(Data.Windows, &newGuiWindow)

	if (newGuiWindow.UiTab == nil) {
		tabnum			:= 0
		newGuiWindow.TabNumber	= &tabnum
	} else {
		tabnum			:= newGuiWindow.UiTab.NumPages()
		newGuiWindow.TabNumber	= &tabnum
	}

	Data.WindowMap[newGuiWindow.Name]    = &newGuiWindow

	var box *GuiBox
	if (axis == Xaxis) {
		box = HardBox(&newGuiWindow, Xaxis, name)
	} else {
		box = HardBox(&newGuiWindow, Yaxis, name)
	}
	log.Println("InitGuiWindow() END *GuiWindow =", &newGuiWindow)
	return box
}

func DeleteWindow(name string) {
	log.Println("gui.DeleteWindow() START name =", name)
	window := Data.WindowMap[name]
	if (window == nil) {
		log.Println("gui.DeleteWindow() NO WINDOW WITH name =", name)
		return
	}

	log.Println("gui.DumpBoxes() MAP: ", name)
	log.Println("gui.DumpBoxes()\tWindow.name =", window.Name)
	if (window.TabNumber == nil) {
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
		if (window.TabNumber == nil) {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
		} else {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
			if (tab < *window.TabNumber) {
				log.Println("gui.DumpBoxes() \tSubtracting 1 from TabNumber")
				*window.TabNumber -= 1
				log.Println("gui.DumpBoxes() \tWindows.TabNumber is now =", *window.TabNumber)
			}
		}
	}
}

// CreateWindow("my title", "my tabname", 300, 200, makeNumbersPagewin2)
func CreateWindow(title string, tabname string, x int, y int, custom func() ui.Control) *ui.Window {
	window := ui.NewWindow(title, x, y, false)
	window.OnClosing(func(*ui.Window) bool {
		log.Println("createWindow().OnClosing()", title)
		return true
	})
	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy()", title)
		window.Destroy()
		return true
	})

	tab := ui.NewTab()
	window.SetChild(tab)
	window.SetMargined(true)

	tab.Append(tabname, custom())
	tab.SetMargined(0, true)

	window.Show()

	return window
}
