package gui

import "log"
import "time"
// import "regexp"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func InitGuiWindow(gw *GuiWindow) *GuiWindow {
	log.Println("InitGuiWindow() START")
	var newGuiWindow GuiWindow
	newGuiWindow.Width	= Config.Width
	newGuiWindow.Height	= Config.Height

	newGuiWindow.Axis	= gw.Axis
	newGuiWindow.MakeWindow	= gw.MakeWindow
	newGuiWindow.UiWindow	= gw.UiWindow
	newGuiWindow.UiTab	= gw.UiTab
	newGuiWindow.Name       = gw.Name

	newGuiWindow.BoxMap	= make(map[string]*GuiBox)
	newGuiWindow.EntryMap	= make(map[string]*GuiEntry)
	newGuiWindow.EntryMap["test"] = nil
	Data.Windows = append(Data.Windows, &newGuiWindow)

	if (gw.UiTab == nil) {
		tabnum			:= 0
		newGuiWindow.TabNumber	= &tabnum
	} else {
		tabnum			:= gw.UiTab.NumPages()
		newGuiWindow.TabNumber	= &tabnum
	}

	Data.WindowMap[newGuiWindow.Name]    = &newGuiWindow

	if (Data.buttonMap == nil) {
		GuiInit()
	}
	log.Println("InitGuiWindow() END *GuiWindow =", &newGuiWindow)
	return &newGuiWindow
}


func StartNewWindow(bg bool, name string, axis int, callback func(*GuiWindow) *GuiBox) {
	log.Println("StartNewWindow() Create a new window")
	var junk GuiWindow
	junk.MakeWindow = callback
	junk.Name = name
	junk.Axis = axis
	window := InitGuiWindow(&junk)
	if (bg) {
		log.Println("StartNewWindow() START NEW GOROUTINE for ui.Main()")
		go ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")
			go initTabWindow(window)
		})
		time.Sleep(2000 * time.Millisecond) // this might make it more stable on windows?
	} else {
		log.Println("StartNewWindow() WAITING for ui.Main()")
		ui.Main(func() {
			log.Println("gui.StartNewWindow() inside ui.Main()")
			initTabWindow(window)
		})
	}
}

// This creates the raw andlabs/ui Window
func initTabWindow(gw *GuiWindow) {
	log.Println("initTabWindow() START. THIS WINDOW IS NOT YET SHOWN")
	log.Println("initTabWindow() START. name =", gw.Name)

	gw.UiWindow = ui.NewWindow(gw.Name, int(gw.Width), int(gw.Height), true)
	gw.UiWindow.SetBorderless(false)

	gw.UiWindow.OnClosing(func(*ui.Window) bool {
		log.Println("initTabWindow() OnClosing() THIS WINDOW IS CLOSING gw=", gw)
                ui.Quit()
		return true
	})

	gw.UiTab = ui.NewTab()
	gw.UiWindow.SetChild(gw.UiTab)
	gw.UiWindow.SetMargined(true)
	tmp := 0
	gw.TabNumber = &tmp

	DumpBoxes()
	// for {}

	box := gw.MakeWindow(gw)
	log.Println("initTabWindow() END box =", box)
	log.Println("initTabWindow() END gw =", gw)
	gw.UiWindow.Show()
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

	// if there is not an account, then go to 'make account'
	gw.Name = name
	var junk GuiWindow
	// junk.MakeWindow	= callback
	junk.Name	= name
	junk.Axis	= axis
	junk.UiWindow	= gw.UiWindow
	junk.UiTab	= gw.UiTab
	newWindow	:= InitGuiWindow(gw)

	var box *GuiBox
	if (axis == Xaxis) {
		box = HardBox(newWindow, Xaxis, name)
	} else {
		box = HardBox(newWindow, Yaxis, name)
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
}
