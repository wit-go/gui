package gui

import "log"

import "github.com/gookit/config"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

func makeCloudInfoBox(custom func(*ButtonMap)) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	if (config.String("debugging") == "true") {
		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		hbox.Append(vbox, false)

		addDebuggingButtons(vbox, custom)

		hbox.Append(ui.NewVerticalSeparator(), false)
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	/*
	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)
	*/

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

	hostnamebox.Append(CreateButton("Edit", "EDIT", custom), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	agrid := ui.NewGrid()
	agrid.SetPadded(true)

	agrid.Append(ui.NewLabel("Accounts:"),   0, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Nickname"),    1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Username"),    2, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Domain Name"), 3, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	row := 1
	for account, _ := range config.StringMap("accounts") {
		// nickname	:= config.String("accounts." + account + ".nickname")
		username	:= config.String("accounts." + account + ".username")
		domainname	:= config.String("accounts." + account + ".domainname")

		hostname	:= config.String("accounts." + account + ".hostname")
		port		:= config.String("accounts." + account + ".port")

		a := account + " " + hostname + " " + domainname + " " + port + " " + username
		log.Println("ACCOUNT: ", a)

		agrid.Append(ui.NewLabel(account),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(username),   2, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(domainname), 3, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		l := CreateLoginButton(account, custom)
		agrid.Append(l, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		b := CreateAccountButton(account, custom)
		agrid.Append(b, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		row += 1
	}
	row += 1
	agrid.Append(ui.NewLabel(""),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	row += 1
	a := CreateButton("Add Account", "ADD", custom)
	agrid.Append(a, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	q := CreateButton("Quit", "QUIT", custom)
	agrid.Append(q, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(agrid, false)
	return hbox
}

func AddVmsTab(name string, count int) *TableData {
	var parts []TableColumnData

	human := 0

	tmp := TableColumnData{}
	tmp.CellType = "BG"
	tmp.Heading  = "background"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "name"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "hostname"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "IPv6"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "cpus"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "memory"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "BUTTON"
	tmp.Heading     = "Details"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	mh := AddTableTab(Data.cloudTab, 1, name, count, parts)
	return mh
}
