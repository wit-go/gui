package gui

// import "github.com/davecgh/go-spew/spew"
// import "time"
// import "fmt"

// import "log"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func ShowSplashBox(vbox *ui.Box, atest chan int, custom func(*ButtonMap)) *ui.Box {
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

	version := "Version: " + Data.Version
	newbox.Append(ui.NewLabel(version), false)

	if (Data.Debug) {
		if (Data.GitCommit != "") {
			tmp := "git rev-list: " + Data.GitCommit
			newbox.Append(ui.NewLabel(tmp), false)
		}
		if (Data.GoVersion != "") {
			tmp := "go build version: " + Data.GoVersion
			newbox.Append(ui.NewLabel(tmp), false)
		}
		if (Data.Buildtime != "") {
			tmp := "build date: " + Data.Buildtime
			newbox.Append(ui.NewLabel(tmp), false)
		}
	}

	okButton := CreateButton("OK", "DONE", custom)
	newbox.Append(okButton, false)

	if (vbox != nil) {
		vbox.Append(newbox, true)
	}

	return newbox
}
