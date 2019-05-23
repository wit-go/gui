package gui

// import "github.com/davecgh/go-spew/spew"
// import "time"
// import "fmt"

// import "log"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func ShowSplashBox(vbox *ui.Box, atest chan int, custom func(int, string)) *ui.Box {
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	makeAttributedString()
	Data.MyArea = makeSplashArea(custom)

	newbox.Append(Data.MyArea, true)

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
