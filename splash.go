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

	return newbox
}

func makeAttributedString() *ui.AttributedString {
	newText := ui.NewAttributedString("")

	appendWithAttributes(newText, "Welcome to the Cloud Control Panel\n", ui.TextSize(16), ui.TextColor{0.0, 0.0, 0.8, .8}) // "RGBT"

	appendWithAttributes(newText, "(alpha)\n\n", ui.TextSize(10))

	appendWithAttributes(newText, "This control panel was designed to be an interface to your 'private' cloud. ", ui.TextWeightBold)
	appendWithAttributes(newText, "The concept of a private cloud means that you can use a providers system, or, seemlessly, use your own hardware in your own datacenter. ", ui.TextWeightBold)

	newText.AppendUnattributed("\n")
	newText.AppendUnattributed("\n")
	appendWithAttributes(newText, "This control panel requires:\n")
	newText.AppendUnattributed("\n")
	appendWithAttributes(newText, "IPv6\n")
	appendWithAttributes(newText, "newText, Your hostname in DNS\n")
	newText.AppendUnattributed("\n\n\n\n\n")

	appendWithAttributes(newText, "<click or press any key>\n", ui.TextSize(10))

	return newText
}
