package gui

import "log"
import "time"
// import "regexp"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func initUI(name string, callback func(*GuiBox) *GuiBox) {
	ui.Main(func() {
		log.Println("gui.initUI() inside ui.Main()")

		box := InitWindow(nil, "StartNewWindow" + name, 0)
		box = callback(box)
		window := box.Window
		log.Println("StartNewWindow() box =", box)

		window.UiWindow.Show()
	})
}

func StartNewWindow(bg bool, name string, axis int, callback func(*GuiBox) *GuiBox) {
	log.Println("StartNewWindow() ui.Main() Create a new window")

	if (bg) {
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
func InitWindow(gw *GuiWindow, name string, axis int) *GuiBox {
	log.Println("InitGuiWindow() START")

	var box *GuiBox
	if (gw == nil) {
		box = mapWindow(nil, name, Config.Height, Config.Width)
	} else {
		box = mapWindow(gw.UiWindow, name, Config.Height, Config.Width)
	}

	// box.Window = &newGuiWindow
	newGuiWindow := box.Window

	// This is the first window. One must create it here
	if (gw == nil) {
		log.Println("initWindow() ADDING ui.NewWindow()")
		newGuiWindow.UiWindow = ui.NewWindow(name, int(newGuiWindow.Width), int(newGuiWindow.Height), true)
		newGuiWindow.UiWindow.SetBorderless(false)

		// newGuiWindow.UiWindow.SetTitle("test")
		newGuiWindow.UiWindow.OnClosing(func(*ui.Window) bool {
			log.Println("initTabWindow() OnClosing() THIS WINDOW IS CLOSING newGuiWindow=", newGuiWindow)
			// newGuiWindow.UiWindow.Destroy()
			if (Config.Exit == nil) {
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

	Data.WindowMap[newGuiWindow.Name]    = newGuiWindow

	log.Println("InitGuiWindow() END *GuiWindow =", newGuiWindow)
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

func CreateWindow(title string, tabname string, x int, y int, custom func() ui.Control) *GuiBox {
	box := CreateBlankWindow(title, x, y)
	box.InitTab(title)
	return box
}

func CreateBlankWindow(title string, x int, y int) *GuiBox {
	window := ui.NewWindow(title, x, y, false)
	window.SetBorderless(false)
	window.OnClosing(func(*ui.Window) bool {
		log.Println("createWindow().OnClosing()", title)
		return true
	})
	ui.OnShouldQuit(func() bool {
		log.Println("createWindow().Destroy()", title)
		window.Destroy()
		return true
	})

	window.SetMargined(true)
	window.Show()

	return mapWindow(window, title, x, y)
}

func initBlankWindow() ui.Control {
        hbox := ui.NewHorizontalBox()
        hbox.SetPadded(true)

	return hbox
}

var master = 0

func mapWindow(window *ui.Window, title string, x int, y int) *GuiBox {
	if (Data.WindowMap[title] != nil) {
		log.Println("Data.WindowMap[title] already exists title =", title)
		title = title + string(master)
	}
	if (Data.WindowMap[title] != nil) {
		log.Println("Data.WindowMap[title] already exists title =", title)
		panic("Data.WindowMap[newGuiWindow.Name] already exists")
		return nil
	}
	var newGuiWindow GuiWindow
	newGuiWindow.Width	= x
	newGuiWindow.Height	= y
	newGuiWindow.Name       = title
	newGuiWindow.UiWindow   = window

	newGuiWindow.BoxMap     = make(map[string]*GuiBox)
        newGuiWindow.EntryMap   = make(map[string]*GuiEntry)

	Data.WindowMap[newGuiWindow.Name]    = &newGuiWindow

	var box GuiBox
	box.Window = &newGuiWindow

	return &box
}
