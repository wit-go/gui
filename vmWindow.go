package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func ShowVM() {
	name := Data.CurrentVM
	log.Println("setupDemoUI() START Data.CurrentVM=", Data.CurrentVM)
	VMwin := ui.NewWindow("VM " + name, 500, 300, false)
	VMwin.OnClosing(func(*ui.Window) bool {
		return true
	})
	ui.OnShouldQuit(func() bool {
		VMwin.Destroy()
		VMwin = nil
		return true
	})

	VMtab := ui.NewTab()
	VMwin.SetChild(VMtab)
	VMwin.SetMargined(true)

//	VMtab.Append("List examples", makeNumbersPage())
//	VMtab.SetMargined(0, true)

	VMwin.Show()
}
