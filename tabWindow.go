package gui

import "log"
import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

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
	Data.cloudTab.Delete(0)

	log.Println("Sleep(1000)")
	time.Sleep(1000 * time.Millisecond)

	Data.smallBox = AddAccountQuestionBox(nil, buttonClick)
	Data.cloudTab.InsertAt("New Account?", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowAccountTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(1000)")
	time.Sleep(1000 * time.Millisecond)

	Data.smallBox = AddAccountBox(buttonClick)
	Data.cloudTab.InsertAt("Add Account", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
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

	Data.cloudTab = ui.NewTab()
	Data.cloudWindow.SetChild(Data.cloudTab)
	Data.cloudWindow.SetMargined(true)

	Data.cloudBox = ShowSplashBox(nil, nil, buttonClick)

	Data.cloudTab.Append("WIT Splash", Data.cloudBox)
	Data.cloudTab.SetMargined(0, true)

	Data.cloudWindow.Show()
	Data.State = "splash"
}
