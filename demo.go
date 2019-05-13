package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var demowin *ui.Window
var demotab *ui.Tab

func setupDemoUI() {
	log.Println("setupDemoUI() START")
	demowin = ui.NewWindow("Demo GUI Widgets", 500, 300, false)
	demowin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		demowin.Destroy()
		return true
	})

	demotab = ui.NewTab()
	demowin.SetChild(demotab)
	demowin.SetMargined(true)

	demotab.Append("List examples", makeNumbersPage())
	tabcount := 0
	demotab.SetMargined(tabcount, true)

	demotab.Append("Choosers examples", makeDataChoosersPage())
	tabcount += 1
	demotab.SetMargined(tabcount, true)

	demotab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	demotab.SetMargined(tabcount, true)

	demowin.Show()
}
