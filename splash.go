package gui

// import "github.com/davecgh/go-spew/spew"
// import "time"
// import "fmt"

import "os"
import "log"
import "runtime"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func ShowSplashBox(gw *GuiWindow, newText *ui.AttributedString) *GuiBox {
	log.Println("ShowSplashBox() START")
	log.Println("ShowSplashBox() START gw =", gw)
	if (gw == nil) {
		log.Println("ShowSplashBox() WE ARE FUCKED BECAUSE WE DON'T KNOW WHAT WINDOW TO DO THIS IN")
		os.Exit(0)
		return nil
	}
	var gb *GuiBox
	gb = new(GuiBox)

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)
	// gw.Box1 = hbox
	gb.UiBox = newbox
	gb.W = gw
	gw.BoxMap["Splash"] = gb

	/*
	// initialize the GuiArea{}
	gb.Area		= new(GuiArea)
	gb.Area.Window	= gw
	gb.Area.UiAttrstr = newText
	*/

	makeSplashArea(gb, newText)

	newbox.Append(gb.Area.UiArea, true)

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

	log.Println("ShowSplashBox() START gb =", gb)

	okButton := CreateButton(gb, nil, nil, "OK", "AREA", nil)
	newbox.Append(okButton.B, false)

	// os.Exit(0)
	return gb
}
