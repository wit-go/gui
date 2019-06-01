package gui

import "log"
// import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import "github.com/davecgh/go-spew/spew"
// var subdomain *ui.Entry

func AddEntry(box *GuiBox, name string) *GuiEntry {
	var ge *GuiEntry
	ge = new(GuiEntry)

	ue := ui.NewEntry()
	ue.SetReadOnly(false)
	ue.OnChanged(func(*ui.Entry) {
		log.Println("gui.AddEntry() OK. ue.Text() =", ue.Text())
	})
	box.UiBox.Append(ue, false)

	ge.UiEntry = ue
	box.EntryMap[name] = ge

	return ge
}

func AddAccountQuestionBox(gw *GuiWindow) *GuiBox {
	var gb *GuiBox
	gb = new(GuiBox)

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	// gw.Box1 = vbox
	gb.UiBox = vbox
	gb.W = gw

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	hbox.Append(ui.NewLabel("Enter your Subdomain or"), false)

	button1 := CreateButton(gb, nil, nil, "Generate", "SUBDOMAIN", generateSubdomain)
	button1.Box = gb
	hbox.Append(button1.B, false)

	AddEntry(gb, "SUBDOMAIN")
	// AddEntry(gb, "USERNAME")

	vbox.Append(ui.NewHorizontalSeparator(), false)

	button2 := CreateButton(gb, nil, nil, "Create Subdomain Account", "ADD", nil)
	button2.Box = gb
	vbox.Append(button2.B, false)

	return gb
}

func generateSubdomain(b *GuiButton) {
	log.Println("generateSubdomain START")
	if (b == nil) {
		log.Println("generateSubdomain ERROR b == nil")
		return
	}
	// subdomain.SetText("cust00013.wit.dev")

	txt := SetText(b.Box, "SUBDOMAIN", "cust001.testing.com.customers.wprod.wit.com")
	log.Println("generateSubdomain subdomain = ", txt)
	log.Println("generateSubdomain END")
}

/*
func addSubdomain(b *GuiButton) {
	log.Println("addSubdomain START")
	// sub := subdomain.Text()
	// log.Println("generateSubdomain subdomain =", sub)
	log.Println("addSubdomain END")
}
*/

func AddAccountBox(gw *GuiWindow) *GuiBox {
	var gb *GuiBox
	gb = new(GuiBox)

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	// gw.Box1 = vbox
	gb.UiBox = vbox

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

	okButton := CreateButton(gb, nil, nil, "Add Account", "ADD", nil)
	hboxButtons.Append(okButton.B, false)

	backButton := CreateButton(gb, nil, nil, "Back", "BACK", nil)
	hboxButtons.Append(backButton.B, false)

	return gb
}
