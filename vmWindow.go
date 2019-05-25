package gui

import "log"
import "fmt"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

import pb "git.wit.com/wit/witProtobuf"

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

	createVmBox(VMtab, mouseClick, Data.CurrentVM)
	VMwin.Show()
}

func AddVmConfigureTab(name string, pbVM *pb.Event_VM) {
	createVmBox(Data.cloudTab, mouseClick, Data.CurrentVM)
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

func createVmBox(tab *ui.Tab, custom func(*ButtonMap), pbVM *pb.Event_VM) {
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

	hboxButtons.Append(CreateButton(nil, pbVM, "Power On",  "POWERON",  custom), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Power Off", "POWEROFF", custom), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Destroy",   "DESTROY",  custom), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "ping",      "PING",     runPingClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Console",   "XTERM",    runTestExecClick), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Save",      "SAVE",     custom), false)
	hboxButtons.Append(CreateButton(nil, pbVM, "Done",      "DONE",     custom), false)

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
