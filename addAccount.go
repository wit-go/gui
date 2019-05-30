package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

var subdomain *ui.Entry

func AddAccountQuestionBox(wm *GuiWindow) *ui.Box {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	wm.Box1 = vbox

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	hbox.Append(ui.NewLabel("Enter your Subdomain or"), false)

	generate := CreateButton(wm, nil, nil, "Generate", "SUBDOMAIN", generateSubdomain)
	hbox.Append(generate, false)

	subdomain = ui.NewEntry()
	subdomain.SetReadOnly(false)
	subdomain.OnChanged(func(*ui.Entry) {
		log.Println("OK. subdomain =", subdomain.Text())
	})
	vbox.Append(subdomain, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	okButton := CreateButton(wm, nil, nil, "Create Subdomain Account", "SUBDOMAIN", addSubdomain)
	vbox.Append(okButton, false)

	return vbox
}

func generateSubdomain(b *ButtonMap) {
	log.Println("generateSubdomain START")
	subdomain.SetText("cust00013.wit.dev")
	log.Println("generateSubdomain END")
}

func addSubdomain(b *ButtonMap) {
	log.Println("generateSubdomain START")
	sub := subdomain.Text()
	log.Println("generateSubdomain subdomain =", sub)
	log.Println("generateSubdomain END")
}

func AddAccountBox(wm *GuiWindow) *ui.Box {
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

	okButton := CreateButton(wm, nil, nil, "Add Account", "ADD", nil)
	hboxButtons.Append(okButton, false)

	backButton := CreateButton(wm, nil, nil, "Back", "BACK", nil)
	hboxButtons.Append(backButton, false)

	return vbox
}
