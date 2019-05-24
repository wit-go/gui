package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func AddAccountQuestionBox(junk *ui.Box, custom func(*ButtonMap, string)) *ui.Box {
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	newButton := CreateButton("Create New Account", "DONE", custom)
	newbox.Append(newButton, false)

	newbox.Append(ui.NewHorizontalSeparator(), false)

	okButton := CreateButton("I Have an Account", "DONE", custom)
	newbox.Append(okButton, false)

	return newbox
}

func AddAccountBox(custom func(*ButtonMap, string)) *ui.Box {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hboxAccount := ui.NewHorizontalBox()
	hboxAccount.SetPadded(true)
	vbox.Append(hboxAccount, false)

	// Start 'Provider' vertical box
	vboxC := ui.NewVerticalBox()
	vboxC.SetPadded(true)
	vboxC.Append(ui.NewLabel("Cloud Provider:"), false)

	cbox := ui.NewCombobox()
	cbox.Append("WIT")
	cbox.Append("Evocative")
	vboxC.Append(cbox, false)
	cbox.SetSelected(0)

	cbox.OnSelected(func(*ui.Combobox) {
		log.Println("OK. Selected Cloud Provider =", cbox.Selected())
	})
	hboxAccount.Append(vboxC, false)
	// End 'Cloud Provider' vertical box

	// Start 'Region' vertical box
	vboxR := ui.NewVerticalBox()
	vboxR.SetPadded(true)
	vboxR.Append(ui.NewLabel("Region:"), false)

	regbox := ui.NewCombobox()
	regbox.Append("Any")
	regbox.Append("SF")
	vboxR.Append(regbox, false)
	regbox.SetSelected(0)

	regbox.OnSelected(func(*ui.Combobox) {
		log.Println("OK. Selected something =", regbox.Selected())
	})
	hboxAccount.Append(vboxR, false)
	// End 'Region' vertical box

	// Start 'Nickname' vertical box
	vboxN := ui.NewVerticalBox()
	vboxN.SetPadded(true)
	vboxN.Append(ui.NewLabel("Account Nickname:"), false)

	entryNick := ui.NewEntry()
	entryNick.SetReadOnly(false)

	vboxN.Append(entryNick, false)

	entryNick.OnChanged(func(*ui.Entry) {
		log.Println("OK. nickname =", entryNick.Text())
		Data.AccNick = entryNick.Text()
	})
	hboxAccount.Append(vboxN, false)
	// End 'Nickname' vertical box

	// Start 'Username' vertical box
	vboxU := ui.NewVerticalBox()
	vboxU.SetPadded(true)
	vboxU.Append(ui.NewLabel("Account Username:"), false)

	entryUser := ui.NewEntry()
	entryUser.SetReadOnly(false)

	vboxU.Append(entryUser, false)

	entryUser.OnChanged(func(*ui.Entry) {
		log.Println("OK. username =", entryUser.Text())
		Data.AccUser = entryUser.Text()
	})
	hboxAccount.Append(vboxU, false)
	// End 'Username' vertical box

	// Start 'Password' vertical box
	vboxP := ui.NewVerticalBox()
	vboxP.SetPadded(true)
	vboxP.Append(ui.NewLabel("Account Password:"), false)

	entryPass := ui.NewEntry()
	entryPass.SetReadOnly(false)

	vboxP.Append(entryPass, false)

	entryPass.OnChanged(func(*ui.Entry) {
		log.Println("OK. password =", entryPass.Text())
		Data.AccPass = entryPass.Text()
	})
	hboxAccount.Append(vboxP, false)
	// End 'Password' vertical box

	vbox.Append(ui.NewHorizontalSeparator(), false)

	hboxButtons := ui.NewHorizontalBox()
	hboxButtons.SetPadded(true)
	vbox.Append(hboxButtons, false)

	okButton := CreateButton("Add Account", "ADD", custom)
	hboxButtons.Append(okButton, false)

	backButton := CreateButton("Back", "BACK", custom)
	hboxButtons.Append(backButton, false)

	return vbox
}
