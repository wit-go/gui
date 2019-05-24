package gui

import "log"
import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

func ShowAccountQuestionTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = AddAccountQuestionBox(nil, mouseClick)
	Data.cloudTab.InsertAt("New Account?", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowAccountTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = AddAccountBox(mouseClick)
	Data.cloudTab.InsertAt("Add Account", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowMainTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = makeCloudInfoBox(mouseClick)
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
		if (Data.MouseClick != nil) {
			log.Println("SIMULATE Data.MouseClick(QUIT)")
			Data.State = "QUIT"
			Data.MouseClick(nil)
		}
		return true
	})
	ui.OnShouldQuit(func() bool {
		if (Data.MouseClick != nil) {
			log.Println("SIMULATE Data.MouseClick(QUIT)")
			Data.State = "QUIT"
			Data.MouseClick(nil)
		}
		return true
	})

	Data.cloudTab = ui.NewTab()
	Data.cloudWindow.SetChild(Data.cloudTab)
	Data.cloudWindow.SetMargined(true)

	Data.cloudBox = ShowSplashBox(nil, nil, mouseClick)

	Data.cloudTab.Append("WIT Splash", Data.cloudBox)
	Data.cloudTab.SetMargined(0, true)

	Data.cloudWindow.Show()
	Data.State = "splash"
}
