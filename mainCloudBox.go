package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// THIS IS NOT CLEAN

func makeCloudInfoBox(gw *GuiWindow) *GuiBox {
	var gb *GuiBox
	gb = new(GuiBox)
	gb.W = gw

	gb.EntryMap = make(map[string]*GuiEntry)
	gb.EntryMap["test"] = nil

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	gb.UiBox = hbox

	if (Data.Debug) {
		log.Println("makeCloudInfoBox() add debugging buttons")
		addDebuggingButtons(gb)
		hbox.Append(ui.NewVerticalSeparator(), false)
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	hostnamebox := ui.NewHorizontalBox()
	hostnamebox.SetPadded(true)
	vbox.Append(hostnamebox, false)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	hostnamebox.Append(entryForm, true)

	hostnameEntry :=  ui.NewEntry()
	entryForm.Append("hostname:", hostnameEntry, false)
	tmp := Data.Hostname + " (" + Data.IPv6 + ")"
	hostnameEntry.SetText(tmp)
	hostnameEntry.SetReadOnly(true)

	anew := CreateButton(gb, nil, nil, "Edit", "EDIT", nil)
	hostnamebox.Append(anew.B, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	agrid := ui.NewGrid()
	agrid.SetPadded(true)

	agrid.Append(ui.NewLabel("Accounts:"),   0, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Domain Name"), 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Email"),       2, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	row := 1

	for key, a := range Data.Config.Accounts {
		log.Println("account =          ", key, a)
		log.Println("Accounts[key] =    ", Data.Config.Accounts[key])
		log.Println("account.Nick =     ", Data.Config.Accounts[key].Nick)
		log.Println("account.Username = ", Data.Config.Accounts[key].Username)
		log.Println("account.Token =    ", Data.Config.Accounts[key].Token)

		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Domain),	1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Email),	2, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name := "Login " + Data.Config.Accounts[key].Nick
		l := CreateButton(gb, Data.Config.Accounts[key], nil, name, "LOGIN", nil)
		agrid.Append(l.B, 3, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name  = "Show " + Data.Config.Accounts[key].Nick
		b := CreateButton(gb, Data.Config.Accounts[key], nil, name, "SHOW", nil)
		agrid.Append(b.B, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		row += 1
	}

	row += 1
	agrid.Append(ui.NewLabel(""),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	row += 1
	a := CreateButton(gb, nil, nil, "Add Account", "ADD TAB", nil)
	agrid.Append(a.B, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	q := CreateButton(gb, nil, nil, "Quit", "QUIT", nil)
	agrid.Append(q.B, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(agrid, false)
	return gb
}
