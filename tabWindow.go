package gui

import "log"
import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

func buttonClick(b *ButtonMap, s string) {
	log.Println("gui.buttonClick() b, s =", b, s)
	log.Println("Figure out what to do here")
	log.Println("Figure out what to do here")
	log.Println("Figure out what to do here")

	if (Data.ButtonClick != nil) {
		log.Println("Data.ButtonClick() START")
		Data.ButtonClick(nil)
	}
}

func ShowAccountQuestionTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = AddAccountQuestionBox(nil, buttonClick)
	Data.cloudTab.InsertAt("New Account?", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowAccountTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = AddAccountBox(buttonClick)
	Data.cloudTab.InsertAt("Add Account", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowMainTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = makeCloudInfoBox(buttonClick)
	Data.cloudTab.InsertAt("Main", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func GoMainWindow() {
	ui.Main(makeCloudWindow)
}

func makeCloudWindow() {
	Data.cloudWindow = ui.NewWindow("", 640, 480, true)
	// cloudWindow.SetBorderless(true)
	Data.cloudWindow.OnClosing(func(*ui.Window) bool {
		if (Data.ButtonClick != nil) {
			log.Println("Data.ButtonClick() START QUIT")
			Data.State = "QUIT"
			Data.ButtonClick(nil)
		}
		// ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		if (Data.ButtonClick != nil) {
			log.Println("Data.ButtonClick() START QUIT")
			Data.State = "QUIT"
			Data.ButtonClick(nil)
		}
		// Data.cloudWindow.Destroy()
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
