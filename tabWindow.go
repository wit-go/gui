package gui

import "log"
import "time"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

var cloudWindow *ui.Window
var cloudTab *ui.Tab
var cloudBox *ui.Box
var smallBox *ui.Box
var state string

func splashClose(a int, b string) {
	log.Println("GOT splashClose(a,b) =", a, b)

	log.Println("cloudBox Delete(0) START")
	cloudBox.Delete(0)
	log.Println("smallBox.Hide() START")
	smallBox.Hide()

	state = "kill"
}

func buttonClick(i int, s string) {
	log.Println("test2 buttonClick() i, s =", i, s)
	cloudTab.Delete(0)

	log.Println("Sleep(2000)")
	time.Sleep(2000 * time.Millisecond)

	smallBox = AddAccountBox(nil, splashClose)
	cloudTab.InsertAt("Intro", 0, smallBox)
	cloudTab.SetMargined(0, true)
}

func ShowAccountTab() {
	cloudTab.Delete(0)

	log.Println("Sleep(1000)")
	time.Sleep(1000 * time.Millisecond)

	smallBox = AddAccountBox(nil, splashClose)
	cloudTab.InsertAt("Intro", 0, smallBox)
	cloudTab.SetMargined(0, true)
}

func GoMainWindow() {
	ui.Main(makeCloudWindow)
}

func makeCloudWindow() {
	cloudWindow := ui.NewWindow("", 640, 480, true)
	// cloudWindow.SetBorderless(true)
	cloudWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		cloudWindow.Destroy()
		return true
	})

//	cloudBox = ui.NewVerticalBox()
//	cloudBox.SetPadded(true)
//	cloudWindow.SetChild(cloudBox)
//	cloudWindow.SetMargined(true)

	cloudTab = ui.NewTab()
	cloudWindow.SetChild(cloudTab)
	cloudWindow.SetMargined(true)

	cloudBox = ShowSplashBox(nil, nil, buttonClick)

	cloudTab.Append("WIT Splash", cloudBox)
	cloudTab.SetMargined(0, true)

	cloudWindow.Show()
	// state = "done"
}
