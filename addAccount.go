package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

func AddAccountQuestionBox() *ui.Box {
	newbox := ui.NewVerticalBox()
	newbox.SetPadded(true)

	newButton := CreateButton(nil, nil, "Create New Account", "AREA", nil)
	newbox.Append(newButton, false)

	newbox.Append(ui.NewHorizontalSeparator(), false)

	okButton := CreateButton(nil, nil, "I Have an Account", "AREA", nil)
	newbox.Append(okButton, false)

	return newbox
}

// func AddAccountBox(aTab *GuiTabStructure) {
func AddAccountBox() *ui.Box {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
//	aTab.firstBox = vbox

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

	Data.EntryNick = ui.NewEntry()
	Data.EntryNick.SetReadOnly(false)

	vboxN.Append(Data.EntryNick, false)

	Data.EntryNick.OnChanged(func(*ui.Entry) {
		log.Println("OK. nickname =", Data.EntryNick.Text())
		// Data.AccNick = entryNick.Text()
	})
	hboxAccount.Append(vboxN, false)
	// End 'Nickname' vertical box

	// Start 'Username' vertical box
	vboxU := ui.NewVerticalBox()
	vboxU.SetPadded(true)
	vboxU.Append(ui.NewLabel("Account Username:"), false)

	Data.EntryUser = ui.NewEntry()
	Data.EntryUser.SetReadOnly(false)

	vboxU.Append(Data.EntryUser, false)

	Data.EntryUser.OnChanged(func(*ui.Entry) {
		log.Println("OK. username =", Data.EntryUser.Text())
		// Data.AccUser = entryUser.Text()
	})
	hboxAccount.Append(vboxU, false)
	// End 'Username' vertical box

	// Start 'Password' vertical box
	vboxP := ui.NewVerticalBox()
	vboxP.SetPadded(true)
	vboxP.Append(ui.NewLabel("Account Password:"), false)

	Data.EntryPass = ui.NewEntry()
	Data.EntryPass.SetReadOnly(false)

	vboxP.Append(Data.EntryPass, false)

	Data.EntryPass.OnChanged(func(*ui.Entry) {
		log.Println("OK. password =", Data.EntryPass.Text())
		// Data.AccPass = entryPass.Text()
	})
	hboxAccount.Append(vboxP, false)
	// End 'Password' vertical box

	vbox.Append(ui.NewHorizontalSeparator(), false)

	hboxButtons := ui.NewHorizontalBox()
	hboxButtons.SetPadded(true)
	vbox.Append(hboxButtons, false)

	okButton := CreateButton(nil, nil, "Add Account", "ADD", nil)
	hboxButtons.Append(okButton, false)

	backButton := CreateButton(nil, nil, "Back", "BACK", nil)
	hboxButtons.Append(backButton, false)

	return vbox
}
