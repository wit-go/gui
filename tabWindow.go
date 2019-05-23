package gui

import "log"
import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

// var cloudWindow *ui.Window
var cloudTab *ui.Tab
var cloudBox *ui.Box
var smallBox *ui.Box
var state string

func buttonClick(i int, s string) {
	log.Println("gui.buttonClick() i, s =", i, s)
	log.Println("Figure out what to do here")
	log.Println("Figure out what to do here")
	log.Println("Figure out what to do here")

	if (Data.ButtonClick != nil) {
		log.Println("Data.ButtonClick() START")
		Data.ButtonClick(i, s)
	}
}

func ShowAccountQuestionTab() {
	cloudTab.Delete(0)

	log.Println("Sleep(1000)")
	time.Sleep(1000 * time.Millisecond)

	smallBox = AddAccountQuestionBox(nil, buttonClick)
	cloudTab.InsertAt("New Account?", 0, smallBox)
	cloudTab.SetMargined(0, true)
}

func ShowAccountTab() {
	cloudTab.Delete(0)

	log.Println("Sleep(1000)")
	time.Sleep(1000 * time.Millisecond)

	smallBox = AddAccountBox(buttonClick)
	cloudTab.InsertAt("Add Account", 0, smallBox)
	cloudTab.SetMargined(0, true)
}

func GoMainWindow() {
	ui.Main(makeCloudWindow)
}

func makeCloudWindow() {
	Data.cloudWindow = ui.NewWindow("", 640, 480, true)
	// cloudWindow.SetBorderless(true)
	Data.cloudWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		Data.cloudWindow.Destroy()
		return true
	})

	cloudTab = ui.NewTab()
	Data.cloudWindow.SetChild(cloudTab)
	Data.cloudWindow.SetMargined(true)

	cloudBox = ShowSplashBox(nil, nil, buttonClick)

	cloudTab.Append("WIT Splash", cloudBox)
	cloudTab.SetMargined(0, true)

	Data.cloudWindow.Show()
	Data.State = "splash"
}
