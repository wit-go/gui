package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"

func AddAccountWindow() {
	accounthWin := ui.NewWindow("Add Account", 400, 300, false)
	accounthWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		accounthWin.Destroy()
		return true
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	accounthWin.SetChild(vbox)
	accounthWin.SetMargined(true)

	// This displays the window
	accounthWin.Show()

	// START create new account button
	newAccountButton := ui.NewButton("Create New Account")
		newAccountButton.OnClicked(func(*ui.Button) {
		log.Println("OK. Closing window.")
		accounthWin.Destroy()
		ui.Quit()
	})
	vbox.Append(newAccountButton, false)
	// END create new account button

	vbox.Append(ui.NewHorizontalSeparator(), false)

	okButton := ui.NewButton("I Have an Account")
	okButton.OnClicked(func(*ui.Button) {
		log.Println("OK. Closing window.")
		accounthWin.Destroy()
		ui.Quit()
	})
	vbox.Append(okButton, false)
	// END add account hbox
}

func AddAccountBox(junk *ui.Box, custom func(int, string)) *ui.Box {
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	// create new account button
	newButton := CreateButton("Create New Account", "CLOSE", custom)
	newbox.Append(newButton, false)

	newbox.Append(ui.NewHorizontalSeparator(), false)

	okButton := CreateButton("I Have an Account", "CLOSE", custom)
	newbox.Append(okButton, false)

	return newbox
}
