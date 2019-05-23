package gui

// import "github.com/davecgh/go-spew/spew"
// import "time"
// import "fmt"

// import "log"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// var splashWin *ui.Window
var MyArea	*ui.Area

func ShowSplash() *ui.Window {
	splashWin := ui.NewWindow("", 640, 480, true)
	splashWin.SetBorderless(true)
	splashWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		splashWin.Destroy()
		return true
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	splashWin.SetChild(vbox)
	splashWin.SetMargined(true)

	// This displays the window
	// splashWin.Show()

	ShowSplashBox(vbox, nil, nil)

	return splashWin
}

func ShowSplashBox(vbox *ui.Box, atest chan int, custom func(int, string)) *ui.Box {
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	makeAttributedString()
	MyArea = makeSplashArea(custom)

	newbox.Append(MyArea, true)

	if runtime.GOOS == "linux" {
		newbox.Append(ui.NewLabel("OS: Linux"), false)
	} else if runtime.GOOS == "windows" {
		newbox.Append(ui.NewLabel("OS: Windows"), false)
	} else {
		newbox.Append(ui.NewLabel("OS: " + runtime.GOOS), false)
	}

	newbox.Append(ui.NewLabel("Version: v0.3"), false)
	okButton := CreateButton("OK", "CLOSE", custom)
	newbox.Append(okButton, false)

	if (vbox != nil) {
		vbox.Append(newbox, true)
	}

	return newbox
}
