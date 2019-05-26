package gui

import "log"
import "time"
import "fmt"
// import "os"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

import "github.com/davecgh/go-spew/spew"

func makeCloudInfoBox(custom func(*ButtonMap)) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	if (Data.Debug) {
		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		hbox.Append(vbox, false)

		addDebuggingButtons(vbox, custom)

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

	hostnamebox.Append(CreateButton(nil, nil, "Edit", "EDIT", custom), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	agrid := ui.NewGrid()
	agrid.SetPadded(true)

	agrid.Append(ui.NewLabel("Accounts:"),   0, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Nickname"),    1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Username"),    2, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Domain Name"), 3, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	row := 1

	for key, foo := range Data.Config.Accounts {
		log.Println("account =          ", key, foo)
		log.Println("Accounts[key] =    ", Data.Config.Accounts[key])
		log.Println("account.Nick =     ", Data.Config.Accounts[key].Nick)
		log.Println("account.Username = ", Data.Config.Accounts[key].Username)
		log.Println("account.Token =    ", Data.Config.Accounts[key].Token)

		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Nick),	1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Username),	2, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Domain),	3, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name := "Login " + Data.Config.Accounts[key].Nick
		l := CreateButton(Data.Config.Accounts[key], nil, name, "LOGIN", custom)
		agrid.Append(l, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name  = "Show " + Data.Config.Accounts[key].Nick
		b := CreateButton(Data.Config.Accounts[key], nil, name, "SHOW", custom)
		agrid.Append(b, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		row += 1
	}

	row += 1
	agrid.Append(ui.NewLabel(""),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	row += 1
	a := CreateButton(nil, nil, "Add Account", "ADD", custom)
	agrid.Append(a, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	q := CreateButton(nil, nil, "Quit", "QUIT", custom)
	agrid.Append(q, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(agrid, false)
	return hbox
}

//
// THIS IS THE STANDARD VM DISPLAY TABLE
// This maps the 'human' indexed cells in the table
// to the machine's andlabs/libui values. That way
// if you want to work against column 4, then you
// can just reference 4 instead of the internal number
// which could be anything since TEXTCOLOR, TEXT, BG, etc
// fields use between 1 and 3 values internally
//
func AddVmsTab(name string, count int, a *pb.Account) *TableData {
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
	tmp.Heading  = "Details"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	mh := AddTableTab(Data.cloudTab, 1, name, count, parts, a)
	return mh
}

func ShowAccountQuestionTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = AddAccountQuestionBox(nil, mouseClick)
	Data.cloudTab.InsertAt("New Account?", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowAccountTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	// Create the things for the Account Tab
	var aTab GuiTabStructure
	Data.CurrentTab = &aTab
	AddAccountBox(&aTab)

	// Set the parents and data structure links
	aTab.me = Data.cloudTab
	aTab.parentWindow = Data.cloudWindow
	aTab.tabOffset = 0

// Data.smallBox = aTab.firstBox
	Data.cloudTab.InsertAt("Add Account", 0, aTab.firstBox)
	Data.cloudTab.SetMargined(0, true)
}

func ShowMainTab() {
	Data.cloudTab.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	Data.smallBox = makeCloudInfoBox(mouseClick)
	Data.cloudTab.InsertAt("Main", 0, Data.smallBox)
	Data.cloudTab.SetMargined(0, true)
}

func GoMainWindow() {
	ui.Main(makeCloudWindow)
}

func makeCloudWindow() {
	Data.cloudWindow = ui.NewWindow("", Data.Width, Data.Height, true)
	// cloudWindow.SetBorderless(true)

        // create a 'fake' button entry for the mouse clicks
	var newButtonMap ButtonMap
	newButtonMap.Action   = "QUIT"
	Data.AllButtons = append(Data.AllButtons, newButtonMap)

	Data.cloudWindow.OnClosing(func(*ui.Window) bool {
		mouseClick(&newButtonMap)
		return true
	})
	ui.OnShouldQuit(func() bool {
		mouseClick(&newButtonMap)
		return true
	})

	Data.cloudTab = ui.NewTab()
	Data.cloudWindow.SetChild(Data.cloudTab)
	Data.cloudWindow.SetMargined(true)

	Data.cloudBox = ShowSplashBox()

	Data.cloudTab.Append("WIT Splash", Data.cloudBox)
	Data.cloudTab.SetMargined(0, true)

	Data.cloudWindow.Show()
	Data.State = "splash"
}

func ShowVM() {
	name := Data.CurrentVM.Name
	log.Println("ShowVM() START Data.CurrentVM=", Data.CurrentVM)
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

	createVmBox(VMtab, Data.CurrentVM)
	VMwin.Show()
}

func AddVmConfigureTab(name string, pbVM *pb.Event_VM) {
	createVmBox(Data.cloudTab, Data.CurrentVM)
}

// makeEntryBox(box, "hostname:", "blah.foo.org") {
func makeEntryVbox(hbox *ui.Box, a string, b string, edit bool) {
	// Start 'Nickname' vertical box
	vboxN := ui.NewVerticalBox()
	vboxN.SetPadded(true)
	vboxN.Append(ui.NewLabel(a), false)

	entryNick := ui.NewEntry()
	entryNick.SetText(b)
	if (edit == false) {
		entryNick.SetReadOnly(true)
	}

	vboxN.Append(entryNick, false)

	entryNick.OnChanged(func(*ui.Entry) {
		log.Println("OK. TEXT WAS CHANGED TO =", entryNick.Text())
		// Data.AccNick = entryNick.Text()
	})
	hbox.Append(vboxN, false)
	// End 'Nickname' vertical box
}

func makeEntryHbox(hbox *ui.Box, a string, b string, edit bool) {
	// Start 'Nickname' vertical box
	hboxN := ui.NewHorizontalBox()
	hboxN.SetPadded(true)
	hboxN.Append(ui.NewLabel(a), false)

	entryNick := ui.NewEntry()
	entryNick.SetText(b)
	if (edit == false) {
		entryNick.SetReadOnly(true)
	}

	hboxN.Append(entryNick, false)

	entryNick.OnChanged(func(*ui.Entry) {
		log.Println("OK. TEXT WAS CHANGED TO =", entryNick.Text())
		// Data.AccNick = entryNick.Text()
	})
	hbox.Append(hboxN, false)
	// End 'Nickname' vertical box
}

func createVmBox(tab *ui.Tab, pbVM *pb.Event_VM) {
	log.Println("createVmBox() START")
	log.Println("createVmBox() pbVM.Name", pbVM.Name)
	spew.Dump(pbVM)
	if (Data.Debug) {
		spew.Dump(pbVM)
	}
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hboxAccount := ui.NewHorizontalBox()
	hboxAccount.SetPadded(true)
	vbox.Append(hboxAccount, false)

	// Add hostname entry box
	makeEntryVbox(hboxAccount, "hostname:",	pbVM.Hostname,			true)
	makeEntryVbox(hboxAccount, "IPv6:",	pbVM.IPv6,			true)
	makeEntryVbox(hboxAccount, "RAM:",	fmt.Sprintf("%d",pbVM.Memory),	true)
	makeEntryVbox(hboxAccount, "CPU:",	fmt.Sprintf("%d",pbVM.Cpus),	true)
	makeEntryVbox(hboxAccount, "Disk (GB):",	fmt.Sprintf("%d",pbVM.Disk),	true)
	makeEntryVbox(hboxAccount, "OS Image:",	pbVM.BaseImage,			true)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	hboxButtons := ui.NewHorizontalBox()
	hboxButtons.SetPadded(true)
	vbox.Append(hboxButtons, false)

	log.Println("pbVM =", pbVM)
	hboxButtons.Append(CreateButton(nil, pbVM, "Power On",  "POWERON",  mouseClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Power Off", "POWEROFF", mouseClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Destroy",   "DESTROY",  mouseClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "ping",      "PING",     runPingClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Console",   "XTERM",    runTestExecClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Save",      "SAVE",     mouseClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Done",      "DONE",     mouseClick), false)

	tab.Append(Data.CurrentVM.Name, vbox)
	tab.SetMargined(0, true)
}

func createAddVmBox(tab *ui.Tab, name string, custom func(*ButtonMap)) {
	log.Println("createAddVmBox() START")
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hboxAccount := ui.NewHorizontalBox()
	hboxAccount.SetPadded(true)
	vbox.Append(hboxAccount, false)

	// Add hostname entry box
	makeEntryHbox(hboxAccount, "hostname:",	"", true)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	hboxButtons := ui.NewHorizontalBox()
	hboxButtons.SetPadded(true)
	vbox.Append(hboxButtons, false)

	hboxButtons.Append(CreateButton(nil, nil, "Add Virtual Machine",	"CREATE", custom), false)
	hboxButtons.Append(CreateButton(nil, nil, "Cancel",		"DONE",   custom), false)

	tab.Append(name, vbox)
	tab.SetMargined(0, true)
}
