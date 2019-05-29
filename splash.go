package gui

// import "github.com/davecgh/go-spew/spew"
// import "time"
// import "fmt"

// import "log"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func ShowSplashBox(newText *ui.AttributedString) *ui.Box {
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	// initialize the AreaHandler{}
	Data.Window1.AH    = new(AreaHandler)
	Data.Window1.AH.Attrstr = newText
	makeSplashArea(Data.Window1.AH)

	newbox.Append(Data.Window1.AH.Area, true)

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

	okButton := CreateButton(nil, nil, "OK", "AREA", nil)
	newbox.Append(okButton, false)
	newbox.Append(CreateButton(nil, nil, "NEWTEXT", "NEWTEXT", nil), false)

	return newbox
}
