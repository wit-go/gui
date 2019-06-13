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
			log.Println("gui.StartNewWindow() inside ui.Main()")
			go runWindow(name, axis, callback)
		})
		time.Sleep(2000 * time.Millisecond) // this might make it more stable on windows?
	} else {
		log.Println("StartNewWindow() WAITING for ui.Main()")
		ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")
			runWindow(name, axis, callback)
		})
	}
}

// This creates the raw andlabs/ui Window
func runWindow(name string, axis int, callback func(*GuiBox) *GuiBox) {
	log.Println("initTabWindow() START. THIS WINDOW IS NOT YET SHOWN")
	log.Println("initTabWindow() START. name =", name)

	box := InitWindow(nil, name, axis)
	box = callback(box)
	window := box.Window
	log.Println("initTabWindow() END box =", box)
	log.Println("initTabWindow() END gw =", window)
	window.UiWindow.Show()
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
	newGuiWindow.Height	= 600
	newGuiWindow.Width	= 800

	// This is the first window. One must create it here
	if (gw == nil) {
		log.Println("initWindow() ADDING ui.NewWindow()")
		newGuiWindow.UiWindow = ui.NewWindow(newGuiWindow.Name, int(newGuiWindow.Width), int(newGuiWindow.Height), true)
		newGuiWindow.UiWindow.SetBorderless(false)

		newGuiWindow.UiWindow.OnClosing(func(*ui.Window) bool {
			log.Println("initTabWindow() OnClosing() THIS WINDOW IS CLOSING newGuiWindow=", newGuiWindow)
			ui.Quit()
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

	newGuiWindow.Axis	= axis
	newGuiWindow.Name       = name

	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	// newGuiWindow.EntryMap["test"] = nil
	Data.Windows = append(Data.Windows, &newGuiWindow)

	if (newGuiWindow.UiTab == nil) {
		tabnum			:= 0
		newGuiWindow.TabNumber	= &tabnum
	} else {
		tabnum			:= newGuiWindow.UiTab.NumPages()
		newGuiWindow.TabNumber	= &tabnum
	}

	Data.WindowMap[newGuiWindow.Name]    = &newGuiWindow

//	if (Data.buttonMap == nil) {
//		GuiInit()
//	}
	log.Println("InitGuiWindow() END *GuiWindow =", &newGuiWindow)

	var box *GuiBox
	if (axis == Xaxis) {
		box = HardBox(&newGuiWindow, Xaxis, name)
	} else {
		box = HardBox(&newGuiWindow, Yaxis, name)
	}
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
	} else {
		log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
	}
}
