package gui

// import "github.com/davecgh/go-spew/spew"
// import "time"
// import "fmt"

import "os"
import "log"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func ShowSplashBox(wm *GuiWindow, newText *ui.AttributedString) *ui.Box {
	log.Println("ShowSplashBox() START")
	log.Println("ShowSplashBox() START wm =", wm)
	if (wm == nil) {
		log.Println("ShowSplashBox() WE ARE FUCKED BECAUSE WE DON'T KNOW WHAT WINDOW TO DO THIS IN")
		os.Exit(0)
		return nil
	}
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	// initialize the AreaHandler{}
	wm.AH    = new(AreaHandler)
	wm.AH.WM = wm
	wm.AH.Attrstr = newText
	makeSplashArea(wm, wm.AH)

	newbox.Append(wm.AH.Area, true)

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

	log.Println("ShowSplashBox() START wm =", wm)
	okButton := CreateButton(wm, nil, nil, "OK", "AREA", nil)
	newbox.Append(okButton, false)
	newbox.Append(CreateButton(wm, nil, nil, "NEWTEXT", "NEWTEXT", nil), false)

	// os.Exit(0)
	return newbox
}
